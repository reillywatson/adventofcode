package main

import (
	"fmt"
	"strconv"
	"strings"
)

func main_part1() {
	var data []int
	for _, s := range strings.Split(input, ",") {
		i, _ := strconv.Atoi(s)
		data = append(data, i)
	}
	pc := 0
	for {
		var shouldContinue bool
		pc, shouldContinue = processOp(data, pc)
		if !shouldContinue {
			break
		}
	}
	var res []string

	for _, d := range data {
		res = append(res, strconv.Itoa(d))
	}
}

func main() {
	var data []int
	for _, s := range strings.Split(input, ",") {
		i, _ := strconv.Atoi(s)
		data = append(data, i)
	}
	pc := 0
	for {
		var shouldContinue bool
		pc, shouldContinue = processOp(data, pc)
		if !shouldContinue {
			break
		}
	}
	var res []string

	for _, d := range data {
		res = append(res, strconv.Itoa(d))
	}
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

func processOp(data []int, pc int) (int, bool) {
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
		fmt.Scanf("%d", addr(data, pc+1, modes[0]))
		return pc + 2, true
	case 4:
		fmt.Println(*addr(data, pc+1, modes[0]))
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

var input = `3,225,1,225,6,6,1100,1,238,225,104,0,1101,33,37,225,101,6,218,224,1001,224,-82,224,4,224,102,8,223,223,101,7,224,224,1,223,224,223,1102,87,62,225,1102,75,65,224,1001,224,-4875,224,4,224,1002,223,8,223,1001,224,5,224,1,224,223,223,1102,49,27,225,1101,6,9,225,2,69,118,224,101,-300,224,224,4,224,102,8,223,223,101,6,224,224,1,224,223,223,1101,76,37,224,1001,224,-113,224,4,224,1002,223,8,223,101,5,224,224,1,224,223,223,1101,47,50,225,102,43,165,224,1001,224,-473,224,4,224,102,8,223,223,1001,224,3,224,1,224,223,223,1002,39,86,224,101,-7482,224,224,4,224,102,8,223,223,1001,224,6,224,1,223,224,223,1102,11,82,225,1,213,65,224,1001,224,-102,224,4,224,1002,223,8,223,1001,224,6,224,1,224,223,223,1001,14,83,224,1001,224,-120,224,4,224,1002,223,8,223,101,1,224,224,1,223,224,223,1102,53,39,225,1101,65,76,225,4,223,99,0,0,0,677,0,0,0,0,0,0,0,0,0,0,0,1105,0,99999,1105,227,247,1105,1,99999,1005,227,99999,1005,0,256,1105,1,99999,1106,227,99999,1106,0,265,1105,1,99999,1006,0,99999,1006,227,274,1105,1,99999,1105,1,280,1105,1,99999,1,225,225,225,1101,294,0,0,105,1,0,1105,1,99999,1106,0,300,1105,1,99999,1,225,225,225,1101,314,0,0,106,0,0,1105,1,99999,1107,677,226,224,1002,223,2,223,1005,224,329,101,1,223,223,8,677,226,224,102,2,223,223,1006,224,344,1001,223,1,223,108,677,677,224,1002,223,2,223,1006,224,359,1001,223,1,223,1108,226,677,224,102,2,223,223,1006,224,374,1001,223,1,223,1008,677,226,224,102,2,223,223,1005,224,389,101,1,223,223,7,226,677,224,102,2,223,223,1005,224,404,1001,223,1,223,1007,677,677,224,1002,223,2,223,1006,224,419,101,1,223,223,107,677,226,224,102,2,223,223,1006,224,434,101,1,223,223,7,677,677,224,1002,223,2,223,1005,224,449,101,1,223,223,108,677,226,224,1002,223,2,223,1006,224,464,101,1,223,223,1008,226,226,224,1002,223,2,223,1006,224,479,101,1,223,223,107,677,677,224,1002,223,2,223,1006,224,494,1001,223,1,223,1108,677,226,224,102,2,223,223,1005,224,509,101,1,223,223,1007,226,677,224,102,2,223,223,1005,224,524,1001,223,1,223,1008,677,677,224,102,2,223,223,1005,224,539,1001,223,1,223,1107,677,677,224,1002,223,2,223,1006,224,554,1001,223,1,223,1007,226,226,224,1002,223,2,223,1005,224,569,1001,223,1,223,7,677,226,224,1002,223,2,223,1006,224,584,1001,223,1,223,108,226,226,224,102,2,223,223,1005,224,599,1001,223,1,223,8,677,677,224,102,2,223,223,1005,224,614,1001,223,1,223,1107,226,677,224,102,2,223,223,1005,224,629,1001,223,1,223,8,226,677,224,102,2,223,223,1006,224,644,1001,223,1,223,1108,226,226,224,1002,223,2,223,1006,224,659,101,1,223,223,107,226,226,224,1002,223,2,223,1006,224,674,1001,223,1,223,4,223,99,226`
