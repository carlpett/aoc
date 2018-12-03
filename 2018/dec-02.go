package main

import (
	"fmt"
	"strings"
	"time"

	"github.com/carlpett/aoc/utils"
)

func main() {
	lines := utils.MustReadStdinAsStringSlice()

	tA := time.Now()
	fmt.Printf("A: %d (in %v)\n", solveA(lines), time.Since(tA))
	tB := time.Now()
	fmt.Printf("B: %s (in %v)\n", solveB(lines), time.Since(tB))
}

func solveA(boxIds []string) int {
	var nTwo, nThree int

	for _, id := range boxIds {
		chars := make(map[rune]int)
		for _, c := range id {
			chars[c]++
		}
		var countedTwo, countedThree bool
		for _, n := range chars {
			if n == 2 && !countedTwo {
				countedTwo = true
				nTwo++
			}
			if n == 3 && !countedThree {
				countedThree = true
				nThree++
			}
		}
	}
	return nTwo * nThree
}

func solveB(boxIds []string) string {
	bestLen := 0
	var bestCommon string
	for idxA, boxA := range boxIds {
		for _, boxB := range boxIds[idxA+1:] {
			common := strings.Builder{}
			for i := 0; i < len(boxA); i++ {
				if boxA[i] == boxB[i] {
					common.WriteByte(boxA[i])
				}
			}
			if common.Len() > bestLen {
				bestCommon = common.String()
				bestLen = len(bestCommon)
			}
		}
	}
	return bestCommon
}
