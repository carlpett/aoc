package main

import (
	"fmt"
)

func main() {
	for _, n := range []int{1, 12, 23, 1024, 347991} {
		fmt.Printf("%d: %d\n", n, manhattan(n))
	}
}

// Bottom right is the last element of each turn
// 1 9 25 49 81 ... (2n+1)^2
func bottomRight(turn int) int {
	return (2*turn + 1) * (2*turn + 1)
}

// Which turn is n on?
func turn(n int) int {
	turn := 0
	for v := 0; v < n; turn++ {
		v = bottomRight(turn)
	}
	return turn - 1
}

func midpoints(turn int) []int {
	br := bottomRight(turn)
	sl := sideLength(turn)
	return []int{
		br - sl/2,
		br - sl*3/2 + 1,
		br - sl*5/2 + 2,
		br - sl*7/2 + 3,
	}
}
func distanceFromMidpoint(n int) int {
	t := turn(n)
	ms := midpoints(t)
	d := n
	for _, m := range ms {
		if abs(n-m) < d {
			d = abs(n - m)
		}
	}
	return d
}

func sideLength(turn int) int {
	return (2*turn + 1)
}

func manhattan(n int) int {
	r := turn(n)
	s := distanceFromMidpoint(n)
	return r + s
}

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}
