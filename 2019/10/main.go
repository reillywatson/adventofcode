package main

import (
	"fmt"
	"math"
	"sort"
	"strings"
)

type point struct {
	x int
	y int
}

func main_part1() {
	asteroids := []point{}
	grid := strings.Split(input, "\n")
	for y, row := range grid {
		for x, c := range row {
			if c == '#' {
				asteroids = append(asteroids, point{x: x, y: y})
			}
		}
	}
	maxVisible := 0
	var maxAsteroid point
	for _, a := range asteroids {
		numVisible := 0
		for _, b := range asteroids {
			if a.x == b.x && a.y == b.y {
				continue
			}
			visible := true
			for _, p := range pointsBetween(a, b) {
				if grid[p.y][p.x] == '#' {
					visible = false
					break
				}
			}
			if visible {
				numVisible++
			}
		}
		if numVisible > maxVisible {
			maxVisible = numVisible
			maxAsteroid = a
		}
	}
	fmt.Println(maxAsteroid, maxVisible)
}

func main() {
	asteroids := []point{}
	grid := strings.Split(input, "\n")
	for y, row := range grid {
		for x, c := range row {
			if c == '#' {
				asteroids = append(asteroids, point{x: x, y: y})
			}
		}
	}
	maxVisible := 0
	var maxAsteroid point
	for _, a := range asteroids {
		numVisible := 0
		for _, b := range asteroids {
			if a.x == b.x && a.y == b.y {
				continue
			}
			visible := true
			for _, p := range pointsBetween(a, b) {
				if grid[p.y][p.x] == '#' {
					visible = false
					break
				}
			}
			if visible {
				numVisible++
			}
		}
		if numVisible > maxVisible {
			maxVisible = numVisible
			maxAsteroid = a
		}
	}
	fmt.Println(maxAsteroid, maxVisible)
	var withoutStation []point
	for _, a := range asteroids {
		if a.x != maxAsteroid.x || a.y != maxAsteroid.y {
			withoutStation = append(withoutStation, a)
		}
	}
	asteroids = withoutStation

	numDestroyed := 0
	for len(asteroids) > 0 {
		sort.Slice(asteroids, func(i, j int) bool {
			angleA := angle(maxAsteroid, asteroids[i])
			angleB := angle(maxAsteroid, asteroids[j])
			if angleA != angleB {
				return angleA < angleB
			}
			return distance(maxAsteroid, asteroids[i]) < distance(maxAsteroid, asteroids[j])
		})
		var newList []point
		for i := range asteroids {
			if i > 0 && angle(maxAsteroid, asteroids[i]) == angle(maxAsteroid, asteroids[i-1]) {
				newList = append(newList, asteroids[i])
				continue
			}
			numDestroyed++
			fmt.Println(numDestroyed, asteroids[i].x, asteroids[i].y)
		}
		asteroids = newList
	}
}

func pointsBetween(a, b point) []point {
	var res []point
	startX, startY, endX, endY := a.x, a.y, b.x, b.y
	if startX > endX {
		startX, endX = endX, startX
	}
	if startY > endY {
		startY, endY = endY, startY
	}
	for y := startY; y <= endY; y++ {
		for x := startX; x <= endX; x++ {
			if (x == a.x && y == a.y) || (x == b.x && y == b.y) {
				continue
			}
			if onLine(a, b, point{x: x, y: y}) {
				res = append(res, point{x: x, y: y})
			}
		}
	}
	return res
}

func angle(a, b point) float64 {
	angle := math.Atan2(float64(b.y-a.y), float64(b.x-a.x))
	if angle < -math.Pi/2 {
		angle = angle + (math.Pi * 2)
	}
	return angle
}

func distance(a, b point) float64 {
	return math.Sqrt(math.Pow(float64(a.x-b.x), 2) + math.Pow(float64(a.y-b.y), 2))
}

func onLine(p1, p2, test point) bool {
	a := float64(p2.y - p1.y)
	b := float64(p1.x - p2.x)
	c := a*float64(p1.x) + b*float64(p1.y)
	return a*float64(test.x)+b*float64(test.y) == c
}

var testinput = `.#..##.###...#######
##.############..##.
.#.######.########.#
.###.#######.####.#.
#####.##.#.##.###.##
..#####..#.#########
####################
#.####....###.#.#.##
##.#################
#####.##.###..####..
..######..##.#######
####.##.####...##..#
.#####..#.######.###
##...#.##########...
#.##########.#######
.####.#.###.###.#.##
....##.##.###..#####
.#.#.###########.###
#.#.#.#####.####.###
###.##.####.##.#..##`

var input = `##.#..#..###.####...######
#..#####...###.###..#.###.
..#.#####....####.#.#...##
.##..#.#....##..##.#.#....
#.####...#.###..#.##.#..#.
..#..#.#######.####...#.##
#...####.#...#.#####..#.#.
.#..#.##.#....########..##
......##.####.#.##....####
.##.#....#####.####.#.####
..#.#.#.#....#....##.#....
....#######..#.##.#.##.###
###.#######.#..#########..
###.#.#..#....#..#.##..##.
#####.#..#.#..###.#.##.###
.#####.#####....#..###...#
##.#.......###.##.#.##....
...#.#.#.###.#.#..##..####
#....#####.##.###...####.#
#.##.#.######.##..#####.##
#.###.##..##.##.#.###..###
#.####..######...#...#####
#..#..########.#.#...#..##
.##..#.####....#..#..#....
.###.##..#####...###.#.#.#
.##..######...###..#####.#`
