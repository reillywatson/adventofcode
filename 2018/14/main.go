package main

import (
	"fmt"
	"reflect"
)

func main_part1() {
	elf1 := 0
	elf2 := 1
	recipes := []int{3, 7}
	for {
		sum := recipes[elf1] + recipes[elf2]
		if sum >= 10 {
			recipes = append(recipes, 1)
			sum = sum - 10
		}
		recipes = append(recipes, sum)
		elf1 = (elf1 + recipes[elf1] + 1) % len(recipes)
		elf2 = (elf2 + recipes[elf2] + 1) % len(recipes)
	}
	fmt.Println(recipes[numRecipes:])
}

var numRecipes = 635041

var goal = []int{6, 3, 5, 0, 4, 1}

func main() {
	elf1 := 0
	elf2 := 1
	recipes := []int{3, 7}
	for {
		sum := recipes[elf1] + recipes[elf2]
		if sum >= 10 {
			recipes = append(recipes, 1)
			if len(recipes) > len(goal) && reflect.DeepEqual(goal, recipes[len(recipes)-len(goal):]) {
				fmt.Println(len(recipes) - len(goal))
				return
			}
			sum = sum - 10
		}
		recipes = append(recipes, sum)
		if len(recipes) > len(goal) && reflect.DeepEqual(goal, recipes[len(recipes)-len(goal):]) {
			fmt.Println(len(recipes) - len(goal))
			return
		}
		elf1 = (elf1 + recipes[elf1] + 1) % len(recipes)
		elf2 = (elf2 + recipes[elf2] + 1) % len(recipes)
	}
}
