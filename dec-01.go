package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"strconv"
	"time"
)

func main() {
	b, err := ioutil.ReadAll(os.Stdin)
	if err != nil {
		panic(err)
	}
	nums := make([]int, len(b))
	for i, n := range b {
		nums[i] = btoi(n)
	}

	tA := time.Now()
	fmt.Printf("A: %d (in %v)\n", solve(nums, 1), time.Since(tA))
	tB := time.Now()
	fmt.Printf("B: %d (in %v)\n", solve(nums, len(nums)/2), time.Since(tB))
}

func btoi(b byte) int {
	i, _ := strconv.Atoi(string(b))
	return i
}

func solve(ns []int, step int) int {
	sum := 0
	for i, n := range ns {
		if n == ns[(i+step)%len(ns)] {
			sum += n
		}
	}
	return sum
}
