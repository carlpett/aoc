package main

import (
	"encoding/json"
	"fmt"
	"time"

	"github.com/carlpett/aoc/utils"
)

func main() {
	input := utils.MustReadStdinAsByteSlice()
	data := make([]interface{}, 0)
	err := json.Unmarshal(input, &data)
	if err != nil {
		panic(err)
	}

	tA := time.Now()
	fmt.Printf("A: %d (in %v)\n", solveA(data), time.Since(tA))
	tB := time.Now()
	fmt.Printf("B: %d (in %v)\n", solveB(data), time.Since(tB))
}

func solveA(data interface{}) int {
	return sumDeep(data, func(v interface{}) bool { return false })
}
func solveB(data interface{}) int {
	filter := func(v interface{}) bool {
		switch val := v.(type) {
		case string:
			if val == "red" {
				return true
			}
		}
		return false
	}
	return sumDeep(data, filter)
}

func sumDeep(data interface{}, filter func(v interface{}) bool) int {
	n := 0
	switch typed := data.(type) {
	case map[string]interface{}:
		for _, v := range typed {
			if filter(v) {
				return 0
			}
			n += sumDeep(v, filter)
		}
	case []interface{}:
		for _, v := range typed {
			n += sumDeep(v, filter)
		}
	case float64:
		n += int(typed)
	}

	return n
}
