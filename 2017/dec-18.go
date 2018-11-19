package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"strconv"
	"strings"
	"sync"
	"time"
)

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
	fmt.Printf("B: %d (in %v)\n", solveB(instructions), time.Since(tB))
}

func (p *program) valueOf(s string) int {
	i, err := strconv.Atoi(s)
	if err == nil {
		return i
	}
	return p.registers[s]
}

func solveB(instructions [][]string) int {
	p0, p1 := newProgramPair()
	deadlock := make(chan bool, 2)
	doneWg := sync.WaitGroup{}

	var p1Sends int
	doneWg.Add(2)
	go func() {
		defer doneWg.Done()
		p0.run(instructions, deadlock)
	}()
	go func() {
		defer doneWg.Done()
		p1Sends = p1.run(instructions, deadlock)
	}()
	go func() {
		for {
			select {
			case <-time.After(time.Millisecond):
				if p0.isBlocked && p1.isBlocked {
					deadlock <- true
					deadlock <- true
				}
			}
		}
	}()
	doneWg.Wait()
	return p1Sends
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
			sends++
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
	lastPlayed := 0
	for {
		if p.pc >= len(instructions) {
			break
		}

		i := instructions[p.pc]
		switch i[0] {
		case "snd":
			r := i[1]
			lastPlayed = p.registers[r]
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
			r := i[1]
			if p.registers[r] != 0 {
				return lastPlayed
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
	return lastPlayed
}
