package main

import (
	"fmt"
	"io"
	"time"
	"unicode"

	"github.com/carlpett/aoc/utils"
)

type aunt struct {
	id         int
	attributes map[string]int
}

func newAunt() aunt {
	return aunt{attributes: make(map[string]int)}
}

func (a *aunt) Scan(s fmt.ScanState, verb rune) error {
	for {
		// Fmt: (%s: %d,?)+
		tok, err := s.Token(true, unicode.IsLetter)
		if err != nil {
			panic(err)
		}
		marker := make([]byte, len(tok))
		copy(marker, tok)

		r, _, err := s.ReadRune()
		if err != nil {
			panic(err)
		}
		if r != ':' {
			panic(fmt.Sprintf("Unexpected non-colon %c", r))
		}

		tok, err = s.Token(true, func(r rune) bool {
			if unicode.IsDigit(r) {
				return true
			}
			if r == '-' {
				return true
			}
			return false
		})
		if err != nil {
			panic(err)
		}
		value := utils.MustAtoi(string(tok))
		a.attributes[string(marker)] = value

		r, _, err = s.ReadRune()
		if err != nil {
			if err == io.EOF {
				return nil
			}
			return err
		}
		if r != ',' {
			panic(fmt.Sprintf("Unexpected non-comma %c", r))
		}
	}
	return nil
}

func main() {
	input := utils.MustReadStdinAsStringSlice()
	aunts := make([]aunt, 0)
	for _, str := range input {
		a := newAunt()
		var id int
		_, err := fmt.Sscanf(str, "Sue %d: %v", &id, &a)
		if err != nil {
			panic(err)
		}
		a.id = id
		aunts = append(aunts, a)
	}

	measured := aunt{
		attributes: map[string]int{
			"children":    3,
			"cats":        7,
			"samoyeds":    2,
			"pomeranians": 3,
			"akitas":      0,
			"vizslas":     0,
			"goldfish":    5,
			"trees":       3,
			"cars":        2,
			"perfumes":    1,
		},
	}

	tA := time.Now()
	fmt.Printf("A: %d (in %v)\n", solveA(aunts, measured), time.Since(tA))
	tB := time.Now()
	fmt.Printf("B: %d (in %v)\n", solveB(aunts, measured), time.Since(tB))
}

func solveA(aunts []aunt, measured aunt) int {
auntiloop:
	for _, aunt := range aunts {
		for k, v := range measured.attributes {
			if candidateVal, hasMeasurement := aunt.attributes[k]; hasMeasurement && candidateVal != v {
				continue auntiloop
			}
		}
		return aunt.id
	}
	return -1
}

func solveB(aunts []aunt, measured aunt) int {
auntiloop:
	for _, aunt := range aunts {
		for k, v := range measured.attributes {
			if candidateVal, hasMeasurement := aunt.attributes[k]; hasMeasurement {
				switch k {
				case "cats", "trees":
					if candidateVal <= v {
						continue auntiloop
					}
				case "pomeranians", "goldfish":
					if candidateVal >= v {
						continue auntiloop
					}
				default:
					if candidateVal != v {
						continue auntiloop
					}
				}
			}
		}
		return aunt.id
	}
	return -1
}
