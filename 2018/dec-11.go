package main

import (
	"fmt"
	"time"

	"github.com/carlpett/aoc/utils"
)

func genGrid(serial int) [][]int {
	grid := make([][]int, 301)
	for x := 0; x < len(grid); x++ {
		grid[x] = make([]int, len(grid))
		for y := 0; y < len(grid); y++ {
			if x > 0 && x < 300 && y > 0 && y < 300 {
				rack := x + 10
				powerLevel := (rack*y + serial) * rack
				grid[x][y] = (powerLevel/100)%10 - 5
			}
		}
	}
	return grid
}

func main() {
	tS := time.Now()
	input := utils.MustReadStdinAsInt()
	grid := genGrid(input)
	fmt.Printf("Setup in %v\n", time.Since(tS))

	tA := time.Now()
	cx, cy := solveA(grid)
	fmt.Printf("A: %d,%d (in %v)\n", cx, cy, time.Since(tA))
	tB := time.Now()
	cx, cy, sz := solveB(grid)
	fmt.Printf("B: %d,%d,%d (in %v)\n", cx, cy, sz, time.Since(tB))
}

func solveA(grid [][]int) (int, int) {
	var bestX, bestY, bestSum int
	for x := 1; x < len(grid)-1; x++ {
		for y := 1; y < len(grid)-1; y++ {
			sum := 0
			for sX := x - 1; sX <= x+1; sX++ {
				for sY := y - 1; sY <= y+1; sY++ {
					sum += grid[sX][sY]
				}
			}
			if sum > bestSum {
				bestSum = sum
				bestX = x
				bestY = y
			}
		}
	}
	return bestX - 1, bestY - 1
}

func solveB(grid [][]int) (int, int, int) {
	sat := make([][]int, len(grid))
	for x := 0; x < len(sat); x++ {
		sat[x] = make([]int, len(grid))
		for y := 0; y < len(sat); y++ {
			sat[x][y] = grid[x][y]
			if x > 0 {
				sat[x][y] += sat[x-1][y]
			}
			if y > 0 {
				sat[x][y] += sat[x][y-1]
			}
			if x > 0 && y > 0 {
				sat[x][y] -= sat[x-1][y-1]
			}
		}
	}

	var bestX, bestY, bestSize, bestSum int
	for sz := 1; sz <= len(sat)-2; sz++ {
		for x := sz; x < len(sat)-1; x++ {
			for y := sz; y < len(sat)-1; y++ {
				sum := sat[x][y] -
					sat[utils.Max(0, x-sz)][y] -
					sat[x][utils.Max(0, y-sz)] +
					sat[utils.Max(0, x-sz)][utils.Max(0, y-sz)]

				if sum > bestSum {
					bestSum = sum
					bestX = x - sz + 1
					bestY = y - sz + 1
					bestSize = sz
				}
			}
		}
	}
	return bestX, bestY, bestSize
}
