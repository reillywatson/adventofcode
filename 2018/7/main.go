package main

import (
	"fmt"
	"sort"
	"strings"
)

func main_part1() {
	type rule struct {
		step string
		req  string
	}
	var rules []rule
	var result string
	incomplete := map[string]bool{}
	for _, line := range strings.Split(input, "\n") {
		r := rule{step: strings.Split(line, " ")[7], req: strings.Split(line, " ")[1]}
		fmt.Println(r)
		rules = append(rules, r)
		incomplete[r.step] = true
		incomplete[r.req] = true
	}
	for len(incomplete) > 0 {
		var potential []string
		for step := range incomplete {
			valid := true
			for _, rule := range rules {
				if rule.step == step {
					if incomplete[rule.req] {
						valid = false
						break
					}
				}
			}
			if valid {
				potential = append(potential, step)
			}
		}
		if len(potential) == 0 {
			panic("no solution")
		}
		sort.Strings(potential)
		result = result + potential[0]
		delete(incomplete, potential[0])
	}
	fmt.Println(result)
}

var numWorkers = 5
var minTime = 60

func main() {
	type rule struct {
		step string
		req  string
	}
	var rules []rule
	var result string
	incomplete := map[string]bool{}
	complete := map[string]bool{}
	for _, line := range strings.Split(input, "\n") {
		r := rule{step: strings.Split(line, " ")[7], req: strings.Split(line, " ")[1]}
		rules = append(rules, r)
		incomplete[r.step] = true
		incomplete[r.req] = true
	}
	type worker struct {
		remainingTime int
		step          string
	}
	workers := []*worker{}
	for i := 0; i < numWorkers; i++ {
		workers = append(workers, &worker{})
	}
	time := 0
	for ; len(incomplete) > 0; time++ {
		idle := map[*worker]bool{}
		for _, worker := range workers {
			if worker.remainingTime > 0 {
				worker.remainingTime--
			}
			if worker.remainingTime <= 0 {
				idle[worker] = true
				complete[worker.step] = true
				worker.step = ""
			}
		}
		for worker := range idle {
			if len(incomplete) == 0 {
				break
			}
			var potential []string
			for step := range incomplete {
				valid := true
				for _, rule := range rules {
					if rule.step == step {
						if !complete[rule.req] {
							valid = false
							break
						}
					}
				}
				if valid {
					potential = append(potential, step)
				}
			}
			if len(potential) == 0 {
				break
			}
			sort.Strings(potential)
			result = result + potential[0]
			delete(incomplete, potential[0])
			worker.step = potential[0]
			worker.remainingTime = int(potential[0][0]) - int('A') + minTime + 1
		}
		fmt.Println(workers[0].step, workers[1].step)
	}
	maxRemaining := 0
	for _, worker := range workers {
		fmt.Println(worker.remainingTime, worker.step)
		if worker.remainingTime > maxRemaining {
			maxRemaining = worker.remainingTime
		}
	}
	time += maxRemaining
	fmt.Println(time - 1)
}

var input = `Step V must be finished before step H can begin.
Step U must be finished before step R can begin.
Step E must be finished before step D can begin.
Step B must be finished before step R can begin.
Step W must be finished before step X can begin.
Step A must be finished before step P can begin.
Step T must be finished before step L can begin.
Step F must be finished before step C can begin.
Step P must be finished before step Y can begin.
Step N must be finished before step G can begin.
Step R must be finished before step S can begin.
Step D must be finished before step C can begin.
Step O must be finished before step K can begin.
Step L must be finished before step J can begin.
Step J must be finished before step H can begin.
Step M must be finished before step I can begin.
Step G must be finished before step K can begin.
Step Z must be finished before step Q can begin.
Step X must be finished before step Q can begin.
Step H must be finished before step I can begin.
Step K must be finished before step Y can begin.
Step Q must be finished before step S can begin.
Step I must be finished before step Y can begin.
Step S must be finished before step Y can begin.
Step C must be finished before step Y can begin.
Step T must be finished before step S can begin.
Step P must be finished before step S can begin.
Step I must be finished before step S can begin.
Step V must be finished before step O can begin.
Step O must be finished before step Q can begin.
Step T must be finished before step R can begin.
Step E must be finished before step J can begin.
Step F must be finished before step S can begin.
Step O must be finished before step H can begin.
Step Z must be finished before step S can begin.
Step D must be finished before step Z can begin.
Step F must be finished before step K can begin.
Step W must be finished before step P can begin.
Step G must be finished before step I can begin.
Step B must be finished before step T can begin.
Step G must be finished before step Y can begin.
Step X must be finished before step S can begin.
Step B must be finished before step K can begin.
Step V must be finished before step A can begin.
Step U must be finished before step N can begin.
Step T must be finished before step P can begin.
Step V must be finished before step D can begin.
Step G must be finished before step X can begin.
Step B must be finished before step D can begin.
Step R must be finished before step J can begin.
Step M must be finished before step Z can begin.
Step U must be finished before step Z can begin.
Step U must be finished before step G can begin.
Step A must be finished before step C can begin.
Step H must be finished before step Q can begin.
Step X must be finished before step K can begin.
Step B must be finished before step S can begin.
Step Q must be finished before step C can begin.
Step Q must be finished before step Y can begin.
Step R must be finished before step I can begin.
Step V must be finished before step Q can begin.
Step A must be finished before step D can begin.
Step D must be finished before step S can begin.
Step K must be finished before step S can begin.
Step G must be finished before step C can begin.
Step D must be finished before step O can begin.
Step R must be finished before step H can begin.
Step K must be finished before step Q can begin.
Step W must be finished before step R can begin.
Step H must be finished before step Y can begin.
Step P must be finished before step J can begin.
Step N must be finished before step Z can begin.
Step J must be finished before step K can begin.
Step W must be finished before step M can begin.
Step A must be finished before step Z can begin.
Step V must be finished before step W can begin.
Step J must be finished before step X can begin.
Step U must be finished before step F can begin.
Step P must be finished before step L can begin.
Step W must be finished before step G can begin.
Step T must be finished before step F can begin.
Step R must be finished before step C can begin.
Step R must be finished before step O can begin.
Step Z must be finished before step C can begin.
Step E must be finished before step S can begin.
Step L must be finished before step I can begin.
Step U must be finished before step O can begin.
Step W must be finished before step K can begin.
Step K must be finished before step I can begin.
Step O must be finished before step M can begin.
Step V must be finished before step M can begin.
Step V must be finished before step Z can begin.
Step A must be finished before step I can begin.
Step F must be finished before step J can begin.
Step F must be finished before step O can begin.
Step M must be finished before step C can begin.
Step Q must be finished before step I can begin.
Step H must be finished before step S can begin.
Step U must be finished before step A can begin.
Step J must be finished before step S can begin.
Step P must be finished before step Z can begin.`
