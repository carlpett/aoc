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
	fmt.Printf("A: %d (in %v)\n", solveA(bs), time.Since(tA))
	tB := time.Now()
	fmt.Printf("B: %d (in %v)\n", solveB(bs), time.Since(tB))
}

func solveA(bs []byte) int {
	var x, y int
	visited := map[string]bool{"0,0": true}
	for _, b := range bs {
		switch b {
		case '<':
			x--
		case '>':
			x++
		case '^':
			y++
		case 'v':
			y--
		}
		visited[fmt.Sprintf("%d,%d", x, y)] = true
	}
	return len(visited)
}

func solveB(bs []byte) int {
	var sx, sy, rx, ry int
	visited := map[string]bool{"0,0": true}
	for idx, b := range bs {
		switch b {
		case '<':
			if idx%2 == 0 {
				sx--
			} else {
				rx--
			}
		case '>':
			if idx%2 == 0 {
				sx++
			} else {
				rx++
			}
		case '^':
			if idx%2 == 0 {
				sy++
			} else {
				ry++
			}
		case 'v':
			if idx%2 == 0 {
				sy--
			} else {
				ry--
			}
		}
		if idx%2 == 0 {
			visited[fmt.Sprintf("%d,%d", sx, sy)] = true
		} else {
			visited[fmt.Sprintf("%d,%d", rx, ry)] = true
		}

	}
	return len(visited)
}
