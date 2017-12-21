package main

import (
	"fmt"
)

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

func main() {
	fmt.Printf("B: %d\n", solveB(347991))
}

var steppings = []coord{{0, 1}, {-1, 0}, {0, -1}, {1, 0}}

func genCoordinates(t int) []coord {
	c := coord{t, -t + 1}
	cs := []coord{c}
	for i := 0; i < len(steppings); i++ {
		for j := 0; abs(c.x+steppings[i].x) <= t && abs(c.y+steppings[i].y) <= t; j++ {
			c.x += steppings[i].x
			c.y += steppings[i].y
			cs = append(cs, coord{c.x, c.y})
		}
	}
	return cs
}

func elemIndex(x, y int) string {
	return fmt.Sprintf("%d,%d", x, y)
}

func solveB(n int) int {
	elems := map[string]int{"0,0": 1}
	i := 2
	x := 0
	y := 0
	for {
		v := get(x, y, elems)
		if v > n {
			return v
		}
		elems[elemIndex(x, y)] = v
		x, y = getNextCoords(i)
		i++
	}
	return 0
}
func get(x, y int, elems map[string]int) int {
	n := 0
	s := []int{-1, 0, 1}
	for _, xS := range s {
		for _, yS := range s {
			if v, ok := elems[fmt.Sprintf("%d,%d", x+xS, y+yS)]; ok {
				n += v
			}
		}
	}
	return n
}

type coord struct {
	x int
	y int
}

var turnCoords = make(map[int][]coord)

func getNextCoords(idx int) (int, int) {
	t := turn(idx)
	if _, ok := turnCoords[t]; !ok {
		turnCoords[t] = genCoordinates(t)
	}
	start := bottomRight(t-1) + 1
	offset := idx - start
	return turnCoords[t][offset].x, turnCoords[t][offset].y
}

func turn(n int) int {
	turn := 0
	for v := 0; v < n; turn++ {
		v = bottomRight(turn)
	}
	return turn - 1
}

func bottomRight(turn int) int {
	return (2*turn + 1) * (2*turn + 1)
}

//////////////////
func sideLength(turn int) int {
	return (2*turn + 1)
}

func side(n int) int {
	t := turn(n)
	br := bottomRight(t)
	turnEndOffset := br - n
	return 4 - turnEndOffset/(sideLength(t)-1)
}
func sideOffset(n int) int {
	t := turn(n)
	br := bottomRight(t)
	turnEndOffset := br - n
	return sideLength(t)/2 - turnEndOffset%sideLength(t)
}

const (
	E int = iota + 1
	N
	W
	S
)
