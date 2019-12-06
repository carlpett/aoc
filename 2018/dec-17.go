package main

import (
	"fmt"
	"time"

	"github.com/carlpett/aoc/utils"
)

func main() {
	tS := time.Now()
	input := utils.MustReadStdinAsStringSlice()
	scanMap := make(map[int]map[int]rune)
	yMin := 1 << 31
	yMax := 0
	xMin := 1 << 31
	xMax := 0
	for _, str := range input {
		switch str[0] {
		case 'x':
			var x, yFrom, yTo int
			fmt.Sscanf(str, "x=%d, y=%d..%d", &x, &yFrom, &yTo)
			for y := yFrom; y <= yTo; y++ {
				if _, ok := scanMap[y]; !ok {
					scanMap[y] = make(map[int]rune)
				}
				scanMap[y][x] = '#'
			}
			if yTo > yMax {
				yMax = yTo
			}
			if yFrom < yMin {
				yMin = yFrom
			}
			if x > xMax {
				xMax = x
			}
			if x < xMin {
				xMin = x
			}
		case 'y':
			var y, xFrom, xTo int
			fmt.Sscanf(str, "y=%d, x=%d..%d", &y, &xFrom, &xTo)
			if _, ok := scanMap[y]; !ok {
				scanMap[y] = make(map[int]rune)
			}
			for x := xFrom; x <= xTo; x++ {
				scanMap[y][x] = '#'
			}
			if y > yMax {
				yMax = y
			}
			if y < yMin {
				yMin = y
			}
			if xTo > xMax {
				xMax = xTo
			}
			if xFrom < xMin {
				xMin = xFrom
			}
		}
	}

	xSpan := xMax - xMin
	scan := make([][]rune, yMax+1)
	for y := 0; y <= yMax; y++ {
		scan[y] = make([]rune, xSpan+3)
		for x := 0; x <= xSpan+2; x++ {
			if y == 0 && x+xMin-1 == 500 {
				scan[y][x] = '+'
				continue
			}
			if r, ok := scanMap[y][x+xMin-1]; ok {
				scan[y][x] = r
			} else {
				scan[y][x] = '.'
			}
		}
	}

	fmt.Printf("Setup in %v\n", time.Since(tS))

	tA := time.Now()
	a, b := solve(scan, yMin, yMax, xMin)
	fmt.Printf("A: %d (in %v)\n", a, time.Since(tA))
	fmt.Printf("B: %d\n", b)
}

func solve(scan [][]rune, yMin, yMax, x0 int) (int, int) {
	reachedSquares := 0
	retainedSquares := 0
	for !flow(scan, 500-x0+1, 1) {
	}
	//printScan(scan)

	for y, row := range scan {
		if y >= yMin && y <= yMax {
			for _, sq := range row {
				if sq == '|' || sq == '~' {
					reachedSquares++
				}
				if sq == '~' {
					retainedSquares++
				}
			}
		}
	}
	return reachedSquares, retainedSquares
}

func printScan(scan [][]rune) {
	for y, row := range scan {
		fmt.Printf("%.04d ", y)
		for _, sq := range row {
			fmt.Print(string(sq))
		}
		fmt.Println()
	}
	fmt.Println()
}

func flow(scan [][]rune, x, y int) bool {
	for ; y < len(scan) && scan[y][x] != '#' && scan[y][x] != '~'; y++ {
		scan[y][x] = '|'
		//printScan(scan)
	}
	y--
	if y == len(scan)-1 {
		return true
	}

	var done bool
	limitedLeft := false
	leftWall := -1
	for left := x - 1; left >= 0; left-- {
		if scan[y+1][left] == '.' || scan[y+1][left] == '|' {
			done = flow(scan, left, y)
			break
		}
		if scan[y][left] == '#' {
			limitedLeft = true
			leftWall = left
			break
		}
		scan[y][left] = '|'
	}
	limitedRight := false
	rightWall := -1
	for right := x + 1; right < len(scan[y]); right++ {
		if scan[y+1][right] == '.' || scan[y+1][right] == '|' {
			done = flow(scan, right, y)
			break
		}
		if scan[y][right] == '#' {
			limitedRight = true
			rightWall = right
			break
		}
		scan[y][right] = '|'
	}
	if limitedLeft && limitedRight {
		for fillX := leftWall + 1; fillX < rightWall; fillX++ {
			scan[y][fillX] = '~'
		}
	}

	//printScan(scan)

	return done
}
