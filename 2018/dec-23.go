package main

import (
	"fmt"
	"time"

	"github.com/carlpett/aoc/utils"
)

type pos struct {
	x, y, z int
}
type bot struct {
	pos pos
	rng int
}

func main() {
	tS := time.Now()
	input := utils.MustReadStdinAsStringSlice()
	bots := make([]bot, len(input))
	var maxRange, maxRangeIdx int
	for idx, s := range input {
		b := bot{pos: pos{}}
		fmt.Sscanf(s, "pos=<%d,%d,%d>, r=%d", &b.pos.x, &b.pos.y, &b.pos.z, &b.rng)
		bots[idx] = b
		if b.rng > maxRange {
			maxRange = b.rng
			maxRangeIdx = idx
		}
	}
	fmt.Printf("Setup in %v\n", time.Since(tS))

	tA := time.Now()
	fmt.Printf("A: %d (in %v)\n", solveA(bots, maxRangeIdx), time.Since(tA))
	tB := time.Now()
	fmt.Printf("B: %d (in %v)\n", solveB(bots), time.Since(tB))
}

func solveA(bots []bot, maxRangeIdx int) int {
	var inRange int
	for _, b := range bots {
		if manhattan(b.pos, bots[maxRangeIdx].pos) <= bots[maxRangeIdx].rng {
			inRange++
		}
	}
	return inRange
}

func solveB(bots []bot) int {
	var maxDfo int
	for _, b := range bots {
		dfo := manhattan(b.pos, pos{})
		if dfo > maxDfo {
			maxDfo = dfo
		}
	}
	p := pos{}
	for r := maxDfo; ; r /= 2 {
		var best int
		for _, c := range mhsp(p, r) {
			//fmt.Println(p, c, coveredBy(bots, c, r))
			if n := coveredBy(bots, c, r); n > best {
				best = n
				p = c
			}
		}
		fmt.Println(p, "won", r)
		if r == 0 {
			break
		}
	}
	fmt.Println(inRange(bots, p), manhattan(p, pos{}))
	return -1
}

/*
func solveB_bak(bots []bot) int {
	overlapping := newGraph()
	for idx, b := range bots {
		overlapping.addNode(idx)
		for oIdx, ob := range bots {
			l := manhattan(b.pos, ob.pos)
			if l <= ob.rng+b.rng {
				overlapping.addEdge(idx, oIdx, l)
			}
		}
	}

	var largestPossibleClique int
	var potentialClique []int
	for c := len(bots); c >= 0; c-- {
		found := 0
		potentialClique = make([]int, 0, c)
		for _, n := range overlapping.Nodes() {
			if len(overlapping.edges[n]) >= c {
				found++
				potentialClique = append(potentialClique, n)
			}
		}
		if found >= c {
			largestPossibleClique = found
			break
		}
	}
	fmt.Println(largestPossibleClique)
	var avgPos pos
	for n, bi := range potentialClique {
		avgPos.x = (n*avgPos.x + bots[bi].pos.x) / (n + 1)
		avgPos.y = (n*avgPos.y + bots[bi].pos.y) / (n + 1)
		avgPos.z = (n*avgPos.z + bots[bi].pos.z) / (n + 1)
	}
	fmt.Println(avgPos, inRange(bots, avgPos))
	for _, p := range gridSixNeighbours(avgPos) {
		fmt.Println(p, inRange(bots, p))
	}

	return -1
}
*/

func mhsp(center pos, r int) []pos {
	return []pos{
		{center.x, center.y, center.z},
		{center.x + r, center.y, center.z},
		{center.x - r, center.y, center.z},
		{center.x, center.y + r, center.z},
		{center.x, center.y - r, center.z},
		{center.x, center.y, center.z + r},
		{center.x, center.y, center.z - r},
	}
}

func manhattan(p1, p2 pos) int {
	return utils.Abs(p1.x-p2.x) + utils.Abs(p1.y-p2.y) + utils.Abs(p1.z-p2.z)
}

func inRange(bots []bot, p pos) int {
	var inRange int
	for _, b := range bots {
		if manhattan(b.pos, p) <= b.rng {
			inRange++
		}
	}
	return inRange
}
func coveredBy(bots []bot, p pos, rng int) int {
	var n int
	for _, b := range bots {
		if manhattan(b.pos, p) <= rng {
			n++
		}
	}
	return n
}

type graph struct {
	edges map[int]map[int]int
	nodes map[int]bool
}

func newGraph() *graph {
	return &graph{
		edges: make(map[int]map[int]int),
		nodes: make(map[int]bool),
	}
}

func (g *graph) addEdge(n, m int, l int) {
	g.addNode(n)
	g.addNode(m)

	g.edges[m][n] = l
	g.edges[n][m] = l
}
func (g *graph) removeEdge(n, m int) {
	delete(g.edges[n], m)
}
func (g *graph) addNode(p int) {
	if _, ok := g.edges[p]; !ok {
		g.edges[p] = make(map[int]int)
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
