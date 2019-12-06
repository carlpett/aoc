package main

import (
	"fmt"
	"math/rand"
	"strings"
	"time"

	"github.com/carlpett/aoc/utils"
)

type conversion struct {
	from, to string
}

func main() {
	tS := time.Now()
	input := utils.MustReadStdinAsStringSlice()
	conversions := make([]conversion, 0, len(input)-1)
	for _, str := range input[:len(input)-2] {
		parts := strings.Split(str, " => ")
		conversions = append(conversions, conversion{parts[0], parts[1]})
	}
	medicine := input[len(input)-1]

	fmt.Printf("Setup in %v\n", time.Since(tS))

	tA := time.Now()
	fmt.Printf("A: %d (in %v)\n", solveA(conversions, medicine), time.Since(tA))
	tB := time.Now()
	fmt.Printf("B: %d (in %v)\n", solveB(conversions, medicine), time.Since(tB))
}

func solveA(conversions []conversion, medicine string) int {
	outputs := make(map[string]bool)
	for _, c := range conversions {
		for _, idx := range StringFindAll(medicine, c.from) {
			output := medicine[:idx] + strings.Replace(medicine[idx:], c.from, c.to, 1)
			outputs[output] = true
		}
	}
	return len(outputs)
}

type foundMolecule struct {
	formula string
	depth   int
}

func solveB(conversions []conversion, medicine string) int {
	var current foundMolecule
	seen := make(map[string]bool)
	maxDepth := 0

	newOrder := rand.Perm(len(conversions))
	convs := make([]conversion, len(conversions))
	for i1, i2 := range newOrder {
		convs[i1] = conversions[i2]
	}

	for candidates := []foundMolecule{{medicine, 0}}; len(candidates) > 0; {
		current, candidates = candidates[0], candidates[1:]
		if current.depth > maxDepth {
			maxDepth = current.depth
		}
		for _, c := range convs {
			for _, idx := range StringFindAll(current.formula, c.to) {
				output := current.formula[:idx] + strings.Replace(current.formula[idx:], c.to, c.from, 1)
				if output == "e" {
					return current.depth + 1
				}
				if !seen[output] {
					seen[output] = true
					candidates = append([]foundMolecule{{output, current.depth + 1}}, candidates...)
				}
			}
		}
	}

	return -1
}

func StringFindAll(haystack string, needle string) []int {
	indices := make([]int, 0)
	for idx := range haystack {
		if strings.HasPrefix(haystack[idx:], needle) {
			indices = append(indices, idx)
		}
	}
	return indices
}
