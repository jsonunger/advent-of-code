package main

import (
	"fmt"
	"sort"

	"github.com/jsonunger/advent-of-code/utils"
)

func part1(lines []string) interface{} {
	current, max := 0, 0
	for _, line := range lines {
		if line == "" {
			if current > max {
				max = current
			}
			current = 0
		} else {
			val := utils.ConvertStringToInt(line)
			current += val
		}
	}
	if current > max {
		max = current
	}
	return max
}

func part2(lines []string) interface{} {
	totals := []int{}
	current := 0
	for _, line := range lines {
		if line == "" {
			totals = append(totals, current)
			current = 0
		} else {
			val := utils.ConvertStringToInt(line)
			current += val
		}
	}
	totals = append(totals, current)

	sort.Sort(sort.Reverse(sort.IntSlice(totals)))

	fmt.Println(totals)

	return utils.SumInts(totals[:3], nil)
}

func main() {
	testCase := utils.ReadLines("./2022/day1/test_case.txt")
	input := utils.ReadLines("./2022/day1/input.txt")

	// PART 1
	fmt.Printf("PART 1 EXAMPLE: %v\n", part1(testCase))
	fmt.Printf("PART 1 RESULT: %v\n", part1(input))

	// PART 2
	fmt.Printf("PART 2 EXAMPLE: %v\n", part2(testCase))
	fmt.Printf("PART 2 RESULT: %v\n", part2(input))
}
