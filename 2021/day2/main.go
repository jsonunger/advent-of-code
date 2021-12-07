package main

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/jsonunger/advent-of-code/utils"
)

func part1(lines []string) int {
	horizontal, depth := 0, 0
	for _, line := range lines {
		split := strings.Split(line, " ")
		direction := split[0]
		val := split[1]
		value, err := strconv.Atoi(val)
		utils.PanicWithMsg(err, fmt.Sprintf("converting %v", val))

		switch direction {
		case "forward":
			horizontal += value
		case "down":
			depth += value
		case "up":
			depth -= value
		}
	}
	return horizontal * depth
}

func part2(lines []string) int {
	horizontal, depth, aim := 0, 0, 0
	for _, line := range lines {
		split := strings.Split(line, " ")
		direction := split[0]
		val := split[1]
		value, err := strconv.Atoi(val)
		utils.PanicWithMsg(err, fmt.Sprintf("converting %v", val))

		switch direction {
		case "forward":
			horizontal += value
			depth += (value * aim)
		case "down":
			aim += value
		case "up":
			aim -= value
		}
	}
	return horizontal * depth
}

func main() {
	testCaseLines := utils.ReadLines("./2021/day2/test_case.txt")
	inputLines := utils.ReadLines("./2021/day2/input.txt")

	// PART 1
	exRes1 := part1(testCaseLines)
	fmt.Printf("PART 1 EXAMPLE: %v\n", exRes1)

	res1 := part1(inputLines)
	fmt.Printf("PART 1 RESULT: %v\n", res1)

	// PART 2
	exRes2 := part2(testCaseLines)
	fmt.Printf("PART 2 EXAMPLE: %v\n", exRes2)

	res2 := part2(inputLines)
	fmt.Printf("PART 2 RESULT: %v\n", res2)
}
