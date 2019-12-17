package main

import (
	"fmt"
	"io"
	"math/rand"
	"os"
	"sort"
	"strconv"
	"strings"
)

const gridWidth = 64
const gridHeight = 64

type computer struct {
	data         []int
	pc           int
	relativeBase int
	in           io.Reader
	out          io.Writer
}

type dir int

const (
	north dir = 1
	south dir = 2
	west  dir = 3
	east  dir = 4
)

type resp int

const (
	wall   resp = 0
	moved  resp = 1
	oxygen resp = 2
)

type monitor struct {
	pixels [][]int
	xpos   int
	ypos   int
	startx int
	starty int
	dir    dir
	oxyx   int
	oxyy   int
	steps  int
}

func (m *monitor) print(costs map[point]int) {
	s := "\033[H\033[2J"
	for y, row := range m.pixels {
		for x, col := range row {
			if x == m.xpos && y == m.ypos {
				s += "D"
				continue
			}
			if x == gridWidth/2 && y == gridHeight/2 {
				s += "X"
				continue
			}
			switch col {
			case 0:
				s += " "
			case 1:
				if c, ok := costs[point{x, y}]; ok {
					s += strconv.Itoa(c)[:1]
				} else {
					s += "?"
				}
			case 2:
				s += "#"
			}
		}
		s += "\n"
	}
	fmt.Println(s)
}

func (m *monitor) Read(p []byte) (n int, err error) {
	m.dir = dir(rand.Intn(4) + 1)
	p[0] = strconv.Itoa(int(m.dir))[0]
	return 1, nil
}

func (m *monitor) Write(p []byte) (n int, err error) {
	m.steps++
	if m.steps == 5000000 {
		path := shortestPath(m.pixels, m.oxyx, m.oxyy, m.startx, m.starty)
		m.print(path)
		fmt.Println("YOOOOOOO", m.oxyx, m.oxyy)
		fmt.Println(path)
		max := 0
		for _, c := range path {
			if c > max {
				max = c
			}
		}
		fmt.Println("MAX:", max)
		os.Exit(0)
	}
	var s string
	for _, b := range p {
		if b == ' ' {
			break
		}
		s += string([]byte{b})
	}
	b, err := strconv.Atoi(s)
	goalX, goalY := m.xpos, m.ypos
	switch m.dir {
	case north:
		goalY--
	case south:
		goalY++
	case west:
		goalX--
	case east:
		goalX++
	}
	switch resp(b) {
	case wall:
		m.pixels[goalY][goalX] = 2
	case moved:
		m.pixels[goalY][goalX] = 1
		m.xpos, m.ypos = goalX, goalY
	case oxygen:
		m.xpos, m.ypos = goalX, goalY
		m.pixels[goalY][goalX] = 1
		m.oxyx, m.oxyy = goalX, goalY
	}
	return len(s), err
}

type point struct {
	x int
	y int
}

func shortestPath(grid [][]int, startx, starty, goalx, goaly int) map[point]int {
	inf := 10000000000
	costs := map[point]int{}
	unvisited := []point{}
	visitedMap := map[point]bool{}
	for y, row := range grid {
		for x, cell := range row {
			if cell == 1 {
				costs[point{x, y}] = inf
				unvisited = append(unvisited, point{x, y})
			}
		}
	}
	costs[point{startx, starty}] = 0
	currentx := startx
	currenty := starty
	for len(unvisited) > 0 {
		p := point{currentx, currenty}
		cost := costs[p]
		neighbours := []point{{currentx - 1, currenty}, {currentx + 1, currenty}, {currentx, currenty - 1}, {currentx, currenty + 1}}
		for _, neighbour := range neighbours {
			if visitedMap[neighbour] {
				continue
			}
			if c, ok := costs[neighbour]; ok && c > cost+1 {
				costs[neighbour] = cost + 1
			}
		}
		for i := 0; i < len(unvisited); i++ {
			if unvisited[i] == p {
				unvisited = unvisited[:i+copy(unvisited[i:], unvisited[i+1:])]
				break
			}
		}
		sort.Slice(unvisited, func(i, j int) bool {
			return costs[unvisited[i]] < costs[unvisited[j]]
		})
		if len(unvisited) > 0 {
			currentx, currenty = unvisited[0].x, unvisited[0].y
		}
	}
	fmt.Println(costs)
	return costs
}

func main() {
	data := make([]int, 65536) // 64K should be enough for anyone!
	for i, s := range strings.Split(input, ",") {
		b, _ := strconv.Atoi(s)
		data[i] = b
	}

	var pixels [][]int
	for i := 0; i < gridHeight; i++ {
		row := make([]int, gridWidth)
		pixels = append(pixels, row)
	}
	m := &monitor{pixels: pixels, xpos: gridWidth / 2, ypos: gridHeight / 2, startx: gridWidth / 2, starty: gridHeight / 2}

	c := &computer{data: data, in: m, out: m}
	for {
		var shouldContinue bool
		c.pc, c.relativeBase, shouldContinue = processOp(c.data, c.pc, c.relativeBase, c.in, c.out)
		if !shouldContinue {
			break
		}
	}
}

