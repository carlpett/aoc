package main

import (
	"bufio"
	"fmt"
	"os"
	"time"
)

func main() {
	tS := time.Now()
	s := bufio.NewScanner(os.Stdin)
	s.Split(bufio.ScanRunes)
	s.Scan() // Get rid of leading {
	groups := readGroup(s, 1)
	fmt.Printf("Setup in %v\n", time.Since(tS))

	tA := time.Now()
	fmt.Printf("A: %d (in %v)\n", groups.sumScore(), time.Since(tA))
	tB := time.Now()
	fmt.Printf("B: %d (in %v)\n", groups.sumGarbage(), time.Since(tB))
}

type group struct {
	data     string
	children []*group
	garbage  int
	score    int
}

func (g *group) sumScore() int {
	n := g.score
	for _, c := range g.children {
		n += c.sumScore()
	}
	return n
}
func (g *group) sumGarbage() int {
	n := g.garbage
	for _, c := range g.children {
		n += c.sumGarbage()
	}
	return n
}

func readGroup(s *bufio.Scanner, n int) *group {
	g := group{}
	g.score = n
scan:
	for s.Scan() {
		if err := s.Err(); err != nil {
			panic(err)
		}
		c := s.Text()
		switch c {
		case "{":
			g.children = append(g.children, readGroup(s, n+1))
		case "}":
			break scan
		case ",":
		case "<":
			g.garbage += readGarbage(s)
		default:
			g.data += c
		}
	}
	return &g
}
func readGarbage(s *bufio.Scanner) int {
	g := 0
scan:
	for s.Scan() {
		if err := s.Err(); err != nil {
			panic(err)
		}
		c := s.Text()
		switch c {
		case "!":
			s.Scan()
		case ">":
			break scan
		default:
			g++
		}
	}
	return g
}
