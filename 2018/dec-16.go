package main

import (
	"fmt"
	"time"

	"github.com/carlpett/aoc/utils"
)

type opcode struct {
	id  int
	inA int
	inB int
	out int
}

type sample struct {
	op            opcode
	before, after []int
}

func main() {
	tS := time.Now()
	input := utils.MustReadStdinAsStringSlice()
	samples := make([]sample, 0)
	for idx := 0; ; idx++ {
		// Crappy end-of-samples detection
		if len(input[idx]) == 0 {
			break
		}

		before := make([]int, 4)
		_, err := fmt.Sscanf(input[idx], "Before: [%d, %d, %d, %d]", &before[0], &before[1], &before[2], &before[3])
		if err != nil {
			panic(err)
		}
		idx++

		var op opcode
		_, err = fmt.Sscanf(input[idx], "%d %d %d %d", &op.id, &op.inA, &op.inB, &op.out)
		if err != nil {
			panic(err)
		}
		idx++

		after := make([]int, 4)
		_, err = fmt.Sscanf(input[idx], "After: [%d, %d, %d, %d]", &after[0], &after[1], &after[2], &after[3])
		if err != nil {
			panic(err)
		}
		idx++

		samples = append(samples, sample{op, before, after})
	}

	program := make([]opcode, len(input[len(samples)*4+2:]))
	for _, step := range input[len(samples)*4+2:] {
		var op opcode
		_, err := fmt.Sscanf(step, "%d %d %d %d", &op.id, &op.inA, &op.inB, &op.out)
		if err != nil {
			panic(err)
		}
		program = append(program, op)
	}

	fmt.Printf("Setup in %v\n", time.Since(tS))

	tA := time.Now()
	fmt.Printf("A: %d (in %v)\n", solveA(samples), time.Since(tA))
	tB := time.Now()
	fmt.Printf("B: %d (in %v)\n", solveB(samples, program), time.Since(tB))
}

func solveA(samples []sample) int {
	threePlus := 0
	for _, s := range samples {
		candidates := 0
		for _, fn := range ops {
			output := fn(s.before, s.op.inA, s.op.inB)
			if output == s.after[s.op.out] {
				candidates++
			}
		}
		if candidates >= 3 {
			threePlus++
		}
	}

	return threePlus
}
func solveB(samples []sample, program []opcode) int {
	possibleOpcodeMappings := make(map[string]map[int]int, 16)
	for n := range ops {
		possibleOpcodeMappings[n] = make(map[int]int, 16)
	}
	for _, s := range samples {
		for name, fn := range ops {
			output := fn(s.before, s.op.inA, s.op.inB)
			if output == s.after[s.op.out] {
				possibleOpcodeMappings[name][s.op.id]++
			}
		}
	}
	solvedIds := make(map[int]string)
	solvedNames := make(map[string]int)
	for len(possibleOpcodeMappings) > 0 {
		for name, poss := range possibleOpcodeMappings {
			if len(poss) == 1 {
				id := mapKeys(poss)[0]
				solvedNames[name] = id
				solvedIds[id] = name

				for n, otherPoss := range possibleOpcodeMappings {
					delete(otherPoss, id)
					if len(otherPoss) == 0 {
						delete(possibleOpcodeMappings, n)
					}
				}
			}
		}
	}

	registers := make([]int, 4)
	for _, step := range program {
		fn := ops[solvedIds[step.id]]
		output := fn(registers, step.inA, step.inB)
		registers[step.out] = output
	}

	return registers[0]
}

var ops = map[string]func(regs []int, a, b int) int{
	"addr": func(regs []int, a, b int) int {
		return regs[a] + regs[b]
	},
	"addi": func(regs []int, a, b int) int {
		return regs[a] + b
	},

	"mulr": func(regs []int, a, b int) int {
		return regs[a] * regs[b]
	},
	"muli": func(regs []int, a, b int) int {
		return regs[a] * b
	},

	"banr": func(regs []int, a, b int) int {
		return regs[a] & regs[b]
	},
	"bani": func(regs []int, a, b int) int {
		return regs[a] & b
	},

	"borr": func(regs []int, a, b int) int {
		return regs[a] | regs[b]
	},
	"bori": func(regs []int, a, b int) int {
		return regs[a] | b
	},

	"setr": func(regs []int, a, b int) int {
		return regs[a]
	},
	"seti": func(regs []int, a, b int) int {
		return a
	},

	"gtir": func(regs []int, a, b int) int {
		if a > regs[b] {
			return 1
		}
		return 0
	},
	"gtri": func(regs []int, a, b int) int {
		if regs[a] > b {
			return 1
		}
		return 0
	},
	"gtrr": func(regs []int, a, b int) int {
		if regs[a] > regs[b] {
			return 1
		}
		return 0
	},

	"eqir": func(regs []int, a, b int) int {
		if a == regs[b] {
			return 1
		}
		return 0
	},
	"eqri": func(regs []int, a, b int) int {
		if regs[a] == b {
			return 1
		}
		return 0
	},
	"eqrr": func(regs []int, a, b int) int {
		if regs[a] == regs[b] {
			return 1
		}
		return 0
	},
}

func mapKeys(m map[int]int) []int {
	ks := make([]int, 0, len(m))
	for k := range m {
		ks = append(ks, k)
	}
	return ks
}
