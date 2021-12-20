package main

import (
	"fmt"
	"strconv"
	"strings"
)

func main() {
	in := input
	numGens := 50
	algorithm := strings.Split(in, "\n")[0]
	var pixels [][]int
	for _, line := range strings.Split(in, "\n")[2:] {
		var row []int
		for _, c := range line {
			if c == '#' {
				row = append(row, 1)
			} else {
				row = append(row, 0)
			}
		}
		pixels = append(pixels, row)
	}
	infinity := 0
	for i := 0; i < numGens; i++ {
		fmt.Println("-------")
		pixels = grow(pixels, infinity)
		pixels = grow(pixels, infinity)
		var newPixels [][]int
		for y := 0; y < len(pixels); y++ {
			var newRow []int
			for x := 0; x < len(pixels[y]); x++ {
				pos, _ := strconv.ParseInt(get(pixels, x-1, y-1, infinity)+
					get(pixels, x, y-1, infinity)+
					get(pixels, x+1, y-1, infinity)+
					get(pixels, x-1, y, infinity)+
					get(pixels, x, y, infinity)+
					get(pixels, x+1, y, infinity)+
					get(pixels, x-1, y+1, infinity)+
					get(pixels, x, y+1, infinity)+
					get(pixels, x+1, y+1, infinity), 2, 64)

				newPixel := newVal(algorithm, int(pos))
				if newPixel == 1 {
					fmt.Print("#")
				} else {
					fmt.Print(".")
				}
				newRow = append(newRow, newPixel)
			}
			newPixels = append(newPixels, newRow)
			fmt.Println()
		}
		if infinity == 1 {
			infinity = newVal(algorithm, 511)
		} else {
			infinity = newVal(algorithm, 0)
		}
		pixels = newPixels
	}
	litPixels := 0
	for _, row := range pixels {
		for _, c := range row {
			if c == 1 {
				litPixels++
			}
		}
	}
	fmt.Println(litPixels)
}

func newVal(algorithm string, pos int) int {
	if algorithm[pos] == '#' {
		return 1
	}
	return 0
}

func get(pixels [][]int, x, y int, infinity int) string {
	if x < 0 || y < 0 || y >= len(pixels) || x >= len(pixels[0]) {
		return strconv.Itoa(infinity)
	}
	return strconv.Itoa(pixels[y][x])
}

// expand by one in every direction
func grow(pixels [][]int, padVal int) [][]int {
	var res [][]int
	var line []int
	for i := 0; i < len(pixels[0])+2; i++ {
		line = append(line, padVal)
	}
	res = append(res, line)
	for _, l := range pixels {
		l = append([]int{padVal}, l...)
		l = append(l, padVal)
		res = append(res, l)
	}
	res = append(res, line)
	return res
}

const testInput = `..#.#..#####.#.#.#.###.##.....###.##.#..###.####..#####..#....#..#..##..###..######.###...####..#..#####..##..#.#####...##.#.#..#.##..#.#......#.###.######.###.####...#.##.##..#..#..#####.....#.#....###..#.##......#.....#..#..#..##..#...##.######.####.####.#.#...#.......#..#.#.#...####.##.#......#..#...##.#.##..#...##.#.##..###.#......#.#.......#.#.#.####.###.##...#.....####.#..#..#.##.#....##..#.####....##...##..#...#......#.#.......#.......##..####..#...#.#.#...##..#.#..###..#####........#..####......#..#

#..#.
#....
##..#
..#..
..###`

