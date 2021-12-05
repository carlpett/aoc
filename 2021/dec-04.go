package main

import (
	"fmt"
	"time"
	"strings"

	"github.com/carlpett/aoc/utils"
)

type board struct{
	fields [][]int
}
func (b *board) mark(n int) {
	for r := range b.fields {
		for c := range b.fields[r] {
			if b.fields[r][c] == n {
				b.fields[r][c] = 0
			}
		}
	}
}
func (b *board) bingo() bool {
	for r := range b.fields {
		rsum := 0
		for c := range b.fields[r] {
			rsum += b.fields[r][c]
		}
		if rsum == 0 {
			return true
		}
	}
	for cidx := 0; cidx < 5; cidx++ {
		csum := 0
		for ridx := 0; ridx < 5; ridx++  {
			csum += b.fields[ridx][cidx]
		}
		if csum == 0 {
			return true
		}
	}
	return false
}
func (b *board) sum() int {
	sum := 0
	for r := range b.fields {
		for c := range b.fields[r] {
			sum += b.fields[r][c]
		}
	}
	return sum
}

func boardFromSlice(ss []string) *board {
	fields := make([][]int, 5)
	for idx := range fields {
		fields[idx] = utils.MustAtoiSlice(strings.Fields(ss[idx]))
	}
	return &board{
		fields: fields,
	}
}

func main() {
	tS := time.Now()
	input := utils.MustReadStdinAsStringSlice()
	numbers := utils.MustAtoiSlice(strings.Split(input[0], ","))

	boards := make([]*board, 0)
	for idx := 2; idx < len(input); idx += 6 {
		boards = append(boards, boardFromSlice(input[idx:idx+5]))
	}
	fmt.Printf("Setup done (in %v)\n", time.Since(tS))

	tA := time.Now()
	fmt.Printf("A: %d (in %v)\n", solveA(numbers, boards), time.Since(tA))
	tB := time.Now()
	fmt.Printf("B: %d (in %v)\n", solveB(numbers, boards), time.Since(tB))
}

func solveA(numbers []int, boards []*board) int {
	for _, n := range numbers {
		for _, b := range boards {
			b.mark(n)
			if b.bingo() {
				return b.sum()*n
			}
		}
	}
	return 0
}

func solveB(numbers []int, boards []*board) int {
	fmt.Println(boards[0])
	bingoTracker := make(map[int]bool)
	bingos := make([]int, 0, len(boards))
	for _, n := range numbers {
		for bidx, b := range boards {
			b.mark(n)
			if _, bingoed := bingoTracker[bidx]; !bingoed && b.bingo() {
				bingoTracker[bidx] = true
				bingos = append(bingos, b.sum()*n)
			}
		}
	}
	return bingos[len(boards)-1]
}
