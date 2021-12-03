package main

import (
	"fmt"
	"time"

	"github.com/carlpett/aoc/utils"
)

func main() {
	instructions := utils.MustReadStdinAsStringSlice()

	tA := time.Now()
	fmt.Printf("A: %d (in %v)\n", solveA(instructions), time.Since(tA))
	tB := time.Now()
	fmt.Printf("B: %d (in %v)\n", solveB(instructions), time.Since(tB))
}

func solveA(instructions []string) int {
	depth := 0
	distance := 0

	var dir string
	var steps int
	for _, i := range instructions {
		fmt.Sscanf(i, "%s %d", &dir, &steps)
		switch dir {
		case "up":
			depth -= steps
		case "down":
			depth += steps
		case "forward":
			distance += steps
		}
	}

	return depth * distance
}

func solveB(instructions []string) int {
	depth := 0
	distance := 0
	aim := 0

	var dir string
	var steps int
	for _, i := range instructions {
		fmt.Sscanf(i, "%s %d", &dir, &steps)
		switch dir {
		case "up":
			aim -= steps
		case "down":
			aim += steps
		case "forward":
			distance += steps
			depth += aim * steps
		}
	}

	return depth * distance
}
