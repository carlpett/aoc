package utils

import (
	"io/ioutil"
	"os"
	"strings"
)

func MustReadStdinAsByteSlice() []byte {
	bs, err := ioutil.ReadAll(os.Stdin)
	if err != nil {
		panic(err)
	}
	return bs
}
func MustReadStdinAsInt() int {
	return MustAtoi(MustReadStdinAsString())
}
func MustReadStdinAsString() string {
	bs := MustReadStdinAsByteSlice()
	return strings.TrimSpace(string(bs))
}
func MustReadStdinAsStringSlice() []string {
	s := MustReadStdinAsString()
	return strings.Split(s, "\n")
}
func MustReadStdinAsSSStringSlice() []string {
	s := MustReadStdinAsString()
	return strings.Split(s, " ")
}
func MustReadStdinAsIntSlice() []int {
	strs := MustReadStdinAsStringSlice()
	ints := make([]int, len(strs))
	for idx, s := range strs {
		ints[idx] = MustAtoi(s)
	}
	return ints
}
func MustReadStdinAsSSIntSlice() []int {
	strs := MustReadStdinAsSSStringSlice()
	ints := make([]int, len(strs))
	for idx, s := range strs {
		ints[idx] = MustAtoi(s)
	}
	return ints
}
