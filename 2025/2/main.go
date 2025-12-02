package main

import (
	"fmt"
	"math"
	"strconv"
	"strings"
)

func main() {
	invalids := []int{}
	for _, entry := range strings.Split(input, ",") {
		parts := strings.Split(entry, "-")
		start, _ := strconv.ParseInt(parts[0], 10, 64)
		end, _ := strconv.ParseInt(parts[1], 10, 64)
		for i := start; i <= end; i++ {
			asStr := strconv.FormatInt(i, 10)
			for sub := 1; sub < len(asStr)/2+1; sub++ {
				prefix := asStr[:sub]
				if strings.Repeat(prefix, int(math.Ceil(float64(len(asStr))/float64(len(prefix))))) == asStr {
					invalids = append(invalids, int(i))
					break
				}
			}
		}
	}
	sum := 0
	for _, s := range invalids {
		sum += s
	}
	fmt.Println(sum)
}

func main_partone() {
	invalids := []int{}
	for _, entry := range strings.Split(input, ",") {
		parts := strings.Split(entry, "-")
		start, _ := strconv.ParseInt(parts[0], 10, 64)
		end, _ := strconv.ParseInt(parts[1], 10, 64)
		for i := start; i <= end; i++ {
			asStr := strconv.FormatInt(i, 10)
			midpoint := len(asStr) / 2
			firstHalf := asStr[:midpoint]
			secondHalf := asStr[midpoint:]
			if firstHalf == secondHalf {
				invalids = append(invalids, int(i))
			}
		}
	}
	sum := 0
	for _, s := range invalids {
		sum += s
	}
	fmt.Println(sum)
}

var sampleInput = `11-22,95-115,998-1012,1188511880-1188511890,222220-222224,1698522-1698528,446443-446449,38593856-38593862,565653-565659,824824821-824824827,2121212118-2121212124`

var input = `328412-412772,1610-2974,163-270,7693600637-7693779967,352-586,65728-111612,734895-926350,68-130,183511-264058,8181752851-8181892713,32291-63049,6658-12472,720-1326,21836182-21869091,983931-1016370,467936-607122,31-48,6549987-6603447,8282771161-8282886238,7659673-7828029,2-18,7549306131-7549468715,3177-5305,20522-31608,763697750-763835073,5252512393-5252544612,6622957-6731483,9786096-9876355,53488585-53570896`
