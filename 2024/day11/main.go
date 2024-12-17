package main

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/jsonunger/advent-of-code/utils"
)

func part1(line string, blinks int) interface{} {
	stones := strings.Fields(line)
	blinkNum := 0

	for blinkNum < blinks {
		newStones := []string{}

		for _, stone := range stones {
			if stone == "0" {
				newStones = append(newStones, "1")
			} else if len(stone)%2 == 0 {
				// Split the stone
				mid := len(stone) / 2
				lhs, rhs := stone[:mid], stone[mid:]
				left, right := utils.ConvertStringToInt(lhs), utils.ConvertStringToInt(rhs)

				newStones = append(newStones, strconv.Itoa(left), strconv.Itoa(right))
			} else {
				stoneNum := utils.ConvertStringToInt(stone)
				newStones = append(newStones, strconv.Itoa(stoneNum*2024))
			}
		}

		stones = newStones
		blinkNum += 1
	}

	return len(stones)
}

func part2(lines []string) interface{} {
	return nil
}

func main() {
	testCaseA := utils.ReadFile("./2024/day11/test_case_a.txt")
	testCaseB := utils.ReadFile("./2024/day11/test_case_b.txt")
	input := utils.ReadFile("./2024/day11/input.txt")

	// PART 1
	fmt.Printf("PART 1 EXAMPLE: %v\n", part1(testCaseA, 1))
	fmt.Printf("PART 1 EXAMPLE: %v\n", part1(testCaseB, 6))
	fmt.Printf("PART 1 EXAMPLE: %v\n", part1(testCaseB, 25))
	fmt.Printf("PART 1 RESULT: %v\n", part1(input, 25))

	// // PART 2
	// fmt.Printf("PART 2 EXAMPLE: %v\n", part2(testCase))
	// fmt.Printf("PART 2 RESULT: %v\n", part1(input, 75))
}
