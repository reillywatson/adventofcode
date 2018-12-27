package main

import (
	"fmt"
	"strings"
)

func main_part1() {
	regs := []int{0, 0, 0, 0}
	result := []int{0, 0, 0, 0}
	in := []int{0, 0, 0, 0}
	ambiguousInstructions := 0
	for _, test := range strings.Split(input, "\n\n") {
		parts := strings.Split(test, "\n")
		fmt.Sscanf(parts[0], "Before: [%d, %d, %d, %d]", &regs[0], &regs[1], &regs[2], &regs[3])
		fmt.Sscanf(parts[1], "%d  %d %d %d", &in[0], &in[1], &in[2], &in[3])
		fmt.Sscanf(parts[2], "After: [%d, %d, %d, %d]", &result[0], &result[1], &result[2], &result[3])
		fmt.Println("REGS:", regs, "in:", in, "RESULT:", result)
		tests := []struct {
			name string
			fn   func([]int, []int) []int
		}{
			{"addr", addr},
			{"addi", addi},
			{"mulr", mulr},
			{"muli", muli},
			{"banr", banr},
			{"bani", bani},
			{"borr", borr},
			{"bori", bori},
			{"setr", setr},
			{"seti", seti},
			{"gtir", gtir},
			{"gtri", gtri},
			{"gtrr", gtrr},
			{"eqir", eqir},
			{"eqri", eqri},
			{"eqrr", eqrr},
		}
		numOpcodes := 0
		for _, test := range tests {
			var testRegs []int
			for _, x := range regs {
				testRegs = append(testRegs, x)
			}
			got := test.fn(testRegs, in)
			eq := true
			for i := 0; i < len(got); i++ {
				if got[i] != result[i] {
					eq = false
					break
				}
			}
			if eq {
				numOpcodes++
			}
		}
		if numOpcodes >= 3 {
			ambiguousInstructions++
		}
	}
	fmt.Println(ambiguousInstructions)
}

func getOpcodes() {
	regs := []int{0, 0, 0, 0}
	result := []int{0, 0, 0, 0}
	in := []int{0, 0, 0, 0}
	potentialOpcodes := map[int]map[string]bool{}
	tests := []struct {
		name string
		fn   func([]int, []int) []int
	}{
		{"addr", addr},
		{"addi", addi},
		{"mulr", mulr},
		{"muli", muli},
		{"banr", banr},
		{"bani", bani},
		{"borr", borr},
		{"bori", bori},
		{"setr", setr},
		{"seti", seti},
		{"gtir", gtir},
		{"gtri", gtri},
		{"gtrr", gtrr},
		{"eqir", eqir},
		{"eqri", eqri},
		{"eqrr", eqrr},
	}
	for i := 0; i < 16; i++ {
		names := map[string]bool{}
		for _, test := range tests {
			names[test.name] = true
		}
		potentialOpcodes[i] = names
	}
	for _, test := range strings.Split(input, "\n\n") {
		parts := strings.Split(test, "\n")
		fmt.Sscanf(parts[0], "Before: [%d, %d, %d, %d]", &regs[0], &regs[1], &regs[2], &regs[3])
		fmt.Sscanf(parts[1], "%d  %d %d %d", &in[0], &in[1], &in[2], &in[3])
		fmt.Sscanf(parts[2], "After: [%d, %d, %d, %d]", &result[0], &result[1], &result[2], &result[3])
		opCode := in[0]
		//fmt.Println("REGS:", regs, "in:", in, "RESULT:", result)
		var valid []string
		for _, test := range tests {
			var testRegs []int
			for _, x := range regs {
				testRegs = append(testRegs, x)
			}
			got := test.fn(testRegs, in)
			eq := true
			for i := 0; i < len(got); i++ {
				if got[i] != result[i] {
					eq = false
					break
				}
			}
			if eq {
				valid = append(valid, test.name)
			}
		}
		fmt.Println(opCode, valid)
		for _, test := range tests {
			potential := false
			for _, p := range valid {
				if p == test.name && potentialOpcodes[opCode][test.name] {
					potential = true
					break
				}
			}
			if !potential {
				delete(potentialOpcodes[opCode], test.name)
			}
		}
		for i := 0; i < 16; i++ {
			if len(potentialOpcodes[i]) == 1 {
				// found it, remove it from the other potentials!
				for j := 0; j < 16; j++ {
					for k := range potentialOpcodes[i] {
						if i != j && potentialOpcodes[j][k] {
							fmt.Println("Removing:", j, k, "Correct:", i, potentialOpcodes[i])
							delete(potentialOpcodes[j], k)
						}
					}
				}
			}
		}
	}
	for k, v := range potentialOpcodes {
		fmt.Println(k, v)
	}
}

func main() {
	regs := []int{0, 0, 0, 0}
	in := []int{0, 0, 0, 0}
	ops := []struct {
		opcode int
		name   string
		fn     func([]int, []int) []int
	}{
		{4, "addr", addr},
		{0, "addi", addi},
		{15, "mulr", mulr},
		{6, "muli", muli},
		{8, "banr", banr},
		{7, "bani", bani},
		{2, "borr", borr},
		{12, "bori", bori},
		{10, "setr", setr},
		{5, "seti", seti},
		{11, "gtir", gtir},
		{3, "gtri", gtri},
		{9, "gtrr", gtrr},
		{14, "eqir", eqir},
		{13, "eqri", eqri},
		{1, "eqrr", eqrr},
	}
	for _, instr := range strings.Split(testProgram, "\n") {
		fmt.Sscanf(instr, "%d  %d %d %d", &in[0], &in[1], &in[2], &in[3])
		opCode := in[0]
		for _, op := range ops {
			if op.opcode == opCode {
				op.fn(regs, in)
				break
			}
		}
	}
	fmt.Println(regs)
}

func addr(regs []int, in []int) []int {
	regs[in[3]] = regs[in[1]] + regs[in[2]]
	return regs
}

func addi(regs []int, in []int) []int {
	regs[in[3]] = regs[in[1]] + in[2]
	return regs
}

func mulr(regs []int, in []int) []int {
	regs[in[3]] = regs[in[1]] * regs[in[2]]
	return regs
}

func muli(regs []int, in []int) []int {
	regs[in[3]] = regs[in[1]] * in[2]
	return regs
}

func banr(regs []int, in []int) []int {
	regs[in[3]] = regs[in[1]] & regs[in[2]]
	return regs
}

func bani(regs []int, in []int) []int {
	regs[in[3]] = regs[in[1]] & in[2]
	return regs
}

func borr(regs []int, in []int) []int {
	regs[in[3]] = regs[in[1]] | regs[in[2]]
	return regs
}

func bori(regs []int, in []int) []int {
	regs[in[3]] = regs[in[1]] | in[2]
	return regs
}

func setr(regs []int, in []int) []int {
	regs[in[3]] = regs[in[1]]
	return regs
}

func seti(regs []int, in []int) []int {
	regs[in[3]] = in[1]
	return regs
}

func gtir(regs []int, in []int) []int {
	if in[1] > regs[in[2]] {
		regs[in[3]] = 1
	} else {
		regs[in[3]] = 0
	}
	return regs
}

func gtri(regs []int, in []int) []int {
	if regs[in[1]] > in[2] {
		regs[in[3]] = 1
	} else {
		regs[in[3]] = 0
	}
	return regs
}

func gtrr(regs []int, in []int) []int {
	if regs[in[1]] > regs[in[2]] {
		regs[in[3]] = 1
	} else {
		regs[in[3]] = 0
	}
	return regs
}

func eqir(regs []int, in []int) []int {
	if in[1] == regs[in[2]] {
		regs[in[3]] = 1
	} else {
		regs[in[3]] = 0
	}
	return regs
}

func eqri(regs []int, in []int) []int {
	if regs[in[1]] == in[2] {
		regs[in[3]] = 1
	} else {
		regs[in[3]] = 0
	}
	return regs
}

func eqrr(regs []int, in []int) []int {
	if regs[in[1]] == regs[in[2]] {
		regs[in[3]] = 1
	} else {
		regs[in[3]] = 0
	}
	return regs
}

