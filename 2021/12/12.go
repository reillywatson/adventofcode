package main

import (
	"fmt"
	"strings"
	"time"
)

func main() {
	var edges [][]string
	for _, line := range strings.Split(input, "\n") {
		nodes := strings.Split(line, "-")
		// add twice so things can be easily bidirectional
		edges = append(edges, []string{nodes[0], nodes[1]})
		edges = append(edges, []string{nodes[1], nodes[0]})
	}
	paths := []string{"start"}
	pathMap := map[string]bool{}
	added := true
	for added {
		start := time.Now()
		added = false
		numPaths := len(paths)
		fmt.Println("num paths:", numPaths)
		for pathNo := 0; pathNo < numPaths; pathNo++ {
			if pathNo > 0 && pathNo%10000 == 0 {
				fmt.Println("path:", pathNo, time.Since(start))
			}
			path := paths[pathNo][:]
			lastIndex := strings.LastIndex(path, "-")
			last := path[lastIndex+1:]
			if last == "end" {
				continue
			}
			for i := 0; i < len(edges); i++ {
				edge := edges[i]
				if edge[0] == last {
					if canVisit(paths, pathMap, path, edge[1]) {
						newPath := path + "-" + edge[1]
						paths = append(paths, newPath)
						pathMap[newPath] = true
						added = true
					}
				}
			}
		}
	}
	_ = added
	numPaths := 0
	for _, path := range paths {
		if strings.HasSuffix(path, "-end") {
			fmt.Println(path)
			numPaths++
		}
	}
	fmt.Println(numPaths)
}

func canVisit(paths []string, pathMap map[string]bool, path string, node string) bool {
	if node == "start" {
		return false
	}
	newPath := path + "-" + node
	if pathMap[newPath] {
		return false
	}
	if node == strings.ToUpper(node) {
		print(path, "uppercase, can visit")
		return true
	}
	lowercases := map[string]bool{}
	anyDoubles := false
	parts := strings.Split(path, "-")
	for _, p := range parts {
		if lowercases[p] {
			anyDoubles = true
			break
		}
		if strings.ToLower(p) == p {
			lowercases[p] = true
		}
	}
	for _, p := range parts {
		if p == node && anyDoubles {
			print(path, "lowercase, already visited", node, "PATH:", path)
			return false
		}
	}
	return true
}

func print(path string, args ...interface{}) {

}

func main_partone() {
	var edges [][]string
	for _, line := range strings.Split(testInput, "\n") {
		nodes := strings.Split(line, "-")
		// add twice so things can be easily bidirectional
		edges = append(edges, []string{nodes[0], nodes[1]})
		edges = append(edges, []string{nodes[1], nodes[0]})
	}
	paths := [][]string{{"start"}}
	added := true
	for added {
		added = false
		numPaths := len(paths)
		for pathNo := 0; pathNo < numPaths; pathNo++ {
			path := paths[pathNo][:]
			last := path[len(path)-1]
			if last == "end" {
				continue
			}
			for i := 0; i < len(edges); i++ {
				edge := edges[i]
				if edge[0] == last {
					if canVisitPartOne(paths, path, edge[1]) {
						var newPath []string
						newPath = append(newPath, path...)
						newPath = append(newPath, edge[1])
						paths = append(paths, newPath)
						added = true
					}
				}
			}
		}
	}
	_ = added
	numPaths := 0
	for _, path := range paths {
		if path[len(path)-1] == "end" {
			numPaths++
		}
	}
	fmt.Println(numPaths)
}

func canVisitPartOne(paths [][]string, path []string, node string) bool {
	var newPath []string
	newPath = append(newPath, path...)
	newPath = append(newPath, node)
	for _, existing := range paths {
		if len(existing) == len(newPath) {
			allSame := true
			for i := 0; i < len(existing); i++ {
				if existing[i] != newPath[i] {
					allSame = false
					break
				}
			}
			if allSame {
				return false
			}
		}
	}
	if node == strings.ToUpper(node) {
		return true
	}
	for _, p := range path {
		if p == node {
			return false
		}
	}
	return true
}

var testInput = `start-A
start-b
A-c
A-b
b-d
A-end
b-end`

var testInput2 = `dc-end
HN-start
start-kj
dc-start
dc-HN
LN-dc
HN-end
kj-sa
kj-HN
kj-dc`

var testInput3 = `fs-end
he-DX
fs-he
start-DX
pj-DX
end-zg
zg-sl
zg-pj
pj-he
RW-he
fs-DX
pj-RW
zg-RW
start-pj
he-WI
zg-he
pj-fs
start-RW`

var input = `RT-start
bp-sq
em-bp
end-em
to-MW
to-VK
RT-bp
start-MW
to-hr
sq-AR
RT-hr
bp-to
hr-VK
st-VK
sq-end
MW-sq
to-RT
em-er
bp-hr
MW-em
st-bp
to-start
em-st
st-end
VK-sq
hr-st`
