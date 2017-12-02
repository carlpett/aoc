package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"strconv"
)

func main() {
	if len(os.Args) != 2 {
		fmt.Fprintln(os.Stderr, "Usage: dec-1 [a|b]")
		os.Exit(1)
	}

	b, err := ioutil.ReadAll(os.Stdin)
	if err != nil {
		panic(err)
	}
	nums := make([]int, len(b))
	for i, n := range b {
		nums[i] = btoi(n)
	}

	var offset int
	switch os.Args[1] {
	case "a":
		offset = 1
	case "b":
		offset = len(nums) / 2
	default:
		fmt.Printf("Unknown subproblem %s", os.Args[1])
		os.Exit(1)
	}

	fmt.Printf("Captcha: %d\n", solve(nums, offset))
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
