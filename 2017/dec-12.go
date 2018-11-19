package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"strconv"
	"strings"
	"time"
)

type program struct {
	id          int
	connections map[int]bool
}

func main() {
	tS := time.Now()
	b, err := ioutil.ReadAll(os.Stdin)
	if err != nil {
		panic(err)
	}
	lines := strings.Split(strings.TrimSpace(string(b)), "\n")
	ps := make(map[int]*program)
	for _, l := range lines {
		p := program{connections: make(map[int]bool)}
		parts := strings.Split(l, " <-> ")
		fmt.Sscanf(parts[0], "%d", &p.id)
		for _, cS := range strings.Split(parts[1], ", ") {
			c, _ := strconv.Atoi(cS)
			p.connections[c] = true
		}
		ps[p.id] = &p
	}
	fmt.Printf("Setup in %v\n", time.Since(tS))

	tA := time.Now()
	fmt.Printf("A: %d (in %v)\n", len(solveA(ps, 0)), time.Since(tA))
	tB := time.Now()
	fmt.Printf("B: %d (in %v)\n", solveB(ps), time.Since(tB))
}

func solveA(ps map[int]*program, root int) map[int]bool {
	seen := make(map[int]bool)
	backlog := []int{root}

	for len(backlog) > 0 {
		pId := backlog[0]
		backlog = backlog[1:]
		seen[pId] = true
		p := ps[pId]
		for c := range p.connections {
			if _, ok := seen[c]; !ok {
				backlog = append(backlog, c)
			}
		}
	}

	return seen
}
func solveB(ps map[int]*program) int {
	n := 0
	for len(ps) > 0 {
		randKey := 0
		for k := range ps {
			randKey = k
			break
		}

		inComponent := solveA(ps, randKey)
		for p := range inComponent {
			delete(ps, p)
		}
		n++
	}
	return n
}
