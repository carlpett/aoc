package main

import (
	"fmt"
	"hash/crc32"
	"strings"
	"time"

	"github.com/carlpett/aoc/utils"
)

func main() {
	done := utils.ProfileCPU()
	defer done()

	tS := time.Now()
	input := utils.MustReadStdinAsStringSlice()
	scanMapA := make([][]rune, len(input))
	scanMapB := make([][]rune, len(input))
	for y, row := range input {
		scanMapA[y] = make([]rune, len(row))
		scanMapB[y] = make([]rune, len(row))
		for x, sq := range row {
			scanMapA[y][x] = sq
			scanMapB[y][x] = sq
		}
	}
	fmt.Printf("Setup in %v\n", time.Since(tS))

	tA := time.Now()
	fmt.Printf("A: %d (in %v)\n", solveA(scanMapA), time.Since(tA))
	tB := time.Now()
	fmt.Printf("B: %d (in %v)\n", solveB(scanMapB), time.Since(tB))
}

func solveA(scan [][]rune) int {
	cur := scan
	var trees, lumberyards int
	for t := 0; t < 10; t++ {
		next := make([][]rune, len(cur))
		for y := range cur {
			next[y] = make([]rune, len(cur[y]))
			for x := range cur[y] {
				adjTrees, adjYards := countNeighbours(cur, x, y)
				if cur[y][x] == '.' && adjTrees >= 3 {
					next[y][x] = '|'
				} else if cur[y][x] == '|' && adjYards >= 3 {
					next[y][x] = '#'
				} else if cur[y][x] == '#' {
					if adjTrees >= 1 && adjYards >= 1 {
						next[y][x] = '#'
					} else {
						next[y][x] = '.'
					}
				} else {
					next[y][x] = cur[y][x]
				}
				//fmt.Print(string(next[y][x]))
			}
			//fmt.Println()
		}
		//fmt.Println()
		cur = next
	}

	for y := range cur {
		for x := range cur[y] {
			switch cur[y][x] {
			case '|':
				trees++
			case '#':
				lumberyards++
			}
		}
	}

	return trees * lumberyards
}

func solveB(scan [][]rune) int {
	cur := scan
	prevval := 0
	seenHashes := make(map[uint32]int)
	foundLength := false
	target := 1000000000
	for t := 0; t < target; t++ {
		next := make([][]rune, len(cur))
		sb := strings.Builder{}

		for y := range cur {
			next[y] = make([]rune, len(cur[y]))
			for x := range cur[y] {
				adjTrees, adjYards := countNeighbours(cur, x, y)
				if cur[y][x] == '.' && adjTrees >= 3 {
					next[y][x] = '|'
				} else if cur[y][x] == '|' && adjYards >= 3 {
					next[y][x] = '#'
				} else if cur[y][x] == '#' {
					if adjTrees >= 1 && adjYards >= 1 {
						next[y][x] = '#'
					} else {
						next[y][x] = '.'
					}
				} else {
					next[y][x] = cur[y][x]
				}
				sb.WriteRune(next[y][x])
				//fmt.Print(string(next[y][x]))
			}
			//fmt.Println()
		}
		//fmt.Println()
		cur = next

		var trees, lumberyards int
		for y := range cur {
			for x := range cur[y] {
				switch cur[y][x] {
				case '|':
					trees++
				case '#':
					lumberyards++
				}
			}
		}

		//fmt.Println(trees*lumberyards - prevval)
		prevval = trees * lumberyards

		chk := crc32.ChecksumIEEE([]byte(sb.String()))
		if seenAt, ok := seenHashes[chk]; !foundLength && ok {
			cycleLength := t - seenAt
			target = t + ((1000000000 - t) % cycleLength)
			foundLength = true
		} else {
			seenHashes[chk] = t
		}
	}

	return prevval
}

type coord struct {
	x, y int
}

func countNeighbours(cur [][]rune, x, y int) (int, int) {
	pos := gridEightNeighbours(coord{x, y}, len(cur[0]), len(cur))
	var trees, lumberyards int
	for _, p := range pos {
		switch cur[p.y][p.x] {
		case '|':
			trees++
		case '#':
			lumberyards++
		}
	}

	return trees, lumberyards
}

func gridEightNeighbours(center coord, sizeX, sizeY int) []coord {
	pos := []coord{{center.x - 1, center.y}, {center.x - 1, center.y - 1}, {center.x, center.y - 1}, {center.x + 1, center.y - 1}, {center.x + 1, center.y}, {center.x - 1, center.y + 1}, {center.x, center.y + 1}, {center.x + 1, center.y + 1}}
	for idx := 0; idx < len(pos); idx++ {
		if pos[idx].x < 0 || pos[idx].x >= sizeX || pos[idx].y < 0 || pos[idx].y >= sizeY {
			pos = append(pos[:idx], pos[idx+1:]...)
			idx--
		}
	}
	return pos
}
