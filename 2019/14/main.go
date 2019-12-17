package main

import (
	"fmt"
	"math"
	"strconv"
	"strings"
)

type reaction struct {
	ins []chem
	out chem
}

type chem struct {
	name   string
	amount int
}

func main_part1() {
	var reactions []reaction
	for _, line := range strings.Split(input, "\n") {
		var r reaction
		ins := strings.Split(strings.Split(line, " => ")[0], ",")
		r.out = parseChem(strings.Split(line, " => ")[1])
		for _, in := range ins {
			r.ins = append(r.ins, parseChem(in))
		}
		reactions = append(reactions, r)
	}
	fmt.Println(oreNeeded(reactions, 1))
}

func main() {
	var reactions []reaction
	for _, line := range strings.Split(input, "\n") {
		var r reaction
		ins := strings.Split(strings.Split(line, " => ")[0], ",")
		r.out = parseChem(strings.Split(line, " => ")[1])
		for _, in := range ins {
			r.ins = append(r.ins, parseChem(in))
		}
		reactions = append(reactions, r)
	}
	l := 0
	h := 10000000
	m := (l + h) / 2
	goal := 1000000000000
	for l < h {
		got := oreNeeded(reactions, m)
		fmt.Println(m, got)
		if got == goal {
			fmt.Println("EXACT!")
			break
		}
		if got < goal {
			fmt.Println("UPDATING L:", l, m)
			l = m
		} else {
			fmt.Println("UPDATING H:", h, m)
			h = m
		}
		m = (l + h) / 2
	}
	fmt.Println(m)
}

