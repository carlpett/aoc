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
	rows := strings.Split(strings.TrimSpace(string(b)), "\n")
	m := make([]int, len(rows))
	for ri, r := range rows {
		m[ri], _ = strconv.Atoi(r)
	}

	mB := make([]int, len(m))
	copy(mB, m)
	fmt.Printf("Setup in %v\n", time.Since(tS))

	tA := time.Now()
	fmt.Printf("A: %d (in %v)\n", solve(m, ppA), time.Since(tA))
	tB := time.Now()
	fmt.Printf("B: %d (in %v)\n", solve(mB, ppB), time.Since(tB))
}

func ppA(i int) int {
	return i + 1
}
func ppB(i int) int {
	if i >= 3 {
		return i - 1
	}
	return i + 1
}

func solve(m []int, pp func(int) int) int {
	steps := 0
	i := 0
	for {
		steps++
		s := m[i]
		m[i] = pp(m[i])
		i += s
		if i >= len(m) {
			break
		}
	}
	return steps
}
