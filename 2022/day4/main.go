package main

import (
	"fmt"
	"strings"

	"github.com/jsonunger/advent-of-code/utils"
)

func part1(lines []string) interface{} {
	containedPairs := 0
	for _, line := range lines {
		parts := strings.Split(line, ",")
		firstRange := utils.ConvertStringsToInts(strings.Split(parts[0], "-"))
		secondRange := utils.ConvertStringsToInts(strings.Split(parts[1], "-"))

		if firstRange[0] < secondRange[0] {
			if firstRange[1] >= secondRange[1] {
				containedPairs++
			}
		} else if firstRange[0] > secondRange[0] {
			if secondRange[1] >= firstRange[1] {
				containedPairs++
			}
		} else {
			containedPairs++
		}
	}
	return containedPairs
}

func part2(lines []string) interface{} {
	overlappingPairs := 0
	for _, line := range lines {
		parts := strings.Split(line, ",")
		firstRange := utils.ConvertStringsToInts(strings.Split(parts[0], "-"))
		secondRange := utils.ConvertStringsToInts(strings.Split(parts[1], "-"))

		if firstRange[0] < secondRange[0] {
			if firstRange[1] >= secondRange[0] {
				overlappingPairs++
			}
		} else if firstRange[0] > secondRange[0] {
			if secondRange[1] >= firstRange[0] {
				overlappingPairs++
			}
		} else {
			overlappingPairs++
		}
	}
	return overlappingPairs
}

func main() {
	testCase := utils.ReadLines("./2022/day4/test_case.txt")
	input := utils.ReadLines("./2022/day4/input.txt")

	// PART 1
	fmt.Printf("PART 1 EXAMPLE: %v\n", part1(testCase))
	fmt.Printf("PART 1 RESULT: %v\n", part1(input))

	// // PART 2
	fmt.Printf("PART 2 EXAMPLE: %v\n", part2(testCase))
	fmt.Printf("PART 2 RESULT: %v\n", part2(input))
}
