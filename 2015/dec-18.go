package main

import (
	"fmt"
	"strings"
	"time"

	"github.com/carlpett/aoc/utils"
)

type lightGrid [][]bool

func newLightGrid(n int) lightGrid {
	lights := make([][]bool, n)
	for idx := range lights {
		lights[idx] = make([]bool, n)
	}
	return lights
}

func main() {
	tS := time.Now()
	input := utils.MustReadStdinAsStringSlice()
	lights := newLightGrid(len(input))
	for x, lineState := range input {
		for y, lightState := range lineState {
			if lightState == '#' {
				lights[x][y] = true
			}
		}
	}
	fmt.Printf("Setup in %v\n", time.Since(tS))

	tA := time.Now()
	fmt.Printf("A: %d (in %v)\n", solveA(lights, 100), time.Since(tA))
	tB := time.Now()
	fmt.Printf("B: %d (in %v)\n", solveB(lights, 100), time.Since(tB))
}

func solveA(lights lightGrid, steps int) int {
	current := lights
	var currentOn int
	for s := 0; s < steps; s++ {
		next := newLightGrid(len(lights))
		currentOn = 0
		for x, lightRow := range current {
			for y := range lightRow {
				n := current.countNeighbours(x, y)
				switch {
				case current[x][y] && n >= 2 && n <= 3,
					!current[x][y] && n == 3:
					next[x][y] = true
					currentOn++
				default:
					next[x][y] = false
				}
			}
		}
		current = next
	}

	return currentOn
}
func solveB(lights lightGrid, steps int) int {
	current := lights
	for s := 0; s < steps; s++ {
		current[0][0] = true
		current[0][len(lights)-1] = true
		current[len(lights)-1][0] = true
		current[len(lights)-1][len(lights)-1] = true

		next := newLightGrid(len(lights))
		for x, lightRow := range current {
			for y := range lightRow {
				n := current.countNeighbours(x, y)
				switch {
				case current[x][y] && n >= 2 && n <= 3,
					!current[x][y] && n == 3:
					next[x][y] = true
				default:
					next[x][y] = false
				}
			}
		}
		current = next
	}

	current[0][0] = true
	current[0][len(lights)-1] = true
	current[len(lights)-1][0] = true
	current[len(lights)-1][len(lights)-1] = true
	var currentOn int
	for _, row := range current {
		for _, state := range row {
			if state {
				currentOn++
			}
		}
	}

	return currentOn
}

func (lights lightGrid) countNeighbours(centX, centY int) int {
	n := 0
	for x := utils.Max(0, centX-1); x <= utils.Min(len(lights)-1, centX+1); x++ {
		for y := utils.Max(0, centY-1); y <= utils.Min(len(lights)-1, centY+1); y++ {
			if x == centX && y == centY {
				continue
			}
			if lights[x][y] {
				n++
			}
		}
	}
	return n
}

func (lights lightGrid) String() string {
	var sb strings.Builder
	for _, row := range lights {
		for _, state := range row {
			if state {
				sb.WriteRune('#')
			} else {
				sb.WriteRune('.')
			}
		}
		sb.WriteRune('\n')
	}
	return sb.String()
}
