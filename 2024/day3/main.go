package main

import (
	"fmt"
	"regexp"
	"strconv"

	"github.com/jsonunger/advent-of-code/utils"
)

var multiplierRegexp, _ = regexp.Compile(`mul\((\d{1,3}),(\d{1,3})\)`)

func multiply(match []string) int {
	lhsStr, rhsStr := match[1], match[2]
	lhs, _ := strconv.Atoi(lhsStr)
	rhs, _ := strconv.Atoi(rhsStr)
	return lhs * rhs
}

func part1(lines []string) interface{} {
	sum := 0

	for _, memory := range lines {
		for _, match := range multiplierRegexp.FindAllStringSubmatch(memory, -1) {
			sum += multiply(match)
		}
	}

	return sum
}

func isEnabled(enabledMap map[int]bool, index int) bool {
	nearestIndex := 0

	for i := range enabledMap {
		if i < index && i > nearestIndex {
			nearestIndex = i
		}
	}

	return enabledMap[nearestIndex]
}

func part2(lines []string) interface{} {
	sum := 0
	doOrDontRegexp, _ := regexp.Compile(`(do|don't)\(\)`)

	startEnabled := true

	for _, memory := range lines {
		enabledMap := map[int]bool{0: startEnabled}
		controlIndices := doOrDontRegexp.FindAllStringSubmatchIndex(memory, -1)
		controlMatches := doOrDontRegexp.FindAllStringSubmatch(memory, -1)
		for i, indexMatch := range controlIndices {
			match := controlMatches[i]
			enabledMap[indexMatch[0]] = match[1] == "do"
			if i == len(controlIndices)-1 {
				startEnabled = match[1] == "do"
			}
		}

		commandIndices := multiplierRegexp.FindAllStringSubmatchIndex(memory, -1)
		commandMatches := multiplierRegexp.FindAllStringSubmatch(memory, -1)

		for i, indexMatch := range commandIndices {
			if isEnabled(enabledMap, indexMatch[0]) {
				match := commandMatches[i]
				sum += multiply(match)
			}
		}

	}

	return sum
}

func main() {
	testCase := utils.ReadLines("./2024/day3/test_case.txt")
	input := utils.ReadLines("./2024/day3/input.txt")

	// PART 1
	fmt.Printf("PART 1 EXAMPLE: %v\n", part1(testCase))
	fmt.Printf("PART 1 RESULT: %v\n", part1(input))

	testCase = utils.ReadLines("./2024/day3/test_case_pt2.txt")

	// PART 2
	fmt.Printf("PART 2 EXAMPLE: %v\n", part2(testCase))
	fmt.Printf("PART 2 RESULT: %v\n", part2(input))
}
