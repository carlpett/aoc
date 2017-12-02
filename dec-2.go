package main

import (
	"fmt"
	"io/ioutil"
	"math"
	"os"
	"strconv"
	"strings"
)

func main() {
	b, err := ioutil.ReadAll(os.Stdin)
	if err != nil {
		panic(err)
	}
	rows := strings.Split(strings.TrimSpace(string(b)), "\n")
	m := make([][]int, len(rows))
	for ri, r := range rows {
		cols := strings.Split(r, "\t")
		m[ri] = make([]int, len(cols))
		for ci, c := range cols {
			m[ri][ci], _ = strconv.Atoi(c)
		}
	}

	fmt.Printf("A: %d\n", solve(m, checksumA))
	fmt.Printf("B: %d\n", solve(m, checksumB))
}

func solve(m [][]int, fn func([]int) int) int {
	sum := 0
	for _, r := range m {
		sum += fn(r)
	}
	return sum
}

func checksumA(r []int) int {
	hi := 0
	lo := math.MaxInt64
	for _, e := range r {
		if e > hi {
			hi = e
		}
		if e < lo {
			lo = e
		}
	}
	return hi - lo
}

func checksumB(r []int) int {
	for i, c := range r {
		for _, c2 := range r[i+1:] {
			if c%c2 == 0 || c2%c == 0 {
				hi := max(c, c2)
				lo := min(c, c2)
				return hi / lo
			}
		}
	}
	fmt.Println("Found no divisors :o")
	return 0
}

func max(lhs, rhs int) int {
	if lhs > rhs {
		return lhs
	}
	return rhs
}
func min(lhs, rhs int) int {
	if lhs < rhs {
		return lhs
	}
	return rhs
}
