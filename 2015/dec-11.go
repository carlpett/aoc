package main

import (
	"fmt"
	"strconv"
	"strings"
	"time"

	"github.com/carlpett/aoc/utils"
)

func main() {
	input := utils.MustReadStdinAsString()

	tA := time.Now()
	slnA := solve(input)
	fmt.Printf("A: %s (in %v)\n", slnA, time.Since(tA))
	tB := time.Now()
	fmt.Printf("B: %s (in %v)\n", solve(slnA), time.Since(tB))
}

func solve(input string) string {
	pw := input
	for {
		numpw, err := strconv.ParseInt(pw, 36, 64)
		if err != nil {
			panic(err)
		}
		numpw++
		pw = strconv.FormatInt(numpw, 36)
		pw = strings.Replace(pw, "0", "a", -1)
		if check(pw) {
			return pw
		}
	}
}

func check(pw string) bool {
	hasStraight := false
	hasOnlyValidChars := true
	hasNonOverlappingPairs := false

	pairs := make(map[string]bool)
	for idx := range pw {
		if idx >= 2 {
			if pw[idx-2]+1 == pw[idx-1] &&
				pw[idx-1]+1 == pw[idx] {
				hasStraight = true
			}
		}
		if strings.IndexByte("iol", pw[idx]) != -1 {
			hasOnlyValidChars = false
		}
		if idx > 1 && pw[idx-1] == pw[idx] {
			lastTwo := pw[idx-1 : idx+1]
			pairs[lastTwo] = true
			if len(pairs) > 1 {
				hasNonOverlappingPairs = true
			}
		}
	}
	return hasStraight && hasOnlyValidChars && hasNonOverlappingPairs
}
