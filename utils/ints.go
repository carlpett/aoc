package utils

import (
	"math"
	"sort"
	"strconv"
)

func MustAtoi(s string) int {
	i, err := strconv.Atoi(s)
	if err != nil {
		panic(err)
	}
	return i
}

func Abs(n int) int {
	if n < 0 {
		return -n
	}
	return n
}
func Max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
func MaxList(is ...int) int {
	best := -(1 << 32)
	for _, d := range is {
		best = Max(best, d)
	}
	return best
}
func Min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
func MinList(is ...int) int {
	best := (1 << 32)
	for _, d := range is {
		best = Min(best, d)
	}
	return best
}
func Pow(a, b int) int {
	return int(math.Pow(float64(a), float64(b)))
}
func Log10(a int) int {
	return int(math.Log10(float64(a)))
}

func IntInSlice(s []int, candidate int) bool {
	for _, v := range s {
		if v == candidate {
			return true
		}
	}
	return false
}

func SumInts(i []int) int {
	s := 0
	for _, v := range i {
		s += v
	}
	return s
}
func MinSlice(i []int) (int, int) {
	var min, minIdx int
	min = 1 << 32
	for idx, val := range i {
		if val < min {
			min = val
			minIdx = idx
		}
	}
	return min, minIdx
}
func MaxSlice(i []int) (int, int) {
	var max, maxIdx int
	for idx, val := range i {
		if val > max {
			max = val
			maxIdx = idx
		}
	}
	return max, maxIdx
}
func UniqIntSlice(is []int) []int {
	out := make([]int, 0)
	seen := make(map[int]bool)
	for _, i := range is {
		if _, ok := seen[i]; !ok {
			out = append(out, i)
			seen[i] = true
		}
	}
	sort.Ints(out)
	return out
}
func IntSqrt(n int) int {
	return int(math.Sqrt(float64(n)))
}

func sieveTo(n int) []int {
	primes := make([]int, 0)
	sieve := make(map[int]bool, n)
	for i := 2; i < n; i++ {
		sieve[i] = true
	}
	for i := 2; i < IntSqrt(n); i++ {
		for mult := 2; mult*i < n; mult++ {
			sieve[mult*i] = false
		}
	}
	for i, isPrime := range sieve {
		if isPrime {
			primes = append(primes, i)
		}
	}

	sort.Ints(primes)
	return primes
}
