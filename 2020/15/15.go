package main

import (
	"fmt"
	"strconv"
	"strings"
)

func main() {
	nums := map[int][]int{}
	var start []int
	for _, num := range strings.Split(data, ",") {
		d, _ := strconv.Atoi(num)
		start = append(start, d)
	}
	var last int
	for i := 0; i < 30000000; i++ {
		var next int
		if i < len(start) {
			next = start[i]
		} else if lasts, ok := nums[last]; ok && len(lasts) > 1 {
			next = lasts[len(lasts)-1] - lasts[len(lasts)-2]
		} else {
			next = 0
		}
		nums[next] = append(nums[next], i)
		fmt.Println(i+1, next)
		last = next
	}
}

var testdata = `0,3,6`

var data = `0,3,1,6,7,5`
