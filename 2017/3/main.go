package main

import (
	"fmt"
	"math"
)

/*
spiral first 100 numbers:
100 99  98  97  96  95  94  93  92  91
65  64  63  62  61  60  59  58  57  90
66  37  36  35  34  33  32  31  56  89
67  38  17  16  15  14  13  30  55  88
68  39  18  5   4   3   12  29  54  87
69  40  19  6   1   2   11  28  53  86
70  41  20  7   8   9   10  27  52  85
71  42  21  22  23  24  25  26  51  84
72  43  44  45  46  47  48  49  50  83
73  74  75  76  77  78  79  80  81  82
*/

const (
	RIGHT = iota
	DOWN
	UP
	LEFT
)

func main() {
	goal := 289326
	var minX, minY, maxX, maxY, x, y int
	dir := RIGHT
	for n := 1; n <= goal; n++ {
		switch dir {
		case RIGHT:
			x++
			if x == maxX+1 {
				maxX++
				dir = UP
			}
		case DOWN:
			y++
			if y == maxY+1 {
				maxY++
				dir = RIGHT
			}
		case UP:
			y--
			if y == minY-1 {
				minY--
				dir = LEFT
			}
		case LEFT:
			x--
			if x == minX-1 {
				minX--
				dir = DOWN
			}
		}
		// for part 2, comment out if part 1
		if sumVal(x, y) > goal {
			fmt.Println(sumVal(x, y))
			return
		}
	}
	fmt.Println(math.Abs(float64(x)) + math.Abs(float64(y)) - 1)
}

func val(goalX, goalY int) int {
	var minX, minY, maxX, maxY, x, y int
	dir := RIGHT
	n := 1
	for {
		if x == goalX && y == goalY {
			return n
		}
		switch dir {
		case RIGHT:
			x++
			if x == maxX+1 {
				maxX++
				dir = UP
			}
		case DOWN:
			y++
			if y == maxY+1 {
				maxY++
				dir = RIGHT
			}
		case UP:
			y--
			if y == minY-1 {
				minY--
				dir = LEFT
			}
		case LEFT:
			x--
			if x == minX-1 {
				minX--
				dir = DOWN
			}
		}
		n++
	}
	return -1
}

// sumVal is slow if you don't memoize!
var memoized = map[[2]int]int{}

func sumVal(x, y int) int {
	if x == 0 && y == 0 {
		return 1
	}
	if cached, ok := memoized[[2]int{x, y}]; ok {
		return cached
	}
	sum := 0
	chk := val(x, y)
	adjacents := [][2]int{
		{x - 1, y - 1},
		{x, y - 1},
		{x + 1, y - 1},
		{x - 1, y},
		{x + 1, y},
		{x - 1, y + 1},
		{x, y + 1},
		{x + 1, y + 1},
	}
	for _, adj := range adjacents {
		if val(adj[0], adj[1]) < chk {
			sum += sumVal(adj[0], adj[1])
		}
	}
	memoized[[2]int{x, y}] = sum
	return sum
}
