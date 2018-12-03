package utils

import (
	"io/ioutil"
	"os"
	"strconv"
	"strings"
)

func MustReadStdinAsByteSlice() []byte {
	bs, err := ioutil.ReadAll(os.Stdin)
	if err != nil {
		panic(err)
	}
	return bs
}
func MustReadStdinAsString() string {
	bs := MustReadStdinAsByteSlice()
	return strings.TrimSpace(string(bs))
}
func MustReadStdinAsStringSlice() []string {
	s := MustReadStdinAsString()
	return strings.Split(s, "\n")
}
func MustReadStdinAsIntSlice() []int {
	strs := MustReadStdinAsStringSlice()
	ints := make([]int, len(strs))
	for idx, s := range strs {
		ints[idx] = MustAtoi(s)
	}
	return ints
}

func MustAtoi(s string) int {
	i, err := strconv.Atoi(s)
	if err != nil {
		panic(err)
	}
	return i
}

func Max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
func MaxList(is []int) int {
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
func MinList(is []int) int {
	best := 1 << 32
	for _, d := range is {
		best = Min(best, d)
	}
	return best
}
