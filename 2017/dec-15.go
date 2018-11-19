package main

import (
	"fmt"
	"math"
	"time"
)

type generator struct {
	factor              int
	value               int
	divisibiliyCriteria int
}

const (
	divider = math.MaxInt32
	bitmask = math.MaxUint16
)

func (g *generator) next(picky bool) int {
	for {
		g.value = (g.value * g.factor) % divider
		if !picky || g.value%g.divisibiliyCriteria == 0 {
			return g.value
		}
	}
}

func main() {
	tS := time.Now()
	var aStart, bStart int
	_, err := fmt.Scanf("%d %d", &aStart, &bStart)
	if err != nil {
		panic(err)
	}

	gA := generator{factor: 16807, value: aStart, divisibiliyCriteria: 4}
	gB := generator{factor: 48271, value: bStart, divisibiliyCriteria: 8}
	fmt.Printf("Setup in %v\n", time.Since(tS))

	tA := time.Now()
	fmt.Printf("A: %d (in %v)\n", solve(gA, gB, 40e6, false), time.Since(tA))
	tB := time.Now()
	fmt.Printf("B: %d (in %v)\n", solve(gA, gB, 5e6, true), time.Since(tB))
}

func solve(gA, gB generator, n int, picky bool) int {
	m := 0
	for i := 0; i < n; i++ {
		if gA.next(picky)&bitmask == gB.next(picky)&bitmask {
			m++
		}
	}
	return m
}
