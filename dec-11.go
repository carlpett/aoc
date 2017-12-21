package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"strings"
	"time"
)

func main() {
	tS := time.Now()
	b, err := ioutil.ReadAll(os.Stdin)
	if err != nil {
		panic(err)
	}
	steps := strings.Split(strings.TrimSpace(string(b)), ",")
	dirs := make(map[string]int)
	for _, s := range steps {
		dirs[s]++
	}
	fmt.Printf("Setup in %v\n", time.Since(tS))

	tA := time.Now()
	fmt.Printf("A: %d (in %v)\n", solveA(dirs), time.Since(tA))
	tB := time.Now()
	fmt.Printf("B: %d (in %v)\n", solveB(steps), time.Since(tB))
}

var cancellations = map[string]map[string]string{
	"n":  {"se": "ne", "s": "-", "sw": "nw"},
	"ne": {"s": "se", "sw": "-", "nw": "n"},
	"se": {"n": "ne", "sw": "s", "nw": "-"},
	"s":  {"n": "-", "ne": "se", "nw": "sw"},
	"sw": {"n": "nw", "ne": "-", "se": "s"},
	"nw": {"ne": "n", "se": "-", "s": "sw"},
}

func solveA(dirs map[string]int) int {
	for dir, cancels := range cancellations {
		for opp := range cancels {
			cancel(dirs, dir, opp)
		}
	}

	return sum(dirs)
}
func solveB(steps []string) int {
	max := 0
	dirs := make(map[string]int)
	for _, s := range steps {
		dirs[s]++
		n := solveA(dirs)
		if n > max {
			max = n
		}
	}
	return max
}

func cancel(dirs map[string]int, a, b string) {
	minDir := a
	maxDir := b
	if dirs[a] > dirs[b] {
		minDir = b
		maxDir = a
	}
	cancelDir := cancellations[maxDir][minDir]
	diff := dirs[minDir]
	dirs[a] -= diff
	dirs[b] -= diff
	if cancelDir != "-" {
		dirs[cancelDir] += diff
	}
}

func sum(dirs map[string]int) int {
	s := 0
	for _, v := range dirs {
		s += v
	}
	return s
}
