package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"regexp"
	"strconv"
	"strings"
	"time"
)

type pos struct {
	x, y int
}
type instruction struct {
	op     string
	c1, c2 pos
}

var instructionPattern = regexp.MustCompile(`^(turn on|toggle|turn off) (\d+),(\d+) through (\d+),(\d+)$`)

const (
	turnOn  = "turn on"
	toggle  = "toggle"
	turnOff = "turn off"
)

func newInstruction(s string) instruction {
	parts := instructionPattern.FindStringSubmatch(s)
	if parts == nil {
		panic("Couldn't match pattern: " + s)
	}
	return instruction{
		op: parts[1],
		c1: pos{mustAtoi(parts[2]), mustAtoi(parts[3])},
		c2: pos{mustAtoi(parts[4]), mustAtoi(parts[5])},
	}
}

func mustAtoi(s string) int {
	i, err := strconv.Atoi(s)
	if err != nil {
		panic(err)
	}
	return i
}

func main() {
	bs, err := ioutil.ReadAll(os.Stdin)
	if err != nil {
		panic(err)
	}
	rawinstrs := strings.Split(strings.TrimSpace(string(bs)), "\n")
	instructions := make([]instruction, len(rawinstrs))
	for idx, i := range rawinstrs {
		instructions[idx] = newInstruction(i)
	}

	tA := time.Now()
	fmt.Printf("A: %d (in %v)\n", solveA(instructions), time.Since(tA))
	tB := time.Now()
	fmt.Printf("B: %d (in %v)\n", solveB(instructions), time.Since(tB))
}

func solveA(instructions []instruction) int {
	lights := make([][]bool, 1000)
	for idx := range lights {
		lights[idx] = make([]bool, 1000)
	}

	for _, i := range instructions {
		for x := i.c1.x; x <= i.c2.x; x++ {
			for y := i.c1.y; y <= i.c2.y; y++ {
				switch i.op {
				case turnOn:
					lights[x][y] = true
				case turnOff:
					lights[x][y] = false
				case toggle:
					lights[x][y] = !lights[x][y]
				}
			}
		}
	}

	on := 0
	for xidx := range lights {
		for yidx := range lights[xidx] {
			if lights[xidx][yidx] {
				on++
			}
		}
	}
	return on
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func solveB(instructions []instruction) int {
	lights := make([][]int, 1000)
	for idx := range lights {
		lights[idx] = make([]int, 1000)
	}

	for _, i := range instructions {
		for x := i.c1.x; x <= i.c2.x; x++ {
			for y := i.c1.y; y <= i.c2.y; y++ {
				switch i.op {
				case turnOn:
					lights[x][y]++
				case turnOff:
					lights[x][y] = max(0, lights[x][y]-1)
				case toggle:
					lights[x][y] += 2
				}
			}
		}
	}

	brightness := 0
	for xidx := range lights {
		for yidx := range lights[xidx] {
			brightness += lights[xidx][yidx]
		}
	}
	return brightness
}
