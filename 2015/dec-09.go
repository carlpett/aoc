package main

import (
	"fmt"
	"time"

	"github.com/carlpett/aoc/utils"
)

type graph struct {
	dists map[string]map[string]int
}

func newGraph() *graph {
	return &graph{
		dists: make(map[string]map[string]int),
	}
}

func (g *graph) addEdge(p1, p2 string, d int) {
	g.addNode(p1)
	g.addNode(p2)

	g.dists[p1][p2] = d
	g.dists[p2][p1] = d
}
func (g *graph) addNode(p string) {
	if _, ok := g.dists[p]; !ok {
		g.dists[p] = make(map[string]int)
	}
}
func (g *graph) nodes() []string {
	ns := make([]string, 0, len(g.dists))
	for n := range g.dists {
		ns = append(ns, n)
	}
	return ns
}

func main() {
	lines := utils.MustReadStdinAsStringSlice()

	g := newGraph()
	for _, str := range lines {
		var p1, p2 string
		var d int
		_, err := fmt.Sscanf(str, "%s to %s = %d\n", &p1, &p2, &d)
		if err != nil {
			panic(err)
		}
		g.addEdge(p1, p2, d)
	}

	tA := time.Now()
	fmt.Printf("A: %d (in %v)\n", solveA(g), time.Since(tA))
	tB := time.Now()
	fmt.Printf("B: %d (in %v)\n", solveB(g), time.Since(tB))
}

func solveA(g *graph) int {
	dist := search(g, []string{}, 0, utils.MinList)
	return dist
}
func solveB(g *graph) int {
	dist := search(g, []string{}, 0, utils.MaxList)
	return dist
}

type routePicker func([]int) int

func search(g *graph, path []string, dist int, rp routePicker) int {
	nodes := g.nodes()
	options := make([]string, 0, len(nodes)-len(path))
	for _, n := range nodes {
		found := false
		for _, e := range path {
			if n == e {
				found = true
			}
		}
		if !found {
			options = append(options, n)
		}
	}
	if len(options) == 0 {
		return dist
	}

	dists := make([]int, len(options))
	for idx, o := range options {
		var distance int
		if len(path) == 0 {
			distance = 0
		} else {
			distance = g.dists[path[len(path)-1]][o]
		}
		dists[idx] = search(g, append(path, o), dist+distance, rp)
	}

	return rp(dists)
}
