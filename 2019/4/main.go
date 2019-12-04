package main

import (
	"fmt"
	"strconv"
)

func main_part1() {
	res := 0
	for i := start; i < end; i++ {
		digits := strconv.Itoa(i)
		if len(digits) != 6 {
			continue
		}
		samePair := false
		increasing := true
		for j := 1; j < len(digits); j++ {
			if digits[j] == digits[j-1] {
				samePair = true
			}
			if digits[j] < digits[j-1] {
				increasing = false
			}
		}
		if samePair && increasing {
			res++
		}
	}
	fmt.Println(res)
}

func main() {
	res := 0
	for i := start; i < end; i++ {
		digits := strconv.Itoa(i)
		if len(digits) != 6 {
			continue
		}
		samePair := false
		increasing := true
		for j := 1; j < len(digits); j++ {
			if digits[j] == digits[j-1] {
				if (j == len(digits)-1 || digits[j+1] != digits[j]) && (j < 2 || digits[j-2] != digits[j]) {
					samePair = true
				}
			}
			if digits[j] < digits[j-1] {
				increasing = false
			}
		}
		if samePair && increasing {
			res++
		}
	}
	fmt.Println(res)
}

var start = 109165
var end = 576723
