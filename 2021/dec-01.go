package main

import (
	"fmt"
	"time"

	"github.com/carlpett/aoc/utils"
)

func main() {
	depths := utils.MustReadStdinAsIntSlice()

	tA := time.Now()
	fmt.Printf("A: %d (in %v)\n", solveA(depths), time.Since(tA))
	tB := time.Now()
	fmt.Printf("B: %d (in %v)\n", solveB(depths), time.Since(tB))
}

func solveA(depths []int) int {
	increases := 0
	last := int(1e6) // "Inf"-ish
	for _, d := range depths {
		if d > last {
			increases += 1
		}
		last = d
	}
	return increases
}

func solveB(depths []int) int {
	increases := 0
	last := depths[0] + depths[1] + depths[2]
	for idx := 1; idx < len(depths)-2; idx += 1 {
		cur := depths[idx] + depths[idx+1] + depths[idx+2]
		if cur > last {
			increases += 1
		}
		last = cur
	}
	return increases
}
