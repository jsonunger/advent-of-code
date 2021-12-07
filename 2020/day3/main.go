package main

import (
	"fmt"

	"github.com/jsonunger/advent-of-code/utils"
)

func part1(lines []string, slopeX, slopeY int) int {
	rowLength := len(lines[0])

	x, y, numTrees := 0, 0, 0
	for {
		if y >= len(lines) {
			break
		}
		if lines[y][x%rowLength] == []byte("#")[0] {
			numTrees++
		}
		x += slopeX
		y += slopeY
	}

	return numTrees
}

func part2(lines []string) int {
	slopes := [][2]int{
		{1, 1},
		{3, 1},
		{5, 1},
		{7, 1},
		{1, 2},
	}

	multiple := 1

	for _, slope := range slopes {
		multiple *= part1(lines, slope[0], slope[1])
	}

	return multiple
}

func main() {
	testCaseLines := utils.ReadLines("./2020/day3/test_case.txt")
	inputLines := utils.ReadLines("./2020/day3/input.txt")

	// PART 1
	fmt.Printf("PART 1 EXAMPLE: %v\n", part1(testCaseLines, 3, 1))
	fmt.Printf("PART 1 RESULT: %v\n", part1(inputLines, 3, 1))

	// // PART 2
	fmt.Printf("PART 2 EXAMPLE: %v\n", part2(testCaseLines))
	fmt.Printf("PART 2 RESULT: %v\n", part2(inputLines))
}
