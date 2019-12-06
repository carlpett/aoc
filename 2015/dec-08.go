package main

import (
	"fmt"
	"strconv"
	"time"

	"github.com/carlpett/aoc/utils"
)

func main() {
	lines := utils.MustReadStdinAsStringSlice()

	tA := time.Now()
	fmt.Printf("A: %d (in %v)\n", solveA(lines), time.Since(tA))
	tB := time.Now()
	fmt.Printf("B: %d (in %v)\n", solveB(lines), time.Since(tB))
}

func solveA(lines []string) int {
	chars := 0
	codes := 0
	for _, c := range lines {
		un, err := strconv.Unquote(c)
		if err != nil {
			panic(err)
		}
		chars += len(c)
		codes += len(un)
	}
	return chars - codes
}

func solveB(lines []string) int {
	chars := 0
	codes := 0
	for _, c := range lines {
		q := strconv.Quote(c)
		chars += len(c)
		codes += len(q)
	}
	return codes - chars
}
