package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"strings"
	"time"
)

func main() {
	bs, err := ioutil.ReadAll(os.Stdin)
	if err != nil {
		panic(err)
	}
	words := strings.Split(string(bs), "\n")

	tA := time.Now()
	fmt.Printf("A: %d (in %v)\n", solve(words, isNiceA), time.Since(tA))
	tB := time.Now()
	fmt.Printf("B: %d (in %v)\n", solve(words, isNiceB), time.Since(tB))
}

func solve(words []string, niceFn func(string) bool) int {
	n := 0
	for _, w := range words {
		if niceFn(w) {
			n++
		}
	}
	return n
}

func isNiceA(word string) bool {
	var hasTwoConsecutive, hasBadSubstring bool
	numVowels := 0
	const vowels = "aeiou"
	badSubstrings := []string{"ab", "cd", "pq", "xy"}
	for idx, _ := range word {
		if strings.IndexByte(vowels, word[idx]) != -1 {
			numVowels++
		}
		if idx > 0 {
			if word[idx] == word[idx-1] {
				hasTwoConsecutive = true
			}
			for _, ss := range badSubstrings {
				if ss == word[idx-1:idx+1] {
					hasBadSubstring = true
				}
			}
		}
	}

	return numVowels >= 3 &&
		hasTwoConsecutive &&
		!hasBadSubstring
}
func isNiceB(word string) bool {
	var hasPair, hasABARepeat bool
	seenPairs := make(map[string]int)
	for idx, _ := range word {
		if idx > 0 {
			lastTwo := word[idx-1 : idx+1]
			if seenIdx, ok := seenPairs[lastTwo]; ok && idx > seenIdx+1 {
				hasPair = true
			}
			seenPairs[lastTwo] = idx
		}
		if idx > 1 {
			if word[idx-2] == word[idx] {
				hasABARepeat = true
			}
		}
	}

	return hasPair && hasABARepeat
}
