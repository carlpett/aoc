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
	fmt.Printf("A: %d (in %v)\n", solve(input, 40), time.Since(tA))
	tB := time.Now()
	fmt.Printf("B: %d (in %v)\n", solve(input, 50), time.Since(tB))
}

func solve(input string, steps int) int {
	data := input
	for i := 0; i < steps; i++ {
		data = lookAndSay(data)
	}
	return len(data)
}

func lookAndSay(s string) string {
	prev := rune(s[0])
	runLen := 1
	output := strings.Builder{}
	for _, c := range s[1:] {
		if c == prev {
			runLen++
		} else {
			output.WriteString(fmt.Sprintf("%d%c", runLen, prev))
			prev = c
			runLen = 1
		}
	}
	output.WriteString(fmt.Sprintf("%d%c", runLen, prev))

	return output.String()
}
