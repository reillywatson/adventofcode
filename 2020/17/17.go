package main

import (
	"fmt"
	"strings"
)

type coord struct {
	w int
	x int
	y int
	z int
}

func main() {
	grid := map[coord]bool{}
	for y, line := range strings.Split(data, "\n") {
		for x, c := range line {
			if c == '#' {
				grid[coord{0, x, y, 0}] = true
				addNeighbours(grid, coord{0, x, y, 0})
			}
		}
	}
	for i := 0; i < 6; i++ {
		for k := range grid {
			addNeighbours(grid, k)
		}
		newGrid := map[coord]bool{}
		for c, wasActive := range grid {
			count := 0
			for _, n := range neighbours(c) {
				if grid[n] {
					count++
				}
			}
			if wasActive {
				if count == 2 || count == 3 {
					newGrid[c] = true
				} else {
					newGrid[c] = false
				}
			} else {
				if count == 3 {
					newGrid[c] = true
				}
			}
		}
		grid = newGrid
	}
	numActive := 0
	for _, v := range grid {
		if v {
			numActive++
		}
	}
	fmt.Println(numActive)
}

func addNeighbours(grid map[coord]bool, c coord) {
	for _, n := range neighbours(c) {
		if _, ok := grid[n]; !ok {
			grid[n] = false
		}
	}
}

func neighbours(c coord) []coord {
	var res []coord
	for _, z := range []int{c.z - 1, c.z, c.z + 1} {
		for _, y := range []int{c.y - 1, c.y, c.y + 1} {
			for _, x := range []int{c.x - 1, c.x, c.x + 1} {
				for _, w := range []int{c.w - 1, c.w, c.w + 1} {
					n := coord{w: w, x: x, y: y, z: z}
					if n != c {
						res = append(res, n)
					}
				}
			}
		}
	}
	return res
}

func main_partone() {
	grid := map[coord]bool{}
	for y, line := range strings.Split(data, "\n") {
		for x, c := range line {
			if c == '#' {
				grid[coord{0, x, y, 0}] = true
				addNeighbours_partone(grid, coord{0, x, y, 0})
			}
		}
	}
	for i := 0; i < 6; i++ {
		for k := range grid {
			addNeighbours_partone(grid, k)
		}
		newGrid := map[coord]bool{}
		for c, wasActive := range grid {
			count := 0
			for _, n := range neighbours_partone(c) {
				if grid[n] {
					count++
				}
			}
			if wasActive {
				if count == 2 || count == 3 {
					newGrid[c] = true
				} else {
					newGrid[c] = false
				}
			} else {
				if count == 3 {
					newGrid[c] = true
				}
			}
		}
		grid = newGrid
	}
	numActive := 0
	for _, v := range grid {
		if v {
			numActive++
		}
	}
	fmt.Println(numActive)
}

func addNeighbours_partone(grid map[coord]bool, c coord) {
	for _, n := range neighbours_partone(c) {
		if _, ok := grid[n]; !ok {
			grid[n] = false
		}
	}
}

func neighbours_partone(c coord) []coord {
	var res []coord
	for _, z := range []int{c.z - 1, c.z, c.z + 1} {
		for _, y := range []int{c.y - 1, c.y, c.y + 1} {
			for _, x := range []int{c.x - 1, c.x, c.x + 1} {
				n := coord{x: x, y: y, z: z}
				if n != c {
					res = append(res, n)
				}
			}
		}
	}
	return res
}

var testdata = `.#.
..#
###`

var data = `##...#.#
####.#.#
#...####
..#.#.#.
####.#..
#.#.#..#
.####.##
..#...##`
