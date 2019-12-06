package main

import (
	"fmt"
	"time"

	"github.com/carlpett/aoc/utils"
)

type character struct {
	hp    int
	dmg   int
	armor int
}

func main() {
	tS := time.Now()
	input := utils.MustReadStdinAsStringSlice()
	boss := character{}
	fmt.Sscanf(input[0], "Hit Points: %d", &boss.hp)
	fmt.Sscanf(input[1], "Damage: %d", &boss.dmg)
	fmt.Sscanf(input[2], "Armor: %d", &boss.armor)
	fmt.Printf("Setup in %v\n", time.Since(tS))

	tA := time.Now()
	fmt.Printf("A: %d (in %v)\n", solveA(boss), time.Since(tA))
	tB := time.Now()
	fmt.Printf("B: %d (in %v)\n", solveB(boss), time.Since(tB))
}

type item struct {
	desc  string
	dmg   int
	armor int
}
type itemClass int

const (
	classWeapon itemClass = iota
	classArmor
	classRing
)

var storeItems = map[itemClass]map[int]item{
	classWeapon: map[int]item{
		8:  item{"Dagger", 4, 0},
		10: item{"Shortsword", 5, 0},
		25: item{"Warhammer", 6, 0},
		40: item{"Longsword", 7, 0},
		74: item{"Greataxe", 8, 0},
	},
	classArmor: map[int]item{
		0:   item{"No armor", 0, 0},
		13:  item{"Leather", 0, 1},
		31:  item{"Chainmail", 0, 2},
		53:  item{"Splintmail", 0, 3},
		75:  item{"Bandedmail", 0, 4},
		102: item{"Platemail", 0, 5},
	},
	classRing: map[int]item{
		0:   item{"No ring", 0, 0},
		25:  item{"Damage +1", 1, 0},
		50:  item{"Damage +2", 2, 0},
		100: item{"Damage +3", 3, 0},
		20:  item{"Defense +1", 0, 1},
		40:  item{"Defense +2", 0, 2},
		80:  item{"Defense +3", 0, 3},
	},
}

func solveA(boss character) int {
	bestCost := 1 << 32
	for weaponCost, weapon := range storeItems[classWeapon] {
		for armorCost, armor := range storeItems[classArmor] {
			for ring1Cost, ring1 := range storeItems[classRing] {
				for ring2Cost, ring2 := range storeItems[classRing] {
					if ring1Cost > 0 && ring1Cost == ring2Cost {
						continue
					}
					c := weaponCost + armorCost + ring1Cost + ring2Cost
					player := character{100, weapon.dmg + ring1.dmg + ring2.dmg, armor.armor + ring1.armor + ring2.armor}
					if simulateOutcome(player, boss) && c < bestCost {
						bestCost = c
					}
				}
			}
		}
	}
	return bestCost
}

func solveB(boss character) int {
	highestCost := 0
	for weaponCost, weapon := range storeItems[classWeapon] {
		for armorCost, armor := range storeItems[classArmor] {
			for ring1Cost, ring1 := range storeItems[classRing] {
				for ring2Cost, ring2 := range storeItems[classRing] {
					if ring1Cost > 0 && ring1Cost == ring2Cost {
						continue
					}
					c := weaponCost + armorCost + ring1Cost + ring2Cost
					player := character{100, weapon.dmg + ring1.dmg + ring2.dmg, armor.armor + ring1.armor + ring2.armor}
					if !simulateOutcome(player, boss) && c > highestCost {
						highestCost = c
					}
				}
			}
		}
	}
	return highestCost
}

func simulateOutcome(player, boss character) bool {
	return (boss.hp / utils.Max(1, player.dmg-boss.armor)) <=
		(player.hp / utils.Max(1, boss.dmg-player.armor))
}
