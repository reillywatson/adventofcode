package main

import (
	"fmt"
	"math"
	"sort"
	"strings"
)

type cellType byte

const (
	elf    byte = 'E'
	goblin byte = 'G'
	wall   byte = '#'
	open   byte = '.'
)

type unit struct {
	x        int
	y        int
	hp       int
	att      int
	unitType byte
}

func main_part1() {
	var grid [][]byte
	for _, line := range strings.Split(input, "\n") {
		grid = append(grid, []byte(line))
	}
	var units []*unit
	for i, row := range grid {
		for j, cell := range row {
			switch cell {
			case elf, goblin:
				units = append(units, &unit{x: j, y: i, hp: 200, att: 3, unitType: cell})
			}
		}
	}
	round := 0
	printGrid(grid, units)
	for {
		sort.Slice(units, func(i, j int) bool { return readingOrderCmp(units, i, j) })
		for i, unit := range units {
			if unit.hp <= 0 {
				continue
			}
			if target := attackTarget(unit, units); target != nil {
				fmt.Println(unit.x, unit.y, target.x, target.y)
				target.hp = target.hp - unit.att
				if target.hp <= 0 {
					grid[target.y][target.x] = open
				}
			} else if next := moveCoord(unit, grid); next != nil {
				fmt.Println(unit.x, unit.y, next)
				grid[unit.y][unit.x] = open
				unit.x = next.x
				unit.y = next.y
				grid[unit.y][unit.x] = unit.unitType
				if target := attackTarget(unit, units); target != nil {
					target.hp = target.hp - unit.att
					if target.hp <= 0 {
						grid[target.y][target.x] = open
					}
				}
			}
			if elfHp, goblinHp := unitHps(units); elfHp == 0 || goblinHp == 0 {
				if i < len(units)-1 {
					round--
					break
				}
			}
		}
		round++
		printGrid(grid, units)
		elfHp, goblinHp := unitHps(units)
		if goblinHp == 0 || elfHp == 0 {
			winner := "Goblins"
			if elfHp > 0 {
				winner = "Elves"
			}
			fmt.Printf("Combat ends after %d full rounds\n%s win with %d total hit points left\nOutcome: %d * %d = %d\n", round, winner, goblinHp+elfHp, round, goblinHp+elfHp, round*(goblinHp+elfHp))
			return
		}
	}
}

func main() {
mainLoop:
	for elfPower := 4; ; elfPower++ {
		var grid [][]byte
		for _, line := range strings.Split(input, "\n") {
			grid = append(grid, []byte(line))
		}
		var units []*unit
		for i, row := range grid {
			for j, cell := range row {
				switch cell {
				case elf:
					units = append(units, &unit{x: j, y: i, hp: 200, att: elfPower, unitType: cell})
				case goblin:
					units = append(units, &unit{x: j, y: i, hp: 200, att: 3, unitType: cell})
				}
			}
		}
		round := 0
		printGrid(grid, units)
		for {
			sort.Slice(units, func(i, j int) bool { return readingOrderCmp(units, i, j) })
			for i, unit := range units {
				if unit.hp <= 0 {
					continue
				}
				if target := attackTarget(unit, units); target != nil {
					target.hp = target.hp - unit.att
					if target.hp <= 0 {
						grid[target.y][target.x] = open
						if target.unitType == elf {
							continue mainLoop
						}
					}
				} else if next := moveCoord(unit, grid); next != nil {
					grid[unit.y][unit.x] = open
					unit.x = next.x
					unit.y = next.y
					grid[unit.y][unit.x] = unit.unitType
					if target := attackTarget(unit, units); target != nil {
						target.hp = target.hp - unit.att
						if target.hp <= 0 {
							grid[target.y][target.x] = open
							if target.unitType == elf {
								continue mainLoop
							}
						}
					}
				}
				if elfHp, goblinHp := unitHps(units); elfHp == 0 || goblinHp == 0 {
					if i < len(units)-1 {
						round--
						break
					}
				}
			}
			round++
			printGrid(grid, units)
			elfHp, goblinHp := unitHps(units)
			if goblinHp == 0 || elfHp == 0 {
				winner := "Goblins"
				if elfHp > 0 {
					winner = "Elves"
				}
				fmt.Println("ELF POWER:", elfPower)
				fmt.Printf("Combat ends after %d full rounds\n%s win with %d total hit points left\nOutcome: %d * %d = %d\n", round, winner, goblinHp+elfHp, round, goblinHp+elfHp, round*(goblinHp+elfHp))
				return
			}
		}
	}
}

