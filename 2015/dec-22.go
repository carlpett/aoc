package main

import (
	"fmt"
	"time"

	"github.com/carlpett/aoc/utils"
)

type character struct {
	hp         int
	dmg        int
	magicArmor int
	mana       int
}

func main() {
	tS := time.Now()
	input := utils.MustReadStdinAsStringSlice()
	boss := character{}
	fmt.Sscanf(input[0], "Hit Points: %d", &boss.hp)
	fmt.Sscanf(input[1], "Damage: %d", &boss.dmg)
	fmt.Printf("Setup in %v\n", time.Since(tS))

	tA := time.Now()
	fmt.Printf("A: %d (in %v)\n", solveA(boss), time.Since(tA))
	//tB := time.Now()
	//fmt.Printf("B: %d (in %v)\n", solveB(boss), time.Since(tB))
}

type effect func(player, boss *character)
type spell struct {
	manaCost int
	duration int
	effect   effect
}

var spells = map[string]spell{
	"Magic Missile": spell{manaCost: 53, duration: -1, effect: func(player, boss *character) { boss.hp -= 4 }},
	"Drain":         spell{manaCost: 73, duration: -1, effect: func(player, boss *character) { player.hp += 2; boss.hp -= 2 }},
	"Shield":        spell{manaCost: 113, duration: 6, effect: func(player, boss *character) { player.magicArmor = 7 }},
	"Poison":        spell{manaCost: 173, duration: 6, effect: func(player, boss *character) { boss.hp -= 3 }},
	"Recharge":      spell{manaCost: 229, duration: 5, effect: func(player, boss *character) { player.mana += 101 }},
}

type activeEffect struct {
	name    string
	apply   effect
	endTime int
}

func solveA(boss character) int {
	// Test-data
	//player := character{hp: 10, mana: 250}
	//boss = character{hp: 14, dmg: 8}
	player := character{hp: 50, mana: 500}
	win, mana := simulate(0, 0, player, boss, nil, nil)
	if !win {
		panic(":O")
	}
	if mana != 953 {
		panic(fmt.Errorf("wrong result %d, expected %d", mana, 953))
	}
	return mana
}

// Yay globals
var bestSoFar = 1 << 32

func simulate(tIn int, spentMana int, player character, boss character, effects []activeEffect, sequence []string) (bool, int) {
	for t := tIn; ; t++ {
		fmt.Printf("t=%d Pointers: %p %p\n", t, &player, &boss)
		if spentMana >= bestSoFar {
			//fmt.Println("Aborting at t=", t, "with", spentMana, "spent casting", sequence)
			return false, spentMana
		}
		//fmt.Println("starting round", t, "Concurrent:", inProgress)
		fmt.Println("Before effects:", player, boss)
		player.magicArmor = 0 // Reset
		for idx := 0; idx < len(effects); idx++ {
			e := effects[idx]
			fmt.Println(t, "Applying", e.name)
			e.apply(&player, &boss)
			if e.endTime <= t {
				fmt.Println(t, "Removing", e.name)
				effects = append(effects[:idx], effects[idx+1:]...)
				idx--
				continue
			}
		}
		fmt.Println("After effects:", player, boss)
		if boss.hp <= 0 {
			// TODO: Why does the boss die without enough spells having been cast?!
			fmt.Println("Boss dies! Spent", spentMana, "in", t, "turns, cast", sequence, "Player status", player, "Boss status", boss)
			return true, spentMana
		}

		if t%2 == 0 {
			// Player's turn
			if player.mana < 53 {
				//fmt.Println("Player dies! No more mana", t)
				return false, spentMana
			}

			bestSpending := 1 << 32
			anyWin := false
		spellSelection:
			for name, spell := range spells {
				for _, ae := range effects {
					if ae.name == name {
						continue spellSelection
					}
				}

				if spell.manaCost <= player.mana {
					//fmt.Println("t=", t, "Casting", name, "after", sequence)
					nplayer := player
					nplayer.mana -= spell.manaCost
					won, spending := simulate(t+1,
						spentMana+spell.manaCost,
						nplayer, boss,
						append(effects, activeEffect{name, spell.effect, t + spell.duration}),
						append(sequence, name),
					)
					if won && spending < bestSpending {
						bestSpending = spending
						anyWin = true
					}
					if won && spending < bestSoFar {
						bestSoFar = spending
					}
				}
			}
			return anyWin, bestSpending
		} else {
			player.hp -= utils.Max(1, boss.dmg-player.magicArmor)
			//fmt.Println("t=", t, "Player is hit for", utils.Max(1, boss.dmg-player.magicArmor), ", hp now", player.hp)
			if player.hp <= 0 {
				//fmt.Println("Player dies after", sequence)
				return false, spentMana
			}
		}
	}
}

func solveB(boss character) int {
	return -1
}
