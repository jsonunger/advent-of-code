package main

import (
	"fmt"

	"github.com/jsonunger/advent-of-code/utils"
)

func maxJoltage(line string, digitsToPick int) int {
	lineLength := len(line)
	currentPositionIdx := 0
	maxJoltageStr := ""

	for digitToPickIdx := 0; digitToPickIdx < digitsToPick; digitToPickIdx++ {
		maxDigitByte := byte('0')
		maxDigitIdx := -1
		maxSearch := lineLength - (digitsToPick - digitToPickIdx)
		for i := currentPositionIdx; i <= maxSearch; i++ {
			if line[i] > maxDigitByte {
				maxDigitByte = line[i]
				maxDigitIdx = i
				if maxDigitByte == '9' {
					break
				}
			}
		}
		maxJoltageStr += string(maxDigitByte)
		currentPositionIdx = maxDigitIdx + 1
	}
	return utils.ConvertStringToInt(maxJoltageStr)
}

func part1(lines []string) interface{} {
	sum := 0

	for _, line := range lines {
		sum += maxJoltage(line, 2)
	}
	return sum
}

func part2(lines []string) interface{} {
	sum := 0

	for _, line := range lines {
		sum += maxJoltage(line, 12)
	}
	return sum
}

func main() {
	testCase := utils.ReadLines("./2025/day3/test_case.txt")
	input := utils.ReadLines("./2025/day3/input.txt")

	// PART 1
	fmt.Printf("PART 1 EXAMPLE: %v\n", part1(testCase))
	fmt.Printf("PART 1 RESULT: %v\n", part1(input))

	// // PART 2
	fmt.Printf("PART 2 EXAMPLE: %v\n", part2(testCase))
	fmt.Printf("PART 2 RESULT: %v\n", part2(input))
}