const input = `#.#..##..##....#.####.#...##..###..#.#.##..##....###..#.##.#.#.#......##...#..##...#####.##..##...##..#.##.##..###.##.##...##....###.##.#...#.#.##..#..###.#.##.#.##.####.###.#..#######.##..##..#.##..#####.#..###.####.##....####.#....#...#..#..#....#..#...####.....#.##.###.##.##.###..###.##.###...#.##..#.###.##..##..#.##...##....##.#...#..#...#.##.#..#.###...#.#.##...#..#......#.#...#######.###.##.####.#.#.#.#.#.#.#######....##.##.##..##.##....##....##.##..####..#.#.##...###.##...#..##...#####.#.#.##.#.####.

##.##.#.#..##...####..##...#....#....##.##..##..##..##.#..###.....#########...#..##.###....#...##...
#...#.#####.#.##.#.#####...#.#####..#.##.##..##.....#..#.......#.###..#..#..#####..#.##.##.#...#..##
###...###..###.#.####....#######..#..#..#..###..#.######...##.#..#..#..##.#....#.....####.#...##....
##.#..###...#..#.#######.#.##..####.###.###...###.#...#.#...######......#.#.#...#.##..........#..#.#
.....#...###.##..##.#.########.##.#..#..###..#..##...##.###.#.#..#######..##.###.#..#..##.##.#...###
.#.#.###....###...##.#...##...#.#.#...#.#.##..##.#..#.##.#..####.###...##....###..#......###.......#
.###.#.##.#.#####..##.###.##....#.##.####...#.#..###.##..####...#..#####.......#.###.###.##....#..#.
##...#.##..##...###.#......###....####..###.......###..####.##.####.##.....#.#.##.###.######..#.....
###..#...#.#.#...#..#...##..###.#..###.####.#..#..##..##.###.####..##...##..###........#...###..##..
...#...#..#..#..######..##..##.#..##...#.#.#...##.#.##.####..#..##.##..#.###.#.#.##.#.##..###.###..#
#...#...#..#######...##.#....##.##...##....#.##....#.#.....#.#..#..#.#..##.#........##.#.##.....#.#.
###.#.....#######.#...####..##...#.#....####.#....#.######.#..###......#.#....##..####.##..#.#.##.##
..#..#...##.##.##...#...####....##.#.####...##...#...#..#....#.#.#####....#.#.##......##..#..#...###
###..###..#.##...#.####.#..#..#...#.##.######.#.##.##.#######.#...#.####.###...#.#............##..##
#.####...#..#.##.#....######.####.##........#.#..#...#..##.....#..#...#..###...###.#..#.#.##..##.#.#
##.##..##.###.#...#.##.....#.#..##...####...#.#..###..###...#....######..##...#.##...#...#...#...#.#
##.#.....##...#.##..#...#.####....#....##.##.#.#.######..##.#.##.####.#.#..#.#.##.#.#...#.....###.##
..##...###..###.......#.....#.#..#.#######..##....##...##...#.#.#.#...##.####.....#.##..###..##.....
#.#..######.####.####.##.##.###.####..#....####..#.#######......##...#...####.##.#######..###..####.
##.###...##.##....######.#.##.#...###.##.##.#......###.#..###...#...........#.......###.####.##..#..
##.#......#.#.....##.##.###.#####.######..#.#...##.#.#.#..#.#...#..#.#.#.####.##.###.##.######.#.#..
.#..#.##.###.####.##.##..##...####..#.##.##.##...##...##.####.#####.#..##.#...##..#..##.#..#.#..#.##
#..###.#.#######.###.......#####.#.#.###.#.##.##.####.#####...#..#....#.#....#...##......#.#..#....#
.#..###.....##..###...##...####..#..##.##.#..#.#.#..####..#.......##.##...#.###..#...#....#........#
#.#.#.#...#.###.........###....#...##.##.##..#..#...####...##.###...#.##.##..##....##.#...####.#.###
..#.#..#.#.......#.#....####.##.#.#........#.#.##..###.#..#.##..##..##.###.#....#.###.#.##.####.#..#
....#.###..#.#######...###.#..#..##.#.#....##.#.#.#.....###.#..###......#....#######......#...#.#.#.
#.#.#.#...#.#...#######...###.###..##.#...####...#####...#..##......#..##.#........#.#..##.#.###....
#...##.#..##..#.#..#.#.....######..##...#.##.#.##.##.###..#.##.#..##.###.####...###...#.###..##.#.#.
#.#.##..#..#..##.....#####........#.##..#.#.#.#.#.#.#.#..#####.#.#.....###..#.###..#.##.##.#######.#
.##.#.#.##.##.#..#.###..#..##..###..#..#..#.###...###.#.#....##..#.##...##..#..##.#.#.#.#.#####.##.#
#.#.......###.#...##.##.##.......##.###.#..###.#..#....#.#.####....#.###.##...##.#...##...#..###.##.
#..###.#.##.#..#####..##.#.########.#.###.####....##.###.##.#..######...#.#..#..###.#...#.#...##...#
.#.#.#..#.##..####...#.##.#..#.#........##.#.#....##.##.#.....###.#.#.#...#.#####..#.#...##.#..####.
#.##...########.##.......##.#.####.#...#.##.#######..#...........#.#.###..###...#.####.##.#####..###
...#..##.###.#.#.#...###..###....#..#.##..##.####..#.###..##...#..##.....###.##.##.#.####.####..#.#.
.#.#.#.....##.#.###...###..#.#...###.#.....##....#.#.###..#....##.....#.#...#.#.##..#..#...#.#.#.#.#
#.#.#.#######.#.##..######...##########.#.##.#.###.##.#..#..#....##......#...##.##..###.##.####.#...
#..##..##.##....#..###..##.....##....##.#.#...#....##...#..#.#..#...#..####....#.##.....#.##.#......
....##..##.#.#........#.#.#...##.##......####.....#.####.#.##....#....#.###.#....#.####.####....#.##
...###.#..#..###.####...#.#...#####..#...##..#.#..#..#.####....#...#####.###.##.#####....###.##.####
###..##.....##.#.#.#.#.#.#.#.###..##..#..........###..#.###.####...##..#.#.#.#.##.#..###....#.####..
...#...#..#.............###.##..##.###.#.######...#.###.##..##.#######...###.###.#..##.####.#..#.###
#..##.##...###..#####.####..##..#.#.#..##..##..#.....#########.##.#.###...##......#...#....#.###...#
#..##.#.###..#...###.#.#...#..###.###..##.##.#.#.#....###.####.#...###..##.#...#.##.###.#..##....###
#.#...####.##.####.#.##...##.##..#.##.#.#..##..##.....#.#.#####.#.##..##..#.##.####.#..##..#.##.##.#
..##.##...###.#.......##.......##.#..##.#.#.#.#..####.##..##..#####.#..#..#.####.#.##.#.###..#...###
##.##.###...#......##.######.#...#....#.#####.##.#.####.##.##.##.#.#...######....##.##.###..##.###..
##.#########..#.##..#..####..#.####.#.#.##..#.##.##.###.####..##..#..#.####..#.#.....#.....#.###.###
...###.#.##.##.##.##...#.#..#.####..###.....#...##.###.###..##.#.#..##.######.....#.#.###.##.##..###
#..#####.#.#....##.#.##..###.#..###...#...####.#####..#.##..#.##..###....#..##.##.###...#.#...#.....
#...#.#..#.######..##......#.####.###....##.#.##..#..#####....#..###.#.#####...###.#..##.##.....#...
##..##.###..#.#..##.##...#######..#...#..#....#.##...##.##.###.#..##.#..##...#.###.....##.#.#.#.####
##...##.###..#...#..#..#...#..##.##...###..##..#..####...#.##......##.##...#.#######.####.###....##.
...#..##.###...##....#######.###...##..##.##.#.##...#..###.#.##.#..#.#.##..###.##....#.....#.#.#.#..
..#..##..#..##.##.#.##..##.####...#.#####..####.#.##.#.#..#......##.#.##.....#.#..##..#.###.#...###.
#.....#..#...#.#.#.##########....##.####...###......#.#..#..#........#.###..#.####.####.#####.##.###
.##..#.###.##..##..######.##..###...#.##...#....#....####.###.#.###...#.###..#......#.##.###....#...
.#..#.##.##..##.###.#..#.#.##.#..##.###.###.#....#.######...#.....###.##..###.####..#.##..#.####.###
#.##.##.#.#.........#.....#.##..#.##..#.##..#.#..##.##.#..##..##########....#.##.#...#.####.#.#.#.##
.#.....###.##.#..##...#.#..##.##.#.#.#..####..#####......###..#####..###...###...##.##..#.#####..#.#
.##..#.###..#.#.###..###..##.#...###.####.######......#..##.#.####.##.#.#....#......#..#..###.##..##
###..##....#.#######.###.###.##.###..##...##.#..#.#..#....#.#.####..##..##.#####.#.#####.#...##..#.#
.##.....##..##.#.#.#..###.#..###.##.##..###.##.####...#.##..#..#.#.#.##.###.#..###....#.###.##.#.#..
####.##.###.#.##..#.###.##.#.....#.###.#.####..#..#.####.....###.#.####.#..#...###....###...##.#.#.#
#.##..#.......###.##.###.##..#..####.###.#..#####.#..##.###.#..###.####....#..#...#.####.#.##....#..
#....#.#....#.#..###.##..###..##..#####.##...#.#....###..##.#....#####.#....###.#..##.#.....####.#.#
##.###.#.##.#.....###.#.......##.###....#.#.##....##.#.##.#.#...#.#.##......#.#.##.#...#...#......##
.##.#..#.###.###..#..##.#..##.###.....#....#####..#.#..#####....####...##.##.#.####.#....##.#####.#.
#...##....#####.#.##.##.#####.#.#.##....#.#.##.#....#.....###.#....##..###.####..#..#...#...#...#..#
..#..##....##.#.##.#.###..####.#..##....#.##.#..#######...#.##.#....##...###..#..#.#.#..#.##.#.#..#.
##..###....#...####.##.###.####.###...##..####..###..##.#.##.....#####...#.#.##...#.#..######.#....#
.#...######..........#......#.##.#...##..#####...###..###..#...##.#..##.##..#..##.....###.##.##.###.
.#..##.###.##...#..#.####..#....#.#.#.#.####.##..#.#..#..###..##.#..##..###......##..#.#.##.#.##.###
.#.#.#..#####..#...#######..####..##.####.#.#####.#...####.####.##.##.#.#...#.....###.##..##.#.##...
.#...######...###.##.#...#...#...##.#..#...#.#..#.#.##.#.#.#...###.####.##..##.##..###..#.#..#.#.#.#
.###.#...#.##.###.########......#.###.###.#.###.....###.#..####..#..#.#.#####....#.#...#.#....###..#
#.#.#...#....#####.##.#.###.##.#..#####.......#..#.##..#.#....###.#..#.##..#..#.#..######...#...###.
.####.#.#.#..####.#####..###.#.##.##..#######.#.#.###..#.##.###.#....##.####.#.#.##.##.#.###....#..#
#.#.##..#..##.####.#.#..####..#..#####.########..###.#.###....##...##.####..#......##.#.#.#.#..#...#
....#.....###..##...##.##..#####.##..###.#...##....##.#####..####.....####..#..####..###.##....#.###
.....###..#......#....#.###....#.###.#.#.#.#.###.##.##...###.##.#####...#....#.##..##..##..#.#..##..
#.##..#.#..#..##..#..###...#..#.######.##.##.#.####...##.#..#..##...##.#.###..###.####.##...#..#....
###...###..#..#####..###.#####....##..#..#...#.###.##..##..########.##....#...#..##......##..#.##...
..#.#..###..#......#.#...####.....#.#...######..#.....#..###.##..##.#...##.#..##.##..##..##.#####.##
#.#..#......###..##.#####..#####...##.#.#####.######.#..##....##.##.#.#..#....#.#...##.#.###.#.##.#.
.####...##.####....##..####.#..####..#.....###.###..#####..###.##.#.####.##.#..#.#.#..#..##..##.#.##
#.####..#.###.#.#...#.###.##..#......######.#...#####..#.###..#..#####.#.#.##.#..#..##...#.....##.##
..#..##......#.#.##.####...#.######..####..##.###.##..#.#..#.########...#.#...###.##..###.#.####..##
..#......#..#..#.#.#.#####.#..###...#.#.#...#.##.##.##.##.#.##.#####.####.#..##.######..#.##..#...##
###.##.#..#.##.#..###..##.##.#..#.####..#....####.##.#.##...#.##......##.#...#....###.#.....####.###
..########.##..##.###..##.###.##..#.#.##.#.#....####.##.####.##.###....#.###.########.#.##.#.#######
..#.#.....#..#....#####.#.#..#...##.#.....#.....#..###.##..##.####..#.##...#####..####.....#..#..###
##...###.######.##.#..#....#.##.#..#..#.#####..#.######..#..##...##.#..#...##.#.#..#.##.#######...##
....#..#.#.#####.###.##.##.#.##..#.#....#.##.###...###.##..#.####..##.#...####.#....#....#.##...##..
.####....####.##.#..#..####.#..#..##....##.#..##.##..##.###....##..##.........##.#.####......####.#.
...#.#.#..####...#.#.#.......#..##.#.#..#.#..##..####..#....###....#..###.##.#..#.###.###.###.##.#..
##.#..#.#.####.###.#.#.###.....###...#.###.#.##.#.#..#..#.#..###.....###.#.#...########...#..#.#..##
..#..##..#...#.#..##.#.#...#..#..#..###.##.##.#.#.#......#..###.#.....##..#.#.#.#.#...###.#..#.#.##.
#.####.#..###.###.#....#.##..##.#.##..#.####....#.##.#...##.....#####..#.#..##..#.##..#......##..##.`
