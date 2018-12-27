package main

import (
	"fmt"
	"strings"
)

func main() {
	rules := map[string]bool{}
	state := []byte(strings.Repeat(".", numGenerations*5) + initialState + strings.Repeat(".", numGenerations*5))
	for _, line := range strings.Split(input, "\n") {
		var rule, result string
		fmt.Sscanf(line, "%s => %s", &rule, &result)
		rules[rule] = (result == "#")
	}
	for gen := 0; gen < numGenerations; gen++ {
		fmt.Println(string(state))
		nextGen := []byte(string(state))
		for i := 2; i < len(state)-2; i++ {
			if rules[string(state[i-2:i+3])] {
				nextGen[i] = '#'
			} else {
				nextGen[i] = '.'
			}
		}
		state = nextGen
	}
	fmt.Println(numPots(state, numGenerations))
}

// it turns out the number of "alive" pots stabilizes after a while and the sum for it just keeps increasing by NUM_ALIVE_POTS since each pot just goes one to the right each generation, so I ran the first 2000 generations and then did this:
// >>> 105872+((50000000000-2000)*52)
// 2600000001872
// Kind of disappointing from a code perspective, but an answer is an answer!

func numPots(s []byte, numGenerations int) int {
	n := 0
	for i, r := range s {
		if r == '#' {
			n += (i - (numGenerations * 5))
		}
	}
	return n
}

const numGenerations = 2000
const initialState = `###.......##....#.#.#..###.##..##.....#....#.#.....##.###...###.#...###.###.#.###...#.####.##.#....#`
const input = `..... => .
#..## => .
..### => #
..#.# => #
.#.#. => .
####. => .
##.## => #
#.... => .
#...# => .
...## => .
##..# => .
.###. => #
##### => #
#.#.. => #
.##.. => #
.#.## => .
...#. => #
#.##. => #
..#.. => #
##... => #
....# => .
###.# => #
#..#. => #
#.### => #
##.#. => .
###.. => #
.#### => .
.#... => #
..##. => .
.##.# => .
#.#.# => #
.#..# => .`
