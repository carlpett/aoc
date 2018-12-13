package main

import (
	"bufio"
	"bytes"
	"fmt"
	"time"

	"github.com/carlpett/aoc/utils"
)

type node struct {
	metadata []int
	children []*node
}

type intScanner struct {
	sc *bufio.Scanner
}

func (is *intScanner) Next() int {
	if !is.sc.Scan() {
		panic("EOF")
	}
	return utils.MustAtoi(is.sc.Text())
}

func newIntScanner(bs []byte) *intScanner {
	sc := bufio.NewScanner(bytes.NewReader(bs))
	sc.Split(bufio.ScanWords)
	return &intScanner{sc}
}

func newTree(sc *intScanner) *node {
	numChildren := sc.Next()
	numMetadata := sc.Next()
	root := &node{
		metadata: make([]int, numMetadata),
		children: make([]*node, numChildren),
	}
	for c := 0; c < numChildren; c++ {
		root.children[c] = newTree(sc)
	}
	for m := 0; m < numMetadata; m++ {
		root.metadata[m] = sc.Next()
	}
	return root
}

func main() {
	tS := time.Now()
	input := utils.MustReadStdinAsByteSlice()
	t := newTree(newIntScanner(input))
	fmt.Printf("Setup in %v\n", time.Since(tS))

	tA := time.Now()
	fmt.Printf("A: %d (in %v)\n", solveA(t), time.Since(tA))
	tB := time.Now()
	fmt.Printf("B: %d (in %v)\n", solveB(t), time.Since(tB))
}

func solveA(t *node) int {
	sum := utils.SumInts(t.metadata)
	for _, c := range t.children {
		sum += solveA(c)
	}
	return sum
}
func solveB(t *node) int {
	if len(t.children) == 0 {
		return utils.SumInts(t.metadata)
	}
	var sum int
	for _, cIdx := range t.metadata {
		if cIdx == 0 || cIdx-1 >= len(t.children) {
			continue
		}
		sum += solveB(t.children[cIdx-1])
	}
	return sum
}
