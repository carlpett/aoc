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
	tS := time.Now()
	b, err := ioutil.ReadAll(os.Stdin)
	if err != nil {
		panic(err)
	}
	rows := strings.Split(strings.TrimSpace(string(b)), "\t")
	m := make([]int, len(rows))
	for ri, r := range rows {
		m[ri], _ = strconv.Atoi(r)
	}
	fmt.Printf("Setup in %v\n", time.Since(tS))

	t := time.Now()
	steps, loopLen := solve(m)
	fmt.Printf("A: %d\n", steps)
	fmt.Printf("B: %d\n", loopLen)
	fmt.Printf("(in %v)\n", time.Since(t))
}

func max(m []int) (int, int) {
	maxVal := 0
	maxIdx := 0
	for idx, val := range m {
		if val > maxVal {
			maxIdx = idx
			maxVal = val
		}
	}
	return maxIdx, maxVal
}

func hash(m []int) int {
	h := 17
	for _, v := range m {
		h = h*19 + v
	}
	return h
}

func solve(m []int) (int, int) {
	seen := make(map[int]int)
	loopLen := 0
	steps := 0
	for {
		h := hash(m)
		if v, ok := seen[h]; ok {
			loopLen = steps - v
			break
		}
		seen[h] = steps
		steps++

		idx, blocks := max(m)
		m[idx] = 0
		i := idx
		for {
			i = (i + 1) % len(m)
			m[i]++
			blocks--
			if blocks == 0 {
				break
			}
		}
	}
	return steps, loopLen
}
