package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"strconv"
	"strings"
	"time"
)

func main() {
	t := time.Now()
	data := make([]int, 256)
	for idx := range data {
		data[idx] = idx
	}

	b, err := ioutil.ReadAll(os.Stdin)
	if err != nil {
		panic(err)
	}
	input := strings.Split(strings.TrimSpace(string(b)), ",")
	lens := make([]int, len(input))
	for idx, l := range input {
		lens[idx], _ = strconv.Atoi(strings.TrimSpace(l))
	}

	n := hash(data, lens)
	fmt.Printf("A: %v\n", n)
	fmt.Println(time.Since(t))

	d := []int{0, 1, 2, 3, 4}
	hash(d, []int{3, 4, 1, 5})
	fmt.Println(d)
}

type circularSlice struct {
	s   []int
	pos int
	l   int
}

func newCircularSlice(s []int, pos, l int) (c circularSlice) {
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

func hash(data []int, lens []int) int {
	pos := 0
	skip := 0
	for _, l := range lens {
		reverse(newCircularSlice(data, pos, l))
		pos = (pos + l + skip) % len(data)
		skip++
	}

	return data[0] * data[1]
}
