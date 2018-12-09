package main

import (
	"container/list"
	"fmt"
	"sort"
)

// uses slices. Too slow for part 2!
func main_part1() {
	numPlayers := 458
	lastMarble := 72019

	circle := make([]int, 0, lastMarble)
	current := 0
	player := 0
	scores := make([]int, numPlayers, numPlayers)
	for marbleNo := 0; marbleNo <= lastMarble; marbleNo++ {
		player++
		if player == len(scores) {
			player = 0
		}
		if marbleNo > 0 && marbleNo%23 == 0 {
			scores[player] += marbleNo
			current = advance(current, len(circle), -7)
			scores[player] += circle[current]
			circle = append(circle[:current], circle[current+1:]...)
			continue
		}
		newLoc := advance(current, len(circle), 2)
		circle = append(circle[:newLoc], append([]int{marbleNo}, circle[newLoc:]...)...)
		current = newLoc
		if len(circle)%1000 == 0 {
			fmt.Println(len(circle))
		}
	}
	sort.Ints(scores)
	fmt.Println(scores)
}

// uses linked lists. Much better!
func main() {
	numPlayers := 458
	lastMarble := 7201900

	circle := list.New()
	player := 0
	var current *list.Element
	scores := make([]int, numPlayers, numPlayers)
	for marbleNo := 0; marbleNo <= lastMarble; marbleNo++ {
		player++
		if player == len(scores) {
			player = 0
		}
		if marbleNo > 0 && marbleNo%23 == 0 {
			scores[player] += marbleNo
			for i := 0; i < 7; i++ {
				current = current.Prev()
				if current == nil {
					current = circle.Back()
				}
			}
			scores[player] += current.Value.(int)
			toRemove := current
			current = current.Next()
			if current == nil {
				current = circle.Front()
			}
			circle.Remove(toRemove)
		} else if current == nil {
			current = circle.PushBack(marbleNo)
		} else {
			current = current.Next()
			if current == nil {
				current = circle.Front()
			}
			current = circle.InsertAfter(marbleNo, current)
		}
	}
	sort.Ints(scores)
	fmt.Println(scores)
}

func advance(n int, length int, num int) int {
	if length == 0 {
		return 0
	}
	for ; num < 0; num++ {
		n--
		if n < 0 {
			n = length - 1
		}
	}
	for ; num > 0; num-- {
		n++
		if n >= length {
			n = 0
		}
	}
	return n
}
