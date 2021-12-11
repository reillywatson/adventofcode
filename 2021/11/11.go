package main

import (
	"fmt"
	"strconv"
	"strings"
)

type grid struct {
	cells              [][]int
	currentStepFlashes map[string]bool
	totalFlashes       int
}

func (g *grid) inc(x, y int) {
	if x < 0 || y < 0 {
		return
	}
	if y >= len(g.cells) {
		return
	}
	if x >= len(g.cells[y]) {
		return
	}
	g.cells[y][x] = g.cells[y][x] + 1
	key := fmt.Sprintf("%d %d", x, y)
	if g.cells[y][x] > 9 && !g.currentStepFlashes[key] {
		g.currentStepFlashes[key] = true
		g.totalFlashes++
		g.inc(x-1, y-1)
		g.inc(x-1, y)
		g.inc(x-1, y+1)
		g.inc(x, y-1)
		g.inc(x, y+1)
		g.inc(x+1, y-1)
		g.inc(x+1, y)
		g.inc(x+1, y+1)
	}
}

func (g *grid) resetFlashes() {
	for k := range g.currentStepFlashes {
		var x, y int
		fmt.Sscanf(k, "%d %d", &x, &y)
		g.cells[y][x] = 0
	}
	g.currentStepFlashes = map[string]bool{}
}

func (g *grid) print() {
	for _, row := range g.cells {
		for _, cell := range row {
			fmt.Printf("%d", cell)
		}
		fmt.Print("\n")
	}
	fmt.Println()
}

func main() {
	grid := grid{currentStepFlashes: map[string]bool{}}
	gridSize := 0
	for _, line := range strings.Split(input, "\n") {
		var row []int
		for _, c := range line {
			n, _ := strconv.Atoi(string(c))
			row = append(row, n)
			gridSize++
		}
		grid.cells = append(grid.cells, row)
	}
	numSteps := 0
	for {
		numSteps++
		for y := 0; y < len(grid.cells); y++ {
			for x := 0; x < len(grid.cells[y]); x++ {
				grid.inc(x, y)
			}
		}
		if len(grid.currentStepFlashes) == gridSize {
			break
		}
		grid.resetFlashes()
		grid.print()
	}
	fmt.Println(numSteps)
}

func main_partone() {
	grid := grid{currentStepFlashes: map[string]bool{}}
	for _, line := range strings.Split(input, "\n") {
		var row []int
		for _, c := range line {
			n, _ := strconv.Atoi(string(c))
			row = append(row, n)
		}
		grid.cells = append(grid.cells, row)
	}
	numSteps := 100
	for i := 0; i < numSteps; i++ {
		for y := 0; y < len(grid.cells); y++ {
			for x := 0; x < len(grid.cells[y]); x++ {
				grid.inc(x, y)
			}
		}
		grid.resetFlashes()
	}
	fmt.Println(grid.totalFlashes)
}

var testInput = `5483143223
2745854711
5264556173
6141336146
6357385478
4167524645
2176841721
6882881134
4846848554
5283751526`

var input = `4836484555
4663841772
3512484556
1481547572
7741183422
8683222882
4215244233
1544712171
5725855786
1717382281`
