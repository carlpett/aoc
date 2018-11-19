package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"strings"
	"time"
)

type component struct {
	a int
	b int
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func main() {
	tS := time.Now()
	b, err := ioutil.ReadAll(os.Stdin)
	if err != nil {
		panic(err)
	}
	lines := strings.Split(string(b), "\n")
	cs := make([]component, len(lines))
	cIdx := make(map[int][]int)
	for idx, l := range lines {
		var a, b int
		_, err := fmt.Sscanf(l, "%d/%d", &a, &b)
		if err != nil {
			panic(err)
		}
		c := component{
			a: min(a, b),
			b: max(a, b),
		}
		cs[idx] = c
		cIdx[c.a] = append(cIdx[c.a], idx)
		if c.a != c.b {
			cIdx[c.b] = append(cIdx[c.b], idx)
		}
	}
	fmt.Printf("Setup in %v\n", time.Since(tS))

	tA := time.Now()
	fmt.Printf("A: %d (in %v)\n", solveA(cs, cIdx), time.Since(tA))
	tB := time.Now()
	fmt.Printf("B: %d (in %v)\n", solveB(cs, cIdx), time.Since(tB))
}

func solveA(cs []component, cIdx map[int][]int) int {
	strength := 0
	for _, starterIdx := range cIdx[0] {
		fmt.Println("Building bridge with", starterIdx)
		s, _ := buildStrongBridge([]int{starterIdx}, cs[starterIdx].b, cs, cIdx)
		fmt.Println("Strength", s)
		if s > strength {
			strength = s
		}
	}
	return strength
}
func solveB(cs []component, cIdx map[int][]int) int {
	strength := 0
	length := 0
	for _, starterIdx := range cIdx[0] {
		fmt.Println("Building bridge with", starterIdx)
		l, s, _ := buildLongBridge([]int{starterIdx}, cs[starterIdx].b, cs, cIdx)
		fmt.Println("Strength", s)
		fmt.Println("Length", l)
		if l > length || (l == length && s > strength) {
			length = l
			strength = s
		}
	}
	return strength
}

func bridgeStrength(bridge []int, cs []component) int {
	s := 0
	for _, i := range bridge {
		s += cs[i].a + cs[i].b
	}
	return s
}

func buildStrongBridge(bridge []int, nextVal int, cs []component, cIdx map[int][]int) (int, []int) {
	candidates := make([]int, 0)
candidateLoop:
	for _, i := range cIdx[nextVal] {
		for _, c := range bridge {
			if i == c {
				continue candidateLoop
			}
		}
		candidates = append(candidates, i)
	}
	if len(candidates) == 0 {
		return bridgeStrength(bridge, cs), bridge
	}

	bestStrength := 0
	var bestBridge []int
	for _, matchIdx := range candidates {
		nv := cs[matchIdx].a
		if nv == nextVal {
			nv = cs[matchIdx].b
		}
		s, b := buildStrongBridge(append(bridge, matchIdx), nv, cs, cIdx)
		if s > bestStrength {
			bestStrength = s
			bestBridge = b
		}
	}

	return bestStrength, bestBridge
}

func buildLongBridge(bridge []int, nextVal int, cs []component, cIdx map[int][]int) (int, int, []int) {
	candidates := make([]int, 0)
candidateLoop:
	for _, i := range cIdx[nextVal] {
		for _, c := range bridge {
			if i == c {
				continue candidateLoop
			}
		}
		candidates = append(candidates, i)
	}
	if len(candidates) == 0 {
		return len(bridge), bridgeStrength(bridge, cs), bridge
	}

	bestStrength := 0
	bestLength := 0
	var bestBridge []int
	for _, matchIdx := range candidates {
		nv := cs[matchIdx].a
		if nv == nextVal {
			nv = cs[matchIdx].b
		}
		l, s, b := buildLongBridge(append(bridge, matchIdx), nv, cs, cIdx)
		if l > bestLength || (l == bestLength && s > bestStrength) {
			bestLength = l
			bestBridge = b
			bestStrength = s
		}
	}

	return bestLength, bridgeStrength(bestBridge, cs), bestBridge
}
