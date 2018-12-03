package main

import (
	"fmt"
	"time"

	"github.com/carlpett/aoc/utils"
)

type claim struct {
	id     int
	top    int
	left   int
	width  int
	height int
}

func main() {
	input := utils.MustReadStdinAsStringSlice()
	claims := make([]claim, 0)
	for _, str := range input {
		c := claim{}
		_, err := fmt.Sscanf(str, "#%d @ %d,%d: %dx%d", &c.id, &c.left, &c.top, &c.width, &c.height)
		if err != nil {
			panic(err)
		}
		claims = append(claims, c)
	}

	tA := time.Now()
	fmt.Printf("A: %d (in %v)\n", solveA(claims), time.Since(tA))
	tB := time.Now()
	fmt.Printf("B: %d (in %v)\n", solveB(claims), time.Since(tB))
}

func solveA(claims []claim) int {
	fabric := make(map[int]map[int]int)
	var overlap int
	for _, c := range claims {
		for x := c.left; x < c.left+c.width; x++ {
			if _, ok := fabric[x]; !ok {
				fabric[x] = make(map[int]int)
			}
			for y := c.top; y < c.top+c.height; y++ {
				fabric[x][y]++
				if fabric[x][y] == 2 {
					overlap++
				}
			}
		}
	}

	return overlap
}
func solveB(claims []claim) int {
	fabric := make(map[int]map[int][]int)
	nonOverlapped := make(map[int]bool)
	for _, c := range claims {
		nonOverlapped[c.id] = true
		for x := c.left; x < c.left+c.width; x++ {
			if _, ok := fabric[x]; !ok {
				fabric[x] = make(map[int][]int)
			}
			for y := c.top; y < c.top+c.height; y++ {
				fabric[x][y] = append(fabric[x][y], c.id)
				if len(fabric[x][y]) > 1 {
					for _, o := range fabric[x][y] {
						delete(nonOverlapped, o)
					}
				}
			}
		}
	}

	if len(nonOverlapped) != 1 {
		panic(nonOverlapped)
	}
	for c := range nonOverlapped {
		return c
	}
	return -1
}
