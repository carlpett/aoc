package main

import (
	"fmt"
	"time"

	"github.com/carlpett/aoc/utils"
)

type region struct {
	geoIdx  int
	erosion int
	rType   int
}

func genGrid(depth, tX, tY int) [][]region {
	grid := make([][]region, tY*2)
	for y := 0; y < len(grid); y++ {
		grid[y] = make([]region, tX*2)
		for x := 0; x < len(grid[y]); x++ {
			var idx int
			if (x == 0 && y == 0) || (x == tX && y == tY) {
				idx = 0
			} else if y == 0 {
				idx = x * 16807
			} else if x == 0 {
				idx = y * 48271
			} else {
				idx = grid[y-1][x].erosion * grid[y][x-1].erosion
			}
			erosion := (idx + depth) % 20183
			grid[y][x] = region{
				geoIdx:  idx,
				erosion: erosion,
				rType:   erosion % 3,
			}
		}
	}

	return grid
}

func printMap(grid [][]region, goal coord) {
	for y := 0; y < len(grid); y++ {
		for x := 0; x < len(grid[y]); x++ {
			if x == 0 && y == 0 {
				fmt.Print("M")
			} else if x == goal.x && y == goal.y {
				fmt.Print("T")
			} else {
				switch grid[y][x].rType {
				case 0:
					fmt.Print(".")
				case 1:
					fmt.Print("=")
				case 2:
					fmt.Print("|")
				}
			}
		}
		fmt.Println()
	}
}

func main() {
	tS := time.Now()
	input := utils.MustReadStdinAsStringSlice()
	var depth, tX, tY int
	fmt.Sscanf(input[0], "depth: %d", &depth)
	fmt.Sscanf(input[1], "target: %d,%d", &tX, &tY)
	grid := genGrid(depth, tX, tY)
	fmt.Printf("Setup in %v\n", time.Since(tS))

	tA := time.Now()
	fmt.Printf("A: %d (in %v)\n", solveA(grid, coord{tX, tY}), time.Since(tA))
	tB := time.Now()
	fmt.Printf("B: %d (in %v)\n", solveB(grid, coord{tX, tY}), time.Since(tB))
}

func solveA(grid [][]region, goal coord) int {
	var sum int
	for y := 0; y <= goal.y; y++ {
		for x := 0; x <= goal.x; x++ {
			sum += grid[y][x].rType
		}
	}
	return sum
}

type tool int

const (
	noTool tool = iota
	torch
	climbingGear
)

var validTools = map[int][]tool{
	0: []tool{torch, climbingGear},
	1: []tool{noTool, climbingGear},
	2: []tool{noTool, torch},
}

type searchState struct {
	passedTime int
	loc        coord
	equipped   tool
}

type checkedState struct {
	visited  bool
	bestTime int
}
type visitLog map[coord]map[tool]checkedState

func (l visitLog) add(state searchState) {
	if _, ok := l[state.loc]; !ok {
		l[state.loc] = make(map[tool]checkedState)
	}
	l[state.loc][state.equipped] = checkedState{true, state.passedTime}
}

func solveB(grid [][]region, goal coord) int {
	state := searchState{0, coord{0, 0}, torch}
	l := make(visitLog)
	return search(grid, goal, state, l)
}

var globalBest = 1 << 12

func search(grid [][]region, goal coord, state searchState, checked visitLog) int {
	if state.loc == goal && state.equipped == torch {
		if state.passedTime < globalBest {
			globalBest = state.passedTime
		}
		return state.passedTime
	}
	if checked[state.loc][state.equipped].visited && checked[state.loc][state.equipped].bestTime <= state.passedTime {
		return 1 << 32
	}
	if state.passedTime >= globalBest {
		return 1 << 32
	}
	checked.add(state)

	newStates := getCandidates(grid, state, checked)

	fastest := 1 << 32
	for _, s := range newStates {
		t := search(grid, goal, s, checked)
		if t < fastest {
			fastest = t
		}
	}

	return fastest
}

type coord struct {
	x, y int
}

func getCandidates(grid [][]region, state searchState, checked visitLog) []searchState {
	newStates := make([]searchState, 0, 5)
	// Moves to adjacent regions
	adj := gridFourNeighbours(state.loc, len(grid[state.loc.y]), len(grid))
	for _, nReg := range adj {
		vts := validTools[grid[nReg.y][nReg.x].rType]
		if (vts[0] == state.equipped || vts[1] == state.equipped) && // Can go to that region with current equipement
			(!checked[nReg][state.equipped].visited || checked[nReg][state.equipped].bestTime > state.passedTime+1) { // Haven't visited yet, or with later time
			newStates = append(newStates, searchState{
				passedTime: state.passedTime + 1,
				loc:        nReg,
				equipped:   state.equipped,
			})
		}
	}

	// Stay and swap gear
	var nEq tool
	vts := validTools[grid[state.loc.y][state.loc.x].rType]
	if vts[0] != state.equipped {
		nEq = vts[0]
	} else {
		nEq = vts[1]
	}
	if !checked[state.loc][nEq].visited || checked[state.loc][nEq].bestTime > state.passedTime+7 {
		newStates = append(newStates, searchState{
			passedTime: state.passedTime + 7,
			loc:        state.loc,
			equipped:   nEq,
		})
	}

	return newStates
}

func gridFourNeighbours(center coord, sizeX, sizeY int) []coord {
	pos := []coord{{center.x - 1, center.y}, {center.x, center.y - 1}, {center.x + 1, center.y}, {center.x, center.y + 1}}
	for idx := 0; idx < len(pos); idx++ {
		if pos[idx].x < 0 || pos[idx].x >= sizeX || pos[idx].y < 0 || pos[idx].y >= sizeY {
			pos = append(pos[:idx], pos[idx+1:]...)
			idx--
		}
	}
	return pos
}
