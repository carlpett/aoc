package main

import (
	"fmt"
	"strings"
	"time"

	"github.com/carlpett/aoc/utils"
)

func patternHash(L2, L1, C, R1, R2 bool) uint8 {
	h := uint8(0)
	if L2 {
		h |= 1 << 4
	}
	if L1 {
		h |= 1 << 3
	}
	if C {
		h |= 1 << 2
	}
	if R1 {
		h |= 1 << 1
	}
	if R2 {
		h |= 1
	}

	return h
}

func main() {
	tS := time.Now()
	input := utils.MustReadStdinAsStringSlice()
	stateStr := strings.Split(input[0], " ")[2]
	state := make(map[int]bool)
	for idx := range stateStr {
		state[idx] = stateStr[idx] == '#'
	}
	patterns := make(map[uint8]bool)
	for _, str := range input[2:] {
		parts := strings.Split(str, " => ")
		prev := parts[0]
		next := parts[1]
		patterns[patternHash(prev[0] == '#', prev[1] == '#', prev[2] == '#', prev[3] == '#', prev[4] == '#')] = next[0] == '#'
	}
	fmt.Printf("Setup in %v\n", time.Since(tS))

	tA := time.Now()
	fmt.Printf("A: %d (in %v)\n", solve(state, patterns, 20), time.Since(tA))
	tB := time.Now()
	fmt.Printf("B: %d (in %v)\n", solve(state, patterns, 50000000000), time.Since(tB))
}

func solve(state map[int]bool, patterns map[uint8]bool, gens int) int {
	minPot := 0
	maxPot := len(state)
	var lastSum, lastDiff int
	sameDiff := 0
	for gen := 0; gen < gens; gen++ {
		next := make(map[int]bool)
		nextMin := minPot
		nextMax := maxPot
		for pot := minPot - 2; pot <= maxPot+2; pot++ {
			h := patternHash(state[pot-2], state[pot-1], state[pot], state[pot+1], state[pot+2])
			next[pot] = patterns[h]

			if next[pot] && pot < minPot {
				nextMin = pot
			}
			if next[pot] && pot > maxPot {
				nextMax = pot
			}
		}
		state = next
		minPot = nextMin
		maxPot = nextMax

		sum := 0
		for pot, hasPlant := range state {
			if hasPlant {
				sum += pot
			}
		}

		// Forecasting
		diff := sum - lastSum
		lastSum = sum
		if diff == lastDiff {
			sameDiff++
			if sameDiff > 10 {
				return lastSum + diff*(gens-gen-1)
			}
		}
		lastDiff = diff
	}

	return lastSum
}
