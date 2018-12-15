package main

import (
	"fmt"
	"strconv"
	"time"

	"github.com/carlpett/aoc/utils"
)

func main() {
	tS := time.Now()
	input := utils.MustReadStdinAsInt()
	fmt.Printf("Setup in %v\n", time.Since(tS))

	tA := time.Now()
	fmt.Printf("A: %d (in %v)\n", solveA(input), time.Since(tA))
	tB := time.Now()
	fmt.Printf("B: %d (in %v)\n", solveB(input), time.Since(tB))
}

func solveA(in int) int {
	board := []int{3, 7}
	e1 := 0
	e2 := 1
	for n := 1; len(board) < in+10; n++ {
		c1 := board[e1]
		c2 := board[e2]
		next := c1 + c2
		nextAsString := strconv.Itoa(next)
		for _, digit := range nextAsString {
			board = append(board, utils.MustAtoi(string(digit)))
		}

		e1 = (e1 + 1 + c1) % len(board)
		e2 = (e2 + 1 + c2) % len(board)
	}

	return numFromIntSlice(board[len(board)-10:])
}
func solveB(in int) int {
	board := []int{3, 7}
	e1 := 0
	e2 := 1
	digitsInInput := len(strconv.Itoa(in))
	for n := 1; ; n++ {
		c1 := board[e1]
		c2 := board[e2]
		next := c1 + c2
		nextAsString := strconv.Itoa(next)
		for _, digit := range nextAsString {
			board = append(board, utils.MustAtoi(string(digit)))

			if len(board) > digitsInInput {
				if numFromIntSlice(board[len(board)-digitsInInput:]) == in {
					return len(board) - digitsInInput
				}
			}
		}

		e1 = (e1 + 1 + c1) % len(board)
		e2 = (e2 + 1 + c2) % len(board)
	}
	return -1
}

func numFromIntSlice(is []int) int {
	s := 0
	for idx, i := range is {
		s += i * utils.Pow(10, len(is)-idx-1)
	}
	return s
}
