package main

import (
	"fmt"
	"time"

	"github.com/carlpett/aoc/utils"
)

func main() {
	diagnostics := utils.MustReadStdinAsStringSlice()

	tA := time.Now()
	fmt.Printf("A: %d (in %v)\n", solveA(diagnostics), time.Since(tA))
	tB := time.Now()
	fmt.Printf("B: %d (in %v)\n", solveB(diagnostics), time.Since(tB))
}

func solveA(diagnostics []string) int {
	mostCommon := 0
	codeLen := len(diagnostics[0])

	for idx := 0; idx < codeLen; idx++ {
		cnt := map[rune]int{'0': 0, '1': 0}
		for _, d := range diagnostics {
			cnt[rune(d[idx])] += 1
		}
		if cnt['1'] > cnt['0'] {
			mostCommon += 1 << (codeLen - 1 - idx)
		}
	}
	mask := 1<<codeLen - 1
	leastCommon := mostCommon ^ mask

	return mostCommon * leastCommon
}

func solveB(diagnostics []string) int {
	oxygenCandidates := make([]string, len(diagnostics))
	co2Candidates := make([]string, len(diagnostics))
	copy(oxygenCandidates, diagnostics)
	copy(co2Candidates, diagnostics)

	keepMatching := func(candidates []string, idx int, val rune) []string {
		ret := make([]string, 0)
		for _, c := range candidates {
			if rune(c[idx]) == val {
				ret = append(ret, c)
			}
		}
		return ret
	}

	filter := func(candidates []string, keepHigh bool) int {
		var keepOnHigh, keepOnLow rune
		if keepHigh {
			keepOnHigh = '1'
			keepOnLow = '0'
		} else {
			keepOnHigh = '0'
			keepOnLow = '1'
		}
		for idx := 0; idx < len(diagnostics[0]); idx++ {
			cnt := map[rune]int{'0': 0, '1': 0}
			for _, d := range candidates {
				cnt[rune(d[idx])] += 1
			}

			switch {
			case cnt['1'] >= cnt['0']:
				candidates = keepMatching(candidates, idx, keepOnHigh)
			case cnt['1'] < cnt['0']:
				candidates = keepMatching(candidates, idx, keepOnLow)
			}

			if len(candidates) == 1 {
				var rating int
				fmt.Sscanf(candidates[0], "%b", &rating)
				return rating
			}
		}
		panic("Never found candidate")
	}

	oxygenRating := filter(oxygenCandidates, true)
	co2Rating := filter(co2Candidates, false)

	return oxygenRating * co2Rating
}
