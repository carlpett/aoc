package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"strings"
	"time"

	"flag"
	"log"
	"runtime"
	"runtime/pprof"
)

type vec2 struct {
	x int
	y int
}

type nodeState string

const (
	clean    nodeState = "."
	weakened           = "W"
	infected           = "#"
	flagged            = "F"
)

type node struct {
	state            nodeState
	initialInfection bool
}

func (n *node) String() string {
	return string(n.state)
}

var profile = flag.Bool("profile", false, "")

func main() {
	flag.Parse()
	if *profile {
		fmt.Println("Profiling")
		f, err := os.Create("dec-22.cpu")
		if err != nil {
			log.Fatal("could not create CPU profile: ", err)
		}
		if err := pprof.StartCPUProfile(f); err != nil {
			log.Fatal("could not start CPU profile: ", err)
		}
	}

	tS := time.Now()
	b, err := ioutil.ReadAll(os.Stdin)
	if err != nil {
		panic(err)
	}
	lines := strings.Split(string(b), "\n")
	nodesA := make(map[vec2]*node, len(lines))
	nodesB := make(map[vec2]*node, len(lines))
	for i, l := range lines {
		r := i - len(lines)/2
		for j, rn := range l {
			c := j - len(l)/2
			if rn == '#' {
				nodesA[vec2{c, r}] = &node{infected, true}
				nodesB[vec2{c, r}] = &node{infected, true}
			} else {
				nodesA[vec2{c, r}] = &node{clean, false}
				nodesB[vec2{c, r}] = &node{clean, false}
			}
		}
	}
	fmt.Printf("Setup in %v\n", time.Since(tS))

	tA := time.Now()
	fmt.Printf("A: %d (in %v)\n", solveA(nodesA, 10000), time.Since(tA))
	tB := time.Now()
	fmt.Printf("B: %d (in %v)\n", solveB(nodesB, 10000000), time.Since(tB))

	if *profile {
		pprof.StopCPUProfile()

		f, err := os.Create("dec-22.mem")
		if err != nil {
			log.Fatal("could not create memory profile: ", err)
		}
		runtime.GC() // get up-to-date statistics
		if err := pprof.WriteHeapProfile(f); err != nil {
			log.Fatal("could not write memory profile: ", err)
		}
		f.Close()
	}
}

func solveA(nodes map[vec2]*node, iter int) int {
	dirs := []vec2{{0, -1}, {1, 0}, {0, 1}, {-1, 0}}
	pos := vec2{}
	dirsIdx := 0
	newInfections := 0
	for i := 0; i < iter; i++ {
		if _, found := nodes[pos]; !found {
			nodes[pos] = &node{clean, false}
		}

		n := nodes[pos]
		if n.state == clean {
			dirsIdx--
			n.state = infected
			newInfections++
		} else {
			dirsIdx++
			n.state = clean
		}

		if dirsIdx >= len(dirs) {
			dirsIdx %= len(dirs)
		} else if dirsIdx < 0 {
			dirsIdx = len(dirs) + dirsIdx
		}

		pos.x, pos.y = pos.x+dirs[dirsIdx].x, pos.y+dirs[dirsIdx].y
	}

	return newInfections
}

func solveB(nodes map[vec2]*node, iter int) int {
	dirs := []vec2{{0, -1}, {1, 0}, {0, 1}, {-1, 0}}
	pos := vec2{}
	dirsIdx := 0
	newInfections := 0
	for i := 0; i < iter; i++ {
		var n *node
		if nf, found := nodes[pos]; !found {
			n = &node{clean, false}
			nodes[pos] = n
		} else {
			n = nf
		}

		switch n.state {
		case clean:
			dirsIdx--
			n.state = weakened
		case weakened:
			n.state = infected
			newInfections++
		case infected:
			dirsIdx++
			n.state = flagged
		case flagged:
			dirsIdx += 2
			n.state = clean
		}

		if dirsIdx >= len(dirs) {
			dirsIdx %= len(dirs)
		} else if dirsIdx < 0 {
			dirsIdx = len(dirs) + dirsIdx
		}

		pos.x, pos.y = pos.x+dirs[dirsIdx].x, pos.y+dirs[dirsIdx].y
	}

	return newInfections
}

func debugMap(nodes map[vec2]*node, pos vec2, size int) {
	for x := -size; x <= size; x++ {
		for y := -size; y <= size; y++ {
			if n, found := nodes[vec2{y, x}]; !found {
				fmt.Print(". ")
			} else {
				v := vec2{y, x}
				if pos == v {
					fmt.Printf("[%s]", n)
				} else {
					fmt.Printf("%s ", n)
				}
			}
		}
		fmt.Println()
	}
	fmt.Println()
}
