package main

import (
	"fmt"
	"strings"
	"time"

	"github.com/carlpett/aoc/utils"
)

type pos struct {
	x, y int
}
type bounds struct {
	minX, minY, maxX, maxY int
}

func main() {
	tS := time.Now()
	input := utils.MustReadStdinAsStringSlice()
	coords := make([]pos, len(input))
	bounds := bounds{
		minX: 1 << 32,
		minY: 1 << 32,
		maxX: 0,
		maxY: 0,
	}
	for idx, s := range input {
		parts := strings.Split(s, ", ")
		p := pos{utils.MustAtoi(parts[0]), utils.MustAtoi(parts[1])}
		coords[idx] = p

		if p.x < bounds.minX {
			bounds.minX = p.x
		}
		if p.y < bounds.minY {
			bounds.minY = p.y
		}
		if p.x > bounds.maxX {
			bounds.maxX = p.x
		}
		if p.y > bounds.maxY {
			bounds.maxY = p.y
		}
	}
	fmt.Printf("Setup in %v\n", time.Since(tS))

	tA := time.Now()
	fmt.Printf("A: %d (in %v)\n", solveA(coords, bounds), time.Since(tA))
	tB := time.Now()
	fmt.Printf("B: %d (in %v)\n", solveB(coords, bounds, 10000), time.Since(tB))
}

func solveA(coords []pos, b bounds) int {
	limited := make([]int, 0)
	for idx, pt := range coords {
		limitedIn := map[string]bool{
			"x+": false,
			"x-": false,
			"y+": false,
			"y-": false,
		}
		for _, cand := range coords {
			if pt.x < cand.x {
				if pt.x-cand.x < Abs(pt.y-cand.y) {
					limitedIn["x+"] = true
					//fmt.Printf("%d limited in x+ by %d\n", idx, cidx)
				}
			} else if pt.x > cand.x {
				if cand.x-pt.x < Abs(pt.y-cand.y) {
					limitedIn["x-"] = true
					//fmt.Printf("%d limited in x- by %d\n", idx, cidx)
				}
			}
			if pt.y < cand.y {
				if pt.y-cand.y < Abs(pt.x-cand.x) {
					limitedIn["y+"] = true
					//fmt.Printf("%d limited in y+ by %d\n", idx, cidx)
				}
			} else if pt.y > cand.y {
				if cand.y-pt.y < Abs(pt.x-cand.x) {
					limitedIn["y-"] = true
					//fmt.Printf("%d limited in y- by %d\n", idx, cidx)
				}
			}
		}
		anyUnlimited := false
		for _, l := range limitedIn {
			if !l {
				anyUnlimited = true
				break
			}
		}
		if !anyUnlimited {
			limited = append(limited, idx)
		}
		//fmt.Println(idx, anyUnlimited, limited)
	}

	areaSizes := make(map[int]int)
	for x := b.minX; x <= b.maxX; x++ {
		for y := b.minY; y <= b.maxY; y++ {
			smallestDist := 1 << 32
			var nearestIdx int
			var tie bool
			for idx, c := range coords {
				d := manhattan(c, pos{x, y})
				if d < smallestDist {
					//fmt.Printf("%d is now closest to %d,%d (%d)\n", idx, x, y, d)
					smallestDist = d
					nearestIdx = idx
					tie = false
				} else if d == smallestDist {
					tie = true
					//fmt.Printf("%d and %d are equally close to %d,%d (%d)\n", idx, nearestIdx, x, y, smallestDist)
				}
			}

			if !tie {
				areaSizes[nearestIdx]++
				if smallestDist == 0 {
					//fmt.Printf("*%d*", nearestIdx)
				} else {
					//fmt.Printf(" %d ", nearestIdx)
				}
			} else {
				//fmt.Print(" . ")
			}
		}
		//fmt.Println()
	}

	largest := 0
	largestIdx := 0
	for _, cand := range limited {
		if areaSizes[cand] > largest {
			largest = areaSizes[cand]
			largestIdx = cand
		}
	}
	//fmt.Println(areaSizes)
	fmt.Println(largestIdx)
	return largest
}

func manhattan(p1, p2 pos) int {
	return Abs(p1.x-p2.x) + Abs(p1.y-p2.y)
}

func Abs(a int) int {
	if a < 0 {
		return -a
	}
	return a
}

func solveB(coords []pos, b bounds, cutoff int) int {
	safeArea := 0
	for x := b.minX; x <= b.maxX; x++ {
		for y := b.minY; y <= b.maxY; y++ {
			cumDist := 0
			for _, c := range coords {
				cumDist += manhattan(c, pos{x, y})
			}

			if cumDist < cutoff {
				safeArea++
			}
		}
	}
	return safeArea
}
