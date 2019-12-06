package main

import (
	"fmt"
	"strings"
	"time"

	"github.com/carlpett/aoc/utils"
)

func main() {
	tS := time.Now()
	input := utils.MustReadStdinAsStringSlice()
	c1 := utils.MustAtoi(strings.Split(input[8], " ")[1])
	c2 := utils.MustAtoi(strings.Split(input[12], " ")[2])
	fmt.Printf("Setup in %v\n", time.Since(tS))

	t := time.Now()
	a, b := solve(c1, c2)
	fmt.Printf("A: %d\n", a)
	fmt.Printf("B: %d (in %v)\n", b, time.Since(t))
}

func solve(c1, c2 int) (int, int) {
	reg := make([]int, 6)
	seen := make(map[int]bool)
	var a, b int
	for {
		reg[2] = reg[3] | 0x10000
		reg[3] = c1
		for {
			reg[3] = ((reg[3] + reg[2]&0xff) & 0xffffff * c2) & 0xffffff
			if reg[2] < 256 {
				break
			}
			reg[2] /= 256
		}
		if len(seen) == 0 {
			a = reg[3]
		}

		if _, ok := seen[reg[3]]; ok {
			break
		}
		seen[reg[3]] = true
		b = reg[3]
	}

	return a, b
}
