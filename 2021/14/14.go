package main

import (
	"fmt"
	"math"
	"strings"
)

const in = input
const numSteps = 40

func main() {
	rules := map[string]string{}
	template := strings.Split(in, "\n")[0]
	for _, line := range strings.Split(in, "\n")[2:] {
		var start, end string
		fmt.Sscanf(line, "%s -> %s", &start, &end)
		rules[start] = end
	}
	pairs := map[string]int64{}
	for i := 1; i < len(template); i++ {
		pair := template[i-1 : i+1]
		pairs[pair] = pairs[pair] + 1
	}
	for i := 0; i < numSteps; i++ {
		newPairs := map[string]int64{}
		for pair, count := range pairs {
			if insertion, ok := rules[pair]; ok {
				newPair := string(pair[0]) + insertion
				newPairs[newPair] = newPairs[newPair] + count
				newPair = insertion + string(pair[1])
				newPairs[newPair] = newPairs[newPair] + count
			}
		}
		pairs = newPairs
	}
	counts := map[rune]int64{}
	for pair, count := range pairs {
		for _, c := range pair {
			counts[c] = counts[c] + count
		}
	}
	counts[rune(template[0])] = counts[rune(template[0])] + 1
	counts[rune(template[len(template)-1])] = counts[rune(template[len(template)-1])] + 1
	max := int64(math.MinInt64)
	min := int64(math.MaxInt64)
	for _, v := range counts {
		v = v / 2
		if v > max {
			max = v
		}
		if v < min {
			min = v
		}
	}
	fmt.Println(max - min)
}

func main_partone() {
	rules := map[string]string{}
	template := strings.Split(in, "\n")[0]
	for _, line := range strings.Split(in, "\n")[2:] {
		var start, end string
		fmt.Sscanf(line, "%s -> %s", &start, &end)
		rules[start] = end
	}
	for i := 0; i < numSteps; i++ {
		var newTemplate string
		for i := 1; i < len(template); i++ {
			pair := template[i-1 : i+1]
			if insertion, ok := rules[pair]; ok {
				newTemplate += string(pair[0]) + insertion
			} else {
				newTemplate += string(pair[0])
			}
		}
		newTemplate += string(template[len(template)-1])
		template = newTemplate
	}
	counts := map[rune]int{}
	for _, c := range template {
		counts[c] = counts[c] + 1
	}
	max := math.MinInt
	min := math.MaxInt
	for _, v := range counts {
		if v > max {
			max = v
		}
		if v < min {
			min = v
		}
	}
	fmt.Println(max, min, max-min)
}

const testInput = `NNCB

CH -> B
HH -> N
CB -> H
NH -> C
HB -> C
HC -> B
HN -> C
NN -> C
BH -> H
NC -> B
NB -> B
BN -> B
BB -> N
BC -> B
CC -> N
CN -> C`

const input = `OFSNKKHCBSNKBKFFCVNB

KC -> F
CO -> S
FH -> K
VP -> P
KF -> S
SV -> O
CB -> H
PN -> F
NC -> N
BC -> F
NP -> O
SK -> F
HS -> C
SN -> V
OP -> F
ON -> N
FK -> N
SH -> B
HN -> N
BO -> V
VK -> H
SC -> K
KP -> O
VO -> V
HC -> P
BK -> B
VH -> N
PV -> O
HB -> H
VS -> F
KK -> B
HH -> B
CF -> F
PH -> C
NS -> V
SO -> P
NV -> K
BP -> N
SF -> V
SS -> K
FP -> N
PC -> S
OH -> B
CH -> H
VV -> S
VN -> O
OB -> K
PF -> H
CS -> C
PP -> O
NF -> H
SP -> P
OS -> V
BB -> P
NO -> F
VB -> V
HK -> C
NK -> O
HP -> B
HV -> V
BF -> V
KO -> F
BV -> H
KV -> B
OF -> V
NB -> F
VF -> C
PB -> B
FF -> H
CP -> C
KH -> H
NH -> P
PS -> P
PK -> P
CC -> K
BS -> V
SB -> K
OO -> B
OK -> F
BH -> B
CV -> F
FN -> V
CN -> P
KB -> B
FO -> H
PO -> S
HO -> H
CK -> B
KN -> C
FS -> K
OC -> P
FV -> N
OV -> K
BN -> H
HF -> V
VC -> S
FB -> S
NN -> P
FC -> B
KS -> N`
