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

func part1(lines []string) interface{} {
	numSafe := 0

	for _, report := range lines {
		prevLevel := 0
		increasing := false
		safe := true
		for i, level := range strings.Fields(report) {
			num, _ := strconv.Atoi(level)
			if i == 0 {
				prevLevel = num
				continue
			}
			
			if i == 1 && num > prevLevel {
				increasing = true
			}

			if increasing && num <= prevLevel {
				safe = false
				break
			} else if !increasing && num >= prevLevel {
				safe = false
				break
			}

			diff := absDiffInt(num, prevLevel)

			if diff < 1 || diff > 3 {
				safe = false
				break
			}

			prevLevel = num
		}
		if safe {
			numSafe++
		}
	}

	return numSafe
}

func part2(lines []string) interface{} {
	return nil
}

func main() {
	testCase := utils.ReadLines("./2024/day2/test_case.txt")
	input := utils.ReadLines("./2024/day2/input.txt")

	// PART 1
	fmt.Printf("PART 1 EXAMPLE: %v\n", part1(testCase))
	fmt.Printf("PART 1 RESULT: %v\n", part1(input))

	// // PART 2
	// fmt.Printf("PART 2 EXAMPLE: %v\n", part2(testCase))
	// fmt.Printf("PART 2 RESULT: %v\n", part2(input))
}
