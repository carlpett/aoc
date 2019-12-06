package main

import (
	"fmt"
	"strings"
	"time"

	"github.com/carlpett/aoc/utils"
)

type graph struct {
	edges map[string]map[string]int
}

func newGraph() *graph {
	return &graph{
		edges: make(map[string]map[string]int),
	}
}

func (g *graph) addEdge(p1, p2 string, d int) {
	g.addNode(p1)
	g.addNode(p2)

	g.edges[p1][p2] = d
}
func (g *graph) addNode(p string) {
	if _, ok := g.edges[p]; !ok {
		g.edges[p] = make(map[string]int)
	}
}
func (g *graph) nodes() []string {
	ns := make([]string, 0, len(g.edges))
	for n := range g.edges {
		ns = append(ns, n)
	}
	return ns
}

func main() {
	input := utils.MustReadStdinAsStringSlice()
	guests := newGraph()
	for _, line := range input {
		var p1, p2, effect string
		var hp int

		_, err := fmt.Sscanf(line, "%s would %s %d happiness units by sitting next to %s", &p1, &effect, &hp, &p2)
		if err != nil {
			panic(err)
		}

		if effect == "lose" {
			hp *= -1
		}
		p2 = strings.TrimSuffix(p2, ".") // Sscanf takes the trailing dot into the person name

		guests.addEdge(p1, p2, hp)
	}

	tA := time.Now()
	fmt.Printf("A: %d (in %v)\n", solveA(guests), time.Since(tA))
	tB := time.Now()
	fmt.Printf("B: %d (in %v)\n", solveB(guests), time.Since(tB))
}

func solveA(g *graph) int {
	dist := search(g, []string{}, 0)
	return dist
}
func solveB(g *graph) int {
	g.addNode("me")
	dist := search(g, []string{}, 0)
	return dist
}

func search(g *graph, seating []string, previousMood int) int {
	nodes := g.nodes()
	options := make([]string, 0, len(nodes)-len(seating))
	for _, n := range nodes {
		found := false
		for _, e := range seating {
			if n == e {
				found = true
			}
		}
		if !found {
			options = append(options, n)
		}
	}
	if len(options) == 0 {
		return previousMood + g.edges[seating[0]][seating[len(seating)-1]] + g.edges[seating[len(seating)-1]][seating[0]]
	}

	moods := make([]int, len(options))
	for idx, o := range options {
		var mood int
		if len(seating) == 0 {
			mood = 0
		} else {
			mood = g.edges[seating[len(seating)-1]][o] + g.edges[o][seating[len(seating)-1]]
		}
		moods[idx] = search(g, append(seating, o), previousMood+mood)
	}

	return utils.MaxList(moods)
}