type mode int

const (
	position  mode = 0
	immediate mode = 1
	relative  mode = 2
)

func parameterModes(op int) (int, map[int]mode) {
	modes := map[int]mode{}
	opStr := strconv.Itoa(op)
	if len(opStr) <= 2 {
		return op, nil
	}
	op, _ = strconv.Atoi(opStr[len(opStr)-2:])
	opNo := 0
	for i := len(opStr) - 3; i >= 0; i-- {
		val, _ := strconv.Atoi(string(opStr[i]))
		modes[opNo] = mode(val)
		opNo++
	}
	return op, modes
}

func addr(data []int, addr int, relativeBase int, mode mode) *int {
	switch mode {
	case position:
		return &data[data[addr]]
	case immediate:
		return &data[addr]
	case relative:
		return &data[data[addr]+relativeBase]
	}
	panic(fmt.Sprintf("Unknown address mode %d", mode))
}

func processOp(data []int, pc int, relativeBase int, in io.Reader, out io.Writer) (int, int, bool) {
	op := data[pc]
	op, modes := parameterModes(op)
	switch op {
	case 1: // add
		*addr(data, pc+3, relativeBase, modes[2]) = *(addr(data, pc+1, relativeBase, modes[0])) + *(addr(data, pc+2, relativeBase, modes[1]))
		return pc + 4, relativeBase, true
	case 2: // mult
		*addr(data, pc+3, relativeBase, modes[2]) = *(addr(data, pc+1, relativeBase, modes[0])) * *(addr(data, pc+2, relativeBase, modes[1]))
		return pc + 4, relativeBase, true
	case 3: // read
		bytes := make([]byte, 2)
		n, err := in.Read(bytes)
		if err == io.EOF {
			return pc, relativeBase, true // read again later!
		}
		asNum, err := strconv.Atoi(string(bytes[:n]))
		*addr(data, pc+1, relativeBase, modes[0]) = asNum
		return pc + 2, relativeBase, true
	case 4: // print
		fmt.Fprintf(out, "%d ", *addr(data, pc+1, relativeBase, modes[0]))
		return pc + 2, relativeBase, true
	case 5: // jump if non-zero
		if *addr(data, pc+1, relativeBase, modes[0]) != 0 {
			return *addr(data, pc+2, relativeBase, modes[1]), relativeBase, true
		}
		return pc + 3, relativeBase, true
	case 6: // jump if zero
		if *addr(data, pc+1, relativeBase, modes[0]) == 0 {
			return *addr(data, pc+2, relativeBase, modes[1]), relativeBase, true
		}
		return pc + 3, relativeBase, true
	case 7: // lt
		if *addr(data, pc+1, relativeBase, modes[0]) < *addr(data, pc+2, relativeBase, modes[1]) {
			*addr(data, pc+3, relativeBase, modes[2]) = 1
		} else {
			*addr(data, pc+3, relativeBase, modes[2]) = 0
		}
		return pc + 4, relativeBase, true
	case 8: // gt
		if *addr(data, pc+1, relativeBase, modes[0]) == *addr(data, pc+2, relativeBase, modes[1]) {
			*addr(data, pc+3, relativeBase, modes[2]) = 1
		} else {
			*addr(data, pc+3, relativeBase, modes[2]) = 0
		}
		return pc + 4, relativeBase, true
	case 9: // set relative base
		relativeBase = relativeBase + *addr(data, pc+1, relativeBase, modes[0])
		return pc + 2, relativeBase, true
	case 99:
		return 0, relativeBase, false
	}
	panic(fmt.Sprintf("unknown op %d at %d", op, pc))
}