const input = `Before: [2, 3, 2, 2]
15 3 2 2
After:  [2, 3, 4, 2]

Before: [3, 2, 2, 1]
3 1 0 1
After:  [3, 1, 2, 1]

Before: [3, 3, 2, 1]
5 3 2 1
After:  [3, 3, 2, 1]

Before: [0, 1, 2, 2]
10 1 0 1
After:  [0, 1, 2, 2]

Before: [0, 1, 2, 1]
8 0 0 3
After:  [0, 1, 2, 0]

Before: [2, 3, 0, 3]
11 0 3 3
After:  [2, 3, 0, 0]

Before: [2, 3, 1, 0]
0 0 2 3
After:  [2, 3, 1, 4]

Before: [2, 0, 1, 1]
7 2 1 2
After:  [2, 0, 1, 1]

Before: [1, 3, 3, 1]
6 0 2 0
After:  [2, 3, 3, 1]

Before: [1, 2, 2, 1]
5 3 2 3
After:  [1, 2, 2, 3]

Before: [1, 0, 1, 2]
13 1 0 3
After:  [1, 0, 1, 1]

Before: [1, 2, 3, 0]
6 1 3 1
After:  [1, 6, 3, 0]

Before: [1, 0, 0, 3]
11 0 3 0
After:  [0, 0, 0, 3]

Before: [0, 3, 2, 1]
5 3 2 2
After:  [0, 3, 3, 1]

Before: [2, 0, 0, 0]
2 3 0 3
After:  [2, 0, 0, 2]

Before: [1, 0, 2, 1]
15 2 2 0
After:  [4, 0, 2, 1]

Before: [0, 1, 2, 3]
4 3 2 3
After:  [0, 1, 2, 5]

Before: [1, 0, 0, 2]
13 1 0 0
After:  [1, 0, 0, 2]

Before: [3, 1, 2, 1]
5 3 2 0
After:  [3, 1, 2, 1]

Before: [1, 1, 3, 0]
12 0 2 2
After:  [1, 1, 3, 0]

Before: [1, 0, 2, 1]
13 1 0 2
After:  [1, 0, 1, 1]

Before: [2, 2, 3, 1]
6 2 3 2
After:  [2, 2, 9, 1]

Before: [2, 1, 2, 3]
4 1 1 0
After:  [2, 1, 2, 3]

Before: [1, 1, 0, 1]
6 3 2 1
After:  [1, 2, 0, 1]

Before: [3, 2, 2, 3]
3 1 0 3
After:  [3, 2, 2, 1]

Before: [0, 1, 2, 2]
8 0 0 3
After:  [0, 1, 2, 0]

Before: [3, 0, 1, 0]
2 1 0 0
After:  [3, 0, 1, 0]

Before: [3, 2, 3, 3]
3 1 0 0
After:  [1, 2, 3, 3]

Before: [0, 1, 0, 3]
12 0 3 2
After:  [0, 1, 3, 3]

Before: [0, 1, 2, 1]
10 1 0 0
After:  [1, 1, 2, 1]

Before: [2, 1, 1, 2]
12 2 3 3
After:  [2, 1, 1, 3]

Before: [1, 0, 2, 2]
7 0 1 3
After:  [1, 0, 2, 1]

Before: [2, 1, 3, 2]
5 3 1 1
After:  [2, 3, 3, 2]

Before: [1, 0, 3, 2]
13 1 0 3
After:  [1, 0, 3, 1]

Before: [3, 2, 2, 0]
3 1 0 2
After:  [3, 2, 1, 0]

Before: [2, 3, 3, 3]
11 0 3 3
After:  [2, 3, 3, 0]

Before: [0, 1, 3, 3]
10 1 0 2
After:  [0, 1, 1, 3]

Before: [1, 1, 3, 3]
4 0 3 1
After:  [1, 4, 3, 3]

Before: [1, 0, 1, 0]
7 0 1 1
After:  [1, 1, 1, 0]

Before: [3, 0, 1, 3]
14 1 2 1
After:  [3, 1, 1, 3]

Before: [1, 1, 1, 2]
4 0 1 3
After:  [1, 1, 1, 2]

Before: [0, 1, 2, 1]
15 2 2 3
After:  [0, 1, 2, 4]

Before: [0, 3, 3, 3]
8 0 0 2
After:  [0, 3, 0, 3]

Before: [0, 0, 1, 2]
1 0 1 0
After:  [1, 0, 1, 2]

Before: [3, 1, 1, 1]
4 3 1 2
After:  [3, 1, 2, 1]

Before: [3, 2, 3, 2]
3 1 0 2
After:  [3, 2, 1, 2]

Before: [1, 1, 3, 2]
12 0 3 3
After:  [1, 1, 3, 3]

Before: [1, 2, 2, 2]
15 1 2 2
After:  [1, 2, 4, 2]

Before: [2, 0, 1, 2]
14 1 2 3
After:  [2, 0, 1, 1]

Before: [1, 0, 2, 3]
7 0 1 3
After:  [1, 0, 2, 1]

Before: [1, 0, 2, 2]
2 2 0 2
After:  [1, 0, 3, 2]

Before: [1, 0, 0, 0]
13 1 0 3
After:  [1, 0, 0, 1]

Before: [1, 0, 1, 0]
7 0 1 0
After:  [1, 0, 1, 0]

Before: [3, 0, 0, 1]
2 2 0 0
After:  [3, 0, 0, 1]

Before: [0, 2, 3, 0]
8 0 0 3
After:  [0, 2, 3, 0]

Before: [0, 1, 1, 0]
4 1 2 0
After:  [2, 1, 1, 0]

Before: [0, 2, 3, 0]
0 1 2 1
After:  [0, 4, 3, 0]

Before: [3, 2, 1, 3]
3 1 0 0
After:  [1, 2, 1, 3]

Before: [0, 2, 3, 0]
2 3 1 3
After:  [0, 2, 3, 2]

Before: [3, 2, 0, 2]
0 1 2 3
After:  [3, 2, 0, 4]

Before: [0, 0, 1, 3]
8 0 0 2
After:  [0, 0, 0, 3]

Before: [1, 0, 1, 0]
7 2 1 0
After:  [1, 0, 1, 0]

Before: [1, 3, 1, 2]
12 2 3 0
After:  [3, 3, 1, 2]

Before: [2, 3, 3, 1]
5 3 2 1
After:  [2, 3, 3, 1]

Before: [2, 3, 1, 0]
0 0 2 1
After:  [2, 4, 1, 0]

Before: [1, 2, 2, 1]
15 2 2 3
After:  [1, 2, 2, 4]

Before: [1, 1, 1, 3]
11 0 3 3
After:  [1, 1, 1, 0]

Before: [3, 2, 3, 1]
5 3 2 0
After:  [3, 2, 3, 1]

Before: [2, 2, 1, 2]
0 3 2 1
After:  [2, 4, 1, 2]

Before: [1, 1, 3, 3]
12 1 3 3
After:  [1, 1, 3, 3]

Before: [3, 0, 3, 2]
4 0 2 3
After:  [3, 0, 3, 6]

Before: [3, 2, 2, 2]
3 1 0 3
After:  [3, 2, 2, 1]

Before: [2, 1, 3, 2]
0 0 2 3
After:  [2, 1, 3, 4]

Before: [0, 3, 1, 1]
8 0 0 2
After:  [0, 3, 0, 1]

Before: [1, 0, 0, 1]
13 1 0 3
After:  [1, 0, 0, 1]

Before: [3, 0, 2, 2]
15 2 2 1
After:  [3, 4, 2, 2]

Before: [0, 0, 1, 1]
7 3 1 0
After:  [1, 0, 1, 1]

Before: [2, 2, 2, 3]
4 1 3 1
After:  [2, 5, 2, 3]

Before: [1, 0, 2, 1]
13 1 0 1
After:  [1, 1, 2, 1]

Before: [0, 0, 1, 1]
14 1 2 3
After:  [0, 0, 1, 1]

Before: [2, 3, 1, 3]
4 1 3 1
After:  [2, 6, 1, 3]

Before: [2, 2, 3, 2]
0 0 2 0
After:  [4, 2, 3, 2]

Before: [1, 0, 3, 3]
11 0 3 2
After:  [1, 0, 0, 3]

Before: [3, 3, 2, 1]
5 3 2 2
After:  [3, 3, 3, 1]

Before: [2, 3, 2, 1]
5 3 2 0
After:  [3, 3, 2, 1]

Before: [3, 2, 3, 1]
3 1 0 1
After:  [3, 1, 3, 1]

Before: [1, 1, 2, 0]
15 2 2 0
After:  [4, 1, 2, 0]

Before: [0, 1, 0, 3]
10 1 0 3
After:  [0, 1, 0, 1]

Before: [3, 0, 1, 2]
0 3 2 2
After:  [3, 0, 4, 2]

Before: [1, 2, 1, 3]
9 1 0 1
After:  [1, 1, 1, 3]

Before: [2, 0, 3, 3]
11 0 3 0
After:  [0, 0, 3, 3]

Before: [0, 0, 1, 3]
1 0 1 1
After:  [0, 1, 1, 3]

Before: [1, 3, 3, 3]
4 1 3 0
After:  [6, 3, 3, 3]

Before: [1, 2, 2, 2]
15 3 2 0
After:  [4, 2, 2, 2]

Before: [2, 0, 0, 1]
14 1 3 0
After:  [1, 0, 0, 1]

Before: [0, 0, 0, 3]
1 0 1 0
After:  [1, 0, 0, 3]

Before: [1, 2, 3, 3]
12 0 3 0
After:  [3, 2, 3, 3]

Before: [1, 0, 2, 3]
13 1 0 1
After:  [1, 1, 2, 3]

Before: [2, 1, 3, 3]
2 1 0 2
After:  [2, 1, 3, 3]

Before: [0, 1, 1, 0]
10 1 0 0
After:  [1, 1, 1, 0]

Before: [1, 0, 2, 1]
7 0 1 3
After:  [1, 0, 2, 1]

Before: [0, 1, 0, 2]
8 0 0 0
After:  [0, 1, 0, 2]

Before: [3, 0, 0, 1]
14 1 3 1
After:  [3, 1, 0, 1]

Before: [2, 1, 1, 1]
12 0 1 3
After:  [2, 1, 1, 3]

Before: [0, 0, 3, 1]
8 0 0 1
After:  [0, 0, 3, 1]

Before: [3, 2, 2, 1]
3 1 0 0
After:  [1, 2, 2, 1]

Before: [2, 1, 2, 3]
11 0 3 3
After:  [2, 1, 2, 0]

Before: [3, 2, 2, 0]
3 1 0 0
After:  [1, 2, 2, 0]

Before: [1, 2, 0, 3]
12 2 3 2
After:  [1, 2, 3, 3]

Before: [1, 2, 1, 2]
12 0 3 0
After:  [3, 2, 1, 2]

Before: [1, 0, 3, 0]
7 0 1 0
After:  [1, 0, 3, 0]

Before: [3, 2, 2, 2]
3 1 0 2
After:  [3, 2, 1, 2]

Before: [0, 1, 3, 1]
5 3 2 2
After:  [0, 1, 3, 1]

Before: [1, 3, 1, 1]
2 2 1 3
After:  [1, 3, 1, 3]

Before: [0, 0, 3, 1]
14 1 3 1
After:  [0, 1, 3, 1]

Before: [3, 2, 1, 1]
3 1 0 3
After:  [3, 2, 1, 1]

Before: [1, 0, 3, 1]
13 1 0 2
After:  [1, 0, 1, 1]

Before: [2, 2, 1, 0]
0 0 2 2
After:  [2, 2, 4, 0]

Before: [0, 0, 1, 1]
14 1 2 0
After:  [1, 0, 1, 1]

Before: [3, 0, 0, 1]
7 3 1 1
After:  [3, 1, 0, 1]

Before: [0, 2, 3, 3]
8 0 0 0
After:  [0, 2, 3, 3]

Before: [3, 2, 2, 0]
15 1 2 0
After:  [4, 2, 2, 0]

Before: [3, 2, 0, 1]
3 1 0 3
After:  [3, 2, 0, 1]

Before: [3, 0, 1, 3]
14 1 2 0
After:  [1, 0, 1, 3]

Before: [2, 0, 2, 3]
11 0 3 2
After:  [2, 0, 0, 3]

Before: [2, 1, 1, 2]
0 3 2 2
After:  [2, 1, 4, 2]

Before: [2, 0, 1, 2]
14 1 2 0
After:  [1, 0, 1, 2]

Before: [1, 2, 0, 0]
9 1 0 2
After:  [1, 2, 1, 0]

Before: [1, 2, 3, 1]
9 1 0 0
After:  [1, 2, 3, 1]

Before: [0, 3, 3, 2]
2 0 1 3
After:  [0, 3, 3, 3]

Before: [1, 3, 1, 3]
11 0 3 0
After:  [0, 3, 1, 3]

Before: [3, 2, 3, 0]
3 1 0 2
After:  [3, 2, 1, 0]

Before: [2, 3, 2, 3]
4 1 3 1
After:  [2, 6, 2, 3]

Before: [2, 1, 3, 3]
12 0 1 2
After:  [2, 1, 3, 3]

Before: [1, 0, 1, 2]
13 1 0 1
After:  [1, 1, 1, 2]

Before: [0, 0, 0, 1]
8 0 0 2
After:  [0, 0, 0, 1]

Before: [1, 0, 3, 0]
13 1 0 0
After:  [1, 0, 3, 0]

Before: [2, 1, 2, 1]
5 3 2 1
After:  [2, 3, 2, 1]

Before: [1, 0, 3, 2]
7 0 1 3
After:  [1, 0, 3, 1]

Before: [0, 3, 2, 3]
8 0 0 1
After:  [0, 0, 2, 3]

Before: [0, 0, 1, 1]
1 0 1 0
After:  [1, 0, 1, 1]

Before: [0, 0, 1, 0]
1 0 1 0
After:  [1, 0, 1, 0]

Before: [3, 2, 3, 3]
3 1 0 2
After:  [3, 2, 1, 3]

Before: [2, 1, 3, 3]
4 1 1 1
After:  [2, 2, 3, 3]

Before: [3, 1, 3, 2]
6 0 3 3
After:  [3, 1, 3, 9]

Before: [0, 1, 0, 1]
10 1 0 0
After:  [1, 1, 0, 1]

Before: [2, 1, 0, 2]
5 3 1 1
After:  [2, 3, 0, 2]

Before: [0, 3, 0, 1]
8 0 0 3
After:  [0, 3, 0, 0]

Before: [0, 1, 3, 3]
10 1 0 1
After:  [0, 1, 3, 3]

Before: [3, 2, 2, 3]
3 1 0 2
After:  [3, 2, 1, 3]

Before: [2, 1, 0, 1]
2 1 0 3
After:  [2, 1, 0, 3]

Before: [1, 2, 3, 3]
9 1 0 1
After:  [1, 1, 3, 3]

Before: [0, 1, 1, 1]
10 1 0 0
After:  [1, 1, 1, 1]

Before: [3, 0, 2, 3]
12 1 3 0
After:  [3, 0, 2, 3]

Before: [0, 1, 0, 3]
10 1 0 2
After:  [0, 1, 1, 3]

Before: [1, 0, 1, 0]
13 1 0 2
After:  [1, 0, 1, 0]

Before: [0, 0, 2, 1]
8 0 0 0
After:  [0, 0, 2, 1]

Before: [1, 1, 2, 3]
11 0 3 0
After:  [0, 1, 2, 3]

Before: [1, 1, 0, 2]
5 3 1 2
After:  [1, 1, 3, 2]

Before: [0, 3, 2, 1]
5 3 2 1
After:  [0, 3, 2, 1]

Before: [2, 2, 1, 3]
4 0 3 2
After:  [2, 2, 5, 3]

Before: [0, 1, 3, 1]
10 1 0 3
After:  [0, 1, 3, 1]

Before: [0, 1, 0, 2]
10 1 0 3
After:  [0, 1, 0, 1]

Before: [2, 0, 2, 1]
14 1 3 2
After:  [2, 0, 1, 1]

Before: [2, 1, 0, 3]
11 0 3 0
After:  [0, 1, 0, 3]

Before: [2, 2, 2, 1]
15 2 2 0
After:  [4, 2, 2, 1]

Before: [0, 2, 1, 2]
0 1 2 0
After:  [4, 2, 1, 2]

Before: [1, 0, 2, 1]
14 1 3 2
After:  [1, 0, 1, 1]

Before: [1, 0, 2, 1]
13 1 0 0
After:  [1, 0, 2, 1]

Before: [2, 3, 1, 3]
2 0 2 1
After:  [2, 3, 1, 3]

Before: [0, 3, 0, 2]
2 0 3 2
After:  [0, 3, 2, 2]

Before: [1, 2, 1, 3]
12 2 3 1
After:  [1, 3, 1, 3]

Before: [1, 0, 2, 1]
13 1 0 3
After:  [1, 0, 2, 1]

Before: [1, 0, 2, 2]
7 0 1 1
After:  [1, 1, 2, 2]

Before: [1, 0, 1, 1]
13 1 0 1
After:  [1, 1, 1, 1]

Before: [1, 1, 2, 1]
12 0 2 3
After:  [1, 1, 2, 3]

Before: [1, 1, 0, 3]
4 3 3 1
After:  [1, 6, 0, 3]

Before: [1, 3, 3, 3]
11 0 3 0
After:  [0, 3, 3, 3]

Before: [0, 2, 2, 1]
15 2 2 1
After:  [0, 4, 2, 1]

Before: [2, 1, 2, 3]
11 0 3 0
After:  [0, 1, 2, 3]

Before: [2, 2, 1, 3]
4 3 3 3
After:  [2, 2, 1, 6]

Before: [3, 2, 2, 2]
15 2 2 0
After:  [4, 2, 2, 2]

Before: [1, 1, 3, 3]
4 2 3 0
After:  [6, 1, 3, 3]

Before: [0, 2, 3, 1]
0 1 2 0
After:  [4, 2, 3, 1]

Before: [0, 2, 0, 2]
0 3 2 2
After:  [0, 2, 4, 2]

Before: [1, 0, 0, 2]
13 1 0 3
After:  [1, 0, 0, 1]

Before: [3, 1, 1, 2]
5 3 1 0
After:  [3, 1, 1, 2]

Before: [2, 2, 1, 1]
2 2 0 2
After:  [2, 2, 3, 1]

Before: [1, 3, 0, 3]
11 0 3 3
After:  [1, 3, 0, 0]

Before: [1, 1, 3, 2]
5 3 1 2
After:  [1, 1, 3, 2]

Before: [0, 2, 0, 1]
2 0 3 1
After:  [0, 1, 0, 1]

Before: [0, 0, 3, 0]
2 0 2 1
After:  [0, 3, 3, 0]

Before: [3, 0, 3, 1]
14 1 3 1
After:  [3, 1, 3, 1]

Before: [0, 3, 2, 0]
8 0 0 0
After:  [0, 3, 2, 0]

Before: [0, 3, 2, 3]
8 0 0 2
After:  [0, 3, 0, 3]

Before: [0, 1, 3, 3]
10 1 0 3
After:  [0, 1, 3, 1]

Before: [0, 0, 0, 0]
1 0 1 3
After:  [0, 0, 0, 1]

Before: [0, 1, 3, 1]
10 1 0 1
After:  [0, 1, 3, 1]

Before: [0, 0, 2, 1]
14 1 3 0
After:  [1, 0, 2, 1]

Before: [2, 2, 1, 0]
2 1 2 1
After:  [2, 3, 1, 0]

Before: [1, 0, 0, 3]
11 0 3 1
After:  [1, 0, 0, 3]

Before: [0, 1, 3, 0]
10 1 0 1
After:  [0, 1, 3, 0]

Before: [3, 3, 1, 1]
2 2 1 0
After:  [3, 3, 1, 1]

Before: [0, 2, 1, 0]
2 2 1 0
After:  [3, 2, 1, 0]

Before: [2, 1, 2, 3]
12 1 3 1
After:  [2, 3, 2, 3]

Before: [0, 0, 1, 2]
8 0 0 0
After:  [0, 0, 1, 2]

Before: [2, 3, 2, 0]
6 2 3 1
After:  [2, 6, 2, 0]

Before: [3, 0, 2, 1]
5 3 2 1
After:  [3, 3, 2, 1]

Before: [2, 1, 1, 3]
2 2 0 1
After:  [2, 3, 1, 3]

Before: [2, 2, 0, 0]
2 3 0 2
After:  [2, 2, 2, 0]

Before: [3, 2, 3, 0]
3 1 0 3
After:  [3, 2, 3, 1]

Before: [0, 1, 3, 2]
10 1 0 2
After:  [0, 1, 1, 2]

Before: [1, 2, 3, 0]
2 0 1 2
After:  [1, 2, 3, 0]

Before: [1, 3, 2, 3]
11 0 3 2
After:  [1, 3, 0, 3]

Before: [0, 0, 2, 1]
15 2 2 1
After:  [0, 4, 2, 1]

Before: [3, 2, 0, 0]
3 1 0 1
After:  [3, 1, 0, 0]

Before: [1, 0, 0, 1]
14 1 3 1
After:  [1, 1, 0, 1]

Before: [2, 0, 1, 3]
7 2 1 2
After:  [2, 0, 1, 3]

Before: [1, 0, 0, 0]
13 1 0 2
After:  [1, 0, 1, 0]

Before: [1, 0, 1, 1]
13 1 0 2
After:  [1, 0, 1, 1]

Before: [1, 0, 2, 0]
13 1 0 2
After:  [1, 0, 1, 0]

Before: [3, 1, 1, 2]
12 1 3 3
After:  [3, 1, 1, 3]

Before: [0, 1, 1, 0]
12 0 1 0
After:  [1, 1, 1, 0]

Before: [1, 0, 3, 2]
13 1 0 0
After:  [1, 0, 3, 2]

Before: [2, 2, 1, 3]
11 0 3 1
After:  [2, 0, 1, 3]

Before: [0, 0, 2, 3]
1 0 1 0
After:  [1, 0, 2, 3]

Before: [1, 2, 0, 3]
11 0 3 2
After:  [1, 2, 0, 3]

Before: [1, 0, 3, 1]
7 3 1 0
After:  [1, 0, 3, 1]

Before: [1, 0, 0, 1]
13 1 0 1
After:  [1, 1, 0, 1]

Before: [1, 2, 3, 0]
9 1 0 2
After:  [1, 2, 1, 0]

Before: [1, 1, 3, 3]
11 0 3 2
After:  [1, 1, 0, 3]

Before: [2, 3, 2, 3]
11 0 3 1
After:  [2, 0, 2, 3]

Before: [0, 1, 2, 2]
10 1 0 2
After:  [0, 1, 1, 2]

Before: [2, 0, 1, 2]
2 0 2 2
After:  [2, 0, 3, 2]

Before: [1, 2, 1, 1]
9 1 0 1
After:  [1, 1, 1, 1]

Before: [0, 0, 2, 0]
12 1 2 0
After:  [2, 0, 2, 0]

Before: [3, 1, 1, 1]
6 0 3 3
After:  [3, 1, 1, 9]

Before: [1, 1, 2, 0]
2 2 0 2
After:  [1, 1, 3, 0]

Before: [0, 0, 1, 3]
8 0 0 1
After:  [0, 0, 1, 3]

Before: [2, 1, 2, 2]
15 3 2 0
After:  [4, 1, 2, 2]

Before: [3, 1, 0, 2]
12 1 3 1
After:  [3, 3, 0, 2]

Before: [0, 0, 1, 1]
14 1 2 1
After:  [0, 1, 1, 1]

Before: [1, 2, 3, 3]
2 1 0 2
After:  [1, 2, 3, 3]

Before: [0, 3, 2, 3]
15 2 2 0
After:  [4, 3, 2, 3]

Before: [2, 0, 0, 2]
2 1 3 0
After:  [2, 0, 0, 2]

Before: [2, 2, 1, 0]
6 0 3 2
After:  [2, 2, 6, 0]

Before: [0, 1, 1, 2]
5 3 1 0
After:  [3, 1, 1, 2]

Before: [2, 3, 3, 3]
4 1 3 3
After:  [2, 3, 3, 6]

Before: [0, 1, 1, 0]
10 1 0 3
After:  [0, 1, 1, 1]

Before: [0, 0, 2, 1]
14 1 3 2
After:  [0, 0, 1, 1]

Before: [2, 0, 1, 2]
0 0 2 0
After:  [4, 0, 1, 2]

Before: [3, 2, 1, 0]
3 1 0 0
After:  [1, 2, 1, 0]

Before: [0, 1, 2, 0]
10 1 0 1
After:  [0, 1, 2, 0]

Before: [1, 1, 3, 3]
12 1 2 1
After:  [1, 3, 3, 3]

Before: [2, 0, 2, 0]
15 2 2 3
After:  [2, 0, 2, 4]

Before: [1, 2, 3, 1]
5 3 2 2
After:  [1, 2, 3, 1]

Before: [1, 0, 3, 1]
14 1 3 2
After:  [1, 0, 1, 1]

Before: [1, 2, 3, 2]
0 3 2 0
After:  [4, 2, 3, 2]

Before: [0, 1, 2, 2]
10 1 0 3
After:  [0, 1, 2, 1]

Before: [1, 2, 0, 2]
9 1 0 3
After:  [1, 2, 0, 1]

Before: [3, 2, 2, 2]
3 1 0 1
After:  [3, 1, 2, 2]

Before: [3, 3, 2, 2]
15 3 2 1
After:  [3, 4, 2, 2]

Before: [2, 0, 2, 3]
11 0 3 0
After:  [0, 0, 2, 3]

Before: [0, 0, 2, 3]
1 0 1 3
After:  [0, 0, 2, 1]

Before: [2, 0, 2, 1]
7 3 1 2
After:  [2, 0, 1, 1]

Before: [0, 0, 2, 0]
1 0 1 2
After:  [0, 0, 1, 0]

Before: [2, 2, 0, 3]
11 0 3 1
After:  [2, 0, 0, 3]

Before: [1, 0, 1, 0]
14 1 2 3
After:  [1, 0, 1, 1]

Before: [0, 0, 1, 3]
7 2 1 3
After:  [0, 0, 1, 1]

Before: [3, 2, 1, 3]
4 1 3 1
After:  [3, 5, 1, 3]

Before: [1, 0, 2, 3]
11 0 3 0
After:  [0, 0, 2, 3]

Before: [1, 0, 0, 0]
13 1 0 1
After:  [1, 1, 0, 0]

Before: [1, 2, 0, 3]
0 1 2 0
After:  [4, 2, 0, 3]

Before: [3, 3, 3, 0]
6 2 3 1
After:  [3, 9, 3, 0]

Before: [1, 1, 2, 1]
15 2 2 2
After:  [1, 1, 4, 1]

Before: [2, 1, 0, 0]
6 1 2 3
After:  [2, 1, 0, 2]

Before: [2, 0, 0, 3]
11 0 3 1
After:  [2, 0, 0, 3]

Before: [0, 3, 1, 3]
12 2 3 0
After:  [3, 3, 1, 3]

Before: [2, 0, 1, 1]
7 2 1 1
After:  [2, 1, 1, 1]

Before: [1, 2, 2, 3]
9 1 0 3
After:  [1, 2, 2, 1]

Before: [3, 0, 2, 0]
6 2 3 3
After:  [3, 0, 2, 6]

Before: [3, 0, 0, 2]
2 2 0 0
After:  [3, 0, 0, 2]

Before: [3, 2, 0, 2]
3 1 0 2
After:  [3, 2, 1, 2]

Before: [0, 1, 1, 0]
10 1 0 2
After:  [0, 1, 1, 0]

Before: [0, 0, 0, 2]
1 0 1 0
After:  [1, 0, 0, 2]

Before: [0, 2, 1, 1]
2 2 1 1
After:  [0, 3, 1, 1]

Before: [0, 0, 3, 1]
8 0 0 3
After:  [0, 0, 3, 0]

Before: [0, 1, 2, 0]
10 1 0 3
After:  [0, 1, 2, 1]

Before: [3, 2, 1, 0]
3 1 0 1
After:  [3, 1, 1, 0]

Before: [3, 3, 0, 1]
6 0 3 1
After:  [3, 9, 0, 1]

Before: [0, 1, 3, 1]
2 0 2 0
After:  [3, 1, 3, 1]

Before: [2, 3, 1, 1]
6 1 3 2
After:  [2, 3, 9, 1]

Before: [1, 0, 3, 3]
2 1 2 3
After:  [1, 0, 3, 3]

Before: [1, 1, 2, 3]
11 0 3 3
After:  [1, 1, 2, 0]

Before: [3, 3, 3, 2]
4 2 2 0
After:  [6, 3, 3, 2]

Before: [3, 1, 0, 0]
2 1 0 1
After:  [3, 3, 0, 0]

Before: [2, 0, 2, 2]
15 2 2 1
After:  [2, 4, 2, 2]

Before: [0, 1, 3, 2]
8 0 0 2
After:  [0, 1, 0, 2]

Before: [1, 2, 2, 0]
9 1 0 3
After:  [1, 2, 2, 1]

Before: [1, 3, 2, 1]
6 1 3 0
After:  [9, 3, 2, 1]

Before: [0, 1, 0, 1]
10 1 0 2
After:  [0, 1, 1, 1]

Before: [3, 0, 1, 1]
7 3 1 1
After:  [3, 1, 1, 1]

Before: [0, 3, 1, 0]
8 0 0 3
After:  [0, 3, 1, 0]

Before: [0, 1, 2, 1]
10 1 0 3
After:  [0, 1, 2, 1]

Before: [0, 1, 1, 3]
10 1 0 3
After:  [0, 1, 1, 1]

Before: [3, 0, 1, 3]
7 2 1 1
After:  [3, 1, 1, 3]

Before: [3, 2, 1, 3]
3 1 0 1
After:  [3, 1, 1, 3]

Before: [0, 1, 0, 0]
10 1 0 3
After:  [0, 1, 0, 1]

Before: [1, 3, 2, 1]
5 3 2 2
After:  [1, 3, 3, 1]

Before: [2, 2, 2, 1]
5 3 2 2
After:  [2, 2, 3, 1]

Before: [2, 0, 2, 1]
15 2 2 1
After:  [2, 4, 2, 1]

Before: [3, 2, 2, 2]
6 0 3 1
After:  [3, 9, 2, 2]

Before: [2, 3, 3, 3]
4 2 2 2
After:  [2, 3, 6, 3]

Before: [0, 0, 3, 2]
1 0 1 1
After:  [0, 1, 3, 2]

Before: [3, 3, 3, 2]
0 3 2 0
After:  [4, 3, 3, 2]

Before: [0, 1, 2, 3]
8 0 0 2
After:  [0, 1, 0, 3]

Before: [1, 2, 1, 3]
0 1 2 0
After:  [4, 2, 1, 3]

Before: [0, 0, 1, 0]
8 0 0 1
After:  [0, 0, 1, 0]

Before: [0, 0, 2, 1]
1 0 1 2
After:  [0, 0, 1, 1]

Before: [1, 0, 1, 1]
13 1 0 0
After:  [1, 0, 1, 1]

Before: [0, 3, 2, 0]
8 0 0 2
After:  [0, 3, 0, 0]

Before: [2, 0, 1, 0]
7 2 1 2
After:  [2, 0, 1, 0]

Before: [0, 1, 3, 0]
10 1 0 0
After:  [1, 1, 3, 0]

Before: [2, 0, 1, 3]
14 1 2 1
After:  [2, 1, 1, 3]

Before: [0, 0, 1, 3]
1 0 1 0
After:  [1, 0, 1, 3]

Before: [0, 1, 3, 1]
10 1 0 2
After:  [0, 1, 1, 1]

Before: [0, 0, 1, 3]
14 1 2 1
After:  [0, 1, 1, 3]

Before: [0, 0, 0, 0]
1 0 1 2
After:  [0, 0, 1, 0]

Before: [0, 1, 3, 2]
10 1 0 1
After:  [0, 1, 3, 2]

Before: [3, 3, 3, 1]
5 3 2 3
After:  [3, 3, 3, 3]

Before: [0, 0, 0, 0]
8 0 0 3
After:  [0, 0, 0, 0]

Before: [0, 1, 2, 3]
8 0 0 3
After:  [0, 1, 2, 0]

Before: [1, 3, 1, 2]
0 3 2 3
After:  [1, 3, 1, 4]

Before: [0, 0, 1, 1]
8 0 0 0
After:  [0, 0, 1, 1]

Before: [0, 0, 3, 1]
1 0 1 1
After:  [0, 1, 3, 1]

Before: [2, 3, 1, 1]
4 3 2 3
After:  [2, 3, 1, 2]

Before: [1, 2, 2, 0]
9 1 0 0
After:  [1, 2, 2, 0]

Before: [3, 3, 0, 0]
6 0 3 2
After:  [3, 3, 9, 0]

Before: [1, 0, 1, 2]
14 1 2 1
After:  [1, 1, 1, 2]

Before: [3, 2, 3, 1]
3 1 0 0
After:  [1, 2, 3, 1]

Before: [0, 1, 0, 0]
10 1 0 2
After:  [0, 1, 1, 0]

Before: [1, 2, 3, 2]
9 1 0 1
After:  [1, 1, 3, 2]

Before: [0, 1, 1, 3]
4 1 3 2
After:  [0, 1, 4, 3]

Before: [3, 0, 1, 2]
14 1 2 0
After:  [1, 0, 1, 2]

Before: [3, 2, 1, 1]
3 1 0 2
After:  [3, 2, 1, 1]

Before: [1, 2, 3, 3]
0 1 2 0
After:  [4, 2, 3, 3]

Before: [3, 1, 1, 3]
12 2 3 0
After:  [3, 1, 1, 3]

Before: [3, 0, 1, 0]
7 2 1 1
After:  [3, 1, 1, 0]

Before: [1, 0, 3, 2]
0 3 2 0
After:  [4, 0, 3, 2]

Before: [0, 2, 3, 3]
0 1 2 2
After:  [0, 2, 4, 3]

Before: [0, 1, 3, 1]
4 2 1 2
After:  [0, 1, 4, 1]

Before: [1, 0, 3, 3]
13 1 0 3
After:  [1, 0, 3, 1]

Before: [3, 1, 3, 0]
4 1 1 1
After:  [3, 2, 3, 0]

Before: [1, 2, 0, 3]
9 1 0 2
After:  [1, 2, 1, 3]

Before: [0, 2, 2, 3]
4 3 2 1
After:  [0, 5, 2, 3]

Before: [1, 1, 1, 1]
4 0 2 3
After:  [1, 1, 1, 2]

Before: [1, 1, 2, 3]
11 0 3 1
After:  [1, 0, 2, 3]

Before: [2, 0, 1, 1]
7 3 1 1
After:  [2, 1, 1, 1]

Before: [2, 3, 0, 0]
2 2 1 3
After:  [2, 3, 0, 3]

Before: [1, 0, 1, 1]
7 0 1 0
After:  [1, 0, 1, 1]

Before: [1, 0, 0, 0]
7 0 1 2
After:  [1, 0, 1, 0]

Before: [1, 0, 3, 0]
12 0 2 1
After:  [1, 3, 3, 0]

Before: [1, 2, 3, 2]
9 1 0 2
After:  [1, 2, 1, 2]

Before: [1, 2, 3, 0]
9 1 0 0
After:  [1, 2, 3, 0]

Before: [2, 2, 1, 3]
0 0 2 2
After:  [2, 2, 4, 3]

Before: [2, 1, 1, 3]
6 3 2 0
After:  [6, 1, 1, 3]

Before: [2, 1, 2, 0]
15 2 2 3
After:  [2, 1, 2, 4]

Before: [0, 0, 0, 0]
1 0 1 0
After:  [1, 0, 0, 0]

Before: [1, 0, 2, 1]
12 1 2 3
After:  [1, 0, 2, 2]

Before: [0, 0, 1, 3]
14 1 2 3
After:  [0, 0, 1, 1]

Before: [1, 1, 0, 3]
11 0 3 1
After:  [1, 0, 0, 3]

Before: [0, 2, 2, 0]
8 0 0 1
After:  [0, 0, 2, 0]

Before: [3, 1, 2, 0]
6 2 3 3
After:  [3, 1, 2, 6]

Before: [1, 2, 0, 0]
2 3 1 2
After:  [1, 2, 2, 0]

Before: [2, 1, 0, 3]
4 3 3 2
After:  [2, 1, 6, 3]

Before: [3, 2, 0, 0]
3 1 0 0
After:  [1, 2, 0, 0]

Before: [3, 0, 1, 1]
14 1 3 1
After:  [3, 1, 1, 1]

Before: [3, 1, 1, 1]
6 0 3 2
After:  [3, 1, 9, 1]

Before: [0, 0, 1, 1]
1 0 1 1
After:  [0, 1, 1, 1]

Before: [2, 1, 3, 3]
0 0 2 2
After:  [2, 1, 4, 3]

Before: [0, 2, 2, 1]
8 0 0 2
After:  [0, 2, 0, 1]

Before: [1, 0, 2, 2]
15 3 2 3
After:  [1, 0, 2, 4]

Before: [0, 1, 3, 1]
2 0 3 0
After:  [1, 1, 3, 1]

Before: [2, 0, 3, 1]
5 3 2 2
After:  [2, 0, 3, 1]

Before: [3, 2, 0, 3]
3 1 0 1
After:  [3, 1, 0, 3]

Before: [1, 3, 3, 3]
11 0 3 2
After:  [1, 3, 0, 3]

Before: [0, 1, 1, 2]
10 1 0 0
After:  [1, 1, 1, 2]

Before: [3, 2, 1, 2]
2 1 2 3
After:  [3, 2, 1, 3]

Before: [1, 0, 0, 1]
13 1 0 2
After:  [1, 0, 1, 1]

Before: [1, 0, 2, 3]
13 1 0 3
After:  [1, 0, 2, 1]

Before: [0, 0, 0, 2]
1 0 1 2
After:  [0, 0, 1, 2]

Before: [2, 1, 0, 2]
5 3 1 2
After:  [2, 1, 3, 2]

Before: [1, 0, 0, 3]
13 1 0 2
After:  [1, 0, 1, 3]

Before: [2, 1, 3, 2]
5 3 1 2
After:  [2, 1, 3, 2]

Before: [1, 2, 3, 0]
2 3 1 3
After:  [1, 2, 3, 2]

Before: [1, 2, 2, 1]
9 1 0 0
After:  [1, 2, 2, 1]

Before: [0, 0, 3, 3]
1 0 1 1
After:  [0, 1, 3, 3]

Before: [3, 2, 2, 3]
4 1 3 0
After:  [5, 2, 2, 3]

Before: [3, 2, 1, 2]
3 1 0 0
After:  [1, 2, 1, 2]

Before: [2, 1, 3, 1]
6 3 2 3
After:  [2, 1, 3, 2]

Before: [3, 1, 2, 3]
15 2 2 2
After:  [3, 1, 4, 3]

Before: [2, 0, 3, 3]
11 0 3 1
After:  [2, 0, 3, 3]

Before: [0, 0, 2, 2]
1 0 1 0
After:  [1, 0, 2, 2]

Before: [3, 2, 2, 1]
15 2 2 3
After:  [3, 2, 2, 4]

Before: [3, 3, 1, 2]
6 0 3 3
After:  [3, 3, 1, 9]

Before: [2, 1, 0, 2]
12 0 1 3
After:  [2, 1, 0, 3]

Before: [2, 1, 1, 2]
12 0 1 3
After:  [2, 1, 1, 3]

Before: [0, 2, 1, 3]
12 2 3 1
After:  [0, 3, 1, 3]

Before: [3, 1, 1, 3]
4 1 3 3
After:  [3, 1, 1, 4]

Before: [0, 0, 3, 0]
1 0 1 0
After:  [1, 0, 3, 0]

Before: [1, 0, 0, 3]
13 1 0 1
After:  [1, 1, 0, 3]

Before: [3, 0, 1, 0]
14 1 2 0
After:  [1, 0, 1, 0]

Before: [3, 2, 0, 1]
3 1 0 1
After:  [3, 1, 0, 1]

Before: [3, 1, 3, 2]
12 1 2 1
After:  [3, 3, 3, 2]

Before: [2, 0, 0, 1]
7 3 1 3
After:  [2, 0, 0, 1]

Before: [3, 3, 3, 1]
5 3 2 1
After:  [3, 3, 3, 1]

Before: [0, 0, 3, 0]
8 0 0 2
After:  [0, 0, 0, 0]

Before: [1, 3, 1, 1]
6 1 2 2
After:  [1, 3, 6, 1]

Before: [1, 2, 2, 0]
9 1 0 2
After:  [1, 2, 1, 0]

Before: [0, 2, 3, 1]
6 1 3 2
After:  [0, 2, 6, 1]

Before: [0, 0, 1, 2]
14 1 2 0
After:  [1, 0, 1, 2]

Before: [2, 1, 3, 1]
0 0 2 0
After:  [4, 1, 3, 1]

Before: [1, 2, 0, 3]
9 1 0 0
After:  [1, 2, 0, 3]

Before: [1, 3, 0, 3]
11 0 3 2
After:  [1, 3, 0, 3]

Before: [2, 2, 0, 3]
11 0 3 2
After:  [2, 2, 0, 3]

Before: [3, 2, 3, 1]
3 1 0 2
After:  [3, 2, 1, 1]

Before: [1, 0, 1, 3]
13 1 0 3
After:  [1, 0, 1, 1]

Before: [1, 0, 1, 0]
13 1 0 1
After:  [1, 1, 1, 0]

Before: [1, 3, 1, 3]
11 0 3 3
After:  [1, 3, 1, 0]

Before: [0, 1, 1, 2]
8 0 0 1
After:  [0, 0, 1, 2]

Before: [2, 2, 0, 2]
0 3 2 3
After:  [2, 2, 0, 4]

Before: [0, 1, 1, 2]
8 0 0 2
After:  [0, 1, 0, 2]

Before: [1, 2, 3, 3]
2 1 0 1
After:  [1, 3, 3, 3]

Before: [1, 0, 1, 1]
7 0 1 3
After:  [1, 0, 1, 1]

Before: [2, 1, 3, 1]
5 3 2 1
After:  [2, 3, 3, 1]

Before: [0, 2, 3, 3]
8 0 0 2
After:  [0, 2, 0, 3]

Before: [3, 0, 0, 3]
2 2 0 3
After:  [3, 0, 0, 3]

Before: [1, 0, 3, 3]
12 0 2 1
After:  [1, 3, 3, 3]

Before: [3, 2, 0, 2]
3 1 0 1
After:  [3, 1, 0, 2]

Before: [3, 1, 2, 1]
4 0 2 0
After:  [5, 1, 2, 1]

Before: [0, 1, 0, 3]
10 1 0 0
After:  [1, 1, 0, 3]

Before: [1, 2, 0, 1]
2 0 1 2
After:  [1, 2, 3, 1]

Before: [1, 3, 0, 3]
11 0 3 0
After:  [0, 3, 0, 3]

Before: [1, 0, 0, 2]
0 3 2 2
After:  [1, 0, 4, 2]

Before: [0, 1, 0, 1]
10 1 0 3
After:  [0, 1, 0, 1]

Before: [1, 0, 1, 3]
13 1 0 1
After:  [1, 1, 1, 3]

Before: [1, 2, 0, 1]
9 1 0 1
After:  [1, 1, 0, 1]

Before: [2, 2, 3, 2]
0 1 2 2
After:  [2, 2, 4, 2]

Before: [0, 0, 1, 3]
14 1 2 2
After:  [0, 0, 1, 3]

Before: [0, 1, 0, 2]
8 0 0 2
After:  [0, 1, 0, 2]

Before: [2, 2, 3, 1]
5 3 2 0
After:  [3, 2, 3, 1]

Before: [1, 0, 0, 3]
13 1 0 3
After:  [1, 0, 0, 1]

Before: [1, 1, 0, 3]
11 0 3 2
After:  [1, 1, 0, 3]

Before: [0, 0, 3, 2]
2 0 2 2
After:  [0, 0, 3, 2]

Before: [1, 2, 2, 2]
9 1 0 0
After:  [1, 2, 2, 2]

Before: [3, 2, 3, 1]
3 1 0 3
After:  [3, 2, 3, 1]

Before: [1, 2, 3, 1]
9 1 0 2
After:  [1, 2, 1, 1]

Before: [0, 1, 0, 1]
8 0 0 3
After:  [0, 1, 0, 0]

Before: [2, 2, 2, 3]
15 2 2 0
After:  [4, 2, 2, 3]

Before: [0, 1, 2, 2]
10 1 0 0
After:  [1, 1, 2, 2]

Before: [1, 1, 2, 1]
5 3 2 3
After:  [1, 1, 2, 3]

Before: [2, 0, 0, 1]
14 1 3 1
After:  [2, 1, 0, 1]

Before: [3, 0, 2, 0]
12 3 2 0
After:  [2, 0, 2, 0]

Before: [0, 0, 3, 3]
12 1 3 3
After:  [0, 0, 3, 3]

Before: [3, 0, 1, 1]
14 1 3 0
After:  [1, 0, 1, 1]

Before: [1, 1, 3, 1]
4 2 1 2
After:  [1, 1, 4, 1]

Before: [0, 2, 0, 0]
0 1 2 0
After:  [4, 2, 0, 0]

Before: [0, 2, 2, 1]
5 3 2 0
After:  [3, 2, 2, 1]

Before: [3, 2, 0, 0]
3 1 0 3
After:  [3, 2, 0, 1]

Before: [3, 2, 0, 2]
3 1 0 0
After:  [1, 2, 0, 2]

Before: [0, 2, 3, 3]
0 1 2 1
After:  [0, 4, 3, 3]

Before: [2, 3, 1, 3]
2 0 2 0
After:  [3, 3, 1, 3]

Before: [1, 2, 0, 2]
9 1 0 0
After:  [1, 2, 0, 2]

Before: [0, 1, 0, 0]
10 1 0 0
After:  [1, 1, 0, 0]

Before: [2, 0, 3, 1]
7 3 1 1
After:  [2, 1, 3, 1]

Before: [3, 2, 3, 2]
3 1 0 1
After:  [3, 1, 3, 2]

Before: [1, 2, 3, 3]
4 3 2 1
After:  [1, 6, 3, 3]

Before: [0, 3, 2, 0]
15 2 2 2
After:  [0, 3, 4, 0]

Before: [0, 3, 3, 2]
0 3 2 3
After:  [0, 3, 3, 4]

Before: [0, 1, 2, 0]
10 1 0 2
After:  [0, 1, 1, 0]

Before: [3, 2, 0, 0]
6 1 3 1
After:  [3, 6, 0, 0]

Before: [0, 1, 1, 3]
10 1 0 0
After:  [1, 1, 1, 3]

Before: [3, 0, 3, 3]
4 3 2 1
After:  [3, 6, 3, 3]

Before: [1, 0, 0, 3]
11 0 3 3
After:  [1, 0, 0, 0]

Before: [1, 0, 3, 2]
13 1 0 1
After:  [1, 1, 3, 2]

Before: [1, 0, 3, 2]
6 2 3 1
After:  [1, 9, 3, 2]

Before: [1, 0, 3, 0]
6 0 2 2
After:  [1, 0, 2, 0]

Before: [1, 2, 1, 0]
9 1 0 3
After:  [1, 2, 1, 1]

Before: [0, 0, 3, 2]
1 0 1 3
After:  [0, 0, 3, 1]

Before: [1, 0, 1, 3]
13 1 0 2
After:  [1, 0, 1, 3]

Before: [0, 2, 2, 2]
15 1 2 3
After:  [0, 2, 2, 4]

Before: [0, 1, 3, 2]
10 1 0 0
After:  [1, 1, 3, 2]

Before: [1, 0, 3, 0]
13 1 0 3
After:  [1, 0, 3, 1]

Before: [0, 1, 1, 2]
10 1 0 3
After:  [0, 1, 1, 1]

Before: [1, 2, 1, 3]
9 1 0 2
After:  [1, 2, 1, 3]

Before: [1, 3, 2, 2]
6 3 3 3
After:  [1, 3, 2, 6]

Before: [2, 1, 2, 3]
15 2 2 1
After:  [2, 4, 2, 3]

Before: [3, 2, 3, 1]
6 2 3 3
After:  [3, 2, 3, 9]

Before: [3, 0, 1, 1]
14 1 2 3
After:  [3, 0, 1, 1]

Before: [1, 1, 3, 2]
12 0 3 2
After:  [1, 1, 3, 2]

Before: [1, 2, 2, 0]
15 2 2 2
After:  [1, 2, 4, 0]

Before: [0, 1, 3, 0]
8 0 0 0
After:  [0, 1, 3, 0]

Before: [0, 0, 0, 3]
1 0 1 1
After:  [0, 1, 0, 3]

Before: [0, 0, 0, 3]
1 0 1 2
After:  [0, 0, 1, 3]

Before: [3, 2, 0, 3]
3 1 0 3
After:  [3, 2, 0, 1]

Before: [2, 1, 2, 0]
12 2 1 3
After:  [2, 1, 2, 3]

Before: [0, 2, 1, 0]
2 3 1 2
After:  [0, 2, 2, 0]

Before: [0, 0, 1, 0]
1 0 1 1
After:  [0, 1, 1, 0]

Before: [3, 2, 3, 1]
5 3 2 2
After:  [3, 2, 3, 1]

Before: [2, 2, 0, 3]
11 0 3 3
After:  [2, 2, 0, 0]

Before: [3, 2, 0, 0]
3 1 0 2
After:  [3, 2, 1, 0]

Before: [1, 0, 2, 0]
13 1 0 1
After:  [1, 1, 2, 0]

Before: [0, 3, 0, 0]
2 0 1 3
After:  [0, 3, 0, 3]

Before: [0, 1, 1, 2]
10 1 0 1
After:  [0, 1, 1, 2]

Before: [0, 0, 0, 2]
1 0 1 1
After:  [0, 1, 0, 2]

Before: [0, 1, 1, 2]
12 0 1 2
After:  [0, 1, 1, 2]

Before: [0, 0, 3, 0]
1 0 1 2
After:  [0, 0, 1, 0]

Before: [0, 1, 2, 1]
10 1 0 1
After:  [0, 1, 2, 1]

Before: [0, 1, 3, 1]
4 2 2 1
After:  [0, 6, 3, 1]

Before: [0, 1, 2, 3]
12 1 2 2
After:  [0, 1, 3, 3]

Before: [1, 0, 1, 1]
14 1 3 2
After:  [1, 0, 1, 1]

Before: [1, 0, 1, 3]
13 1 0 0
After:  [1, 0, 1, 3]

Before: [1, 1, 3, 1]
5 3 2 1
After:  [1, 3, 3, 1]

Before: [1, 2, 2, 3]
9 1 0 0
After:  [1, 2, 2, 3]

Before: [0, 0, 2, 1]
14 1 3 3
After:  [0, 0, 2, 1]

Before: [0, 3, 2, 0]
8 0 0 1
After:  [0, 0, 2, 0]

Before: [1, 0, 3, 1]
12 0 2 2
After:  [1, 0, 3, 1]

Before: [3, 0, 1, 2]
0 3 2 3
After:  [3, 0, 1, 4]

Before: [0, 2, 2, 2]
15 1 2 2
After:  [0, 2, 4, 2]

Before: [0, 1, 1, 1]
10 1 0 1
After:  [0, 1, 1, 1]

Before: [1, 0, 0, 2]
13 1 0 2
After:  [1, 0, 1, 2]

Before: [0, 0, 1, 1]
14 1 3 0
After:  [1, 0, 1, 1]

Before: [3, 2, 0, 1]
3 1 0 0
After:  [1, 2, 0, 1]

Before: [0, 0, 2, 0]
1 0 1 1
After:  [0, 1, 2, 0]

Before: [0, 2, 0, 2]
2 0 3 0
After:  [2, 2, 0, 2]

Before: [0, 0, 1, 2]
14 1 2 2
After:  [0, 0, 1, 2]

Before: [1, 1, 0, 1]
4 1 1 1
After:  [1, 2, 0, 1]

Before: [0, 1, 2, 1]
5 3 2 2
After:  [0, 1, 3, 1]

Before: [0, 3, 1, 3]
8 0 0 3
After:  [0, 3, 1, 0]

Before: [0, 0, 2, 1]
7 3 1 3
After:  [0, 0, 2, 1]

Before: [0, 1, 1, 2]
0 3 2 0
After:  [4, 1, 1, 2]

Before: [1, 2, 2, 2]
9 1 0 1
After:  [1, 1, 2, 2]

Before: [2, 3, 1, 2]
0 3 2 1
After:  [2, 4, 1, 2]

Before: [2, 0, 2, 2]
15 3 2 0
After:  [4, 0, 2, 2]

Before: [1, 2, 1, 3]
9 1 0 3
After:  [1, 2, 1, 1]

Before: [3, 2, 2, 3]
15 1 2 1
After:  [3, 4, 2, 3]

Before: [1, 0, 2, 2]
13 1 0 3
After:  [1, 0, 2, 1]

Before: [3, 0, 1, 0]
7 2 1 2
After:  [3, 0, 1, 0]

Before: [0, 2, 2, 3]
8 0 0 0
After:  [0, 2, 2, 3]

Before: [2, 0, 1, 3]
14 1 2 3
After:  [2, 0, 1, 1]

Before: [0, 1, 1, 1]
4 1 1 2
After:  [0, 1, 2, 1]

Before: [1, 0, 1, 3]
11 0 3 3
After:  [1, 0, 1, 0]

Before: [1, 3, 2, 1]
5 3 2 0
After:  [3, 3, 2, 1]

Before: [1, 0, 2, 1]
15 2 2 2
After:  [1, 0, 4, 1]

Before: [3, 2, 2, 3]
3 1 0 1
After:  [3, 1, 2, 3]

Before: [0, 1, 2, 3]
10 1 0 0
After:  [1, 1, 2, 3]

Before: [2, 0, 0, 1]
14 1 3 2
After:  [2, 0, 1, 1]

Before: [0, 1, 2, 3]
4 1 3 0
After:  [4, 1, 2, 3]

Before: [3, 2, 0, 2]
0 1 2 1
After:  [3, 4, 0, 2]

Before: [1, 0, 2, 3]
12 0 2 2
After:  [1, 0, 3, 3]

Before: [3, 2, 1, 0]
3 1 0 3
After:  [3, 2, 1, 1]

Before: [1, 1, 3, 1]
6 0 2 2
After:  [1, 1, 2, 1]

Before: [2, 0, 3, 0]
0 0 2 0
After:  [4, 0, 3, 0]

Before: [0, 0, 3, 3]
8 0 0 0
After:  [0, 0, 3, 3]

Before: [0, 0, 0, 1]
14 1 3 3
After:  [0, 0, 0, 1]

Before: [1, 0, 2, 2]
13 1 0 0
After:  [1, 0, 2, 2]

Before: [0, 0, 2, 2]
15 2 2 1
After:  [0, 4, 2, 2]

Before: [2, 2, 3, 2]
0 3 2 2
After:  [2, 2, 4, 2]

Before: [1, 2, 0, 0]
9 1 0 3
After:  [1, 2, 0, 1]

Before: [3, 2, 1, 3]
3 1 0 3
After:  [3, 2, 1, 1]

Before: [0, 0, 3, 1]
1 0 1 2
After:  [0, 0, 1, 1]

Before: [1, 0, 1, 1]
4 3 2 1
After:  [1, 2, 1, 1]

Before: [0, 1, 2, 3]
10 1 0 2
After:  [0, 1, 1, 3]

Before: [1, 1, 3, 2]
12 1 3 1
After:  [1, 3, 3, 2]

Before: [0, 0, 0, 3]
8 0 0 0
After:  [0, 0, 0, 3]

Before: [0, 2, 0, 0]
8 0 0 2
After:  [0, 2, 0, 0]

Before: [1, 0, 1, 0]
13 1 0 3
After:  [1, 0, 1, 1]

Before: [1, 0, 2, 2]
7 0 1 2
After:  [1, 0, 1, 2]

Before: [3, 0, 2, 3]
15 2 2 2
After:  [3, 0, 4, 3]

Before: [3, 2, 3, 3]
3 1 0 3
After:  [3, 2, 3, 1]

Before: [3, 2, 1, 2]
3 1 0 1
After:  [3, 1, 1, 2]

Before: [0, 0, 3, 2]
1 0 1 0
After:  [1, 0, 3, 2]

Before: [1, 2, 0, 1]
9 1 0 3
After:  [1, 2, 0, 1]

Before: [2, 0, 2, 3]
11 0 3 1
After:  [2, 0, 2, 3]

Before: [1, 1, 1, 2]
5 3 1 1
After:  [1, 3, 1, 2]

Before: [3, 2, 2, 1]
3 1 0 2
After:  [3, 2, 1, 1]

Before: [0, 2, 2, 2]
15 3 2 2
After:  [0, 2, 4, 2]

Before: [0, 2, 2, 3]
8 0 0 1
After:  [0, 0, 2, 3]

Before: [3, 2, 2, 3]
3 1 0 0
After:  [1, 2, 2, 3]

Before: [2, 2, 1, 3]
4 3 3 2
After:  [2, 2, 6, 3]

Before: [2, 3, 1, 2]
12 2 3 1
After:  [2, 3, 1, 2]

Before: [3, 0, 1, 1]
14 1 2 2
After:  [3, 0, 1, 1]

Before: [1, 1, 3, 0]
4 2 1 1
After:  [1, 4, 3, 0]

Before: [0, 0, 0, 1]
14 1 3 0
After:  [1, 0, 0, 1]

Before: [1, 1, 3, 0]
4 1 1 0
After:  [2, 1, 3, 0]

Before: [0, 1, 0, 0]
10 1 0 1
After:  [0, 1, 0, 0]

Before: [3, 2, 3, 2]
3 1 0 3
After:  [3, 2, 3, 1]

Before: [2, 0, 3, 2]
2 1 0 3
After:  [2, 0, 3, 2]

Before: [1, 2, 3, 3]
11 0 3 2
After:  [1, 2, 0, 3]

Before: [1, 2, 0, 3]
9 1 0 1
After:  [1, 1, 0, 3]

Before: [3, 2, 0, 3]
4 3 3 1
After:  [3, 6, 0, 3]

Before: [2, 3, 1, 0]
2 3 0 3
After:  [2, 3, 1, 2]

Before: [0, 1, 2, 3]
8 0 0 1
After:  [0, 0, 2, 3]

Before: [2, 3, 3, 3]
4 3 2 1
After:  [2, 6, 3, 3]

Before: [0, 0, 2, 0]
8 0 0 1
After:  [0, 0, 2, 0]

Before: [1, 0, 0, 2]
13 1 0 1
After:  [1, 1, 0, 2]

Before: [1, 0, 2, 3]
11 0 3 3
After:  [1, 0, 2, 0]

Before: [1, 2, 1, 3]
9 1 0 0
After:  [1, 2, 1, 3]

Before: [2, 0, 0, 3]
11 0 3 2
After:  [2, 0, 0, 3]

Before: [1, 0, 0, 0]
13 1 0 0
After:  [1, 0, 0, 0]

Before: [0, 2, 1, 3]
0 1 2 1
After:  [0, 4, 1, 3]

Before: [0, 3, 1, 2]
6 3 3 2
After:  [0, 3, 6, 2]

Before: [1, 2, 0, 0]
9 1 0 0
After:  [1, 2, 0, 0]

Before: [1, 0, 3, 0]
13 1 0 1
After:  [1, 1, 3, 0]

Before: [0, 0, 1, 3]
1 0 1 3
After:  [0, 0, 1, 1]

Before: [0, 2, 1, 1]
8 0 0 0
After:  [0, 2, 1, 1]

Before: [3, 0, 1, 1]
14 1 3 2
After:  [3, 0, 1, 1]

Before: [3, 2, 0, 3]
3 1 0 0
After:  [1, 2, 0, 3]

Before: [1, 2, 1, 0]
9 1 0 1
After:  [1, 1, 1, 0]

Before: [1, 2, 3, 0]
0 1 2 2
After:  [1, 2, 4, 0]

Before: [1, 2, 2, 2]
9 1 0 2
After:  [1, 2, 1, 2]

Before: [0, 1, 3, 1]
10 1 0 0
After:  [1, 1, 3, 1]

Before: [0, 2, 2, 0]
2 0 1 0
After:  [2, 2, 2, 0]

Before: [0, 0, 2, 2]
1 0 1 2
After:  [0, 0, 1, 2]

Before: [1, 2, 3, 0]
9 1 0 3
After:  [1, 2, 3, 1]

Before: [0, 0, 2, 0]
1 0 1 3
After:  [0, 0, 2, 1]

Before: [0, 3, 2, 0]
8 0 0 3
After:  [0, 3, 2, 0]

Before: [0, 0, 1, 2]
1 0 1 3
After:  [0, 0, 1, 1]

Before: [3, 2, 3, 0]
3 1 0 1
After:  [3, 1, 3, 0]

Before: [1, 0, 3, 1]
5 3 2 2
After:  [1, 0, 3, 1]

Before: [1, 0, 1, 0]
14 1 2 2
After:  [1, 0, 1, 0]

Before: [2, 3, 3, 3]
11 0 3 2
After:  [2, 3, 0, 3]

Before: [1, 2, 1, 0]
2 3 1 2
After:  [1, 2, 2, 0]

Before: [0, 2, 3, 2]
0 3 2 2
After:  [0, 2, 4, 2]

Before: [0, 3, 3, 0]
6 2 3 0
After:  [9, 3, 3, 0]

Before: [1, 0, 1, 1]
14 1 3 3
After:  [1, 0, 1, 1]

Before: [0, 1, 0, 3]
10 1 0 1
After:  [0, 1, 0, 3]

Before: [3, 0, 1, 2]
14 1 2 2
After:  [3, 0, 1, 2]

Before: [1, 3, 2, 3]
4 3 3 0
After:  [6, 3, 2, 3]

Before: [2, 3, 2, 2]
15 3 2 3
After:  [2, 3, 2, 4]

Before: [2, 0, 1, 1]
14 1 2 1
After:  [2, 1, 1, 1]

Before: [3, 2, 1, 1]
3 1 0 1
After:  [3, 1, 1, 1]

Before: [1, 2, 3, 3]
4 3 2 0
After:  [6, 2, 3, 3]

Before: [1, 0, 1, 1]
7 3 1 0
After:  [1, 0, 1, 1]

Before: [1, 3, 3, 1]
5 3 2 3
After:  [1, 3, 3, 3]

Before: [2, 0, 3, 1]
14 1 3 1
After:  [2, 1, 3, 1]

Before: [1, 0, 3, 3]
11 0 3 3
After:  [1, 0, 3, 0]

Before: [2, 0, 2, 1]
14 1 3 3
After:  [2, 0, 2, 1]

Before: [0, 0, 0, 0]
8 0 0 0
After:  [0, 0, 0, 0]

Before: [3, 1, 2, 1]
5 3 2 2
After:  [3, 1, 3, 1]

Before: [0, 2, 3, 0]
8 0 0 1
After:  [0, 0, 3, 0]

Before: [2, 1, 1, 3]
4 3 3 3
After:  [2, 1, 1, 6]

Before: [2, 1, 0, 2]
0 3 2 0
After:  [4, 1, 0, 2]

Before: [0, 0, 2, 3]
1 0 1 1
After:  [0, 1, 2, 3]

Before: [2, 1, 2, 3]
15 2 2 0
After:  [4, 1, 2, 3]

Before: [1, 2, 1, 0]
9 1 0 2
After:  [1, 2, 1, 0]

Before: [1, 2, 1, 2]
9 1 0 0
After:  [1, 2, 1, 2]

Before: [1, 0, 3, 3]
13 1 0 1
After:  [1, 1, 3, 3]

Before: [0, 1, 1, 2]
8 0 0 3
After:  [0, 1, 1, 0]

Before: [1, 0, 3, 3]
13 1 0 2
After:  [1, 0, 1, 3]

Before: [0, 1, 2, 1]
5 3 2 0
After:  [3, 1, 2, 1]

Before: [1, 3, 2, 1]
6 2 3 2
After:  [1, 3, 6, 1]

Before: [3, 2, 2, 1]
3 1 0 3
After:  [3, 2, 2, 1]

Before: [2, 0, 1, 1]
14 1 2 3
After:  [2, 0, 1, 1]

Before: [2, 1, 0, 2]
2 1 0 2
After:  [2, 1, 3, 2]

Before: [0, 2, 1, 0]
2 1 2 0
After:  [3, 2, 1, 0]

Before: [3, 2, 2, 0]
4 0 2 0
After:  [5, 2, 2, 0]

Before: [1, 0, 1, 3]
7 0 1 3
After:  [1, 0, 1, 1]

Before: [0, 2, 1, 2]
2 0 2 0
After:  [1, 2, 1, 2]

Before: [3, 0, 2, 3]
12 1 2 0
After:  [2, 0, 2, 3]

Before: [2, 2, 1, 0]
2 0 2 2
After:  [2, 2, 3, 0]

Before: [3, 1, 2, 2]
15 3 2 0
After:  [4, 1, 2, 2]

Before: [0, 3, 2, 2]
15 3 2 2
After:  [0, 3, 4, 2]

Before: [1, 0, 3, 2]
13 1 0 2
After:  [1, 0, 1, 2]

Before: [1, 1, 0, 2]
0 3 2 1
After:  [1, 4, 0, 2]

Before: [0, 1, 0, 2]
10 1 0 0
After:  [1, 1, 0, 2]

Before: [3, 1, 2, 1]
5 3 2 3
After:  [3, 1, 2, 3]

Before: [3, 1, 3, 3]
4 2 2 0
After:  [6, 1, 3, 3]

Before: [3, 3, 1, 0]
6 1 3 2
After:  [3, 3, 9, 0]

Before: [2, 1, 1, 3]
4 3 1 3
After:  [2, 1, 1, 4]

Before: [0, 0, 2, 3]
8 0 0 1
After:  [0, 0, 2, 3]

Before: [1, 3, 2, 3]
11 0 3 1
After:  [1, 0, 2, 3]

Before: [2, 2, 1, 0]
0 0 2 3
After:  [2, 2, 1, 4]

Before: [1, 0, 2, 1]
14 1 3 3
After:  [1, 0, 2, 1]

Before: [0, 1, 2, 2]
5 3 1 3
After:  [0, 1, 2, 3]

Before: [1, 0, 2, 3]
7 0 1 1
After:  [1, 1, 2, 3]

Before: [2, 2, 1, 3]
11 0 3 3
After:  [2, 2, 1, 0]

Before: [0, 1, 2, 3]
10 1 0 3
After:  [0, 1, 2, 1]

Before: [1, 0, 0, 3]
7 0 1 0
After:  [1, 0, 0, 3]

Before: [0, 1, 0, 3]
8 0 0 0
After:  [0, 1, 0, 3]

Before: [2, 1, 0, 0]
2 1 0 2
After:  [2, 1, 3, 0]

Before: [0, 0, 3, 1]
1 0 1 3
After:  [0, 0, 3, 1]

Before: [1, 0, 0, 1]
14 1 3 3
After:  [1, 0, 0, 1]

Before: [2, 0, 1, 2]
2 1 3 2
After:  [2, 0, 2, 2]

Before: [1, 0, 2, 2]
2 1 3 0
After:  [2, 0, 2, 2]

Before: [3, 2, 0, 1]
3 1 0 2
After:  [3, 2, 1, 1]

Before: [1, 0, 2, 3]
11 0 3 2
After:  [1, 0, 0, 3]

Before: [2, 2, 1, 2]
12 2 3 0
After:  [3, 2, 1, 2]

Before: [2, 0, 1, 3]
11 0 3 2
After:  [2, 0, 0, 3]

Before: [2, 0, 3, 1]
14 1 3 3
After:  [2, 0, 3, 1]

Before: [1, 2, 3, 0]
9 1 0 1
After:  [1, 1, 3, 0]

Before: [0, 0, 2, 2]
15 2 2 0
After:  [4, 0, 2, 2]

Before: [0, 1, 0, 2]
10 1 0 1
After:  [0, 1, 0, 2]

Before: [0, 1, 0, 3]
12 0 1 2
After:  [0, 1, 1, 3]

Before: [2, 1, 0, 3]
11 0 3 2
After:  [2, 1, 0, 3]

Before: [2, 3, 0, 3]
6 3 2 1
After:  [2, 6, 0, 3]

Before: [0, 0, 2, 3]
15 2 2 2
After:  [0, 0, 4, 3]

Before: [3, 2, 1, 1]
3 1 0 0
After:  [1, 2, 1, 1]

Before: [0, 2, 2, 0]
15 2 2 2
After:  [0, 2, 4, 0]

Before: [0, 3, 0, 1]
2 2 1 1
After:  [0, 3, 0, 1]

Before: [0, 3, 0, 1]
8 0 0 0
After:  [0, 3, 0, 1]

Before: [2, 1, 1, 3]
11 0 3 0
After:  [0, 1, 1, 3]

Before: [1, 0, 3, 1]
14 1 3 3
After:  [1, 0, 3, 1]

Before: [1, 1, 2, 3]
4 0 3 3
After:  [1, 1, 2, 4]

Before: [2, 0, 2, 1]
15 0 2 0
After:  [4, 0, 2, 1]

Before: [1, 0, 3, 2]
12 0 2 3
After:  [1, 0, 3, 3]

Before: [3, 2, 2, 2]
6 0 3 3
After:  [3, 2, 2, 9]

Before: [1, 0, 1, 2]
13 1 0 0
After:  [1, 0, 1, 2]

Before: [2, 3, 1, 3]
4 3 3 2
After:  [2, 3, 6, 3]

Before: [3, 0, 3, 1]
7 3 1 1
After:  [3, 1, 3, 1]

Before: [1, 0, 3, 1]
13 1 0 1
After:  [1, 1, 3, 1]

Before: [0, 0, 1, 2]
1 0 1 1
After:  [0, 1, 1, 2]

Before: [1, 0, 0, 3]
13 1 0 0
After:  [1, 0, 0, 3]

Before: [0, 1, 3, 2]
10 1 0 3
After:  [0, 1, 3, 1]

Before: [3, 0, 0, 3]
6 3 2 2
After:  [3, 0, 6, 3]

Before: [2, 2, 0, 2]
6 0 3 3
After:  [2, 2, 0, 6]

Before: [1, 2, 1, 3]
2 1 2 2
After:  [1, 2, 3, 3]

Before: [3, 0, 1, 3]
14 1 2 3
After:  [3, 0, 1, 1]

Before: [2, 1, 0, 2]
5 3 1 0
After:  [3, 1, 0, 2]

Before: [1, 0, 2, 3]
11 0 3 1
After:  [1, 0, 2, 3]

Before: [0, 3, 1, 1]
6 1 3 2
After:  [0, 3, 9, 1]

Before: [1, 1, 0, 3]
11 0 3 3
After:  [1, 1, 0, 0]

Before: [3, 2, 0, 2]
3 1 0 3
After:  [3, 2, 0, 1]

Before: [1, 0, 1, 2]
7 2 1 3
After:  [1, 0, 1, 1]

Before: [2, 0, 1, 0]
14 1 2 2
After:  [2, 0, 1, 0]

Before: [1, 0, 3, 3]
13 1 0 0
After:  [1, 0, 3, 3]

Before: [1, 2, 0, 3]
11 0 3 1
After:  [1, 0, 0, 3]

Before: [0, 1, 1, 1]
10 1 0 2
After:  [0, 1, 1, 1]

Before: [1, 0, 1, 0]
13 1 0 0
After:  [1, 0, 1, 0]

Before: [1, 1, 0, 1]
4 0 1 1
After:  [1, 2, 0, 1]

Before: [0, 0, 3, 2]
0 3 2 0
After:  [4, 0, 3, 2]

Before: [0, 3, 2, 3]
15 2 2 2
After:  [0, 3, 4, 3]

Before: [0, 1, 0, 0]
8 0 0 0
After:  [0, 1, 0, 0]

Before: [2, 0, 2, 3]
12 1 3 3
After:  [2, 0, 2, 3]

Before: [2, 0, 3, 2]
0 3 2 2
After:  [2, 0, 4, 2]

Before: [0, 0, 1, 2]
7 2 1 2
After:  [0, 0, 1, 2]

Before: [3, 1, 0, 2]
5 3 1 2
After:  [3, 1, 3, 2]

Before: [0, 0, 3, 2]
1 0 1 2
After:  [0, 0, 1, 2]

Before: [1, 1, 1, 2]
5 3 1 0
After:  [3, 1, 1, 2]

Before: [3, 2, 1, 0]
6 0 3 0
After:  [9, 2, 1, 0]

Before: [0, 0, 0, 1]
8 0 0 3
After:  [0, 0, 0, 0]

Before: [2, 0, 1, 2]
14 1 2 1
After:  [2, 1, 1, 2]

Before: [1, 0, 2, 0]
13 1 0 3
After:  [1, 0, 2, 1]

Before: [0, 0, 3, 3]
1 0 1 2
After:  [0, 0, 1, 3]

Before: [0, 2, 2, 1]
5 3 2 2
After:  [0, 2, 3, 1]

Before: [3, 1, 0, 1]
4 3 1 3
After:  [3, 1, 0, 2]

Before: [2, 2, 2, 1]
15 0 2 3
After:  [2, 2, 2, 4]

Before: [0, 1, 1, 0]
4 1 1 3
After:  [0, 1, 1, 2]

Before: [2, 2, 3, 3]
0 1 2 0
After:  [4, 2, 3, 3]

Before: [1, 0, 1, 2]
14 1 2 3
After:  [1, 0, 1, 1]

Before: [3, 3, 1, 1]
6 1 3 1
After:  [3, 9, 1, 1]

Before: [1, 2, 2, 1]
15 1 2 1
After:  [1, 4, 2, 1]

Before: [1, 1, 2, 3]
4 3 2 1
After:  [1, 5, 2, 3]

Before: [1, 0, 2, 3]
13 1 0 0
After:  [1, 0, 2, 3]

Before: [0, 1, 0, 2]
8 0 0 1
After:  [0, 0, 0, 2]

Before: [0, 1, 2, 1]
10 1 0 2
After:  [0, 1, 1, 1]

Before: [3, 0, 3, 1]
7 3 1 2
After:  [3, 0, 1, 1]

Before: [2, 2, 3, 3]
11 0 3 2
After:  [2, 2, 0, 3]

Before: [3, 0, 1, 0]
14 1 2 3
After:  [3, 0, 1, 1]

Before: [0, 0, 1, 2]
12 2 3 0
After:  [3, 0, 1, 2]

Before: [0, 0, 1, 1]
1 0 1 2
After:  [0, 0, 1, 1]

Before: [1, 2, 0, 3]
9 1 0 3
After:  [1, 2, 0, 1]

Before: [0, 1, 3, 0]
10 1 0 3
After:  [0, 1, 3, 1]

Before: [2, 2, 1, 3]
2 2 0 2
After:  [2, 2, 3, 3]

Before: [0, 3, 3, 0]
8 0 0 0
After:  [0, 3, 3, 0]

Before: [1, 2, 0, 1]
9 1 0 2
After:  [1, 2, 1, 1]

Before: [0, 1, 1, 3]
10 1 0 1
After:  [0, 1, 1, 3]

Before: [2, 1, 1, 3]
11 0 3 3
After:  [2, 1, 1, 0]

Before: [3, 0, 1, 1]
7 3 1 3
After:  [3, 0, 1, 1]

Before: [1, 0, 2, 2]
13 1 0 2
After:  [1, 0, 1, 2]

Before: [3, 2, 1, 2]
6 0 3 0
After:  [9, 2, 1, 2]

Before: [0, 2, 3, 0]
2 0 1 3
After:  [0, 2, 3, 2]

Before: [2, 2, 1, 2]
6 3 3 1
After:  [2, 6, 1, 2]

Before: [0, 1, 1, 2]
5 3 1 1
After:  [0, 3, 1, 2]

Before: [1, 2, 2, 1]
6 1 3 3
After:  [1, 2, 2, 6]

Before: [1, 2, 0, 2]
9 1 0 1
After:  [1, 1, 0, 2]

Before: [1, 0, 3, 1]
4 2 2 2
After:  [1, 0, 6, 1]

Before: [1, 0, 2, 2]
13 1 0 1
After:  [1, 1, 2, 2]

Before: [0, 1, 2, 3]
10 1 0 1
After:  [0, 1, 2, 3]

Before: [1, 2, 0, 2]
9 1 0 2
After:  [1, 2, 1, 2]

Before: [3, 2, 2, 2]
15 3 2 2
After:  [3, 2, 4, 2]

Before: [3, 2, 0, 3]
3 1 0 2
After:  [3, 2, 1, 3]

Before: [2, 1, 3, 1]
5 3 2 2
After:  [2, 1, 3, 1]

Before: [1, 0, 3, 1]
13 1 0 0
After:  [1, 0, 3, 1]

Before: [1, 0, 2, 0]
13 1 0 0
After:  [1, 0, 2, 0]

Before: [0, 0, 1, 3]
7 2 1 1
After:  [0, 1, 1, 3]

Before: [1, 0, 3, 0]
13 1 0 2
After:  [1, 0, 1, 0]

Before: [0, 2, 1, 0]
8 0 0 1
After:  [0, 0, 1, 0]

Before: [0, 0, 1, 2]
0 3 2 2
After:  [0, 0, 4, 2]

Before: [0, 0, 3, 3]
1 0 1 3
After:  [0, 0, 3, 1]

Before: [2, 3, 0, 2]
6 1 2 2
After:  [2, 3, 6, 2]

Before: [3, 1, 2, 2]
4 0 2 1
After:  [3, 5, 2, 2]

Before: [3, 2, 3, 0]
2 3 1 1
After:  [3, 2, 3, 0]

Before: [1, 0, 2, 1]
7 0 1 1
After:  [1, 1, 2, 1]

Before: [0, 0, 2, 1]
6 2 3 1
After:  [0, 6, 2, 1]`

