package main

import (
	"fmt"
	"time"

	"github.com/carlpett/aoc/utils"
)

type deer struct {
	name     string
	speed    int
	duration int
	restTime int
}

func main() {
	input := utils.MustReadStdinAsStringSlice()
	deers := make([]deer, 0)
	for _, str := range input {
		d := deer{}
		_, err := fmt.Sscanf(str, "%s can fly %d km/s for %d seconds, but then must rest for %d seconds.\n", &d.name, &d.speed, &d.duration, &d.restTime)
		if err != nil {
			panic(err)
		}
		deers = append(deers, d)
	}

	const seconds = 2503
	tA := time.Now()
	fmt.Printf("A: %d (in %v)\n", solveA(deers, seconds), time.Since(tA))
	tB := time.Now()
	fmt.Printf("B: %d (in %v)\n", solveB(deers, seconds), time.Since(tB))
}

func solveA(deers []deer, checkTime int) int {
	dists := distanceFlown(deers, checkTime)
	longest := 0
	for _, d := range dists {
		if d > longest {
			longest = d
		}
	}
	return longest
}

func solveB(deers []deer, checkTime int) int {
	score := make(map[string]int)
	topScore := 0

	for t := 1; t <= checkTime; t++ {
		dists := distanceFlown(deers, t)
		scorers := make([]string, 0)
		longest := 0
		for deer, dist := range dists {
			if dist > longest {
				longest = dist
				scorers = []string{deer}
			} else if dist == longest {
				scorers = append(scorers, deer)
			}
		}
		for _, deer := range scorers {
			score[deer]++
			if score[deer] > topScore {
				topScore = score[deer]
			}
		}
	}

	return topScore
}

func distanceFlown(deers []deer, time int) map[string]int {
	dist := make(map[string]int)
	for _, d := range deers {
		period := d.duration + d.restTime
		fullCycles := int(time / period)
		timeInCycle := time % period

		if timeInCycle <= d.duration {
			dist[d.name] = (fullCycles*d.duration + timeInCycle) * d.speed
		} else {
			dist[d.name] = (fullCycles + 1) * d.duration * d.speed
		}
	}
	return dist
}
