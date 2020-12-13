package main

import (
	"fmt"
	"math/big"
	"strconv"
	"strings"
)

func main() {
	lines := strings.Split(data, "\n")
	var times []int
	for _, t := range strings.Split(lines[1], ",") {
		if t == "x" {
			times = append(times, 1)
			continue
		}
		time, _ := strconv.Atoi(t)
		times = append(times, time)
	}
	var a, n []*big.Int
	for i, t := range times {
		if t != 1 {
			a = append(a, big.NewInt(int64(t-i)))
			n = append(n, big.NewInt(int64(t)))
		}
	}
	answer, err := crt(a, n)
	if err != nil {
		panic(err)
	}
	fmt.Println(answer)
}

// chinese remainder theorem code stolen from https://rosettacode.org/wiki/Chinese_remainder_theorem#Go.
var one = big.NewInt(1)

func crt(a, n []*big.Int) (*big.Int, error) {
	p := new(big.Int).Set(n[0])
	for _, n1 := range n[1:] {
		p.Mul(p, n1)
	}
	var x, q, s, z big.Int
	for i, n1 := range n {
		q.Div(p, n1)
		z.GCD(nil, &s, n1, &q)
		if z.Cmp(one) != 0 {
			return nil, fmt.Errorf("%d not coprime", n1)
		}
		x.Add(&x, s.Mul(a[i], s.Mul(&s, &q)))
	}
	return x.Mod(&x, p), nil
}

func main_partone() {
	lines := strings.Split(data, "\n")
	start, _ := strconv.Atoi(lines[0])
	waitTime := 9999999999 // arbitrary "too big" number
	bus := 0
	for _, t := range strings.Split(lines[1], ",") {
		if t == "x" {
			continue
		}
		time, _ := strconv.Atoi(t)
		wait := time - (start % time)
		if wait < waitTime {
			waitTime = wait
			bus = time
		}
	}
	fmt.Println(waitTime * bus)
}

var testdata = `939
7,13,x,x,59,x,31,19`

var data = `1013728
23,x,x,x,x,x,x,x,x,x,x,x,x,41,x,x,x,x,x,x,x,x,x,733,x,x,x,x,x,x,x,x,x,x,x,x,13,17,x,x,x,x,19,x,x,x,x,x,x,x,x,x,29,x,449,x,x,x,x,x,x,x,x,x,x,x,x,x,x,x,x,x,x,x,x,x,x,x,x,x,x,x,x,x,x,x,x,x,x,x,x,37`
