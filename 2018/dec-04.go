package main

import (
	"fmt"
	"sort"
	"strings"
	"time"

	"github.com/carlpett/aoc/utils"
)

func main() {
	tS := time.Now()
	input := utils.MustReadStdinAsStringSlice()
	sort.Strings(input)
	fmt.Printf("Setup in %v\n", time.Since(tS))

	tA := time.Now()
	fmt.Printf("A: %d (in %v)\n", solveA(input), time.Since(tA))
	tB := time.Now()
	fmt.Printf("B: %d (in %v)\n", solveB(input), time.Since(tB))
}

func hasSuffix(s string) func(x string) bool {
	return func(x string) bool {
		return strings.HasSuffix(x, s)
	}
}

func solveA(input []string) int {
	var currentGuard, lastSeenMinute int
	timesSleptPerMinute := make(map[int][]int)
	timeSlept := make(map[int]int)
	for _, str := range input {
		minute := utils.MustAtoi(str[15:17])
		utils.Match(str).
			When(hasSuffix("begins shift"), func() {
				id := strings.TrimPrefix(strings.Split(str, " ")[3], "#")
				currentGuard = utils.MustAtoi(id)
				if timesSleptPerMinute[currentGuard] == nil {
					timesSleptPerMinute[currentGuard] = make([]int, 60)
				}
			}).
			When(hasSuffix("falls asleep"), func() {}).
			When(hasSuffix("wakes up"), func() {
				timeSlept[currentGuard] += minute - lastSeenMinute
				for i := lastSeenMinute; i < minute; i++ {
					timesSleptPerMinute[currentGuard][i]++
				}
			}).
			OtherwiseThrow()

		lastSeenMinute = minute
	}

	var guardIdx, topSleptMinutes, mostSleptMinute int
	for idx, sleptMinutes := range timesSleptPerMinute {
		sm := utils.SumInts(sleptMinutes)
		if sm > topSleptMinutes {
			topSleptMinutes = sm
			guardIdx = idx
			_, mostSleptMinute = utils.MaxSlice(sleptMinutes)
		}
	}

	return guardIdx * mostSleptMinute
}

func solveB(input []string) int {
	var currentGuard, lastSeenMinute int
	timesSleptPerMinute := make(map[int][]int)
	for _, str := range input {
		minute := utils.MustAtoi(str[15:17])
		utils.Match(str).
			When(hasSuffix("begins shift"), func() {
				id := strings.TrimPrefix(strings.Split(str, " ")[3], "#")
				currentGuard = utils.MustAtoi(id)
				if timesSleptPerMinute[currentGuard] == nil {
					timesSleptPerMinute[currentGuard] = make([]int, 60)
				}
			}).
			When(hasSuffix("falls asleep"), func() {}).
			When(hasSuffix("wakes up"), func() {
				for i := lastSeenMinute; i < minute; i++ {
					timesSleptPerMinute[currentGuard][i]++
				}
			}).
			OtherwiseThrow()

		lastSeenMinute = minute
	}

	var highestFrequency, minuteMostSlept, guardIdx int
	for idx, sleptMinutes := range timesSleptPerMinute {
		times, topMinute := utils.MaxSlice(sleptMinutes)
		if times > highestFrequency {
			highestFrequency = times
			guardIdx = idx
			minuteMostSlept = topMinute
		}
	}

	return guardIdx * minuteMostSlept
}