func oreNeeded(reactions []reaction, fuel int) int {
	neededOre := 0
	needed := map[string]int{"FUEL": fuel}
	reserves := map[string]int{}
	for len(needed) > 0 {
		for neededName, neededAmount := range needed {
			if reserves[neededName] > 0 {
				amt := min(reserves[neededName], neededAmount)
				needed[neededName] = needed[neededName] - amt
				if needed[neededName] == 0 {
					delete(needed, neededName)
				}
				reserves[neededName] = reserves[neededName] - amt
				break
			}
			for _, r := range reactions {
				if r.out.name == neededName {
					numReactions := int(math.Ceil(float64(neededAmount) / float64(r.out.amount)))
					needed[neededName] = neededAmount - (r.out.amount * numReactions)
					for _, in := range r.ins {
						if in.name == "ORE" {
							neededOre = neededOre + (in.amount * numReactions)
						} else {
							needed[in.name] = needed[in.name] + (in.amount * numReactions)
						}
					}
					if needed[neededName] <= 0 {
						reserves[neededName] = reserves[neededName] - needed[neededName]
						delete(needed, neededName)
						break
					}
				}
			}
		}
	}
	return neededOre
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func parseChem(s string) chem {
	s = strings.TrimSpace(s)
	amt, _ := strconv.Atoi(strings.Split(s, " ")[0])
	return chem{
		name:   strings.Split(s, " ")[1],
		amount: amt,
	}
}

var testinput = `171 ORE => 8 CNZTR
7 ZLQW, 3 BMBT, 9 XCVML, 26 XMNCP, 1 WPTQ, 2 MZWV, 1 RJRHP => 4 PLWSL
114 ORE => 4 BHXH
14 VRPVC => 6 BMBT
6 BHXH, 18 KTJDG, 12 WPTQ, 7 PLWSL, 31 FHTLT, 37 ZDVW => 1 FUEL
6 WPTQ, 2 BMBT, 8 ZLQW, 18 KTJDG, 1 XMNCP, 6 MZWV, 1 RJRHP => 6 FHTLT
15 XDBXC, 2 LTCX, 1 VRPVC => 6 ZLQW
13 WPTQ, 10 LTCX, 3 RJRHP, 14 XMNCP, 2 MZWV, 1 ZLQW => 1 ZDVW
5 BMBT => 4 WPTQ
189 ORE => 9 KTJDG
1 MZWV, 17 XDBXC, 3 XCVML => 2 XMNCP
12 VRPVC, 27 CNZTR => 2 XDBXC
15 KTJDG, 12 BHXH => 5 XCVML
3 BHXH, 2 VRPVC => 7 MZWV
121 ORE => 7 VRPVC
7 XCVML => 6 RJRHP
5 BHXH, 4 VRPVC => 5 LTCX`

var input = `2 MLVWS, 8 LJNWK => 1 TNFQ
1 BWXQJ => 2 BMWK
1 JMGP, 3 WMJW => 9 JQCF
8 BWXQJ, 10 BJWR => 6 QWSLS
3 PLSH, 1 TNFQ => 6 CTPTW
11 GQDJG, 5 BMWK, 1 FZCK => 7 RQCNC
1 VWSRH => 7 PTGXM
104 ORE => 7 VWSRH
1 PTGXM, 13 WMJW, 1 BJGD => 7 KDHF
12 QWSLS, 3 PLSH, 4 HFBPX, 2 DFTH, 11 BCTRK, 4 JPKWB, 4 MKMRC, 3 XQJZQ => 6 BDJK
1 JQCF, 3 CVSC => 2 KRQHC
128 ORE => 7 QLRXZ
32 CXLWB, 18 TZWD => 1 HFQBG
31 KDHF => 9 BWXQJ
21 MLVWS => 9 LJNWK
3 QLRXZ => 5 CXLWB
3 LQWDR, 2 WSDH, 5 JPKWB, 1 RSTQC, 2 BJWR, 1 ZFNR, 16 QWSLS => 4 JTDT
3 BWXQJ, 14 JMGP => 9 MSTS
1 KXMKM, 2 LFCR => 9 DKWLT
6 CVSC => 3 FWQVP
6 XBVH, 1 HFBPX, 2 FZCK => 9 DFTH
9 MSTS => 2 BCTRK
1 PLSH, 28 MSTS => 2 FDKZ
10 XBVH, 5 BJWR, 2 FWQVP => 6 ZFNR
2 CVSC => 6 XBVH
1 BWXQJ, 2 KXMKM => 3 XQJZQ
1 VWSRH, 1 TZWD => 4 WMJW
14 CTPTW, 19 JMGP => 8 GRWK
13 NLGS, 1 PTGXM, 3 HFQBG => 5 BLVK
2 PTGXM => 7 NLGS
123 ORE => 3 DLPZ
2 ZNRPX, 35 DKWLT => 3 WSDH
1 TZWD, 1 BLVK, 9 BWXQJ => 2 MKDQF
2 DLPZ => 2 MLVWS
8 MKDQF, 4 JQCF, 12 VLMQJ => 8 VKCL
1 KRQHC => 7 BJWR
1 GRWK, 2 FWQVP => 9 LFCR
2 MSTS => 2 GQDJG
132 ORE => 9 TZWD
1 FWQVP => 8 RHKZW
43 FDKZ, 11 BJWR, 63 RHKZW, 4 PJCZB, 1 BDJK, 13 RQCNC, 8 JTDT, 3 DKWLT, 13 JPKWB => 1 FUEL
1 LFCR, 5 DFTH => 1 RSTQC
10 GQDJG => 8 KPTF
4 BWXQJ, 1 MKDQF => 7 JMGP
10 FGNPM, 23 DFTH, 2 CXLWB, 6 KPTF, 3 DKWLT, 10 MKDQF, 1 MJSG, 6 RSTQC => 8 PJCZB
8 VWSRH, 1 DLPZ => 7 BJGD
2 BLVK => 9 HBKH
16 LQWDR, 3 MSTS => 9 HFBPX
1 TNFQ, 29 HFQBG, 4 BLVK => 2 KXMKM
11 CVSC => 8 MJSG
3 LFCR => 6 FGNPM
11 HFQBG, 13 MKDQF => 1 FZCK
11 BWXQJ, 1 QLRXZ, 1 TNFQ => 9 KBTWZ
7 XQJZQ, 6 VKCL => 7 LQWDR
1 LJNWK, 4 HBKH => 1 CVSC
4 PLSH, 2 WSDH, 2 KPTF => 5 JPKWB
1 KPTF => 8 MKMRC
5 NLGS, 2 KDHF, 1 KBTWZ => 2 VLMQJ
4 MLVWS, 1 WMJW, 8 LJNWK => 1 PLSH
3 VKCL => 7 ZNRPX`
