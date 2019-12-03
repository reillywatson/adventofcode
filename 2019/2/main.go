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
	data[1] = 12
	data[2] = 2
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
	fmt.Println(strings.Join(res, ","))
}

func main() {
	for noun := 0; noun < 100; noun++ {
		for verb := 0; verb < 100; verb++ {
			var data []int
			for _, s := range strings.Split(input, ",") {
				i, _ := strconv.Atoi(s)
				data = append(data, i)
			}
			data[1] = noun
			data[2] = verb
			pc := 0
			for {
				var shouldContinue bool
				pc, shouldContinue = processOp(data, pc)
				if !shouldContinue {
					break
				}
			}
			if data[0] == 19690720 {
				fmt.Println(100*noun + verb)
				return
			}
		}
	}
	fmt.Println("NOT FOUND!")
}

func processOp(data []int, pc int) (int, bool) {
	switch data[pc] {
	case 1:
		data[data[pc+3]] = data[data[pc+1]] + data[data[pc+2]]
	case 2:
		data[data[pc+3]] = data[data[pc+1]] * data[data[pc+2]]
	case 99:
		return 0, false
	}
	return pc + 4, true
}

var input = `1,0,0,3,1,1,2,3,1,3,4,3,1,5,0,3,2,6,1,19,1,19,5,23,2,9,23,27,1,5,27,31,1,5,31,35,1,35,13,39,1,39,9,43,1,5,43,47,1,47,6,51,1,51,13,55,1,55,9,59,1,59,13,63,2,63,13,67,1,67,10,71,1,71,6,75,2,10,75,79,2,10,79,83,1,5,83,87,2,6,87,91,1,91,6,95,1,95,13,99,2,99,13,103,1,103,9,107,1,10,107,111,2,111,13,115,1,10,115,119,1,10,119,123,2,13,123,127,2,6,127,131,1,13,131,135,1,135,2,139,1,139,6,0,99,2,0,14,0`
