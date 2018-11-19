package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"strings"
	"time"
)

type program struct {
	name     string
	weight   int
	children []*program
	parent   *program
}
type pendingRelation struct {
	parent    *program
	childName string
}

func main() {
	tS := time.Now()
	b, err := ioutil.ReadAll(os.Stdin)
	if err != nil {
		panic(err)
	}
	lines := strings.Split(strings.TrimSpace(string(b)), "\n")
	ps := make(map[string]*program)
	pending := make([]pendingRelation, 0)
	for _, l := range lines {
		p := program{}
		parts := strings.Split(l, " -> ")
		fmt.Sscanf(parts[0], "%s (%d)", &p.name, &p.weight)
		if len(parts) > 1 {
			for _, c := range strings.Split(parts[1], ", ") {
				if child, ok := ps[c]; ok {
					p.addChild(child)
				} else {
					pending = append(pending, pendingRelation{
						childName: c,
						parent:    &p,
					})
				}
			}
		}
		ps[p.name] = &p
	}
	for _, rel := range pending {
		rel.parent.addChild(ps[rel.childName])
	}
	fmt.Printf("Setup in %v\n", time.Since(tS))

	tA := time.Now()
	root := solveA(ps)
	fmt.Printf("A: %s (in %v)\n", root, time.Since(tA))
	tB := time.Now()
	fmt.Printf("B: %d (in %v)\n", solveB(ps[root]), time.Since(tB))
}

func (p *program) addChild(child *program) {
	p.children = append(p.children, child)
	child.parent = p
}

func solveA(ps map[string]*program) string {
	var randKey string
	for k := range ps {
		randKey = k
		break
	}

	p := ps[randKey]
	for p.parent != nil {
		p = p.parent
	}

	return p.name
}

func solveB(root *program) int {
	u := findUnbalanced(root)
	if u == root || u == nil {
		siblWeight := 0
		for _, s := range u.parent.children {
			if s != u {
				siblWeight = weight(s)
			}
		}
		diff := siblWeight - weight(u)
		return u.weight + diff
	} else {
		return solveB(u)
	}
}

func findUnbalanced(p *program) *program {
	ws := make(map[int][]*program)
	for _, c := range p.children {
		ws[weight(c)] = append(ws[weight(c)], c)
	}

	if len(ws) == 1 {
		return p
	} else {
		for _, cs := range ws {
			if len(cs) == 1 {
				return cs[0]
			}
		}
		return nil
	}
}

func weight(p *program) int {
	w := p.weight
	for _, c := range p.children {
		w += weight(c)
	}
	return w
}
