package main

import (
	"fmt"
	"math"
	"strings"
)

type coord struct {
	id int
	x  int
	y  int
}

func main_part1() {
	var minX, minY, maxX, maxY, id int
	var coords []coord
	coordMap := map[int]coord{}
	for _, line := range strings.Split(input, "\n") {
		var x, y int
		fmt.Sscanf(line, "%d, %d", &x, &y)
		if x > maxX {
			maxX = x
		}
		if y > maxY {
			maxY = y
		}
		if y < minY || minY == 0 {
			minY = y
		}
		if x < minX || minX == 0 {
			minX = x
		}
		id++
		coords = append(coords, coord{id, x, y})
		coordMap[id] = coord{id, x, y}
	}

	mins := map[int]int{}
	infinite := map[int]bool{}
	for y := minY - 1; y < maxY+1; y++ {
		for x := minX - 1; x < maxX+1; x++ {
			minId := closestId(x, y, coords)
			if minId > -1 {
				mins[minId]++
				if closestId(x+10000, y, coords) == minId || closestId(x-10000, y, coords) == minId || closestId(x, y-10000, coords) == minId || closestId(x, y+10000, coords) == minId {
					infinite[minId] = true
				}
			}
		}
	}
	max := 0
	maxK := 0
	for k, v := range mins {
		if infinite[k] {
			continue
		}
		if v > max && coordMap[k].x != minX && coordMap[k].x != maxX && coordMap[k].y != minY && coordMap[k].y != maxY {
			maxK = k
			max = v
		}
	}
	fmt.Println(maxK, max)
}

func main() {
	var minX, minY, maxX, maxY, id int
	var coords []coord
	coordMap := map[int]coord{}
	for _, line := range strings.Split(input, "\n") {
		var x, y int
		fmt.Sscanf(line, "%d, %d", &x, &y)
		if x > maxX {
			maxX = x
		}
		if y > maxY {
			maxY = y
		}
		if y < minY || minY == 0 {
			minY = y
		}
		if x < minX || minX == 0 {
			minX = x
		}
		id++
		coords = append(coords, coord{id, x, y})
		coordMap[id] = coord{id, x, y}
	}

	regionSize := 0
	maxDist := 10000
	for y := minY - maxDist; y < maxY+maxDist; y++ {
		for x := minX - maxDist; x < maxX+maxDist; x++ {
			totalDist := 0
			for _, coord := range coords {
				dist := int(math.Abs(float64(coord.x-x)) + math.Abs(float64(coord.y-y)))
				totalDist += dist
				if totalDist > maxDist {
					break
				}
			}
			if totalDist < maxDist {
				fmt.Println(x, y, totalDist)
				regionSize++
			}
		}
	}
	fmt.Println(regionSize)
}

func closestId(x, y int, coords []coord) int {
	dists := map[int]int{}
	for _, coord := range coords {
		dists[coord.id] = int(math.Abs(float64(coord.x-x)) + math.Abs(float64(coord.y-y)))
	}
	minDist := 100000
	minId := -1
	for k, v := range dists {
		if v < minDist {
			minDist = v
			minId = k
		} else if v == minDist {
			minId = -1
		}
	}
	return minId
}

var input = `195, 221
132, 132
333, 192
75, 354
162, 227
150, 108
46, 40
209, 92
153, 341
83, 128
256, 295
311, 114
310, 237
99, 240
180, 337
332, 176
212, 183
84, 61
275, 341
155, 89
169, 208
105, 78
151, 318
92, 74
146, 303
184, 224
285, 348
138, 163
216, 61
277, 270
130, 155
297, 102
197, 217
72, 276
299, 89
357, 234
136, 342
346, 221
110, 188
82, 183
271, 210
46, 198
240, 286
128, 95
111, 309
108, 54
258, 305
241, 157
117, 162
96, 301`
