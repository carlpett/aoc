package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"strings"
	"time"

	"flag"
	"log"
	"runtime"
	"runtime/pprof"
)

type layer struct {
	d int
	r int
	p int
}

var profile = flag.Bool("profile", false, "")

func main() {
	flag.Parse()
	if *profile {
		fmt.Println("Profiling")
		f, err := os.Create("dec-13.cpu")
		if err != nil {
			log.Fatal("could not create CPU profile: ", err)
		}
		if err := pprof.StartCPUProfile(f); err != nil {
			log.Fatal("could not start CPU profile: ", err)
		}
	}

	tS := time.Now()
	b, err := ioutil.ReadAll(os.Stdin)
	if err != nil {
		panic(err)
	}
	lines := strings.Split(strings.TrimSpace(string(b)), "\n")
	layers := make([]layer, len(lines))
	for idx, line := range lines {
		l := layer{}
		fmt.Sscanf(line, "%d: %d", &l.d, &l.r)
		l.p = period(l.r)
		layers[idx] = l
	}
	fmt.Printf("Setup in %v\n", time.Since(tS))

	tA := time.Now()
	fmt.Printf("A: %d (in %v)\n", solveA(layers), time.Since(tA))
	tB := time.Now()
	fmt.Printf("B: %d (in %v)\n", solveB(layers), time.Since(tB))

	if *profile {
		pprof.StopCPUProfile()

		f, err := os.Create("dec-13.mem")
		if err != nil {
			log.Fatal("could not create memory profile: ", err)
		}
		runtime.GC() // get up-to-date statistics
		if err := pprof.WriteHeapProfile(f); err != nil {
			log.Fatal("could not write memory profile: ", err)
		}
		f.Close()
	}

}

func period(r int) int {
	return 2 * (r - 1)
}

var positions = make(map[int][]int)

func genPos(r int) {
	if _, ok := positions[r]; ok {
		return
	}

	ps := make([]int, 2*r-1)
	for i := 0; i < 2*r-1; i++ {
		if i >= r {
			ps[i] = r - i%r - 2
		} else {
			ps[i] = i
		}
	}
	positions[r] = ps
}

func solveA(layers []layer) int {
	severity := 0
	for _, l := range layers {
		if l.d%l.p == 0 {
			severity += l.r * l.d
		}
	}
	return severity
}

func solveB(layers []layer) int {
	t := 0
timeloop:
	for {
		t++
		for _, l := range layers {
			if (t+l.d)%l.p == 0 {
				continue timeloop
			}
		}
		return t
	}
}
