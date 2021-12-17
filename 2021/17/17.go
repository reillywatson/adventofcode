package main

import (
	"fmt"
	"math"
)

func main() {
	var x1, x2, y1, y2 int
	fmt.Sscanf(input, "target area: x=%d..%d, y=%d..%d", &x1, &x2, &y1, &y2)
	fmt.Println(x1, x2, y1, y2)
	maxY := math.MinInt
	numInBox := 0
	for dxtest := -1000; dxtest < 1000; dxtest++ {
		for dytest := -1000; dytest < 1000; dytest++ {
			dx, dy := dxtest, dytest
			x, y := 0, 0
			inBox := false
			maxYStep := math.MinInt
			for {
				if x >= min(x1, x2) && x <= max(x1, x2) && y >= min(y1, y2) && y <= max(y1, y2) {
					inBox = true
				}
				if y < min(y1, y2) && dy < 0 {
					break
				}
				if x > max(x1, x2) && dx > 0 {
					break
				}
				x += dx
				y += dy
				if y > maxYStep {
					maxYStep = y
				}
				if dx < 0 {
					dx++
				}
				if dx > 0 {
					dx--
				}
				dy--
			}
			if inBox {
				numInBox++
				fmt.Println("IN BOX:", dxtest, dytest)
			}
			if inBox && maxYStep > maxY {
				maxY = maxYStep
			}
		}
	}
	fmt.Println(maxY, numInBox)
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

const testInput = `target area: x=20..30, y=-10..-5`

const input = `target area: x=111..161, y=-154..-101`
