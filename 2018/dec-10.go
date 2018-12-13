package main

import (
	"fmt"
	"time"

	"github.com/carlpett/aoc/utils"
)

type star struct {
	x, y   int
	dx, dy int
}

func main() {
	tS := time.Now()
	input := utils.MustReadStdinAsStringSlice()
	stars := make([]star, len(input))
	for idx, str := range input {
		var s star
		_, err := fmt.Sscanf(str, "position=<%d, %d> velocity=<%d, %d>", &s.x, &s.y, &s.dx, &s.dy)
		if err != nil {
			panic(err)
		}
		stars[idx] = s
	}
	fmt.Printf("Setup in %v\n", time.Since(tS))

	tA := time.Now()
	fmt.Printf("A: %s (in %v)\n", solveA(stars), time.Since(tA))
	//tB := time.Now()
	//fmt.Printf("B: %s (in %v)\n", solve(players, lastValue*100), time.Since(tB))
}

func solveA(stars []star) string {
	var xMin, xMax, yMin, yMax, t int
	for t = 0; utils.Abs(yMin-yMax) != 9; t++ {
		yMin = 1 << 31
		yMax = -1 << 31
		xMin = 1 << 31
		xMax = -1 << 31
		for idx := range stars {
			stars[idx].x += stars[idx].dx
			stars[idx].y += stars[idx].dy
			if stars[idx].x > xMax {
				xMax = stars[idx].x
			}
			if stars[idx].x < xMin {
				xMin = stars[idx].x
			}
			if stars[idx].y > yMax {
				yMax = stars[idx].y
			}
			if stars[idx].y < yMin {
				yMin = stars[idx].y
			}
		}
	}
	for y := yMin; y <= yMax; y++ {
	xLoop:
		for x := xMin; x < xMax; x++ {
			for _, s := range stars {
				if s.x == x && s.y == y {
					fmt.Print("#")
					continue xLoop
				}
			}
			fmt.Print(".")
		}
		fmt.Println("")
	}
	fmt.Println(t)
	return "hi"
}