func unitHps(units []*unit) (int, int) {
	var elfHp, goblinHp int
	for _, unit := range units {
		if unit.hp > 0 {
			if unit.unitType == elf {
				elfHp += unit.hp
			} else {
				goblinHp += unit.hp
			}
		}
	}
	return elfHp, goblinHp
}

type coord struct {
	x int
	y int
}

// breadth-first search to find the next move coordinate
func moveCoord(unit *unit, grid [][]byte) *coord {
	unitCoord := coord{unit.x, unit.y}
	parents := map[coord]coord{}
	seen := map[coord]bool{}
	openList := []coord{unitCoord}
	for len(openList) > 0 {
		last := openList[0]
		openList = openList[1:]
		for _, next := range []coord{{last.x, last.y - 1}, {last.x - 1, last.y}, {last.x + 1, last.y}, {last.x, last.y + 1}} {
			if next.y < 0 || next.y >= len(grid) {
				continue
			}
			if next.x < 0 || next.x >= len(grid[next.y]) {
				continue
			}
			if (grid[next.y][next.x] == elf && unit.unitType == goblin) || (grid[next.y][next.x] == goblin && unit.unitType == elf) {
				var result coord
				parents[next] = last
				for result = parents[next]; result != unitCoord; result = parents[result] {
					last = result
				}
				return &last
			}
			if grid[next.y][next.x] != open {
				continue
			}
			if seen[next] {
				continue
			}
			inOpen := false
			for _, o := range openList {
				if next == o {
					inOpen = true
					break
				}
			}
			if !inOpen {
				parents[next] = last
				openList = append(openList, next)
			}
		}
		seen[last] = true
	}
	return nil
}

func printGrid(grid [][]byte, units []*unit) {
	for y, row := range grid {
		line := string(row)
		var unitStrings []string
		for x := range row {
			for _, unit := range units {
				if unit.x == x && unit.y == y && unit.hp > 0 {
					unitStrings = append(unitStrings, fmt.Sprintf("%s(%d)", string([]byte{unit.unitType}), unit.hp))
				}
			}
		}
		fmt.Println(line, " ", strings.Join(unitStrings, ", "))
	}
}

func attackTarget(u *unit, targets []*unit) *unit {
	var potentials []*unit
	for _, pot := range targets {
		if dist(u, pot) == 1 && u.unitType != pot.unitType && pot.hp > 0 {
			potentials = append(potentials, pot)
		}
	}
	sort.Slice(potentials, func(i, j int) bool {
		if potentials[i].hp < potentials[j].hp {
			return true
		}
		if potentials[i].hp > potentials[j].hp {
			return false
		}
		return readingOrderCmp(potentials, i, j)
	})
	if len(potentials) == 0 {
		return nil
	}
	return potentials[0]
}

func dist(u, target *unit) int {
	return int(math.Abs(float64(target.x-u.x))) + int(math.Abs(float64(target.y-u.y)))
}

func readingOrderCmp(units []*unit, i, j int) bool {
	if units[i].y < units[j].y {
		return true
	}
	if units[i].y > units[j].y {
		return false
	}
	return units[i].x < units[j].x
}

var input2 = `#######
#.G...#
#...EG#
#.#.#G#
#..G#E#
#.....#
#######`

var input = `################################
######......###...##..##########
######....#G###G..##.G##########
#####...G##.##.........#########
##....##..#.##...........#######
#....#G.......##.........G.#####
##..##GG....G.................##
##.......G............#.......##
###.....G.....G#......E.......##
##......##....................##
#.....####......G.....#...######
#.#########.G....G....#E.#######
###########...#####......#######
###########..#######..E.......##
###########.#########......#.###
########..#.#########.........##
#######G....#########........###
##.##.#.....#########...EE#..#.#
#...GG......#########.#...##..E#
##...#.......#######..#...#....#
###.##........#####......##...##
###.........................#..#
####.............##........###.#
####............##.........#####
####..##....###.#...#.....######
########....###..............###
########..G...##.###...E...E.###
#########...G.##.###.E....E.####
#########...#.#######.......####
#############..########...######
##############.########.########
################################`
