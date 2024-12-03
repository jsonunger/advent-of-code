package main

import (
	"fmt"
	"strconv"
	"strings"

	"slices"

	"github.com/jsonunger/advent-of-code/utils"
)

func absDiffInt(x, y int) int {
	if x < y {
	   return y - x
	}
	return x - y
 }

func part1(lines []string) interface{} {
	leftSide, rightSide := []int{}, []int{}

	for _, line := range lines {
		digits := strings.Fields(line)
		for i, digit := range digits {
			num, _ := strconv.Atoi(digit)
			if i == 0 {
				leftSide = append(leftSide, num)
			} else {
				rightSide = append(rightSide, num)
			}
		}
	}

	slices.Sort(leftSide)
	slices.Sort(rightSide)

	sum := 0

	for i, lhs := range leftSide {
		rhs := rightSide[i]
		sum += absDiffInt(lhs, rhs)
	}

	return sum
}

func part2(lines []string) interface{} {
	leftSide := []int{}
	rightSide := map[int]int{}

	for _, line := range lines {
		digits := strings.Fields(line)
		for i, digit := range digits {
			num, _ := strconv.Atoi(digit)
			if i == 0 {
				leftSide = append(leftSide, num)
			} else {
				rightSide[num] += 1
			}
		}
	}

	sum := 0

	for _, lhs := range leftSide {
		if rhs, ok := rightSide[lhs]; ok {
			sum += lhs * rhs
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

	// PART 2
	fmt.Printf("PART 2 EXAMPLE: %v\n", part2(testCase))
	fmt.Printf("PART 2 RESULT: %v\n", part2(input))
}
