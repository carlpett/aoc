package main

import (
	"fmt"
	"time"

	"github.com/carlpett/aoc/utils"
)

type ingredient struct {
	name       string
	capacity   int
	durability int
	flavor     int
	texture    int
	calories   int
}

func score(ingredients []ingredient, amounts []int) int {
	scoreParts := make(map[string]int)
	for idx := range ingredients {
		scoreParts["capacity"] += ingredients[idx].capacity * amounts[idx]
		scoreParts["durability"] += ingredients[idx].durability * amounts[idx]
		scoreParts["flavor"] += ingredients[idx].flavor * amounts[idx]
		scoreParts["texture"] += ingredients[idx].texture * amounts[idx]
	}
	s := 1
	for _, sp := range scoreParts {
		s *= sp
	}
	return s
}

func main() {
	input := utils.MustReadStdinAsStringSlice()
	ingredients := make([]ingredient, 0)
	for _, str := range input {
		i := ingredient{}
		_, err := fmt.Sscanf(str, "%s capacity %d, durability %d, flavor %d, texture %d, calories %d",
			&i.name, &i.capacity, &i.durability, &i.flavor, &i.texture, &i.calories)
		if err != nil {
			panic(err)
		}
		ingredients = append(ingredients, i)
	}

	tA := time.Now()
	fmt.Printf("A: %d (in %v)\n", solveA(ingredients), time.Since(tA))
	//tB := time.Now()
	//fmt.Printf("B: %d (in %v)\n", solveB(deers, seconds), time.Since(tB))
}

func solveA(ingredients []ingredient) int {
	return score(ingredients, []int{44, 56})
}
