package main

import (
	"fmt"
	"time"

	"github.com/carlpett/aoc/utils"
)

func main() {
	tS := time.Now()
	input := utils.MustReadStdinAsInt()
	fmt.Printf("Setup in %v\n", time.Since(tS))

	tA := time.Now()
	fmt.Printf("A: %d (in %v)\n", solveA(input), time.Since(tA))
	tB := time.Now()
	fmt.Printf("B: %d (in %v)\n", solveB(input), time.Since(tB))
}

func solveA(input int) int {
	p := make([]int, input/10)
	for e := 1; e < input/10; e++ {
		for h := e; h < input/10; h += e {
			p[h] += e * 10
		}
	}
	for h, np := range p {
		if np > input {
			return h
		}
	}
	return -1
}

func solveB(input int) int {
	p := make([]int, input/11)
	for e := 1; e < input/11; e++ {
		deliveries := 0
		for h := e; h < input/11; h += e {
			deliveries++
			p[h] += e * 11
			if deliveries == 50 {
				break
			}
		}
	}
	for h, np := range p {
		if np > input {
			return h
		}
	}
	return -1
}
