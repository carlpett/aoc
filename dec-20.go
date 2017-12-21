package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"strings"
	"time"
)

type vec3 struct {
	x int
	y int
	z int
}
type particle struct {
	pos vec3
	vel vec3
	acc vec3
}

func main() {
	tS := time.Now()
	b, err := ioutil.ReadAll(os.Stdin)
	if err != nil {
		panic(err)
	}
	lines := strings.Split(string(b), "\n")
	ps := make(map[int]*particle, len(lines))
	for idx, l := range lines {
		p := particle{}
		_, err := fmt.Sscanf(l, "p=<%d,%d,%d>, v=<%d,%d,%d>, a=<%d,%d,%d>",
			&p.pos.x, &p.pos.y, &p.pos.z,
			&p.vel.x, &p.vel.y, &p.vel.z,
			&p.acc.x, &p.acc.y, &p.acc.z,
		)
		if err != nil {
			panic(err)
		}
		ps[idx] = &p
	}
	fmt.Printf("Setup in %v\n", time.Since(tS))

	tA := time.Now()
	fmt.Printf("A: %d (in %v)\n", solveA(ps), time.Since(tA))
	tB := time.Now()
	fmt.Printf("A: %d (in %v)\n", solveB(ps), time.Since(tB))
}
func distance(p *particle) int {
	return abs(p.pos.x) + abs(p.pos.y) + abs(p.pos.z)
}
func (p *particle) update() {
	p.vel.x += p.acc.x
	p.vel.y += p.acc.y
	p.vel.z += p.acc.z

	p.pos.x += p.vel.x
	p.pos.y += p.vel.y
	p.pos.z += p.vel.z
}

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

type candidate struct {
	idx int
	acc int
}

func solveA(ps map[int]*particle) int {
	candidates := make([]candidate, 0)

	minAcc := 1 << 32
	for i := 0; i < len(ps); i++ {
		acc := abs(ps[i].acc.x) + abs(ps[i].acc.y) + abs(ps[i].acc.z)
		if acc == minAcc {
			candidates = append(candidates, candidate{i, acc})
		}
		if acc < minAcc {
			minAcc = acc
			candidates = []candidate{{i, acc}}
		}
	}

	minPos := 1 << 32
	minPosIdx := -1
	for _, c := range candidates {
		if distance(ps[c.idx]) < minPos {
			minPos = distance(ps[c.idx])
			minPosIdx = c.idx
		}
	}
	return minPosIdx
}

func solveB(ps map[int]*particle) int {
	occupants := make(map[vec3][]int)
	for t := 0; t < 100; t++ {
		for i, p := range ps {
			p.update()
			occupants[p.pos] = append(occupants[p.pos], i)
		}
		for pos, occs := range occupants {
			if len(occs) > 1 {
				for _, i := range occs {
					delete(ps, i)
				}
			}
			delete(occupants, pos)
		}
	}
	return len(ps)
}
