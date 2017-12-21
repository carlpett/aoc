package main

import (
	"fmt"
	"time"
)

func main() {
	tS := time.Now()
	var step int
	_, err := fmt.Scanf("%d", &step)
	if err != nil {
		panic(err)
	}
	fmt.Printf("Setup in %v\n", time.Since(tS))

	tA := time.Now()
	fmt.Printf("A: %d (in %v)\n", solveA(step), time.Since(tA))
	tB := time.Now()
	fmt.Printf("B: %d (in %v)\n", solveB(step), time.Since(tB))
}

func solveA(step int) int {
	const iters = 2017
	b := make([]int, 1, iters)
	b[0] = 0
	p := 0
	for i := 1; i <= iters; i++ {
		p = (p+step)%len(b) + 1
		b = append(b, 0)
		copy(b[p+1:], b[p:])
		b[p] = i
	}
	return b[(p+1)%len(b)]
}
func solveB(step int) int {
	const iters = 50e6
	p := 0
	lastPOne := 0
	for i := 1; i <= iters; i++ {
		p = (p+step)%i + 1
		if p == 1 {
			lastPOne = i
		}
	}
	return lastPOne
}
