package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"strings"
	"time"

	"gopkg.in/yaml.v2"
)

func main() {
	tS := time.Now()

	var startState string
	_, err := fmt.Scanf("Begin in state %1s.\n", &startState)
	if err != nil {
		panic(err)
	}
	var iters int
	_, err = fmt.Scanf("Perform a diagnostic checksum after %d steps.", &iters)
	if err != nil {
		panic(err)
	}

	b, err := ioutil.ReadAll(os.Stdin)
	if err != nil {
		panic(err)
	}
	var input map[string]map[string][]string
	err = yaml.Unmarshal(b, &input)
	if err != nil {
		panic(err)
	}
	fmt.Printf("Setup in %v\n", time.Since(tS))

	tA := time.Now()
	fmt.Printf("A: %d (in %v)\n", solveA(startState, iters, input), time.Since(tA))
	//tB := time.Now()
	//fmt.Printf("A: %d (in %v)\n", solveB(ps), time.Since(tB))
}

func solveA(startState string, iters int, input map[string]map[string][]string) int {
	s := startState
	pos := 0
	tape := make(map[int]bool)

	//seenStates := make(map[string]bool)

	for i := 0; i < iters; i++ {
		/*if _, ok := seenStates[fmt.Sprintf("%s-%d-%v", s, pos, tape)]; ok {
			fmt.Println("We have a loop after", i)
		} else {
			seenStates[fmt.Sprintf("%s-%d-%v", s, pos, tape)] = true
		}*/
		if i%5000 == 0 {
			fmt.Println(i)
		}

		v := 0
		if tape[pos] {
			v = 1
		}
		instructionList, ok := input[fmt.Sprintf("In state %s", s)][fmt.Sprintf("If the current value is %d", v)]
		if !ok {
			panic(s)
		}
		for _, instr := range instructionList {
			switch strings.Split(instr, " ")[0] {
			case "Write":
				var w string
				fmt.Sscanf(instr, "Write the value %1s.", &w)
				switch w {
				case "0":
					tape[pos] = false
				case "1":
					tape[pos] = true
				default:
					panic(w)
				}
			case "Move":
				var dir string
				fmt.Sscanf(instr, "Move one slot to the %s.", &dir)
				switch strings.Trim(dir, ".") {
				case "right":
					pos++
				case "left":
					pos--
				default:
					panic(dir)
				}
			case "Continue":
				fmt.Sscanf(instr, "Continue with state %1s.", &s)
			default:
				panic(instr)
			}
		}
	}

	ones := 0
	for _, v := range tape {
		if v {
			ones++
		}
	}
	return ones
}
