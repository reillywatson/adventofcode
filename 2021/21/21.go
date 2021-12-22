package main

import (
	"fmt"
	"strconv"
	"strings"
	"time"
)

type player struct {
	playerNo int
	pos      int
	score    int
}

func main() {
	var players []player
	for _, l := range strings.Split(input, "\n") {
		var p player
		fmt.Sscanf(l, "Player %d starting position: %d", &p.playerNo, &p.pos)
		players = append(players, p)
	}
	start := time.Now()
	a, b := wins(players[0].pos, players[1].pos, 0, 0, true)
	fmt.Println(a, b, len(cache), time.Since(start), calls)
}

const winCondition = 21

var cache = map[string][2]int{}

var calls = 0

func wins(pos1, pos2, score1, score2 int, turn1 bool) (int, int) {
	calls++
	if score1 >= winCondition {
		return 1, 0
	}
	if score2 >= winCondition {
		return 0, 1
	}
	key := strconv.Itoa(pos1) + "-" + strconv.Itoa(pos2) + "-" + strconv.Itoa(score1) + "-" + strconv.Itoa(score2) + strconv.FormatBool(turn1)
	if cached, ok := cache[key]; ok {
		return cached[0], cached[1]
	}
	var wins1, wins2 int
	for i := 1; i <= 3; i++ {
		for j := 1; j <= 3; j++ {
			for k := 1; k <= 3; k++ {
				if turn1 {
					newPos, newScore := inc(pos1, score1, i+j+k)
					a1, a2 := wins(newPos, pos2, newScore, score2, !turn1)
					wins1 += a1
					wins2 += a2
				} else {
					newPos, newScore := inc(pos2, score2, i+j+k)
					a1, a2 := wins(pos1, newPos, score1, newScore, !turn1)
					wins1 += a1
					wins2 += a2
				}
			}
		}
	}
	cache[key] = [2]int{wins1, wins2}
	return wins1, wins2
}

func inc(pos, score, amt int) (int, int) {
	pos += amt
	for pos > 10 {
		pos -= 10
	}
	score += pos
	return pos, score
}

func main_partone() {
	var players []player
	for _, l := range strings.Split(input, "\n") {
		var p player
		fmt.Sscanf(l, "Player %d starting position: %d", &p.playerNo, &p.pos)
		players = append(players, p)
	}
	die := 0
	turn := 0
	numRolls := 0
	for {
		roll := 0
		for i := 0; i < 3; i++ {
			die++
			if die > 100 {
				die = 1
			}
			roll += die
			numRolls++
		}
		players[turn].pos = players[turn].pos + roll
		for players[turn].pos > 10 {
			players[turn].pos -= 10
		}
		players[turn].score += players[turn].pos
		if players[turn].score >= 1000 {
			break
		}
		turn++
		if turn == len(players) {
			turn = 0
		}
	}

	for _, p := range players {
		if p.score < 1000 {
			fmt.Println(numRolls * p.score)
		}
	}
}

const testInput = `Player 1 starting position: 4
Player 2 starting position: 8`

const input = `Player 1 starting position: 4
Player 2 starting position: 3`
