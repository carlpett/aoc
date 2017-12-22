package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"strconv"
	"strings"
	"time"
)

const (
	dataSize  = 256
	blockSize = 16
	rounds    = 64
)

func main() {
	tS := time.Now()
	dataA := make([]byte, dataSize)
	dataB := make([]byte, dataSize)
	dataT := make([]int, dataSize)
	for idx := 0; idx < dataSize; idx++ {
		dataA[idx] = byte(idx)
		dataB[idx] = byte(idx)
		dataT[idx] = idx
	}

	lens, err := ioutil.ReadAll(os.Stdin)
	if err != nil {
		panic(err)
	}
	fmt.Printf("Setup in %v\n", time.Since(tS))

	tA := time.Now()
	fmt.Printf("A: %d (in %v)\n", solveA(dataA, lens), time.Since(tA))
	tB := time.Now()
	fmt.Printf("B: %x (in %v)\n", solveB(dataB, lens), time.Since(tB))
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

func xor(bs []byte) byte {
	var sum byte = 0
	for _, b := range bs {
		sum ^= b
	}
	return sum
}

func atob(s string) byte {
	i, err := strconv.Atoi(s)
	if err != nil {
		panic(err)
	}
	return byte(i)
}

func solveA(data []byte, lens []byte) int {
	bs := make([]byte, 0)
	for _, l := range strings.Split(strings.TrimSpace(string(lens)), ",") {
		b := atob(strings.TrimSpace(l))
		bs = append(bs, b)
	}
	hasher := KnotHasher{data: data}

	hasher.hash(bs)
	return int(hasher.data[0]) * int(hasher.data[1])
}

func solveB(data []byte, lens []byte) []byte {
	magic := []byte{17, 31, 73, 47, 23}
	lens = append(lens, magic...)
	hasher := KnotHasher{data: data}

	for r := 0; r < rounds; r++ {
		hasher.hash(lens)
	}
	denseHash := make([]byte, dataSize/blockSize)
	for block := 0; block < dataSize/blockSize; block++ {
		denseHash[block] = xor(hasher.data[block*blockSize : (block+1)*blockSize])
	}

	return denseHash
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
