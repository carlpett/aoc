package main

import (
	"fmt"
	"regexp/syntax"
	"time"

	"github.com/carlpett/aoc/utils"
)

func main() {
	tS := time.Now()
	input := utils.MustReadStdinAsString()
	re, err := syntax.Parse(input, 0)
	if err != nil {
		panic(err)
	}
	fmt.Printf("Setup in %v\n", time.Since(tS))

	t := time.Now()
	a, b := solve(re)
	fmt.Printf("A: %d\n", a)
	fmt.Printf("B: %d (in %v)\n", b, time.Since(t))
}

type coord struct {
	x, y int
}

var dir = map[rune]coord{
	'N': coord{0, -1},
	'S': coord{0, 1},
	'E': coord{1, 0},
	'W': coord{-1, 0},
}

func solve(re *syntax.Regexp) (int, int) {
	m := make(map[coord]int)
	walkRegexp(re, m, coord{}, 0)

	var a, b int
	for _, d := range m {
		if d > a {
			a = d
		}
		if d >= 918 {
			b++
		}
	}
	return a, b
}

func walkRegexp(re *syntax.Regexp, roomDist map[coord]int, cur coord, curDist int) (coord, int) {
	switch re.Op {
	case syntax.OpLiteral:
		for _, d := range re.String() {
			cur.x += dir[d].x
			cur.y += dir[d].y
			curDist++

			dist, visited := roomDist[cur]
			if !visited || dist > curDist {
				roomDist[cur] = curDist
			} else {
				curDist = dist
			}
		}
		//fmt.Println("End at", cur)
		//printMap(roomDist, cur)
	case syntax.OpConcat:
		next := cur
		nextDist := curDist
		for _, s := range re.Sub {
			next, nextDist = walkRegexp(s, roomDist, next, nextDist)
		}
	case syntax.OpAlternate, syntax.OpCapture:
		for _, s := range re.Sub {
			walkRegexp(s, roomDist, cur, curDist)
		}
	}

	return cur, curDist
}

func printMap(m map[coord]int, cur coord) {
	minX := 1 << 32
	minY := 1 << 32
	maxX := -1 << 32
	maxY := -1 << 32
	for k := range m {
		if k.x < minX {
			minX = k.x
		}
		if k.y < minY {
			minY = k.y
		}
		if k.x > maxX {
			maxX = k.x
		}
		if k.y > maxY {
			maxY = k.y
		}
	}
	for y := minY; y <= maxY; y++ {
		for x := minX; x <= maxX; x++ {
			if x == 0 && y == 0 {
				fmt.Print("S")
			} else if x == cur.x && y == cur.y {
				fmt.Print("C")
			} else {
				if _, ok := m[coord{x, y}]; ok {
					fmt.Print(".")
				} else {
					fmt.Print(" ")
				}
			}
		}
		fmt.Println()
	}
}
