package main

import (
	"fmt"
	"strings"
	"time"
	"unicode"

	"github.com/carlpett/aoc/utils"
)

func main() {
	tS := time.Now()
	input := utils.MustReadStdinAsString()
	fmt.Printf("Setup in %v\n", time.Since(tS))

	tA := time.Now()
	fmt.Printf("A: %d (in %v)\n", solveA(input), time.Since(tA))
	tB := time.Now()
	fmt.Printf("B: %d (in %v)\n", solveB(input), time.Since(tB))
}

func solveA(input string) int {
	for i := 0; i < len(input); {
		if i+1 < len(input) && input[i] != input[i+1] && unicode.ToUpper(rune(input[i])) == unicode.ToUpper(rune(input[i+1])) {
			input = input[:i] + input[i+2:]
			i = utils.Max(i-1, 0)
		} else {
			i++
		}
	}
	return len(input)
}

func solveB(input string) int {
	best := 1 << 32
	var bestP rune
	for p := 'a'; p <= 'z'; p++ {
		data := strings.Replace(input, string(p), "", -1)
		data = strings.Replace(data, string(unicode.ToUpper(p)), "", -1)
		l := solveA(data)
		if l < best {
			best = l
			bestP = p
		}
	}

	data := strings.Replace(input, string(bestP), "", -1)
	data = strings.Replace(data, string(unicode.ToUpper(bestP)), "", -1)

	return solveA(data)
}
