package main

import (
	"fmt"
	"sort"
	"time"

	"github.com/carlpett/aoc/utils"
)

type pos struct {
	X, Y int
}

func (p pos) Add(v velocity) pos {
	return pos{p.X + v.dX, p.Y + v.dY}
}

type velocity struct {
	dX, dY int
}
type cart struct {
	id       int
	pos      pos
	vel      velocity
	nextTurn int // left=0, straight=1, right=2
}
type carts []*cart

func (cs carts) Len() int           { return len(cs) }
func (cs carts) Swap(i, j int)      { cs[i], cs[j] = cs[j], cs[i] }
func (cs carts) Less(i, j int) bool { return cs[i].pos.Y < cs[j].pos.Y && cs[i].pos.X < cs[j].pos.X }

var headings = map[rune]velocity{
	'<': velocity{-1, 0},
	'^': velocity{0, -1},
	'>': velocity{1, 0},
	'v': velocity{0, 1},
}

func main() {
	tS := time.Now()
	input := utils.MustReadStdinAsStringSlice()
	cartsA := make(carts, 0)
	cartsB := make(carts, 0) // BLEH
	tracks := make([][]rune, len(input))
	cartIdx := 0
	for y, row := range input {
		tracks[y] = make([]rune, len(row))
		for x, ch := range row {
			switch ch {
			case '<', '^', '>', 'v':
				cartsA = append(cartsA, &cart{cartIdx, pos{x, y}, headings[ch], 0})
				cartsB = append(cartsB, &cart{cartIdx, pos{x, y}, headings[ch], 0})
				cartIdx++
				if headings[ch].dX == 0 {
					tracks[y][x] = '|'
				} else {
					tracks[y][x] = '-'
				}
			default:
				tracks[y][x] = ch
			}
		}
	}
	sort.Sort(cartsA)
	sort.Sort(cartsB)

	fmt.Printf("Setup in %v\n", time.Since(tS))

	tA := time.Now()
	fmt.Printf("A: %v (in %v)\n", solveA(cartsA, tracks), time.Since(tA))
	tB := time.Now()
	fmt.Printf("B: %v (in %v)\n", solveB(cartsB, tracks), time.Since(tB))
}

func solveA(carts carts, tracks [][]rune) pos {
	for t := 0; ; t++ {
		//debugMap(carts, tracks)
		for _, cart := range carts {
			switch tracks[cart.pos.Y][cart.pos.X] {
			case '/', '\\':
				cart.vel = turns[cart.vel][tracks[cart.pos.Y][cart.pos.X]]
			case '+':
				switch cart.nextTurn {
				case 0: // Left turn
					cart.vel = turns[cart.vel]['l']
				case 1: // Straight on
				case 2: // Right turn
					cart.vel = turns[cart.vel]['r']
				}
				cart.nextTurn = (cart.nextTurn + 1) % 3
			}

			cart.pos = cart.pos.Add(cart.vel)
			for _, otherCart := range carts {
				if cart.id != otherCart.id && cart.pos == otherCart.pos {
					return cart.pos
				}
			}
		}

		sort.Sort(carts)
	}
}
func solveB(carts carts, tracks [][]rune) pos {
	for t := 0; ; t++ {
		//debugMap(carts, tracks)
		crashed := make(map[int]bool)
		for _, cart := range carts {
			if _, hasCrashed := crashed[cart.id]; hasCrashed {
				continue
			}

			switch tracks[cart.pos.Y][cart.pos.X] {
			case '/', '\\':
				cart.vel = turns[cart.vel][tracks[cart.pos.Y][cart.pos.X]]
			case '+':
				switch cart.nextTurn {
				case 0: // Left turn
					cart.vel = turns[cart.vel]['l']
				case 1: // Straight on
				case 2: // Right turn
					cart.vel = turns[cart.vel]['r']
				}
				cart.nextTurn = (cart.nextTurn + 1) % 3
			}

			cart.pos = cart.pos.Add(cart.vel)
			for _, otherCart := range carts {
				if cart.id != otherCart.id && cart.pos == otherCart.pos {
					crashed[cart.id] = true
					crashed[otherCart.id] = true
				}
			}
		}

		remainingCarts := make([]*cart, 0, len(carts)-len(crashed))
		for _, cart := range carts {
			if _, hasCrashed := crashed[cart.id]; hasCrashed {
				continue
			}
			remainingCarts = append(remainingCarts, cart)
		}
		carts = remainingCarts
		if len(carts) == 1 {
			return carts[0].pos
		}

		sort.Sort(carts)
	}
}

func debugMap(carts carts, tracks [][]rune) {
	for y, row := range tracks {
		for x, ch := range row {
			foundCart := false
			for _, cart := range carts {
				if cart.pos.X == x && cart.pos.Y == y {
					if cart.id == 0 {
						fmt.Print("&")
					} else {
						fmt.Print("*")
					}
					foundCart = true
					break
				}
			}
			if !foundCart {
				fmt.Print(string(ch))
			}
		}
		fmt.Println()
	}
}

var turns = map[velocity]map[rune]velocity{
	// >
	velocity{1, 0}: map[rune]velocity{
		'/':  velocity{0, -1}, // Left turn
		'\\': velocity{0, 1},  // Right turn
		'l':  velocity{0, -1},
		'r':  velocity{0, 1},
	},
	// v
	velocity{0, 1}: map[rune]velocity{
		'/':  velocity{-1, 0}, // Right turn
		'\\': velocity{1, 0},  // Left turn
		'r':  velocity{-1, 0},
		'l':  velocity{1, 0},
	},
	// <
	velocity{-1, 0}: map[rune]velocity{
		'/':  velocity{0, 1},  // Left turn
		'\\': velocity{0, -1}, // Right turn
		'l':  velocity{0, 1},
		'r':  velocity{0, -1},
	},
	// ^
	velocity{0, -1}: map[rune]velocity{
		'/':  velocity{1, 0},  // Right turn
		'\\': velocity{-1, 0}, // Left turn
		'r':  velocity{1, 0},
		'l':  velocity{-1, 0},
	},
}
