package main

import (
	"fmt"
	"sort"
	"time"

	"github.com/carlpett/aoc/utils"
)

func main() {
	tS := time.Now()
	input := utils.MustReadStdinAsStringSlice()
	state := createState(input)
	fmt.Printf("Setup in %v\n", time.Since(tS))

	tA := time.Now()
	fmt.Printf("A: %d (in %v)\n", solveA(state.copy()), time.Since(tA))
	tB := time.Now()
	fmt.Printf("B: %d (in %v)\n", solveB(state), time.Since(tB))
}

func debugMap(world [][]*square) {
	for _, row := range world {
		for _, sq := range row {
			fmt.Print(sq)
		}
		fmt.Println()
	}
}

type pos struct {
	x, y int
}
type positions []pos

func (ps positions) Len() int           { return len(ps) }
func (ps positions) Swap(i, j int)      { ps[i], ps[j] = ps[j], ps[i] }
func (ps positions) Less(i, j int) bool { return ps[i].Less(ps[j]) }

func (p1 pos) Less(p2 pos) bool {
	if p1.y < p2.y {
		return true
	}
	if p1.y == p2.y && p1.x < p2.x {
		return true
	}

	return false
}

type unit struct {
	id    int
	elf   bool
	power int
	hp    int
	pos   pos
	dead  bool
}
type unitSlice []*unit

func (us unitSlice) Len() int           { return len(us) }
func (us unitSlice) Swap(i, j int)      { us[i], us[j] = us[j], us[i] }
func (us unitSlice) Less(i, j int) bool { return us[i].pos.Less(us[j].pos) }

type square struct {
	isWall bool
	unit   *unit
}
type worldGrid [][]*square

type state struct {
	world worldGrid
	units unitSlice
}

func (s *state) copy() *state {
	ns := &state{
		world: make(worldGrid, len(s.world)),
		units: make(unitSlice, 0, len(s.units)),
	}
	for y, row := range s.world {
		ns.world[y] = make([]*square, len(row))
		for x, sq := range row {
			var u *unit
			if sq.unit != nil {
				u = &unit{
					id:    sq.unit.id,
					elf:   sq.unit.elf,
					power: sq.unit.power,
					hp:    sq.unit.hp,
					pos:   sq.unit.pos,
					dead:  sq.unit.dead,
				}
				ns.units = append(ns.units, u)
			}
			ns.world[y][x] = &square{
				isWall: sq.isWall,
				unit:   u,
			}
		}
	}
	return ns
}

func (s *square) String() string {
	if s.isWall {
		return "#"
	}
	if s.unit == nil {
		return "."
	}
	return fmt.Sprintf("%d", s.unit.id)
	if s.unit.elf {
		return "E"
	}
	return "G"
}

func createState(input []string) *state {
	state := &state{
		world: make([][]*square, len(input)),
		units: make(unitSlice, 0),
	}
	for y, row := range input {
		state.world[y] = make([]*square, len(row))
		for x, content := range row {
			switch content {
			case '#':
				state.world[y][x] = &square{isWall: true}
			case '.':
				state.world[y][x] = &square{isWall: false}
			case 'E', 'G':
				u := &unit{
					id:    len(state.units),
					elf:   content == 'E',
					power: 3,
					hp:    200,
					dead:  false,
					pos:   pos{x, y},
				}
				state.units = append(state.units, u)
				state.world[y][x] = &square{isWall: false, unit: u}
			}
		}
	}
	return state
}

