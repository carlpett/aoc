package main

import (
	"fmt"
	"time"

	"github.com/carlpett/aoc/utils"
)

func main() {
	tS := time.Now()
	input := utils.MustReadStdinAsString()
	var players, lastValue int
	_, err := fmt.Sscanf(input, "%d players; last marble is worth %d points", &players, &lastValue)
	if err != nil {
		panic(err)
	}
	fmt.Printf("Setup in %v\n", time.Since(tS))

	tA := time.Now()
	fmt.Printf("A: %d (in %v)\n", solve(players, lastValue), time.Since(tA))
	tB := time.Now()
	fmt.Printf("B: %d (in %v)\n", solve(players, lastValue*100), time.Since(tB))
}

func solve(players, lastValue int) int {
	scores := make([]int, players)
	var currentPlayer int
	n0 := utils.NewCircularLinkedList(0)
	currentMarble := n0

	for n := 1; n <= lastValue; n++ {
		if n%23 == 0 {
			b := currentMarble.Skip(-7)
			scores[currentPlayer] += n + b.Value
			currentMarble = b.Next
			b.Remove()
		} else {
			currentMarble = currentMarble.Next.InsertAfter(n)
		}
		if lastValue < 30 {
			fmt.Printf("[%d]%s\n", currentPlayer+1, n0.StringMarkCurrent(currentMarble))
		}
		currentPlayer = (currentPlayer + 1) % players
	}

	winner, _ := utils.MaxSlice(scores)
	return winner
}
