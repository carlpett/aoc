package main

import (
	"fmt"
	"regexp"
	"sort"
	"strings"
	"time"

	"github.com/carlpett/aoc/utils"
)

const (
	sideImmuneSystem = "Immune System"
	sideInfection    = "Infection"
)

func main() {
	tS := time.Now()
	input := utils.MustReadStdinAsStringSlice()

	armies := readGroups(input[1:], sideImmuneSystem, 0)
	offset := len(armies) + 3
	armies = append(armies, readGroups(input[offset:], sideInfection, len(armies))...)

	armiesA := make(groups, len(armies))
	armiesB := make(groups, len(armies))
	copy(armiesA, armies)
	copy(armiesB, armies)

	fmt.Printf("Setup in %v\n", time.Since(tS))

	tA := time.Now()
	fmt.Printf("A: %d (in %v)\n", solveA(armiesA), time.Since(tA))
	tB := time.Now()
	fmt.Printf("B: %d (in %v)\n", solveB(armiesB), time.Since(tB))
}

var pattern = regexp.MustCompile(`(?P<unitCount>\d+) units each with (?P<hitPoints>\d+) hit points (?:\((?P<weaknessesAndImmunities>[^)]+)\) )?with an attack that does (?P<damage>\d+) (?P<damageType>\w+) damage at initiative (?P<initiative>\d+)`)

func readGroups(input []string, side string, startIdx int) groups {
	groups := make(groups, 0)
	for idx, str := range input {
		if strings.TrimSpace(str) == "" {
			break
		}
		m := pattern.FindStringSubmatch(str)
		var weaknesses, immunities []string
		if m[3] != "" {
			for _, effect := range strings.Split(m[3], "; ") {
				parts := strings.SplitN(effect, " ", 3)
				types := strings.Split(parts[2], ", ")
				switch parts[0] {
				case "immune":
					immunities = types
				case "weak":
					weaknesses = types
				}
			}
		}
		groups = append(groups, group{
			id:         startIdx + idx,
			side:       side,
			units:      utils.MustAtoi(m[1]),
			hp:         utils.MustAtoi(m[2]),
			weaknesses: weaknesses,
			immunities: immunities,
			damage:     utils.MustAtoi(m[4]),
			damageType: m[5],
			initiative: utils.MustAtoi(m[6]),
		})
	}
	return groups
}

const debug = false

func Debug(s string) {
	if debug {
		fmt.Print(s)
	}
}

type group struct {
	id         int
	side       string
	units      int
	hp         int
	weaknesses []string
	immunities []string
	damage     int
	damageType string
	initiative int
}
type groups []group

func (g group) effectivePower() int {
	return g.damage * g.units
}
func (g group) calcDamageAgainst(g2 group) int {
	for _, im := range g2.immunities {
		if im == g.damageType {
			return 0
		}
	}
	dmg := g.effectivePower()
	for _, wk := range g2.weaknesses {
		if wk == g.damageType {
			return dmg * 2
		}
	}
	return dmg
}

func solveA(gs groups) int {
	rem, _ := simulate(gs)
	return rem
}

func solveB(gs groups) int {
	for boost := 1; ; boost++ {
		boosted := make(groups, len(gs))
		copy(boosted, gs)
		for idx := range boosted {
			if boosted[idx].side == sideImmuneSystem {
				boosted[idx].damage += boost
			}
		}
		ep, side := simulate(boosted)
		if side == sideImmuneSystem {
			return ep
		}
	}
}

func simulate(gs groups) (int, string) {
	groupIdxs := make([]int, len(gs))
	for idx := range groupIdxs {
		groupIdxs[idx] = idx
	}
	for t := 0; ; t++ {
		Debug(fmt.Sprintln("Immune System:"))
		for _, g := range gs {
			if g.units <= 0 {
				continue
			}
			if g.side == sideImmuneSystem {
				Debug(fmt.Sprintf("Group %d contains %d units (ep %d)\n", g.id, g.units, g.effectivePower()))
			}
		}
		Debug(fmt.Sprintln("Infection:"))
		for _, g := range gs {
			if g.units <= 0 {
				continue
			}
			if g.side == sideInfection {
				Debug(fmt.Sprintf("Group %d contains %d units (ep %d)\n", g.id, g.units, g.effectivePower()))
			}
		}
		Debug(fmt.Sprintln())

		// Target selection
		sort.Slice(groupIdxs, func(i, j int) bool {
			return gs[groupIdxs[i]].effectivePower() > gs[groupIdxs[j]].effectivePower() ||
				(gs[groupIdxs[i]].effectivePower() == gs[groupIdxs[j]].effectivePower() &&
					gs[groupIdxs[i]].initiative > gs[groupIdxs[j]].initiative)
		})
		targets := make(map[int]int)
		targeted := make(map[int]bool)
		for _, grpIdx := range groupIdxs {
			grp := gs[grpIdx]
			if grp.units <= 0 {
				continue
			}

			var bestTarget, bestDamage, bestEffectivePower, bestInitiative int
			var foundTarget bool
			for _, potTarget := range gs {
				if potTarget.units <= 0 || targeted[potTarget.id] || potTarget.side == grp.side {
					continue
				}
				d := grp.calcDamageAgainst(potTarget)
				Debug(fmt.Sprintf("%s group %d would deal defending group %d %d damage\n", grp.side, grp.id, potTarget.id, d))
				if d > bestDamage ||
					d == bestDamage && potTarget.effectivePower() > bestEffectivePower ||
					d == bestDamage && potTarget.effectivePower() > bestEffectivePower && potTarget.initiative == bestInitiative {
					bestDamage = d
					bestTarget = potTarget.id
					bestEffectivePower = potTarget.effectivePower()
					bestInitiative = potTarget.initiative
					foundTarget = true
				}
			}
			if foundTarget && bestDamage > 0 {
				targets[grp.id] = bestTarget
				targeted[bestTarget] = true
			}
		}
		Debug(fmt.Sprintln())

		// Attacking
		sort.Slice(groupIdxs, func(i, j int) bool {
			return gs[groupIdxs[i]].initiative > gs[groupIdxs[j]].initiative
		})
		totalKills := 0
		for _, grpIdx := range groupIdxs {
			grp := gs[grpIdx]
			if grp.units <= 0 {
				continue
			}
			tgt, hasTarget := targets[grp.id]
			if !hasTarget {
				continue
			}

			t := gs[tgt]
			kills := utils.Min(t.units, grp.calcDamageAgainst(t)/t.hp)
			gs[targets[grp.id]].units -= kills
			Debug(fmt.Sprintf("%s group %d attacks defending group %d, killing %d units\n", grp.side, grp.id, targets[grp.id], kills))
			totalKills += kills
		}

		if totalKills == 0 {
			return 0, "Stalemate"
		}

		epSum := make(map[string]int)
		for _, g := range gs {
			if g.units <= 0 {
				continue
			}
			epSum[g.side] += g.units
		}
		switch {
		case epSum[sideImmuneSystem] == 0:
			return epSum[sideInfection], sideInfection
		case epSum[sideInfection] == 0:
			return epSum[sideImmuneSystem], sideImmuneSystem
		}

		Debug(fmt.Sprintln("\n---"))
	}
}