func solveA(s *state) int {
	units := s.units
	world := s.world
	t := 0
combat:
	for {
		sort.Sort(units)
		//fmt.Println(t)
		//debugMap(world)
		//for _, u := range units {
		//	fmt.Println(u)
		//}
		for _, u := range units {
			//fmt.Println("Active", *u)
			if u.dead {
				continue
			}
			// Identify targets
			targets := make(unitSlice, 0)
			inAttackRange := false
			targetAdjacentSquares := make(map[pos]bool)
			for _, ou := range units {
				if u.id == ou.id || ou.dead || u.elf == ou.elf {
					continue
				}
				targets = append(targets, ou)

				if manhattan(u.pos, ou.pos) == 1 {
					inAttackRange = true
					break
				}
				for _, p := range adjacentCoords(ou.pos) {
					if !world[p.y][p.x].isWall && world[p.y][p.x].unit == nil {
						targetAdjacentSquares[p] = true
					}
				}
			}
			sort.Sort(targets)
			if len(targets) == 0 {
				break combat
			}

			if !inAttackRange && len(targetAdjacentSquares) == 0 {
				continue
			}

			if !inAttackRange {
				// Move
				var bestNext pos
				bestDistance := 1 << 31
				adjList := make(positions, 0, len(targetAdjacentSquares))
				for s := range targetAdjacentSquares {
					adjList = append(adjList, s)
				}
				sort.Sort(adjList)
				for _, p := range adjList {
					l, next := shortestPathLength(u.pos, p, world)
					//fmt.Println("Considering moving from", u.pos, "to", p, "distance is", l, "next move would be", next)
					if l == 1<<31 {
						//fmt.Println("Ignoring, no valid path")
						continue
					}
					if l < bestDistance {
						//fmt.Println("Best so far")
						bestNext = next
						bestDistance = l
					} else {
						//fmt.Println("Not best path found")
					}
				}
				if bestDistance == 1<<31 {
					//fmt.Println(u.id, "could not find a valid move")
					continue
				}
				world[u.pos.y][u.pos.x].unit = nil
				//fmt.Println(u.id, "moving from", u.pos, "to", bestNext)
				u.pos = bestNext
				world[u.pos.y][u.pos.x].unit = u

				if bestDistance == 1 {
					inAttackRange = true
				}
			}

			if inAttackRange {
				// Attack
				var target *unit
				leastHp := 201
				for _, ou := range units {
					if ou.dead || ou.id == u.id || ou.elf == u.elf {
						continue
					}
					if manhattan(u.pos, ou.pos) > 1 {
						continue
					}
					if ou.hp < leastHp {
						target = ou
						leastHp = ou.hp
					} else if ou.hp == leastHp && ou.pos.Less(target.pos) {
						target = ou
					}
				}

				//fmt.Println(u.id, "attacks", target.id)
				target.hp -= u.power
				if target.hp <= 0 {
					//fmt.Println(target.id, "dies")
					target.dead = true
					world[target.pos.y][target.pos.x].unit = nil
				}
			}
		}

		t++
	}

	hpSum := 0
	for _, u := range units {
		if !u.dead {
			hpSum += u.hp
		}
	}

	return hpSum * t
}

func solveB(s *state) int {
	for pow := 3; ; pow++ {
		sCur := s.copy()

		for _, u := range sCur.units {
			if u.elf {
				u.power = pow
			}
		}

		outcome := solveA(sCur)
		elvesDied := false
		for _, u := range sCur.units {
			if u.dead && u.elf {
				elvesDied = true
				break
			}
		}

		if !elvesDied {
			return outcome
		}
	}
}

func manhattan(p1, p2 pos) int {
	return utils.Abs(p1.x-p2.x) + utils.Abs(p1.y-p2.y)
}

func adjacentCoords(p pos) positions {
	ps := make(positions, 0, 4)

	if p.x > 0 {
		ps = append(ps, pos{p.x - 1, p.y})
	}
	if p.y > 0 {
		ps = append(ps, pos{p.x, p.y - 1})
	}
	ps = append(ps, pos{p.x + 1, p.y})
	ps = append(ps, pos{p.x, p.y + 1})

	sort.Sort(ps)

	return ps
}

type node struct {
	dist int
	prev *node
	pos  pos
}

func (n *node) String() string {
	return fmt.Sprintf("  %d  ", n.dist)
}

// TODO: Replace with BFS
func shortestPathLength(start, end pos, world [][]*square) (int, pos) {
	dists := make([][]*node, len(world))
	q := make([]*node, 0)
	for y, row := range world {
		dists[y] = make([]*node, len(row))
		for x, sq := range row {
			if sq.isWall || sq.unit != nil {
				continue
			}
			v := &node{dist: 1 << 31, prev: nil, pos: pos{x, y}}
			dists[y][x] = v

			q = append(q, v)
		}
	}
	dists[start.y][start.x] = &node{dist: 0, pos: start}
	q = append(q, dists[start.y][start.x])

	for len(q) > 0 {
		bestIdx := -1
		bestDist := 1 << 32
		for idx, u := range q {
			if u.dist < bestDist {
				bestIdx = idx
				bestDist = u.dist
			}
		}
		u := q[bestIdx]
		q = append(q[:bestIdx], q[bestIdx+1:]...)

		for _, pV := range adjacentCoords(u.pos) {
			if v := dists[pV.y][pV.x]; v != nil {
				inQ := false
				for _, pot := range q {
					if pot == v {
						inQ = true
						break
					}
				}
				if inQ {
					d := u.dist + manhattan(u.pos, v.pos)
					if d < v.dist {
						v.dist = d
						v.prev = u
					} else if d == v.dist && u.pos.Less(v.prev.pos) {
						v.prev = u
					}
				}
			}
		}
	}

	n := dists[end.y][end.x]
	endNode := n
	for n.prev != nil && n.prev != dists[start.y][start.x] {
		n = n.prev
	}

	return endNode.dist, n.pos
}
