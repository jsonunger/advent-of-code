package main

import (
	"fmt"

	"github.com/jsonunger/advent-of-code/utils"
)

func part1(lines []string) interface{} {
	ints := utils.ConvertStringsToInts(lines)

	for i, valI := range ints {
		for _, valJ := range ints[i+1:] {
			if valI+valJ == 2020 {
				return valI * valJ
			}
		}
	}
	return nil
}

func part2(lines []string) interface{} {
	ints := utils.ConvertStringsToInts(lines)

	for i, valI := range ints {
		for j, valJ := range ints[i+1:] {
			for _, valK := range ints[j+1:] {
				if valI+valJ+valK == 2020 {
					return valI * valJ * valK
				}
			}
		}
	}
	return nil
}

func main() {
	testCaseLines := utils.ReadLines("./2020/day1/test_case.txt")
	inputLines := utils.ReadLines("./2020/day1/input.txt")

	// PART 1
	fmt.Printf("PART 1 EXAMPLE: %v\n", part1(testCaseLines))
	fmt.Printf("PART 1 RESULT: %v\n", part1(inputLines))

	// // PART 2
	fmt.Printf("PART 2 EXAMPLE: %v\n", part2(testCaseLines))
	fmt.Printf("PART 2 RESULT: %v\n", part2(inputLines))
}
