package main

import (
	"bytes"
	"fmt"
	"io"
	"strconv"
	"strings"
)

func main_part1() {
	in := bytes.NewBuffer(nil)
	max := 0
	var maxSeq []int
	for _, settings := range allSettings_partone() {
		lastOut := 0
		for i := 0; i < 5; i++ {
			var data []int
			for _, s := range strings.Split(input, ",") {
				i, _ := strconv.Atoi(s)
				data = append(data, i)
			}
			fmt.Fprintf(in, "%d ", settings[i])
			fmt.Fprintf(in, "%d ", lastOut)
			pc := 0
			out := bytes.NewBuffer(nil)
			for {
				var shouldContinue bool
				pc, shouldContinue = processOp(i, data, pc, in, out)
				if !shouldContinue {
					break
				}
			}
			outstr := out.String()
			lastOut, _ = strconv.Atoi(strings.TrimSpace(outstr))
		}
		if lastOut > max {
			maxSeq = settings
			max = lastOut
		}
	}
	fmt.Println(max, maxSeq)
}

func main() {
	max := 0
	var maxSeq []int

	for _, settings := range allSettings() {

		type computer struct {
			data   []int
			pc     int
			in     *bytes.Buffer
			out    *bytes.Buffer
			halted bool
		}
		var computers []*computer
		for i := 0; i < 5; i++ {
			var data []int
			for _, s := range strings.Split(input, ",") {
				i, _ := strconv.Atoi(s)
				data = append(data, i)
			}
			c := &computer{data: data, in: bytes.NewBuffer(nil), out: bytes.NewBuffer(nil)}
			if i > 0 {
				c.in = computers[i-1].out
			}
			computers = append(computers, c)
		}
		computers[0].in = computers[4].out

		lastOut := 0
		for i, computer := range computers {
			fmt.Fprintf(computer.in, "%d ", settings[i])
		}
		fmt.Fprintf(computers[0].in, "%d ", 0)

		for {
			anyRunning := false
			for i, computer := range computers {
				if computer.halted {
					continue
				}
				anyRunning = true
				var shouldContinue bool
				computer.pc, shouldContinue = processOp(i, computer.data, computer.pc, computer.in, computer.out)
				if !shouldContinue {
					computer.halted = true
					break
				}
			}
			if !anyRunning {
				break
			}
		}
		lastOut, _ = strconv.Atoi(strings.TrimSpace(computers[4].out.String()))
		if lastOut > max {
			maxSeq = settings
			max = lastOut
		}
	}
	fmt.Println(max, maxSeq)
}

// so lazy!
func allSettings_partone() [][]int {
	var res [][]int
	for i := 0; i < 5; i++ {
		for j := 0; j < 5; j++ {
			if j == i {
				continue
			}
			for k := 0; k < 5; k++ {
				if k == i || k == j {
					continue
				}
				for l := 0; l < 5; l++ {
					if l == i || l == j || l == k {
						continue
					}
					for m := 0; m < 5; m++ {
						if m == i || m == j || m == k || m == l {
							continue
						}
						res = append(res, []int{i, j, k, l, m})
					}
				}
			}
		}
	}
	return res
}

// even lazier!
func allSettings() [][]int {
	var res [][]int
	for i := 5; i < 10; i++ {
		for j := 5; j < 10; j++ {
			if j == i {
				continue
			}
			for k := 5; k < 10; k++ {
				if k == i || k == j {
					continue
				}
				for l := 5; l < 10; l++ {
					if l == i || l == j || l == k {
						continue
					}
					for m := 5; m < 10; m++ {
						if m == i || m == j || m == k || m == l {
							continue
						}
						res = append(res, []int{i, j, k, l, m})
					}
				}
			}
		}
	}
	return res
}

type mode int

