package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"time"
)

func main() {
	bs, err := ioutil.ReadAll(os.Stdin)
	if err != nil {
		panic(err)
	}

	tA := time.Now()
	slnA, slnB := solve(bs)
	fmt.Printf("A: %d, B: %d (in %v)\n", slnA, slnB, time.Since(tA))
}

func solve(bs []byte) (int, int) {
	floor := 0
	basement := 0
	for step, b := range bs {
		switch b {
		case '(':
			floor++
		case ')':
			floor--
		}
		if floor < 0 && basement == 0 {
			basement = step + 1
		}
	}
	return floor, basement
}
