package main

import (
	"fmt"
	"regexp"

	"github.com/jsonunger/advent-of-code/utils"
)

func rotate(line string, position, numZeroes int) (int, int) {
	regexpToMatch := regexp.MustCompile(`(?P<direction>L|R)(?P<steps>\d+)`)
	matches := regexpToMatch.FindStringSubmatch(line)
	direction, steps := matches[1], matches[2]
	stepsNum := utils.ConvertStringToInt(steps)

	for i := 0; i < stepsNum; i++ {
		switch direction {
		case "L":
			position -= 1
		case "R":
			position += 1
		}

		if position == -1 {
			position = 99
		} else if position == 100 {
			position = 0
		}

		if position == 0 {
			numZeroes += 1
		}
	}
	return position, numZeroes
}

func part1(lines []string) interface{} {
	position, numZeroes := 50, 0

	for _, line := range lines {
		position, _ = rotate(line, position, 0)

		if position == 0 {
			numZeroes += 1
		}
	}

	return numZeroes
}

func part2(lines []string) interface{} {
	position, numZeroes := 50, 0

	for _, line := range lines {
		position, numZeroes = rotate(line, position, numZeroes)
	}

	return numZeroes
}

func main() {
	testCase := utils.ReadLines("./2025/day1/test_case.txt")
	input := utils.ReadLines("./2025/day1/input.txt")

	// PART 1
	fmt.Printf("PART 1 EXAMPLE: %v\n", part1(testCase))
	fmt.Printf("PART 1 RESULT: %v\n", part1(input))

	// // PART 2
	fmt.Printf("PART 2 EXAMPLE: %v\n", part2(testCase))
	fmt.Printf("PART 2 RESULT: %v\n", part2(input))
}
