package main

import (
	"fmt"
	"sort"
	"strings"
	"time"

	"github.com/carlpett/aoc/utils"
)

type graph struct {
	dists map[string]map[string]bool
	nodes map[string]bool
}

func newGraph() *graph {
	return &graph{
		dists: make(map[string]map[string]bool),
		nodes: make(map[string]bool),
	}
}

func (g *graph) addEdge(dependency, dependent string) {
	g.addNode(dependency)
	g.addNode(dependent)

	g.dists[dependent][dependency] = true
}
func (g *graph) removeEdge(n, m string) {
	delete(g.dists[n], m)
}
func (g *graph) addNode(p string) {
	if _, ok := g.dists[p]; !ok {
		g.dists[p] = make(map[string]bool)
	}
	g.nodes[p] = true
}
func (g *graph) Nodes() []string {
	ns := make([]string, 0, len(g.nodes))
	for n := range g.nodes {
		ns = append(ns, n)
	}
	return ns
}
func (g *graph) inEdges(p string) []string {
	es := make([]string, 0)
	for e := range g.dists[p] {
		es = append(es, e)
	}
	return es
}
func (g *graph) outEdges(p string) []string {
	es := make([]string, 0)
	for n, in := range g.dists {
		for e := range in {
			if e == p {
				es = append(es, n)
			}
		}
	}
	return es
}

func (g *graph) copy() *graph {
	c := newGraph()
	for k, v := range g.dists {
		c.addNode(k)
		for n := range v {
			c.dists[k][n] = true
			c.addNode(n)
		}
	}
	return c
}

func main() {
	tS := time.Now()
	lines := utils.MustReadStdinAsStringSlice()

	g := newGraph()
	for _, str := range lines {
		var dependency, dependent string
		_, err := fmt.Sscanf(str, "Step %s must be finished before step %s can begin.", &dependency, &dependent)
		if err != nil {
			panic(err)
		}
		g.addEdge(dependency, dependent)
	}
	fmt.Printf("Setup in %v\n", time.Since(tS))

	tA := time.Now()
	fmt.Printf("A: %s (in %v)\n", solveA(g), time.Since(tA))
	tB := time.Now()
	fmt.Printf("B: %d (in %v)\n", solveB(g, 5), time.Since(tB))
}

func solveA(g *graph) string {
	ns := g.topoSort()
	return strings.Join(ns, "")
}

type wip struct {
	item      string
	readyAt   int
	workerIdx int
}

var timePerNode = map[string]int{
	"A": 1,
	"B": 2,
	"C": 3,
	"D": 4,
	"E": 5,
	"F": 6,
	"G": 7,
	"H": 8,
	"I": 9,
	"J": 10,
	"K": 11,
	"L": 12,
	"M": 13,
	"N": 14,
	"O": 15,
	"P": 16,
	"Q": 17,
	"R": 18,
	"S": 19,
	"T": 20,
	"U": 21,
	"V": 22,
	"W": 23,
	"X": 24,
	"Y": 25,
	"Z": 26,
}

func solveB(g *graph, numWorkers int) int {
	mg := g.copy()
	finished := make([]string, 0, len(mg.Nodes()))

	workQueue := make([]string, 0)
	inProgress := make([]wip, 0)
	for _, n := range mg.Nodes() {
		if len(g.inEdges(n)) == 0 {
			workQueue = append(workQueue, n)
		}
	}

	workers := make([]bool, numWorkers)
	for i := 0; i < numWorkers; i++ {
		workers[i] = true
	}

	var t int
	for {
		// Check if any in-progress work is done
		for idx, work := range inProgress {
			if work.readyAt == t {
				inProgress = append(inProgress[:idx], inProgress[idx+1:]...)
				finished = append(finished, work.item)
				workers[work.workerIdx] = true

				// Queue any newly available work
				for _, m := range mg.outEdges(work.item) {
					mg.removeEdge(m, work.item)
					if len(mg.inEdges(m)) == 0 {
						workQueue = append(workQueue, m)
					}
				}

				idx--
			}
		}
		// Assign any available work
		for workerIdx, ready := range workers {
			if len(workQueue) > 0 {
				if ready {
					var item string
					item, workQueue = workQueue[0], workQueue[1:]
					finishAt := t + 60 + timePerNode[item]
					inProgress = append(inProgress, wip{item, finishAt, workerIdx})
					workers[workerIdx] = false
				}
			} else {
				break
			}
		}

		// Check if finished
		if len(workQueue) == 0 && len(inProgress) == 0 {
			break
		}

		sort.Strings(workQueue)
		t++
	}

	return t
}

func (g *graph) topoSort() []string {
	mg := g.copy()
	sorted := make([]string, len(mg.Nodes()))

	workSet := make([]string, 0)
	for _, n := range mg.Nodes() {
		if len(g.inEdges(n)) == 0 {
			workSet = append(workSet, n)
		}
	}

	var n string
	for {
		if len(workSet) == 0 {
			break
		}
		n, workSet = workSet[0], workSet[1:]
		sorted = append(sorted, n)
		for _, m := range mg.outEdges(n) {
			mg.removeEdge(m, n)
			if len(mg.inEdges(m)) == 0 {
				workSet = append(workSet, m)
			}
		}
		sort.Strings(workSet)
	}

	return sorted
}