const testProgram = `5 0 2 3
5 1 3 1
5 3 2 2
14 3 2 2
6 2 3 2
4 0 2 0
5 2 3 3
5 0 0 2
14 2 3 3
6 3 3 3
4 3 0 0
10 0 3 2
5 2 1 3
6 1 0 0
0 0 2 0
5 2 0 1
1 0 3 1
6 1 3 1
4 2 1 2
10 2 1 1
5 3 3 0
6 3 0 2
0 2 0 2
6 3 0 3
0 3 1 3
13 2 0 2
6 2 2 2
4 1 2 1
5 0 1 3
5 3 3 2
5 1 2 0
14 3 2 0
6 0 1 0
4 1 0 1
10 1 3 0
6 0 0 1
0 1 1 1
5 2 3 2
5 3 2 3
6 3 1 3
4 3 0 0
10 0 0 1
5 2 1 0
5 3 0 2
5 3 1 3
8 3 0 0
6 0 1 0
4 1 0 1
10 1 3 2
5 1 2 1
6 2 0 0
0 0 2 0
5 2 2 3
1 0 3 1
6 1 2 1
4 1 2 2
10 2 3 3
6 1 0 0
0 0 1 0
5 3 0 2
5 0 3 1
6 0 2 1
6 1 2 1
6 1 2 1
4 3 1 3
10 3 0 0
5 1 2 2
5 0 3 1
6 1 0 3
0 3 0 3
5 3 1 1
6 1 2 1
4 1 0 0
10 0 1 3
5 3 1 1
5 0 1 0
7 1 2 0
6 0 3 0
4 3 0 3
6 0 0 0
0 0 1 0
5 2 0 2
0 0 1 1
6 1 2 1
4 3 1 3
10 3 0 1
5 1 1 3
5 2 2 0
9 0 3 3
6 3 3 3
4 1 3 1
5 2 1 3
1 0 3 0
6 0 3 0
4 0 1 1
10 1 2 0
5 3 1 1
5 0 3 3
6 2 0 2
0 2 0 2
7 1 2 1
6 1 1 1
6 1 3 1
4 1 0 0
10 0 0 1
5 1 2 0
5 3 3 3
5 2 3 3
6 3 3 3
4 3 1 1
10 1 3 2
5 3 1 3
5 3 3 0
5 2 3 1
8 0 1 0
6 0 1 0
4 2 0 2
10 2 1 3
5 0 2 2
6 0 0 0
0 0 3 0
5 1 0 1
13 2 0 2
6 2 2 2
6 2 2 2
4 2 3 3
10 3 1 0
5 0 0 2
6 3 0 3
0 3 1 3
5 3 1 1
6 3 2 2
6 2 2 2
4 0 2 0
10 0 0 3
5 2 2 1
5 1 0 2
5 3 3 0
5 2 0 1
6 1 3 1
4 3 1 3
10 3 3 1
5 2 1 2
5 1 3 0
5 2 0 3
10 0 2 2
6 2 1 2
4 1 2 1
5 2 3 2
6 0 0 3
0 3 0 3
5 0 1 0
11 3 2 2
6 2 2 2
4 1 2 1
10 1 3 0
6 2 0 1
0 1 1 1
5 3 2 3
6 0 0 2
0 2 0 2
7 3 2 2
6 2 1 2
6 2 2 2
4 0 2 0
5 3 0 1
5 0 1 2
6 3 0 3
0 3 1 3
4 3 3 3
6 3 2 3
4 3 0 0
10 0 3 3
5 3 3 2
5 3 2 0
5 2 1 1
2 1 0 1
6 1 1 1
6 1 1 1
4 1 3 3
10 3 2 0
6 2 0 2
0 2 1 2
6 1 0 3
0 3 3 3
5 1 1 1
7 3 2 2
6 2 1 2
6 2 3 2
4 0 2 0
10 0 2 2
5 2 2 3
5 0 2 1
5 2 3 0
1 0 3 1
6 1 1 1
4 2 1 2
10 2 1 0
6 2 0 2
0 2 2 2
5 3 2 1
3 2 1 2
6 2 2 2
4 2 0 0
10 0 2 3
5 1 0 1
5 3 1 2
5 2 1 0
2 0 2 0
6 0 2 0
6 0 3 0
4 3 0 3
10 3 3 0
5 2 1 1
5 2 0 2
5 0 3 3
11 3 2 2
6 2 1 2
4 0 2 0
10 0 0 1
5 0 1 2
5 3 0 0
13 2 0 0
6 0 3 0
4 1 0 1
6 2 0 2
0 2 2 2
5 1 0 0
5 3 1 3
10 0 2 2
6 2 2 2
4 1 2 1
5 3 1 2
5 0 1 3
6 2 0 0
0 0 2 0
13 0 2 2
6 2 2 2
4 1 2 1
10 1 1 3
6 3 0 1
0 1 0 1
6 0 0 0
0 0 1 0
5 2 0 2
10 0 2 0
6 0 2 0
4 0 3 3
10 3 0 1
5 2 0 3
5 3 1 0
5 0 1 2
14 2 3 0
6 0 2 0
4 1 0 1
10 1 3 0
5 3 1 2
6 3 0 1
0 1 1 1
15 1 3 1
6 1 1 1
6 1 3 1
4 1 0 0
10 0 1 3
5 3 1 1
5 2 1 0
13 0 2 1
6 1 1 1
4 1 3 3
10 3 0 1
5 1 1 3
13 0 2 0
6 0 3 0
4 1 0 1
10 1 3 0
5 3 3 1
5 0 2 3
5 2 1 2
11 3 2 1
6 1 3 1
6 1 3 1
4 1 0 0
10 0 2 3
5 3 3 1
5 0 2 0
5 0 3 2
7 1 2 2
6 2 3 2
4 2 3 3
10 3 1 1
5 3 1 0
6 2 0 3
0 3 2 3
5 0 1 2
14 2 3 3
6 3 2 3
4 1 3 1
10 1 2 3
5 1 0 1
5 1 1 2
7 0 2 1
6 1 3 1
4 1 3 3
10 3 2 1
6 2 0 3
0 3 2 3
5 0 2 2
5 2 0 0
1 0 3 0
6 0 1 0
4 1 0 1
6 2 0 2
0 2 3 2
5 2 2 0
1 0 3 0
6 0 2 0
6 0 3 0
4 1 0 1
10 1 2 2
5 2 3 0
5 1 3 1
1 0 3 3
6 3 3 3
4 2 3 2
10 2 0 1
5 2 2 2
5 0 1 3
12 2 3 0
6 0 2 0
4 1 0 1
10 1 1 0
5 3 0 2
5 3 1 1
14 3 2 1
6 1 3 1
6 1 1 1
4 1 0 0
10 0 1 1
6 2 0 2
0 2 1 2
5 2 2 0
5 2 1 3
1 0 3 2
6 2 1 2
4 1 2 1
10 1 2 2
5 3 0 0
5 1 1 3
5 0 3 1
4 3 3 0
6 0 1 0
4 2 0 2
10 2 0 3
5 3 2 0
5 3 3 1
5 0 1 2
13 2 0 1
6 1 2 1
4 3 1 3
5 3 1 2
5 2 3 1
2 1 2 2
6 2 3 2
4 3 2 3
10 3 2 0
5 0 0 2
5 1 0 3
5 3 1 1
7 1 2 2
6 2 2 2
4 0 2 0
5 2 3 1
5 2 0 2
5 0 0 3
11 3 2 1
6 1 3 1
4 1 0 0
5 1 0 1
11 3 2 2
6 2 1 2
6 2 2 2
4 2 0 0
5 0 1 1
5 1 3 2
5 3 0 3
7 3 2 1
6 1 1 1
4 1 0 0
10 0 0 2
5 1 0 3
5 1 2 1
5 2 3 0
9 0 3 3
6 3 2 3
4 3 2 2
10 2 2 3
5 1 1 0
5 0 3 2
5 3 1 1
0 0 1 2
6 2 3 2
6 2 1 2
4 3 2 3
10 3 0 1
5 3 3 2
6 3 0 3
0 3 0 3
5 2 0 0
14 3 2 2
6 2 1 2
4 1 2 1
10 1 2 3
5 3 2 1
5 1 1 0
6 2 0 2
0 2 2 2
10 0 2 1
6 1 1 1
4 1 3 3
10 3 0 2
5 0 1 3
5 2 3 0
6 1 0 1
0 1 3 1
12 0 3 1
6 1 2 1
4 2 1 2
10 2 1 0
5 1 1 3
5 0 3 2
5 3 3 1
6 3 2 1
6 1 2 1
4 1 0 0
10 0 1 1
5 0 1 3
6 2 0 2
0 2 2 2
5 0 0 0
11 3 2 0
6 0 1 0
6 0 2 0
4 1 0 1
5 0 3 2
5 3 1 3
5 3 2 0
5 2 0 2
6 2 2 2
4 2 1 1
6 1 0 0
0 0 2 0
5 1 1 3
5 2 1 2
15 3 0 2
6 2 1 2
4 1 2 1
10 1 1 3
5 3 1 1
6 3 0 2
0 2 2 2
3 0 1 2
6 2 3 2
4 2 3 3
10 3 1 1
5 3 0 3
5 3 0 2
13 0 2 0
6 0 3 0
4 0 1 1
10 1 1 2
5 1 1 0
6 2 0 3
0 3 2 3
6 1 0 1
0 1 0 1
15 0 3 0
6 0 3 0
6 0 3 0
4 2 0 2
5 1 2 3
6 1 0 1
0 1 1 1
5 2 2 0
15 1 0 3
6 3 1 3
6 3 2 3
4 3 2 2
10 2 3 1
5 3 2 2
5 2 1 3
2 0 2 0
6 0 2 0
4 0 1 1
10 1 3 2
6 2 0 0
0 0 2 0
5 2 3 1
5 3 1 3
8 3 1 3
6 3 3 3
4 2 3 2
5 1 2 1
5 1 3 3
9 0 3 1
6 1 3 1
6 1 3 1
4 2 1 2
10 2 3 3
5 3 1 2
5 2 3 1
5 3 0 0
2 1 0 2
6 2 3 2
4 2 3 3
10 3 3 2
5 2 1 0
5 3 3 3
5 1 1 1
15 1 0 0
6 0 2 0
4 2 0 2
10 2 2 3
5 0 0 1
5 0 0 2
5 1 0 0
0 0 1 2
6 2 3 2
4 2 3 3
10 3 2 1
6 3 0 2
0 2 3 2
5 0 0 3
5 3 0 0
14 3 2 3
6 3 2 3
4 3 1 1
10 1 3 3
5 0 2 1
5 0 3 2
13 2 0 1
6 1 3 1
6 1 2 1
4 3 1 3
10 3 1 2
5 2 0 0
6 3 0 1
0 1 3 1
5 1 0 3
3 0 1 1
6 1 3 1
6 1 1 1
4 1 2 2
5 3 0 1
5 0 3 3
3 0 1 3
6 3 2 3
4 2 3 2
10 2 1 1
5 3 0 2
5 1 2 0
5 2 1 3
15 0 3 0
6 0 1 0
6 0 2 0
4 1 0 1
10 1 3 2
5 3 2 1
5 1 3 3
5 1 3 0
4 3 3 3
6 3 2 3
4 3 2 2
10 2 0 1
6 0 0 2
0 2 1 2
5 1 0 3
6 3 0 0
0 0 2 0
15 3 0 2
6 2 2 2
4 1 2 1
10 1 2 3
5 2 3 2
5 3 0 0
5 2 1 1
2 1 0 0
6 0 2 0
4 0 3 3
10 3 1 1
5 1 0 2
5 3 0 0
5 1 0 3
7 0 2 0
6 0 2 0
4 1 0 1
5 2 1 0
5 2 1 3
5 3 3 2
13 0 2 2
6 2 3 2
4 1 2 1
10 1 1 0
5 2 0 2
5 0 0 3
5 1 0 1
11 3 2 3
6 3 3 3
4 0 3 0
10 0 2 1
5 1 2 0
5 1 3 3
10 0 2 0
6 0 1 0
4 0 1 1
10 1 0 0
5 1 3 1
5 3 1 2
5 2 1 3
6 1 2 1
6 1 1 1
6 1 3 1
4 1 0 0
10 0 2 1
5 0 0 3
5 1 2 0
5 2 1 2
10 0 2 0
6 0 3 0
4 0 1 1
10 1 3 2
5 0 2 1
6 1 0 3
0 3 1 3
5 2 2 0
15 3 0 0
6 0 2 0
6 0 3 0
4 0 2 2
5 2 2 1
5 1 1 0
4 3 0 0
6 0 2 0
4 0 2 2
10 2 3 1
6 0 0 2
0 2 2 2
5 1 2 0
10 0 2 3
6 3 1 3
6 3 2 3
4 3 1 1
10 1 2 2
5 3 1 1
5 2 1 0
5 3 3 3
8 3 0 1
6 1 3 1
4 1 2 2
10 2 2 1
6 1 0 2
0 2 0 2
5 2 1 3
14 2 3 3
6 3 1 3
4 3 1 1
10 1 3 3
5 2 3 1
5 3 3 2
13 0 2 0
6 0 2 0
6 0 1 0
4 3 0 3
10 3 2 1
5 2 0 2
5 3 2 0
5 1 2 3
3 2 0 2
6 2 2 2
4 1 2 1
10 1 3 2
5 0 2 1
5 2 1 0
9 0 3 3
6 3 1 3
6 3 2 3
4 2 3 2
10 2 3 0
5 3 1 3
5 3 1 1
5 2 3 2
3 2 1 3
6 3 2 3
4 3 0 0
10 0 3 2
6 3 0 1
0 1 1 1
5 1 0 3
5 3 1 0
4 3 3 3
6 3 2 3
4 3 2 2
10 2 1 3
5 2 3 0
6 2 0 1
0 1 2 1
5 3 0 2
2 0 2 0
6 0 1 0
4 3 0 3
10 3 0 1
6 3 0 0
0 0 2 0
5 2 2 3
5 0 0 2
1 0 3 3
6 3 1 3
6 3 3 3
4 1 3 1
10 1 3 3
5 1 2 0
6 2 0 2
0 2 2 2
5 3 2 1
10 0 2 0
6 0 2 0
4 0 3 3
10 3 1 2
6 3 0 0
0 0 1 0
5 0 2 3
5 2 3 1
12 1 3 3
6 3 1 3
6 3 2 3
4 3 2 2
10 2 0 3
5 3 3 2
5 2 1 0
5 3 3 1
3 0 1 2
6 2 1 2
4 2 3 3
10 3 0 1
6 3 0 0
0 0 3 0
6 2 0 3
0 3 1 3
6 0 0 2
0 2 3 2
6 3 2 0
6 0 2 0
4 0 1 1
5 2 3 0
6 2 0 2
0 2 1 2
6 2 0 3
0 3 2 3
12 0 3 3
6 3 2 3
4 1 3 1
10 1 1 0
5 3 0 1
5 3 0 2
5 2 2 3
8 1 3 1
6 1 3 1
6 1 3 1
4 1 0 0
10 0 1 2
5 1 1 1
6 3 0 3
0 3 1 3
5 2 3 0
15 3 0 3
6 3 3 3
6 3 2 3
4 2 3 2
10 2 0 1
6 1 0 3
0 3 3 3
5 3 2 0
5 2 3 2
3 2 0 2
6 2 2 2
6 2 3 2
4 1 2 1
5 2 0 2
5 0 0 3
3 2 0 3
6 3 3 3
4 3 1 1
10 1 0 0
5 3 2 2
5 3 0 3
6 0 0 1
0 1 2 1
2 1 2 2
6 2 1 2
4 0 2 0
5 2 1 2
6 1 0 3
0 3 0 3
6 0 0 1
0 1 0 1
11 3 2 2
6 2 3 2
4 0 2 0
10 0 3 3
5 3 0 2
5 1 2 1
5 1 1 0
6 0 2 1
6 1 1 1
6 1 1 1
4 3 1 3
10 3 3 1
5 0 3 2
5 2 1 0
5 1 3 3
9 0 3 3
6 3 1 3
4 1 3 1
10 1 1 3
6 2 0 1
0 1 1 1
5 3 3 2
13 0 2 1
6 1 1 1
4 1 3 3
5 1 2 1
5 2 2 2
5 1 3 0
10 0 2 2
6 2 1 2
4 2 3 3
5 3 2 0
5 1 3 2
5 3 0 1
7 0 2 2
6 2 1 2
4 2 3 3
10 3 2 1
5 0 3 3
6 3 0 2
0 2 2 2
6 2 0 0
0 0 1 0
10 0 2 2
6 2 1 2
4 1 2 1
5 2 0 2
5 2 0 0
5 2 3 3
12 2 3 0
6 0 3 0
4 1 0 1
5 0 2 0
5 0 2 2
5 1 0 3
6 3 2 3
6 3 2 3
6 3 2 3
4 1 3 1
10 1 3 3
5 3 2 1
5 2 1 0
6 0 0 2
0 2 3 2
3 0 1 1
6 1 2 1
4 1 3 3
10 3 3 1
5 1 1 2
5 3 3 0
6 3 0 3
0 3 2 3
8 0 3 2
6 2 1 2
6 2 2 2
4 1 2 1
10 1 3 0
5 1 0 3
6 0 0 1
0 1 2 1
5 2 2 2
4 3 3 3
6 3 2 3
4 0 3 0
10 0 1 2
6 3 0 1
0 1 3 1
5 2 1 0
5 1 1 3
15 3 0 1
6 1 3 1
6 1 3 1
4 2 1 2
10 2 3 1
5 2 0 2
5 2 0 3
6 3 0 0
0 0 1 0
4 0 0 3
6 3 1 3
4 1 3 1
5 3 3 0
5 2 1 3
5 0 1 2
14 2 3 2
6 2 2 2
4 1 2 1
10 1 1 2
5 2 1 1
5 2 0 0
12 1 3 0
6 0 1 0
4 0 2 2
5 2 2 0
1 0 3 3
6 3 1 3
4 3 2 2
10 2 2 3
5 1 2 0
6 3 0 1
0 1 0 1
5 0 3 2
6 0 2 2
6 2 2 2
4 2 3 3
10 3 1 1
6 1 0 2
0 2 3 2
5 0 3 3
14 3 2 3
6 3 2 3
6 3 1 3
4 3 1 1
10 1 2 0
6 0 0 3
0 3 3 3
6 0 0 2
0 2 0 2
5 3 0 1
5 2 1 3
6 3 1 3
4 0 3 0
10 0 0 2
5 1 1 3
5 0 2 1
5 0 0 0
0 3 1 3
6 3 1 3
4 2 3 2
10 2 0 1
6 2 0 3
0 3 2 3
6 0 0 2
0 2 3 2
5 2 0 0
1 0 3 2
6 2 1 2
4 2 1 1
10 1 3 2
5 2 1 1
5 1 0 3
5 1 3 0
4 3 0 3
6 3 2 3
4 2 3 2
10 2 2 0
5 3 0 2
5 0 0 3
5 1 0 1
5 2 3 2
6 2 1 2
4 2 0 0
5 1 2 3
5 3 1 2
6 1 2 1
6 1 1 1
4 1 0 0
6 3 0 1
0 1 0 1
6 3 2 3
6 3 2 3
6 3 3 3
4 0 3 0
10 0 0 1
5 2 2 3
5 3 3 0
5 2 0 3
6 3 3 3
4 3 1 1
10 1 0 2
5 1 0 1
5 1 1 3
5 2 1 0
4 1 3 3
6 3 1 3
4 2 3 2
6 3 0 3
0 3 2 3
6 3 0 1
0 1 3 1
1 0 3 0
6 0 1 0
4 0 2 2
10 2 2 3
5 1 3 1
6 1 0 0
0 0 1 0
6 3 0 2
0 2 3 2
6 1 2 2
6 2 3 2
4 3 2 3
5 2 1 2
5 3 3 1
5 0 0 0
3 2 1 0
6 0 2 0
4 0 3 3
10 3 0 1
5 2 2 3
5 2 2 0
5 0 1 2
14 2 3 2
6 2 2 2
6 2 3 2
4 2 1 1
10 1 1 3
5 0 0 1
5 2 0 2
5 3 0 0
3 2 0 0
6 0 2 0
6 0 1 0
4 0 3 3
10 3 1 1
5 0 2 3
5 1 2 0
11 3 2 2
6 2 3 2
4 2 1 1
10 1 3 2
5 3 1 1
5 2 3 0
5 2 2 3
8 1 0 3
6 3 2 3
6 3 2 3
4 3 2 2
5 3 0 3
8 1 0 0
6 0 3 0
4 0 2 2
10 2 3 1
5 0 3 3
6 1 0 0
0 0 1 0
5 2 1 2
10 0 2 3
6 3 3 3
4 1 3 1
5 2 1 0
5 0 1 3
11 3 2 3
6 3 2 3
4 3 1 1
10 1 1 0`
