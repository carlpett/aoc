package main

import (
	"fmt"
	"time"

	"github.com/carlpett/aoc/utils"
)

func main() {
	input := utils.MustReadStdinAsIntSlice()

	tA := time.Now()
	fmt.Printf("A: %d (in %v)\n", solveA(input), time.Since(tA))
	tB := time.Now()
	fmt.Printf("B: %d (in %v)\n", solveB(input), time.Since(tB))
}

func solveA(containers []int) int {
	return subsetSum(containers, 0, 150)
}

func solveB(containers []int) int {
	trackingSubsetSum(containers, 0, 150, []int{})
	lowest := 1 << 32
	numLow := 0
	for _, p := range paths {
		if len(p) < lowest {
			lowest = len(p)
			numLow = 1
		} else if len(p) == lowest {
			numLow++
		}
	}
	return numLow
}

func subsetSum(containers []int, startIdx int, volume int) int {
	options := 0
	for idx, c := range containers[startIdx:] {
		if c == volume {
			options++
		} else if c < volume {
			options += subsetSum(containers, startIdx+idx+1, volume-c)
		}
	}
	return options
}

var paths = make([][]int, 0)

func trackingSubsetSum(containers []int, startIdx int, volume int, path []int) {
	for idx, c := range containers[startIdx:] {
		if c == volume {
			paths = append(paths, append(path, startIdx+idx))
		} else if c < volume {
			trackingSubsetSum(containers, startIdx+idx+1, volume-c, append(path, startIdx+idx))
		}
	}
}
