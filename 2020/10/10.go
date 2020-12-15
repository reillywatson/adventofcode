package main

import (
	"fmt"
	"sort"
	"strconv"
	"strings"
)

type node struct {
	N        int
	Children []*node
}

func main() {
	var nums []int
	for _, line := range strings.Split(fulldata, "\n") {
		n, _ := strconv.Atoi(line)
		nums = append(nums, n)
	}
	nums = append(nums, 0)
	sort.Ints(nums)
	nums = append(nums, nums[len(nums)-1]+3)

	var nodes []*node
	for _, num := range nums {
		n := &node{N: num}
		for _, parent := range nodes {
			if num-parent.N <= 3 {
				parent.Children = append(parent.Children, n)
			}
		}
		nodes = append(nodes, n)
	}
	fmt.Println(visit(nodes[0], nodes[len(nodes)-1]))
}

var visited = map[*node]int{}

func visit(node, goal *node) int {
	if node == goal {
		return 1
	}
	if _, ok := visited[node]; ok {
		return visited[node]
	}
	childCount := 0
	for _, child := range node.Children {
		childCount += visit(child, goal)
	}
	visited[node] = childCount
	return childCount
}

func main_partone() {
	var nums []int
	for _, line := range strings.Split(fulldata, "\n") {
		n, _ := strconv.Atoi(line)
		nums = append(nums, n)
	}
	nums = append(nums, 0)
	sort.Ints(nums)
	nums = append(nums, nums[len(nums)-1]+3)
	diffs := map[int]int{}
	for i := 1; i < len(nums); i++ {
		diff := nums[i] - nums[i-1]
		diffs[diff] = diffs[diff] + 1
	}
	fmt.Println(diffs[1] * diffs[3])
}

var data = `28
33
18
42
31
14
46
20
48
47
24
23
49
45
19
38
39
11
1
32
25
35
8
17
7
9
4
2
34
10
3`

var fulldata = `44
41
48
17
35
146
73
3
16
159
11
29
32
63
65
62
126
151
6
124
87
115
122
43
12
85
2
98
59
156
149
66
10
82
26
79
56
22
74
49
25
69
54
19
108
18
55
131
140
15
125
37
129
91
51
158
117
136
142
109
64
36
160
150
42
118
101
78
28
105
110
40
157
70
97
139
152
47
104
81
27
116
132
143
1
80
75
141
133
9
50
153
123
111
119
130
112
94
90
86`