const (
	position  mode = 0
	immediate mode = 1
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

func addr(data []int, addr int, mode mode) *int {
	if mode == position {
		return &data[data[addr]]
	}
	return &data[addr]
}

func processOp(pcNo int, data []int, pc int, in io.Reader, out io.Writer) (int, bool) {
	op := data[pc]
	op, modes := parameterModes(op)
	switch op {
	case 1:
		*addr(data, pc+3, modes[2]) = *(addr(data, pc+1, modes[0])) + *(addr(data, pc+2, modes[1]))
		return pc + 4, true
	case 2:
		*addr(data, pc+3, modes[2]) = *(addr(data, pc+1, modes[0])) * *(addr(data, pc+2, modes[1]))
		return pc + 4, true
	case 3:
		_, err := fmt.Fscanf(in, "%d", addr(data, pc+1, modes[0]))
		if err == io.EOF {
			return pc, true // read again later!
		}
		fmt.Println(pcNo, "GOT IN:", *addr(data, pc+1, modes[0]))
		return pc + 2, true
	case 4:
		fmt.Println(pcNo, "OUTPUTTING:", *addr(data, pc+1, modes[0]))
		fmt.Fprintf(out, "%d ", *addr(data, pc+1, modes[0]))
		return pc + 2, true
	case 5:
		if *addr(data, pc+1, modes[0]) != 0 {
			return *addr(data, pc+2, modes[1]), true
		}
		return pc + 3, true
	case 6:
		if *addr(data, pc+1, modes[0]) == 0 {
			return *addr(data, pc+2, modes[1]), true
		}
		return pc + 3, true
	case 7:
		if *addr(data, pc+1, modes[0]) < *addr(data, pc+2, modes[1]) {
			*addr(data, pc+3, modes[2]) = 1
		} else {
			*addr(data, pc+3, modes[2]) = 0
		}
		return pc + 4, true
	case 8:
		if *addr(data, pc+1, modes[0]) == *addr(data, pc+2, modes[1]) {
			*addr(data, pc+3, modes[2]) = 1
		} else {
			*addr(data, pc+3, modes[2]) = 0
		}
		return pc + 4, true
	case 99:
		return 0, false
	}
	panic(fmt.Sprintf("unknown op %d at %d", data[pc], pc))
}

//var input = `3,26,1001,26,-4,26,3,27,1002,27,2,27,1,27,26,27,4,27,1001,28,-1,28,1005,28,6,99,0,0,5`
var input = `3,8,1001,8,10,8,105,1,0,0,21,38,59,84,97,110,191,272,353,434,99999,3,9,1002,9,2,9,101,4,9,9,1002,9,2,9,4,9,99,3,9,102,5,9,9,1001,9,3,9,1002,9,5,9,101,5,9,9,4,9,99,3,9,102,5,9,9,101,5,9,9,1002,9,3,9,101,2,9,9,1002,9,4,9,4,9,99,3,9,101,3,9,9,1002,9,3,9,4,9,99,3,9,102,5,9,9,1001,9,3,9,4,9,99,3,9,101,2,9,9,4,9,3,9,102,2,9,9,4,9,3,9,101,1,9,9,4,9,3,9,1002,9,2,9,4,9,3,9,1002,9,2,9,4,9,3,9,101,2,9,9,4,9,3,9,1001,9,2,9,4,9,3,9,102,2,9,9,4,9,3,9,101,2,9,9,4,9,3,9,1002,9,2,9,4,9,99,3,9,1002,9,2,9,4,9,3,9,1002,9,2,9,4,9,3,9,1002,9,2,9,4,9,3,9,101,1,9,9,4,9,3,9,1002,9,2,9,4,9,3,9,101,1,9,9,4,9,3,9,101,2,9,9,4,9,3,9,1001,9,1,9,4,9,3,9,1001,9,1,9,4,9,3,9,1001,9,2,9,4,9,99,3,9,1001,9,2,9,4,9,3,9,1002,9,2,9,4,9,3,9,102,2,9,9,4,9,3,9,1002,9,2,9,4,9,3,9,1001,9,2,9,4,9,3,9,1001,9,2,9,4,9,3,9,1001,9,1,9,4,9,3,9,1002,9,2,9,4,9,3,9,102,2,9,9,4,9,3,9,1002,9,2,9,4,9,99,3,9,101,2,9,9,4,9,3,9,101,1,9,9,4,9,3,9,102,2,9,9,4,9,3,9,101,1,9,9,4,9,3,9,101,2,9,9,4,9,3,9,101,1,9,9,4,9,3,9,102,2,9,9,4,9,3,9,1001,9,2,9,4,9,3,9,1002,9,2,9,4,9,3,9,1002,9,2,9,4,9,99,3,9,1001,9,1,9,4,9,3,9,102,2,9,9,4,9,3,9,102,2,9,9,4,9,3,9,1001,9,2,9,4,9,3,9,101,1,9,9,4,9,3,9,1002,9,2,9,4,9,3,9,1001,9,1,9,4,9,3,9,102,2,9,9,4,9,3,9,1001,9,2,9,4,9,3,9,101,1,9,9,4,9,99`
