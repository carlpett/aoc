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
	tA := time.Now()
	fmt.Printf("A: %d\n", solve(m, ppA))
	fmt.Println(time.Since(tA))
	tB := time.Now()
	fmt.Printf("B: %d\n", solve(mB, ppB))
	fmt.Println(time.Since(tB))
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
