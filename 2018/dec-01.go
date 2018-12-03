package main

import (
	"fmt"
	"time"

	"github.com/carlpett/aoc/utils"
)

func main() {
	changes := utils.MustReadStdinAsIntSlice()

	tA := time.Now()
	fmt.Printf("A: %d (in %v)\n", solveA(changes), time.Since(tA))
	tB := time.Now()
	fmt.Printf("B: %d (in %v)\n", solveB(changes), time.Since(tB))
}

func solveA(changes []int) int {
	freq := 0
	for _, c := range changes {
		freq += c
	}
	return freq
}

func solveB(changes []int) int {
	freq := 0
	seen := make(map[int]bool)
	for {
		for _, c := range changes {
			freq += c
			if _, ok := seen[freq]; ok {
				return freq
			}
			seen[freq] = true
		}
	}
	return freq
}
