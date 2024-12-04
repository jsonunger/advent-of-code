package main

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/jsonunger/advent-of-code/utils"
)

func absDiffInt(x, y int) int {
	if x < y {
		return y - x
	}
	return x - y
}

func reportIsSafe(levels []string) bool {
	prevLevel := 0
	increasing := false
	for i, level := range levels {
		num, _ := strconv.Atoi(level)
		if i == 0 {
			prevLevel = num
			continue
		}

		if i == 1 && num > prevLevel {
			increasing = true
		}

		if increasing && num <= prevLevel {
			return false
		} else if !increasing && num >= prevLevel {
			return false
		}

		diff := absDiffInt(num, prevLevel)

		if diff < 1 || diff > 3 {
			return false
		}

		prevLevel = num
	}
	return true
}

func part1(lines []string) interface{} {
	numSafe := 0

	for _, report := range lines {
		if reportIsSafe(strings.Fields(report)) {
			numSafe++
		}
	}

	return numSafe
}

func remove[T any](slice []T, idx int) []T {
	newSlice := []T{}
	for i, val := range slice {
		if i == idx {
			continue
		}
		newSlice = append(newSlice, val)
	}
	return newSlice
}

func part2(lines []string) interface{} {
	numSafe := 0

	for _, report := range lines {
		fullReport := strings.Fields(report)
		if reportIsSafe(fullReport) {
			numSafe++
			continue
		}
		anySafe := false

		for index := range fullReport {
			slicedReport := remove(fullReport, index)
			if reportIsSafe(slicedReport) {
				anySafe = true
				break
			}
		}

		if anySafe {
			numSafe++
		}
	}
	return numSafe
}

func main() {
	testCase := utils.ReadLines("./2024/day2/test_case.txt")
	input := utils.ReadLines("./2024/day2/input.txt")

	// PART 1
	fmt.Printf("PART 1 EXAMPLE: %v\n", part1(testCase))
	fmt.Printf("PART 1 RESULT: %v\n", part1(input))

	// PART 2
	fmt.Printf("PART 2 EXAMPLE: %v\n", part2(testCase))
	fmt.Printf("PART 2 RESULT: %v\n", part2(input))
}
