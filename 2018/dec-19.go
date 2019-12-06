package main

import (
	"fmt"
	"time"

	"github.com/carlpett/aoc/utils"
)

type opcode struct {
	name string
	inA  int
	inB  int
	out  int
}

func main() {
	tS := time.Now()
	input := utils.MustReadStdinAsStringSlice()
	program := make([]opcode, 0, len(input)-1)
	var ipreg int
	fmt.Sscanf(input[0], "#ip %d", &ipreg)
	for _, step := range input[1:] {
		var op opcode
		_, err := fmt.Sscanf(step, "%s %d %d %d", &op.name, &op.inA, &op.inB, &op.out)
		if err != nil {
			panic(err)
		}
		program = append(program, op)
	}

	fmt.Printf("Setup in %v\n", time.Since(tS))

	//tA := time.Now()
	//fmt.Printf("A: %d (in %v)\n", solve(ipreg, make([]int, 6), program), time.Since(tA))
	tB := time.Now()
	fmt.Printf("B: %d (in %v)\n", solve(ipreg, []int{1, 0, 0, 0, 0, 0}, program), time.Since(tB))
}

func solve(ipreg int, registers []int, program []opcode) int {
	for ; ; registers[ipreg]++ {
		fmt.Print(registers[ipreg], registers)
		step := program[registers[ipreg]]
		fn := ops[step.name]
		output := fn(registers, step.inA, step.inB)
		registers[step.out] = output

		fmt.Println(" =>", registers[ipreg], registers)
		if registers[ipreg] < 0 || registers[ipreg]+1 >= len(program) {
			break
		}
	}

	return registers[0]
}

func solveB(ipreg int) int {
	registers := make([]int, 6)
	registers[0] = 1

mainloop:
	for {
		fmt.Println(registers[ipreg])
		switch registers[ipreg] {
		case 0:
			// addi 4 16 4		#  0		r4 = r4+16		<-  0
			registers[4] += 16
		case 1:
			// seti 1 3 5		#  1		r5 = 1
			registers[5] = 1
		case 2:
			// seti 1 1 3		#  2		r3 = 1
			registers[3] = 1
		case 3:
			// mulr 5 3 1		#  3		r1 = r5*r3
			//registers[1] = registers[5] * registers[3]
			for registers[3] > registers[2] {
				if registers[3]*registers[5] == registers[2] {
					registers[3]++
				} else {
					registers[0] = registers[0] + registers[5]
				}
			}
		case 4:
			// eqrr 1 2 1		#  4		r1 = r1==r2
			if registers[1] == registers[2] {
				registers[1] = 1
			} else {
				registers[1] = 0
			}
		case 5:
			// addr 1 4 4		#  5		r4 = r1+4		<-
			registers[4] = registers[1] + 4
		case 6:
			// addi 4 1 4		#  6		r4 = r4+1		<-
			registers[4] += 1
		case 7:
			// addr 5 0 0		#  7		r0 = r5
			registers[0] = registers[5]
		case 8:
			// addi 3 1 3		#  8		r3 = r3+1
			registers[3] += 1
		case 9:
			// gtrr 3 2 1		#  9		r1 = r3>r2
			if registers[3] > registers[2] {
				registers[1] = 1
			} else {
				registers[1] = 0
			}
		case 10:
			// addr 4 1 4		# 10		r4 = r4+r1		<-
			registers[4] += registers[1]
		case 11:
			// seti 2 8 4		# 11		r4 = 2			<-
			registers[4] = 2
		case 12:
			// addi 5 1 5		# 12		r5 = r5+1
			registers[5] += 1
		case 13:
			// gtrr 5 2 1		# 13		r1 = r5>r2
			if registers[5] > registers[2] {
				registers[1] = 1
			} else {
				registers[1] = 0
			}
		case 14:
			// addr 1 4 4		# 14		r4 = r1+r4		<-
			registers[4] += registers[1]
		case 15:
			// seti 1 3 4		# 15		r4 = r1+3		<- 1+3=4
			registers[4] = registers[1] + 3
		case 16:
			// mulr 4 4 4		# 16		r4 = r4*r4		<-
			registers[4] *= registers[4]
		case 17:
			// addi 2 2 2		# 17		r2 = r2+2
			registers[2] += 2
		case 18:
			// mulr 2 2 2		# 18		r2 = r2*r2
			registers[2] *= registers[2]
		case 19:
			// mulr 4 2 2		# 19		r2 = r4*2
			registers[2] = registers[4] * 2
		case 20:
			// muli 2 11 2		# 20		r2 = r2*11
			registers[2] *= 11
		case 21:
			// addi 1 6 1		# 21		r1 = r1+6
			registers[1] += 6
		case 22:
			// mulr 1 4 1		# 22		r1 = r1*4
			registers[1] *= 4
		case 23:
			// addi 1 18 1		# 23		r1 = r1+18
			registers[1] += 18
		case 24:
			// addr 2 1 2		# 24		r2 = r2+1
			registers[2] += 1
		case 25:
			// addr 4 0 4		# 25		r4 = r4+0		<-
			registers[4] += 0
		case 26:
			// seti 0 3 4		# 26		r4 = 0			<-
			registers[4] = 0
		case 27:
			// setr 4 5 1		# 27		r1 = r4
			registers[1] = registers[4]
		case 28:
			// mulr 1 4 1		# 28		r1 = r1*r4
			registers[1] *= registers[4]
		case 29:
			// addr 4 1 1		# 29		r1 = r4+r1
			registers[1] += registers[4]
		case 30:
			// mulr 4 1 1		# 30		r1 = r4*r1
			registers[1] *= registers[4]
		case 31:
			// muli 1 14 1		# 31		r1 = r1*14
			registers[1] *= 14
		case 32:
			// mulr 1 4 1		# 32		r1 = r1*r4
			registers[1] *= registers[4]
		case 33:
			// addr 2 1 2		# 33		r2 = r2+r1
			registers[2] += registers[1]
		case 34:
			// seti 0 1 0		# 34		r0 = 0
			registers[0] = 0
		case 35:
			// seti 0 4 4		# 35		r4 = 0			<-
			registers[4] = 0
		default:
			break mainloop
		}
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
