package main

import (
	"fmt"
)

func main_part1() {
	serialNumber := 5468
	max := -1000000
	maxX := 0
	maxY := 0
	for y := 1; y <= 297; y++ {
		for x := 1; x <= 297; x++ {
			sum := 0
			for i := 0; i < 3; i++ {
				for j := 0; j < 3; j++ {
					sum += power(serialNumber, x+i, y+j)
				}
			}
			if sum > max {
				max = sum
				maxX = x
				maxY = y
			}
		}
	}
	fmt.Println(maxX, maxY, max)
}

func main() {
	serialNumber := 5468
	max := -1000000
	maxX := 0
	maxY := 0
	maxGridSize := 1
	for gridSize := 1; gridSize < 300; gridSize++ {
		fmt.Println(gridSize)
		for y := 1; y <= (300 - gridSize); y++ {
			for x := 1; x <= (300 - gridSize); x++ {
				sum := 0
				for i := 0; i < gridSize; i++ {
					for j := 0; j < gridSize; j++ {
						sum += power(serialNumber, x+i, y+j)
					}
				}
				if sum > max {
					max = sum
					maxX = x
					maxY = y
					maxGridSize = gridSize
				}
			}
		}
	}
	fmt.Println(maxX, maxY, maxGridSize)
}

func power(serialNumber, x, y int) int {
	rackId := x + 10
	power := rackId * y
	power += serialNumber
	power *= rackId
	power = int((power % 1000) / 100)
	power -= 5
	return power
}
