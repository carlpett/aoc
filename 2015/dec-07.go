package main

import (
	"fmt"
	"strconv"
	"strings"
	"time"

	"github.com/carlpett/aoc/utils"
)

type nodeFunc func(inVals []uint16) uint16
type node struct {
	inputs []string
	ready  bool
	output uint16
	fn     nodeFunc
}

func andNode(inVals []uint16) uint16 {
	return inVals[0] & inVals[1]
}
func orNode(inVals []uint16) uint16 {
	return inVals[0] | inVals[1]
}
func notNode(inVals []uint16) uint16 {
	return ^inVals[0]
}
func constInputNode(n uint16) nodeFunc {
	return func(_ []uint16) uint16 {
		return n
	}
}
func passThroughNode(inVals []uint16) uint16 {
	return inVals[0]
}
func lshiftNode(n uint16) nodeFunc {
	return func(inVals []uint16) uint16 {
		return inVals[0] << n
	}
}
func rshiftNode(n uint16) nodeFunc {
	return func(inVals []uint16) uint16 {
		return inVals[0] >> n
	}
}

func allInputsReady(inputs []string, nodes map[string]*node) bool {
	for _, i := range inputs {
		defer func() {
			if r := recover(); r != nil {
				fmt.Println("Died at ", i)
				panic("aaah")
			}
		}()
		if !nodes[i].ready {
			fmt.Printf("%#v\n", nodes[i])
			fmt.Printf("%s is not ready\n", i)
			return false
		}
	}
	return true
}
func readInputs(inputs []string, nodes map[string]*node) []uint16 {
	out := make([]uint16, len(inputs))
	for idx, i := range inputs {
		out[idx] = nodes[i].output
	}
	return out
}

func mustAtoi(s string) uint16 {
	i, err := strconv.Atoi(s)
	if err != nil {
		panic(err)
	}
	return uint16(i)
}

// 123 -> x
// 456 -> y
// x AND y -> d
// x OR y -> e
// x LSHIFT 2 -> f
// y RSHIFT 2 -> g
// NOT x -> h
// NOT y -> i
func main() {
	lines := utils.MustReadStdinAsStringSlice()
	nodes := make(map[string]*node)
	depends := make(map[string][]string)
	for _, line := range lines {
		parts := strings.Split(line, " -> ")
		desc := strings.Split(parts[0], " ")
		key := parts[1]
		switch len(desc) {
		case 1:
			n, err := strconv.Atoi(desc[0])
			if err == nil {
				nodes[key] = &node{fn: constInputNode(uint16(n))}
			} else {
				nodes[key] = &node{fn: passThroughNode}
				depends[desc[0]] = append(depends[desc[0]], key)
			}
		case 2:
			nodes[key] = &node{inputs: []string{desc[1]}, fn: notNode}
			depends[desc[1]] = append(depends[desc[1]], key)
		case 3:
			switch desc[1] {
			case "AND":
				nodes[key] = &node{inputs: []string{desc[0], desc[2]}, fn: andNode}
				depends[desc[0]] = append(depends[desc[0]], key)
				depends[desc[2]] = append(depends[desc[2]], key)
			case "OR":
				nodes[key] = &node{inputs: []string{desc[0], desc[2]}, fn: orNode}
				depends[desc[0]] = append(depends[desc[0]], key)
				depends[desc[2]] = append(depends[desc[2]], key)
			case "LSHIFT":
				nodes[key] = &node{inputs: []string{desc[0]}, fn: lshiftNode(mustAtoi(desc[2]))}
				depends[desc[0]] = append(depends[desc[0]], key)
			case "RSHIFT":
				nodes[key] = &node{inputs: []string{desc[0]}, fn: rshiftNode(mustAtoi(desc[2]))}
				depends[desc[0]] = append(depends[desc[0]], key)
			}
		}
	}

	for _, n := range nodes {
		if allInputsReady(n.inputs, nodes) {
			n.output = n.fn(readInputs(n.inputs, nodes))
			n.ready = true
		}
	}
	for key, n := range nodes {
		fmt.Printf("%s: %d\n", key, n.output)
	}

	tA := time.Now()
	fmt.Printf("A: %d (in %v)\n", 1, time.Since(tA))
	//tB := time.Now()
	//fmt.Printf("B: %d (in %v)\n", solveB(), time.Since(tB))
}
