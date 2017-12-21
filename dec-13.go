package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"strings"
	"time"
)

type layer struct {
	d int
	r int
}

func main() {
	tS := time.Now()
	b, err := ioutil.ReadAll(os.Stdin)
	if err != nil {
		panic(err)
	}
	lines := strings.Split(strings.TrimSpace(string(b)), "\n")
	layers := make([]layer, len(lines))
	for idx, line := range lines {
		l := layer{}
		fmt.Sscanf(line, "%d: %d", &l.d, &l.r)
		layers[idx] = l
	}
	fmt.Printf("Setup in %v\n", time.Since(tS))

	tA := time.Now()
	s, _ := solveA(layers, 0)
	fmt.Printf("A: %d (in %v)\n", s, time.Since(tA))
	tB := time.Now()
	fmt.Printf("B: %d (in %v)\n", solveB(layers), time.Since(tB))
}

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

var positions = make(map[int][]int)

func scannerPos(t, r int) int {
	var ps []int
	if _, ok := positions[r]; !ok {
		ps = make([]int, 2*r-1)
		for i := 0; i < 2*r-1; i++ {
			if i >= r {
				ps[i] = r - i%r - 2
			} else {
				ps[i] = i
			}
		}
		positions[r] = ps
	} else {
		ps = positions[r]
	}
	return ps[t%(2*(r-1))]
}

func solveA(layers []layer, t int) (int, bool) {
	caught := false
	severity := 0
	lastDepth := 0
	for _, l := range layers {
		t += l.d - lastDepth
		if scannerPos(t, l.r) == 0 {
			caught = true
			severity += l.r * l.d
		}
		lastDepth = l.d
	}
	return severity, caught
}

func solveB(layers []layer) int {
	t := 0
	for {
		_, caught := solveA(layers, t)
		if !caught {
			return t
		}
		t++
	}
}
