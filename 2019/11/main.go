package main

import (
	"fmt"
	"io"
	"strconv"
	"strings"
)

var gridSize = 200

type computer struct {
	data         []int
	pc           int
	relativeBase int
	in           io.Reader
	out          io.Writer
}

type dir int

const (
	up    dir = 0
	right dir = 1
	down  dir = 2
	left  dir = 3
)

type robot struct {
	dir       dir
	x         int
	y         int
	panels    [][]byte
	painted   map[string]bool
	readColor bool
}

func (r *robot) print() {
	s := "\033[H\033[2J" // clear screen
	for y := 0; y < gridSize; y++ {
		for x := 0; x < gridSize; x++ {
			if r.x == x && r.y == y {
				switch r.dir {
				case up:
					s += "^"
				case right:
					s += ">"
				case down:
					s += "v"
				case left:
					s += "<"
				}
			} else if r.panels[y][x] == '1' {
				s += "#"
			} else {
				s += "."
			}
		}
		s += "\n"
	}
	fmt.Println(s)

}

func (r *robot) move() {
	switch r.dir {
	case up:
		r.y--
	case down:
		r.y++
	case left:
		r.x--
	case right:
		r.x++
	}
	r.print()
}

// confusingly, this is how the robot writes its output
func (r *robot) Read(p []byte) (n int, err error) {
	p[0] = r.panels[r.y][r.x]
	return 1, nil
}

// confusingly, this is how the robot gets its input
func (r *robot) Write(p []byte) (n int, err error) {
	if !r.readColor {
		r.panels[r.y][r.x] = p[0]
		r.painted[fmt.Sprintf("%dx%d", r.x, r.y)] = true
		r.readColor = true
		return 1, nil
	}
	if p[0] == '0' {
		switch r.dir {
		case up:
			r.dir = left
		case left:
			r.dir = down
		case down:
			r.dir = right
		case right:
			r.dir = up
		}
		r.move()
		r.readColor = false
		return 1, nil
	} else if p[0] == '1' {
		switch r.dir {
		case up:
			r.dir = right
		case right:
			r.dir = down
		case down:
			r.dir = left
		case left:
			r.dir = up
		}
		r.move()
		r.readColor = false
		return 1, nil
	}
	return 1, nil
}

func main_partone() {
	data := make([]int, 65536) // 64K should be enough for anyone!
	for i, s := range strings.Split(input, ",") {
		b, _ := strconv.Atoi(s)
		data[i] = b
	}
	robot := &robot{x: gridSize / 2, y: gridSize / 2, dir: up, painted: map[string]bool{}}
	for i := 0; i < gridSize; i++ {
		row := make([]byte, gridSize)
		for j := 0; j < gridSize; j++ {
			row[j] = '0'
		}
		robot.panels = append(robot.panels, row)
	}

	c := &computer{data: data, in: robot, out: robot}
	for {
		var shouldContinue bool
		c.pc, c.relativeBase, shouldContinue = processOp(c.data, c.pc, c.relativeBase, c.in, c.out)
		if !shouldContinue {
			break
		}
	}
	fmt.Println(len(robot.painted), "PANELS PAINTED")
}

