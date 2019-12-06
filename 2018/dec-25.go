package main

import (
	"fmt"
	"time"

	"github.com/carlpett/aoc/utils"
)

type pos struct {
	x, y, z, t int
}

func main() {
	done := utils.ProfileCPU()
	defer done()

	tS := time.Now()
	input := utils.MustReadStdinAsStringSlice()
	points := make([]pos, len(input))
	for idx, s := range input {
		p := pos{}
		fmt.Sscanf(s, "%d,%d,%d,%d", &p.x, &p.y, &p.z, &p.t)
		points[idx] = p
	}
	fmt.Printf("Setup in %v\n", time.Since(tS))

	tA := time.Now()
	fmt.Printf("A: %d (in %v)\n", solveA(points), time.Since(tA))
}

func solveA(points []pos) int {
	g := newGraph()
	for idx := range points {
		for oIdx := range points {
			if idx == oIdx {
				continue
			}
			g.addNode(idx)
			if manhattan(points[idx], points[oIdx]) <= 3 {
				g.addEdge(idx, oIdx)
			}
		}
	}

	visited := make(map[int]bool)
	constellationIdx := 0
	constellations := make(map[int][]int)
	for _, v := range g.Nodes() {
		if visited[v] {
			continue
		}
		visited[v] = true

		constellations[constellationIdx] = []int{v}
		e := g.outEdges(v)
		for eIdx := 0; eIdx < len(e); eIdx++ {
			u := e[eIdx]
			if !visited[u] {
				visited[u] = true
				constellations[constellationIdx] = append(constellations[constellationIdx], u)
				e = append(e, g.outEdges(u)...)
			}
		}
		constellationIdx++
	}

	return len(constellations)
}

func manhattan(p1, p2 pos) int {
	return utils.Abs(p1.x-p2.x) + utils.Abs(p1.y-p2.y) + utils.Abs(p1.z-p2.z) + utils.Abs(p1.t-p2.t)
}

type graph struct {
	edges map[int]map[int]bool
	nodes map[int]bool
}

func newGraph() *graph {
	return &graph{
		edges: make(map[int]map[int]bool),
		nodes: make(map[int]bool),
	}
}

func (g *graph) addEdge(a, b int) {
	g.addNode(a)
	g.addNode(b)

	g.edges[a][b] = true
	g.edges[b][a] = true
}
func (g *graph) removeEdge(n, m int) {
	delete(g.edges[n], m)
}
func (g *graph) addNode(p int) {
	if _, ok := g.edges[p]; !ok {
		g.edges[p] = make(map[int]bool)
	}
	g.nodes[p] = true
}
func (g *graph) Nodes() []int {
	ns := make([]int, 0, len(g.nodes))
	for n := range g.nodes {
		ns = append(ns, n)
	}
	return ns
}
func (g *graph) inEdges(p int) []int {
	es := make([]int, 0)
	for e := range g.edges[p] {
		es = append(es, e)
	}
	return es
}
func (g *graph) outEdges(p int) []int {
	es := make([]int, 0)
	for n, in := range g.edges {
		for e := range in {
			if e == p {
				es = append(es, n)
			}
		}
	}
	return es
}