var input = `3,1033,1008,1033,1,1032,1005,1032,31,1008,1033,2,1032,1005,1032,58,1008,1033,3,1032,1005,1032,81,1008,1033,4,1032,1005,1032,104,99,101,0,1034,1039,101,0,1036,1041,1001,1035,-1,1040,1008,1038,0,1043,102,-1,1043,1032,1,1037,1032,1042,1105,1,124,101,0,1034,1039,102,1,1036,1041,1001,1035,1,1040,1008,1038,0,1043,1,1037,1038,1042,1106,0,124,1001,1034,-1,1039,1008,1036,0,1041,102,1,1035,1040,1001,1038,0,1043,1001,1037,0,1042,1106,0,124,1001,1034,1,1039,1008,1036,0,1041,1001,1035,0,1040,1001,1038,0,1043,1002,1037,1,1042,1006,1039,217,1006,1040,217,1008,1039,40,1032,1005,1032,217,1008,1040,40,1032,1005,1032,217,1008,1039,7,1032,1006,1032,165,1008,1040,5,1032,1006,1032,165,1102,1,2,1044,1105,1,224,2,1041,1043,1032,1006,1032,179,1101,0,1,1044,1105,1,224,1,1041,1043,1032,1006,1032,217,1,1042,1043,1032,1001,1032,-1,1032,1002,1032,39,1032,1,1032,1039,1032,101,-1,1032,1032,101,252,1032,211,1007,0,27,1044,1106,0,224,1102,1,0,1044,1106,0,224,1006,1044,247,101,0,1039,1034,101,0,1040,1035,102,1,1041,1036,1001,1043,0,1038,102,1,1042,1037,4,1044,1106,0,0,13,3,18,86,2,10,5,16,95,16,54,4,23,63,70,10,21,20,26,99,85,9,96,3,83,5,9,91,14,1,4,78,11,15,53,10,35,13,7,17,30,90,23,65,65,67,16,4,65,39,11,57,13,36,22,95,53,63,22,47,12,47,2,12,3,71,92,17,55,16,51,79,6,3,92,15,17,15,18,63,8,12,3,49,6,69,32,1,25,83,17,12,1,76,23,95,17,13,92,13,56,16,69,94,11,20,31,83,30,21,88,22,61,45,6,70,12,3,30,23,86,6,93,4,24,9,73,72,7,72,83,9,30,6,24,86,99,11,11,96,16,68,10,35,19,23,6,79,51,8,3,8,75,2,32,26,73,23,80,30,86,25,64,46,24,81,20,18,85,7,94,28,37,93,18,12,77,99,14,22,19,50,2,18,45,63,8,2,89,79,79,7,33,77,18,20,22,12,58,61,20,4,58,20,51,79,14,32,19,87,21,19,76,8,81,7,13,72,75,22,28,22,14,92,30,18,90,10,6,97,25,34,9,20,26,52,45,6,4,97,4,46,26,86,61,20,25,28,26,22,54,69,16,51,3,58,5,23,75,92,18,98,12,11,55,38,22,87,14,20,17,52,73,9,91,30,14,26,12,56,81,54,9,72,18,12,47,93,22,54,21,59,73,7,78,12,87,26,5,39,45,4,55,16,21,86,62,20,98,61,14,20,70,14,25,92,32,44,2,3,15,32,23,23,97,76,78,15,23,95,21,11,69,34,12,89,3,95,24,15,59,38,39,72,14,15,55,48,18,2,43,26,13,58,68,11,22,89,33,79,22,43,40,14,26,5,50,11,28,9,36,33,2,22,43,21,90,15,92,14,14,49,9,80,14,85,99,70,8,16,14,15,70,1,39,32,45,5,57,12,12,4,99,75,28,14,2,28,71,5,69,61,4,28,98,97,87,10,80,2,65,93,6,21,81,7,95,22,35,18,38,23,11,53,14,5,2,84,3,70,33,19,8,52,10,99,14,58,36,1,3,30,53,4,7,47,10,93,2,32,17,40,68,43,20,41,4,16,21,29,23,82,2,18,37,37,15,19,26,41,28,9,95,17,17,52,25,13,49,28,47,22,5,52,14,21,72,83,7,17,86,20,3,18,58,14,19,25,56,65,65,26,53,8,20,75,31,21,40,17,6,33,20,95,47,24,75,26,17,96,24,48,65,97,4,52,20,78,47,14,23,77,32,8,18,98,43,7,61,25,84,40,6,36,24,87,24,71,77,13,20,49,16,60,35,9,64,48,21,2,74,25,1,2,57,11,58,7,45,35,26,13,74,92,2,9,82,9,20,23,15,33,94,7,10,48,78,16,24,94,33,11,21,5,89,47,15,52,12,51,51,81,9,18,39,14,2,97,79,33,23,12,99,3,16,11,79,83,45,18,23,78,86,69,10,25,98,62,62,18,7,44,47,1,3,92,8,22,81,9,3,29,8,81,21,13,95,6,5,99,5,29,16,3,53,72,26,14,44,97,7,43,12,42,65,17,8,12,88,55,18,20,34,13,39,10,72,58,15,11,69,17,94,20,22,52,28,13,30,65,8,2,63,18,4,36,17,8,71,16,71,15,64,14,31,51,75,1,12,92,14,35,23,40,45,1,5,87,28,18,83,43,9,90,2,3,50,18,61,68,5,89,16,44,7,34,82,74,15,83,15,70,13,80,20,43,8,35,14,58,50,75,20,50,9,68,46,52,2,73,11,60,32,61,25,40,9,31,21,73,0,0,21,21,1,10,1,0,0,0,0,0,0`
