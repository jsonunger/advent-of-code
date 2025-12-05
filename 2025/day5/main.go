package main

import (
	"fmt"
	"strings"

	"github.com/jsonunger/advent-of-code/utils"
)

type Range struct {
	Start int
	End   int
}

func (r *Range) Contains(id int) bool {
	return r.Start <= id && r.End >= id
}

type Ranges []Range

func (rs Ranges) AddRange(r Range) Ranges {
	if len(rs) == 0 {
		return append(rs, r)
	}
	// Find the place to insert r so that rs stays sorted by Start
	insertIdx := len(rs)
	for i, existing := range rs {
		if r.Start < existing.Start {
			insertIdx = i
			break
		}
	}
	rs = append(rs, Range{}) // Make space
	copy(rs[insertIdx+1:], rs[insertIdx:])
	rs[insertIdx] = r
	return rs
}

func createRange(line string) Range {
	parts := strings.Split(line, "-")
	start := utils.ConvertStringToInt(parts[0])
	end := utils.ConvertStringToInt(parts[1])
	return Range{Start: start, End: end}
}

func part1(lines []string) interface{} {
	ranges := Ranges{}
	numFresh := 0
	checkID := false
	for _, line := range lines {
		if len(line) == 0 {
			checkID = true
			continue
		}
		if !checkID {
			ranges = ranges.AddRange(createRange(line))
			continue
		}

		for _, r := range ranges {
			if r.Contains(utils.ConvertStringToInt(line)) {
				numFresh++
				break
			}
		}
	}
	return numFresh
}

func (rs Ranges) NumFreshIDs() int {
	if len(rs) == 0 {
		return 0
	}

	ranges := make([]Range, len(rs))
	copy(ranges, rs)

	merged := Ranges{rs[0]}
	for i := 1; i < len(ranges); i++ {
		last := &merged[len(merged)-1]
		if ranges[i].Start <= last.End+1 {
			if ranges[i].End > last.End {
				last.End = ranges[i].End
			}
		} else {
			merged = append(merged, ranges[i])
		}
	}

	total := 0
	for _, r := range merged {
		total += r.End - r.Start + 1
	}
	return total
}

func part2(lines []string) interface{} {
	ranges := Ranges{}
	for _, line := range lines {
		if len(line) == 0 {
			return ranges.NumFreshIDs()
		}

		ranges = ranges.AddRange(createRange(line))
	}
	return nil
}

func main() {
	testCase := utils.ReadLines("./2025/day5/test_case.txt")
	input := utils.ReadLines("./2025/day5/input.txt")

	// PART 1
	fmt.Printf("PART 1 EXAMPLE: %v\n", part1(testCase))
	fmt.Printf("PART 1 RESULT: %v\n", part1(input))

	// // PART 2
	fmt.Printf("PART 2 EXAMPLE: %v\n", part2(testCase))
	fmt.Printf("PART 2 RESULT: %v\n", part2(input))
}
