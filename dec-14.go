package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"time"
)

const (
	dataSize  = 256
	blockSize = 16
	rounds    = 64
	gridSize  = 128
)

func main() {
	tS := time.Now()

	key, err := ioutil.ReadAll(os.Stdin)
	if err != nil {
		panic(err)
	}

	g := makeGrid(key)
	fmt.Printf("Setup in %v\n", time.Since(tS))

	tA := time.Now()
	fmt.Printf("A: %d (in %v)\n", solveA(g), time.Since(tA))
	tB := time.Now()
	fmt.Printf("B: %d (in %v)\n", solveB(g), time.Since(tB))
}

type grid struct {
	data    [gridSize][gridSize]bool
	regions [gridSize][gridSize]int
}

func makeGrid(input []byte) *grid {
	g := grid{}
	for i := 0; i < gridSize; i++ {
		rowKey := append(input, []byte(fmt.Sprintf("-%d", i))...)
		rowHash := hash(rowKey)
		for hi := 0; hi < len(rowHash); hi++ {
			for bi := 7; bi >= 0; bi-- {
				g.data[i][hi*8+bi] = rowHash[hi]&1 == 1
				rowHash[hi] = rowHash[hi] >> 1
			}
		}
	}

	return &g
}

func solveA(g *grid) int {
	n := 0
	for i := 0; i < len(g.data); i++ {
		for j := 0; j < len(g.data[i]); j++ {
			if g.data[i][j] {
				n++
			}
		}
	}
	return n
}

func neighbours(i, j int) []coord {
	ns := make([]coord, 0, 4)
	steps := [][]int{{1, 0}, {0, 1}, {-1, 0}, {0, -1}}
	for _, s := range steps {
		cI := i + s[0]
		cJ := j + s[1]
		if cI >= 0 && cI < gridSize && cJ >= 0 && cJ < gridSize {
			ns = append(ns, coord{cI, cJ})
		}
	}
	return ns
}

type coord struct {
	i int
	j int
}

func solveB(g *grid) int {
	region := 1
	for i := 0; i < gridSize; i++ {
		for j := 0; j < gridSize; j++ {
			if g.data[i][j] && g.regions[i][j] == 0 {
				markRegion(coord{i, j}, g, region)
				region++
			}
		}
	}
	return region - 1
}

func markRegion(start coord, g *grid, region int) {
	q := []coord{start}
	var c coord
	for len(q) > 0 {
		c, q = q[0], q[1:]
		g.regions[c.i][c.j] = region
		ns := neighbours(c.i, c.j)
		for _, n := range ns {
			if g.data[n.i][n.j] && g.regions[n.i][n.j] == 0 {
				q = append(q, n)
			}
		}
	}
}

func xor(bs []byte) byte {
	var sum byte = 0
	for _, b := range bs {
		sum ^= b
	}
	return sum
}

type circularSlice struct {
	s   []byte
	pos int
	l   int
}

func newCircularSlice(s []byte, pos, l int) (c circularSlice) {
	return circularSlice{
		s:   s,
		pos: pos,
		l:   l,
	}
}

func (p circularSlice) mapIndex(i int) int { return (p.pos + i) % len(p.s) }
func (p circularSlice) Len() int           { return p.l }
func (p circularSlice) Less(i, j int) bool {
	return p.s[p.mapIndex(i)] < p.s[p.mapIndex(j)]
}
func (p circularSlice) Swap(i, j int) {
	p.s[p.mapIndex(i)], p.s[p.mapIndex(j)] = p.s[p.mapIndex(j)], p.s[p.mapIndex(i)]
}

func reverse(cs circularSlice) {
	for i := cs.Len()/2 - 1; i >= 0; i-- {
		opp := cs.Len() - 1 - i
		cs.Swap(i, opp)
	}
}

type KnotHasher struct {
	pos  int
	skip int
	data []byte
}

func (kh *KnotHasher) hash(bs []byte) {
	for _, b := range bs {
		reverse(newCircularSlice(kh.data, kh.pos, int(b)))
		kh.pos = (kh.pos + int(b) + kh.skip) % len(kh.data)
		kh.skip++
	}
}

func hash(input []byte) []byte {
	data := make([]byte, dataSize)
	for idx := range data {
		data[idx] = byte(idx)
	}
	hasher := KnotHasher{data: data}

	magic := []byte{17, 31, 73, 47, 23}
	input = append(input, magic...)

	for r := 0; r < rounds; r++ {
		hasher.hash(input)
	}
	denseHash := make([]byte, dataSize/blockSize)
	for block := 0; block < dataSize/blockSize; block++ {
		denseHash[block] = xor(hasher.data[block*blockSize : (block+1)*blockSize])
	}

	return denseHash
}
