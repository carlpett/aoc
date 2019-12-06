package main

import (
	"fmt"
	"io"
	"sort"
	"time"
)

type box struct {
	w, l, h int
}

func main() {
	boxes := make([]box, 0)
readboxes:
	for {
		var w, l, h int
		_, err := fmt.Scanf("%dx%dx%d\n", &w, &l, &h)
		switch {
		case err == io.EOF:
			break readboxes
		case err != nil:
			panic(err)
		}

		boxes = append(boxes, box{w, l, h})
	}

	tA := time.Now()
	fmt.Printf("A: %d (in %v)\n", solveA(boxes), time.Since(tA))
	tB := time.Now()
	fmt.Printf("B: %d (in %v)\n", solveB(boxes), time.Since(tB))
}

func solveA(boxes []box) int {
	sum := 0
	for _, box := range boxes {
		sum += 2*box.l*box.w + 2*box.w*box.h + 2*box.h*box.l
		s := sortedSideLengths(box)
		sum += s[0] * s[1]
	}
	return sum
}

func solveB(boxes []box) int {
	sum := 0
	for _, box := range boxes {
		s := sortedSideLengths(box)
		sum += (2*s[0] + 2*s[1]) + (s[0] * s[1] * s[2])
	}
	return sum
}

func sortedSideLengths(b box) []int {
	sides := []int{b.w, b.l, b.h}
	sort.Ints(sides)
	return sides
}
