package main

import (
	"fmt"
	"strconv"
	"strings"
)

var fishCache = map[string]int{}

func numFish(i int, gens int) int {
	key := strconv.Itoa(i) + "_" + strconv.Itoa(gens)
	if cached, ok := fishCache[key]; ok {
		return cached
	}
	var val int
	if gens == 0 {
		val = 1
	} else if i == 0 {
		val = numFish(6, gens-1) + numFish(8, gens-1)
	} else {
		val = numFish(i-1, gens-1)
	}
	fishCache[key] = val
	return val
}

func main() {
	fish := 0
	gens := 256
	for _, n := range strings.Split(input, ",") {
		i, _ := strconv.Atoi(n)
		fish += numFish(i, gens)
	}
	fmt.Println(fish)
}

func main_partone() {
	var fish [569844575390]int8
	numFish := 0
	for _, n := range strings.Split(input, ",") {
		i, _ := strconv.Atoi(n)
		fish[numFish] = int8(i)
		numFish++
	}
	for i := 0; i < 256; i++ {
		fmt.Printf("day %d, num fish: %d\n", i, numFish)
		for i := 0; i < numFish; i++ {
			fish[i] = fish[i] - 1
			if fish[i] < 0 {
				fish[i] = 6
				fish[numFish] = 9
				numFish++
			}
		}
	}
	fmt.Println(numFish)
}

var testInput = `3,4,3,1,2`

var input = `1,1,3,1,3,2,1,3,1,1,3,1,1,2,1,3,1,1,3,5,1,1,1,3,1,2,1,1,1,1,4,4,1,2,1,2,1,1,1,5,3,2,1,5,2,5,3,3,2,2,5,4,1,1,4,4,1,1,1,1,1,1,5,1,2,4,3,2,2,2,2,1,4,1,1,5,1,3,4,4,1,1,3,3,5,5,3,1,3,3,3,1,4,2,2,1,3,4,1,4,3,3,2,3,1,1,1,5,3,1,4,2,2,3,1,3,1,2,3,3,1,4,2,2,4,1,3,1,1,1,1,1,2,1,3,3,1,2,1,1,3,4,1,1,1,1,5,1,1,5,1,1,1,4,1,5,3,1,1,3,2,1,1,3,1,1,1,5,4,3,3,5,1,3,4,3,3,1,4,4,1,2,1,1,2,1,1,1,2,1,1,1,1,1,5,1,1,2,1,5,2,1,1,2,3,2,3,1,3,1,1,1,5,1,1,2,1,1,1,1,3,4,5,3,1,4,1,1,4,1,4,1,1,1,4,5,1,1,1,4,1,3,2,2,1,1,2,3,1,4,3,5,1,5,1,1,4,5,5,1,1,3,3,1,1,1,1,5,5,3,3,2,4,1,1,1,1,1,5,1,1,2,5,5,4,2,4,4,1,1,3,3,1,5,1,1,1,1,1,1`
