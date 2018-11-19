package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"strings"
	"time"
)

func main() {
	tS := time.Now()
	b, err := ioutil.ReadAll(os.Stdin)
	if err != nil {
		panic(err)
	}
	lines := strings.Split(string(b), "\n")
	puzzle := make([][]string, len(lines))
	for idx, l := range lines {
		puzzle[idx] = strings.Split(l, "")
	}
	start := coord{}
	start.x = strings.IndexRune(lines[0], '|')

	fmt.Printf("Setup in %v\n", time.Since(tS))

	t := time.Now()
	chars, steps := solve(puzzle, start)
	fmt.Printf("A: %s\n", chars)
	fmt.Printf("B: %d\n", steps)
	fmt.Printf("(in %v)\n", time.Since(t))
}

type coord struct {
	x int
	y int
}

func solve(puzzle [][]string, pos coord) (string, int) {
	chars := ""
	dir := coord{0, 1}
	steps := 0
	for {
		switch puzzle[pos.y][pos.x] {
		case " ":
			return chars, steps
		case "|":
		case "-":
		case "+":
			dir = findDir(puzzle, pos, dir)
		default:
			chars += puzzle[pos.y][pos.x]
		}
		pos.x += dir.x
		pos.y += dir.y
		steps++
	}

	return chars, steps
}

var neighbours = []coord{{1, 0}, {0, 1}, {-1, 0}, {0, -1}}

func findDir(puzzle [][]string, pos coord, dir coord) coord {
	for _, n := range neighbours {
		if n.x == dir.x && n.y == dir.y {
			// Same dir as current, skip
			continue
		} else if n.x == -dir.x && n.y == -dir.y {
			// Don't reverse
			continue
		} else if pos.y+n.y < 0 || pos.y+n.y >= len(puzzle) ||
			pos.x+n.x < 0 || pos.x+n.x >= len(puzzle[0]) {
			// Out of bounds
			continue
		}
		if puzzle[pos.y+n.y][pos.x+n.x] != " " {
			return n
		}
	}
	panic("No direction found")
}
