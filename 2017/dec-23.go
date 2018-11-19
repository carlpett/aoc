package main

import (
	"fmt"
	"io"
	"io/ioutil"
	"os"
	"strconv"
	"strings"
	"time"
)

const debug = true

func debugLogStdout(s string) {
	if debug {
		fmt.Println(s)
	}
}
func debugLog(f io.Writer, s string) {
	if debug {
		fmt.Fprintln(f, s)
	}
}

type program struct {
	id        int
	pc        int
	registers map[string]int
	rcvQ      <-chan int
	sndQ      chan<- int
	isBlocked bool
}

func newProgramPair() (program, program) {
	p0ToP1 := make(chan int, 1e5)
	p1ToP0 := make(chan int, 1e5)
	p0 := program{
		id:        0,
		pc:        0,
		registers: map[string]int{"p": 0},
		sndQ:      p0ToP1,
		rcvQ:      p1ToP0,
		isBlocked: false,
	}
	p1 := program{
		id:        1,
		pc:        0,
		registers: map[string]int{"p": 1},
		sndQ:      p1ToP0,
		rcvQ:      p0ToP1,
		isBlocked: false,
	}
	return p0, p1
}

func main() {
	tS := time.Now()
	b, err := ioutil.ReadAll(os.Stdin)
	if err != nil {
		panic(err)
	}
	lines := strings.Split(strings.TrimSpace(string(b)), "\n")
	instructions := make([][]string, len(lines))
	for idx, line := range lines {
		i := strings.Split(line, " ")
		instructions[idx] = i
	}
	fmt.Printf("Setup in %v\n", time.Since(tS))

	tA := time.Now()
	fmt.Printf("A: %d (in %v)\n", solveA(instructions), time.Since(tA))
	tB := time.Now()
	fmt.Printf("B: %d (in %v)\n", solveB(), time.Since(tB))
}

func (p *program) valueOf(s string) int {
	i, err := strconv.Atoi(s)
	if err == nil {
		return i
	}
	return p.registers[s]
}

func (p *program) run(instructions [][]string, deadlock <-chan bool) int {
	sends := 0
	for {
		if p.pc >= len(instructions) {
			break
		}

		i := instructions[p.pc]
		switch i[0] {
		case "snd":
			v := p.valueOf(i[1])
			p.sndQ <- v
		case "set":
			r, v := i[1], p.valueOf(i[2])
			p.registers[r] = v
		case "add":
			r, v := i[1], p.valueOf(i[2])
			p.registers[r] += v
		case "mul":
			r, v := i[1], p.valueOf(i[2])
			p.registers[r] *= v
		case "mod":
			r, v := i[1], p.valueOf(i[2])
			p.registers[r] %= v
		case "rcv":
			p.isBlocked = true
			select {
			case v := <-p.rcvQ:
				p.isBlocked = false
				r := i[1]
				p.registers[r] = v
			case <-deadlock:
				return sends
			}
		case "jgz":
			gz, v := p.valueOf(i[1]), p.valueOf(i[2])
			if gz > 0 {
				p.pc += v
				continue
			}
		default:
			panic(i)
		}
		p.pc++
	}
	p.isBlocked = true
	return sends
}

func solveA(instructions [][]string) int {
	p, _ := newProgramPair()
	muls := 0
	for {
		if p.pc >= len(instructions) {
			break
		}
		i := instructions[p.pc]
		switch i[0] {
		case "set":
			r, v := i[1], p.valueOf(i[2])
			p.registers[r] = v
		case "add":
			r, v := i[1], p.valueOf(i[2])
			p.registers[r] += v
		case "sub":
			r, v := i[1], p.valueOf(i[2])
			p.registers[r] -= v
		case "mul":
			r, v := i[1], p.valueOf(i[2])
			p.registers[r] *= v
			muls++
		case "mod":
			r, v := i[1], p.valueOf(i[2])
			p.registers[r] %= v
		case "jgz":
			gz, v := p.valueOf(i[1]), p.valueOf(i[2])
			if gz > 0 {
				p.pc += v
				continue
			}
		case "jnz":
			nz, v := p.valueOf(i[1]), p.valueOf(i[2])
			if nz != 0 {
				p.pc += v
				continue
			}
		default:
			panic(i)
		}
		p.pc++
	}
	return muls
}

func solveB() int {
	var a, b, c, d, e, f, g, h int
	muls := 0
	a = 1
	b = 81
	c = b

	if a != 0 {
		b = b*100 + 100000
		muls++
		c = b + 17000
	} else {
	}

	for { // subf
		f = 1
		d = 2
		for { // e2
			e = 2
			muls += b - e
			e = b
			g = 0
			if b%d == 0 {
				f = 0
			}

			d++
			g = d - b
			if g != 0 {
				continue // e2
			}
			if f == 0 {
				h++
			}
			g = b - c
			if g == 0 {
				return h
			}
			b = b + 17
			break // break e2
		}
	}
	fmt.Println(a, b, c, d, e, f, g, h)
	return h
}

func solveBdebug() int {
	logFile, err := os.OpenFile("b.log", os.O_TRUNC|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		panic(err)
	}
	defer logFile.Close()

	var a, b, c, d, e, f, g, h int
	muls := 0
	a = 0 // Set to 1
	debugLog(logFile, "0")
	b = 81
	debugLog(logFile, "1")
	c = b

	debugLog(logFile, "2")
	if a != 0 {
		debugLog(logFile, "4")
		debugLog(logFile, "5")
		b = b*100 + 100000
		muls++
		debugLog(logFile, "6")
		debugLog(logFile, "7")
		c = b + 17000
	} else {
		debugLog(logFile, "3")
	}

	for { // subf
		debugLog(logFile, "8")
		f = 1
		debugLog(logFile, "9")
		d = 2
		for { // e2
			debugLog(logFile, "10")
			e = 2
			for { // gd
				debugLog(logFile, "11")
				debugLog(logFile, "12")
				debugLog(logFile, "13")
				g = d*e - b
				muls++
				debugLog(logFile, "14")
				if g == 0 { // !sube
					debugLog(logFile, "15")
					f = 0
				}
				debugLog(logFile, "16")
				e++
				debugLog(logFile, "17")
				debugLog(logFile, "18")
				g = e - b
				debugLog(logFile, "19")
				if g == 0 {
					break
				}
			}
			debugLog(logFile, "20")
			d++
			debugLog(logFile, "21")
			debugLog(logFile, "22")
			g = d - b
			debugLog(logFile, "23")
			if g != 0 {
				continue // e2
			}
			debugLog(logFile, "24")
			if f == 0 {
				debugLog(logFile, "25")
				h++
			}
			debugLog(logFile, "26")
			g = b
			debugLog(logFile, "27")
			g = g - c
			debugLog(logFile, "28")
			if g == 0 {
				debugLog(logFile, "29")
				return h
			}
			debugLog(logFile, "30")
			b = b + 17
			break // break e2
		}
	}
}
