package main

import (
	"fmt"
	"time"
	"strings"

	"github.com/carlpett/aoc/utils"
)

type point struct { x,y int }
type line struct {
	a, b point
}
type grid map[int]map[int]int

func newGrid(min, max point) grid {
	g := make(map[int]map[int]int, max.x-min.x)
	for x := min.x; x <= max.x; x++ {
		g[x] = make(map[int]int, max.y-min.y)
	}
	return g
}
func (g grid) print(min, max point) {
	for col := min.y; col <= max.y; col++ {
		for row := min.x; row <= max.x; row++ {
			if g[row][col] > 0 {
				fmt.Print(g[row][col])
			} else {
				fmt.Print(".")
			}
		}
		fmt.Println()
	}
}

func main() {
	tS := time.Now()
	input := utils.MustReadStdinAsStringSlice()
	lines := make([]line, 0, len(input))
	minBound := point{x:int(1e6),y:int(1e6)}
	maxBound := point{x:-int(1e6),y:-int(1e6)}
	for _, l := range input {
		parts := strings.Split(l, " -> ")
		p1 := strings.Split(parts[0], ",")
		p2 := strings.Split(parts[1], ",")
		l := line{
			a: point {x: utils.MustAtoi(p1[0]), y: utils.MustAtoi(p1[1])},
			b: point {x: utils.MustAtoi(p2[0]), y: utils.MustAtoi(p2[1])},
		}
		lines = append(lines, l)
		minBound.x = utils.MinList(minBound.x, l.a.x, l.b.x)
		minBound.y = utils.MinList(minBound.y, l.a.y, l.b.y)
		maxBound.x = utils.MaxList(maxBound.x, l.a.x, l.b.x)
		maxBound.y = utils.MaxList(maxBound.y, l.a.y, l.b.y)
	}
	fmt.Printf("Setup done (in %v)\n", time.Since(tS))

	tA := time.Now()
	fmt.Printf("A: %d (in %v)\n", solveA(lines, minBound, maxBound), time.Since(tA))
	tB := time.Now()
	fmt.Printf("B: %d (in %v)\n", solveB(lines, minBound, maxBound), time.Since(tB))
}

func dir(i int) int {
	if i > 0 {
		return 1
	} else if i < 0 {
		return -1
	} else {
		return 0
	}
}
func gridpointsOn(l line) []point {
	cur := l.a
	pts := []point{cur}
	dx := dir(l.b.x-l.a.x)
	dy := dir(l.b.y-l.a.y)
	for cur != l.b {
		cur.x += dx
		cur.y += dy
		pts = append(pts, cur)
	}
	return pts
}

func solveA(lines []line, min, max point) int {
	vents := newGrid(min, max)
	danger := 0
	for _, line := range lines {
		// Skip diagonals
		if line.a.x != line.b.x && line.a.y != line.b.y {
			continue
		}
		for _, p := range gridpointsOn(line) {
			vents[p.x][p.y] = vents[p.x][p.y]+1
			if vents[p.x][p.y] == 2 {
				danger++
			}
		}
	}
	return danger
}

func solveB(lines []line, min, max point) int {
	vents := newGrid(min, max)
	danger := 0
	for _, line := range lines {
		for _, p := range gridpointsOn(line) {
			vents[p.x][p.y] = vents[p.x][p.y]+1
			if vents[p.x][p.y] == 2 {
				danger++
			}
		}
	}
	return danger
}
