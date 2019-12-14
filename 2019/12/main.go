package main

import "fmt"

type moon struct {
	x       int
	y       int
	z       int
	dx      int
	dy      int
	dz      int
	xperiod int
	yperiod int
	zperiod int
}

func main() {
	moons := []*moon{
		/*
			{x: -1, y: 0, z: 2},
			{x: 2, y: -10, z: -7},
			{x: 4, y: -8, z: 8},
			{x: 3, y: 5, z: -1},*/

		{x: -8, y: -10, z: 0},
		{x: 5, y: 5, z: 10},
		{x: 2, y: -7, z: 3},
		{x: 9, y: -8, z: -3},
		/*
			{x: 17, y: 5, z: 1},
			{x: -2, y: -8, z: 8},
			{x: 7, y: -6, z: 14},
			{x: 1, y: -10, z: 4},*/
	}
	init := []moon{
		*moons[0],
		*moons[1],
		*moons[2],
		*moons[3],
	}
	numWithPeriods := 0
	for iters := 0; numWithPeriods < len(moons); iters++ {
		numWithPeriods = 0
		for i, a := range moons {
			if a.xperiod == 0 && a.x == init[i].x && a.dx == init[i].dx {
				a.xperiod = iters
			}
			if a.yperiod == 0 && a.y == init[i].y && a.dy == init[i].dy {
				a.yperiod = iters
			}
			if a.zperiod == 0 && a.z == init[i].z && a.dz == init[i].dz {
				a.zperiod = iters
			}
			if a.xperiod > 0 && a.yperiod > 0 && a.zperiod > 0 {
				numWithPeriods++
			}

			for j, b := range moons {
				if i == j {
					continue
				}
				if b.x > a.x {
					a.dx++
				}
				if b.x < a.x {
					a.dx--
				}
				if b.y > a.y {
					a.dy++
				}
				if b.y < a.y {
					a.dy--
				}
				if b.z > a.z {
					a.dz++
				}
				if b.z < a.z {
					a.dz--
				}
			}
		}
		for _, a := range moons {
			a.x += a.dx
			a.y += a.dy
			a.z += a.dz
		}
	}
	periods := []int64{}
	for _, a := range moons {
		periods = append(periods, int64(a.xperiod), int64(a.yperiod), int64(a.zperiod))
		fmt.Printf("pos=<x=%2d, y=%2d,z=%2d>, vel=<x=%2d,y=%2d,z=%2d>\n", a.x, a.y, a.z, a.dx, a.dy, a.dz)
	}
	fmt.Println(periods)
	fmt.Println(lcm(periods...))
}

func lcm(nums ...int64) int64 {
	fmt.Println(nums)
	if len(nums) == 1 {
		return nums[0]
	}
	if len(nums) > 2 {
		return lcm(nums[0], lcm(nums[1:]...))
	}
	return nums[0] / gcd(nums[0], nums[1]) * nums[1]
}

func gcd(a, b int64) int64 {
	for b != 0 {
		tmp := b
		b = a % b
		a = tmp
	}
	return a
}

func main_partone() {
	moons := []*moon{
		{x: 17, y: 5, z: 1},
		{x: -2, y: -8, z: 8},
		{x: 7, y: -6, z: 14},
		{x: 1, y: -10, z: 4},
	}
	for step := 0; step < 1000; step++ {
		for i, a := range moons {
			for j, b := range moons {
				if i == j {
					continue
				}
				if b.x > a.x {
					a.dx++
				}
				if b.x < a.x {
					a.dx--
				}
				if b.y > a.y {
					a.dy++
				}
				if b.y < a.y {
					a.dy--
				}
				if b.z > a.z {
					a.dz++
				}
				if b.z < a.z {
					a.dz--
				}
			}
		}
		for _, a := range moons {
			a.x += a.dx
			a.y += a.dy
			a.z += a.dz
		}
	}
	energy := 0
	for _, a := range moons {
		fmt.Printf("pos=<x=%2d, y=%2d,z=%2d>, vel=<x=%2d,y=%2d,z=%2d>\n", a.x, a.y, a.z, a.dx, a.dy, a.dz)
		energy += (abs(a.x) + abs(a.y) + abs(a.z)) * (abs(a.dx) + abs(a.dy) + abs(a.dz))
	}
	fmt.Println("total energy:", energy)
}

func abs(a int) int {
	if a < 0 {
		return -a
	}
	return a
}
