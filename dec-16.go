package main

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"os"
	"strings"
	"time"
)

func main() {
	tS := time.Now()
	b, err := ioutil.ReadAll(os.Stdin)
	if err != nil {
		panic(err)
	}
	moves := strings.Split(strings.TrimSpace(string(b)), ",")
	programs := []byte("abcdefghijklmnop")
	//programs := []byte("abcde")
	fmt.Printf("Setup in %v\n", time.Since(tS))

	tA := time.Now()
	fmt.Printf("A: %s (in %v)\n", solveA(moves, programs), time.Since(tA))
	tB := time.Now()
	fmt.Printf("B: %s (in %v)\n", solveB(moves, programs, 100000), time.Since(tB))
}

func solveA(moves []string, programs []byte) []byte {
	for _, m := range moves {
		switch m[0] {
		case 'p':
			i0 := bytes.IndexByte(programs, m[1])
			i1 := bytes.IndexByte(programs, m[3])
			programs[i0], programs[i1] = programs[i1], programs[i0]
		case 's':
			var n int
			fmt.Sscanf(m, "s%d", &n)
			for i := 0; i < n; i++ {
				programs = append([]byte{programs[len(programs)-1]}, programs[:len(programs)-1]...)
			}
		case 'x':
			var i0, i1 int
			fmt.Sscanf(m, "x%d/%d", &i0, &i1)
			programs[i0], programs[i1] = programs[i1], programs[i0]
		default:
			panic(m)
		}
	}
	return programs
}

func solveB(moves []string, programs []byte, n int) string {
	p := programs
	seen := make(map[string]bool)
	order := make([]string, 0, 100)
	for {
		if _, ok := seen[string(p)]; ok {
			break
		} else {
			seen[string(p)] = true
			order = append(order, string(p))
		}
		p = solveA(moves, p)
	}

	return order[n%len(order)]
}
