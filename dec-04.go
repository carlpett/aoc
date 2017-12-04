package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"sort"
	"strings"
)

func main() {
	b, err := ioutil.ReadAll(os.Stdin)
	if err != nil {
		panic(err)
	}
	phrases := strings.Split(strings.TrimSpace(string(b)), "\n")

	fmt.Printf("A: %d\n", solve(phrases, identity))
	fmt.Printf("B: %d\n", solve(phrases, sorted))
}

func identity(s string) string {
	return s
}
func sorted(s string) string {
	chars := strings.Split(s, "")
	sort.Strings(chars)
	return strings.Join(chars, "")
}

func solve(phrases []string, fn func(string) string) int {
	n := 0
	for _, p := range phrases {
		used := make(map[string]bool)
		valid := true
		for _, w := range strings.Split(p, " ") {
			s := fn(w)
			if _, present := used[s]; present {
				valid = false
				break
			}
			used[s] = true
		}
		if valid {
			n++
		}
	}
	return n
}