func main() {
	data := make([]int, 65536) // 64K should be enough for anyone!
	for i, s := range strings.Split(input, ",") {
		b, _ := strconv.Atoi(s)
		data[i] = b
	}
	robot := &robot{x: gridSize / 2, y: gridSize / 2, dir: up, painted: map[string]bool{}}
	for i := 0; i < gridSize; i++ {
		row := make([]byte, gridSize)
		for j := 0; j < gridSize; j++ {
			row[j] = '0'
		}
		robot.panels = append(robot.panels, row)
	}
	robot.panels[robot.y][robot.x] = '1'

	c := &computer{data: data, in: robot, out: robot}
	for {
		var shouldContinue bool
		c.pc, c.relativeBase, shouldContinue = processOp(c.data, c.pc, c.relativeBase, c.in, c.out)
		if !shouldContinue {
			break
		}
	}
	fmt.Println(len(robot.painted), "PANELS PAINTED")
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
		bytes := make([]byte, 1)
		_, err := in.Read(bytes)
		if err == io.EOF {
			return pc, relativeBase, true // read again later!
		}
		asNum, err := strconv.Atoi(string(bytes[:1]))
		*addr(data, pc+1, relativeBase, modes[0]) = asNum
		return pc + 2, relativeBase, true
	case 4: // print
		fmt.Fprintf(out, "%d", *addr(data, pc+1, relativeBase, modes[0]))
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

var input = `3,8,1005,8,329,1106,0,11,0,0,0,104,1,104,0,3,8,102,-1,8,10,1001,10,1,10,4,10,1008,8,0,10,4,10,1002,8,1,29,2,1102,1,10,1,1009,16,10,2,4,4,10,1,9,5,10,3,8,1002,8,-1,10,101,1,10,10,4,10,108,0,8,10,4,10,101,0,8,66,2,106,7,10,1006,0,49,3,8,1002,8,-1,10,101,1,10,10,4,10,108,1,8,10,4,10,1002,8,1,95,1006,0,93,3,8,102,-1,8,10,1001,10,1,10,4,10,108,1,8,10,4,10,102,1,8,120,1006,0,61,2,1108,19,10,2,1003,2,10,1006,0,99,3,8,1002,8,-1,10,1001,10,1,10,4,10,1008,8,0,10,4,10,101,0,8,157,3,8,102,-1,8,10,1001,10,1,10,4,10,1008,8,1,10,4,10,1001,8,0,179,2,1108,11,10,1,1102,19,10,3,8,102,-1,8,10,1001,10,1,10,4,10,1008,8,1,10,4,10,101,0,8,209,2,108,20,10,3,8,1002,8,-1,10,101,1,10,10,4,10,108,1,8,10,4,10,101,0,8,234,3,8,102,-1,8,10,101,1,10,10,4,10,108,0,8,10,4,10,1002,8,1,256,2,1102,1,10,1006,0,69,2,108,6,10,2,4,13,10,3,8,102,-1,8,10,101,1,10,10,4,10,1008,8,0,10,4,10,1002,8,1,294,1,1107,9,10,1006,0,87,2,1006,8,10,2,1001,16,10,101,1,9,9,1007,9,997,10,1005,10,15,99,109,651,104,0,104,1,21101,387395195796,0,1,21101,346,0,0,1105,1,450,21101,0,48210129704,1,21101,0,357,0,1105,1,450,3,10,104,0,104,1,3,10,104,0,104,0,3,10,104,0,104,1,3,10,104,0,104,1,3,10,104,0,104,0,3,10,104,0,104,1,21101,0,46413147328,1,21102,404,1,0,1106,0,450,21102,179355823323,1,1,21101,415,0,0,1105,1,450,3,10,104,0,104,0,3,10,104,0,104,0,21102,1,838345843476,1,21101,0,438,0,1105,1,450,21101,709475709716,0,1,21101,449,0,0,1105,1,450,99,109,2,22102,1,-1,1,21102,40,1,2,21101,0,481,3,21101,0,471,0,1105,1,514,109,-2,2105,1,0,0,1,0,0,1,109,2,3,10,204,-1,1001,476,477,492,4,0,1001,476,1,476,108,4,476,10,1006,10,508,1101,0,0,476,109,-2,2106,0,0,0,109,4,2101,0,-1,513,1207,-3,0,10,1006,10,531,21101,0,0,-3,21201,-3,0,1,21201,-2,0,2,21101,1,0,3,21101,550,0,0,1105,1,555,109,-4,2106,0,0,109,5,1207,-3,1,10,1006,10,578,2207,-4,-2,10,1006,10,578,21201,-4,0,-4,1105,1,646,22101,0,-4,1,21201,-3,-1,2,21202,-2,2,3,21101,597,0,0,1105,1,555,22102,1,1,-4,21101,0,1,-1,2207,-4,-2,10,1006,10,616,21101,0,0,-1,22202,-2,-1,-2,2107,0,-3,10,1006,10,638,22102,1,-1,1,21101,638,0,0,106,0,513,21202,-2,-1,-2,22201,-4,-2,-4,109,-5,2106,0,0`
