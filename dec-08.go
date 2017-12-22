package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"strings"
	"time"
)

type compOp func(int, int) bool

type instruction struct {
	reg     string
	op      string
	rhs     int
	condReg string
	condOp  compOp
	condRhs int
}

func main() {
	tS := time.Now()
	ops := map[string]compOp{
		"==": func(a, b int) bool { return a == b },
		"!=": func(a, b int) bool { return a != b },
		">":  func(a, b int) bool { return a > b },
		">=": func(a, b int) bool { return a >= b },
		"<":  func(a, b int) bool { return a < b },
		"<=": func(a, b int) bool { return a <= b },
	}
	b, err := ioutil.ReadAll(os.Stdin)
	if err != nil {
		panic(err)
	}
	lines := strings.Split(string(b), "\n")
	instructions := make([]instruction, len(lines))
	for idx, l := range lines {
		i := instruction{}
		var condOpStr string
		fmt.Sscanf(l, "%s %s %d if %s %s %d", &i.reg, &i.op, &i.rhs, &i.condReg, &condOpStr, &i.condRhs)
		i.condOp = ops[condOpStr]
		instructions[idx] = i
	}
	fmt.Printf("Setup in %v\n", time.Since(tS))

	t := time.Now()
	max, maxAfter, maxReg := solve(instructions)
	fmt.Printf("A: %d\n", max)
	fmt.Printf("B: %d (%s)\n", maxAfter, maxReg)
	fmt.Printf("(in %v)\n", time.Since(t))
}

func solve(instructions []instruction) (int, int, string) {
	regs := make(map[string]int)
	max := 0
	for _, i := range instructions {
		if i.condOp(regs[i.condReg], i.condRhs) {
			switch i.op {
			case "inc":
				regs[i.reg] = regs[i.reg] + i.rhs
				if regs[i.reg] > max {
					max = regs[i.reg]
				}
			case "dec":
				regs[i.reg] = regs[i.reg] - i.rhs
			}
		}
	}

	maxAfter := 0
	maxReg := ""
	for r, v := range regs {
		if v > maxAfter {
			maxAfter = v
			maxReg = r
		}
	}

	return max, maxAfter, maxReg
}
