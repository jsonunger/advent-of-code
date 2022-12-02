package main

import (
	"fmt"
	"math"
	"sort"

	"github.com/jsonunger/advent-of-code/utils"
)

func getAdjacentValues(x, y, totalX, totalY int) [][2]int {
	var adjacent [][2]int

	for _, val := range []int{x - 1, x + 1} {
		if val >= 0 && val != totalX {
			adjacent = append(adjacent, [2]int{val, y})
		}
	}

	for _, val := range []int{y - 1, y + 1} {
		if val >= 0 && val != totalY {
			adjacent = append(adjacent, [2]int{x, val})
		}
	}

	return adjacent
}

func lessThan(a int, b byte) bool {
	return a < utils.ConvertStringToInt(string(b))
}

func part1(lines []string) int {
	totalY := len(lines)
	totalX := len(lines[0])
	sum := 0
	for y, line := range lines {
		for x, char := range line {
			charVal := utils.ConvertStringToInt(string(char))
			allLessThan := true
			for _, coords := range getAdjacentValues(x, y, totalX, totalY) {
				if !lessThan(charVal, lines[coords[1]][coords[0]]) {
					allLessThan = false
					break
				}
			}
			if allLessThan {
				sum += charVal + 1
			}
		}
	}
	return sum
}

func part2(lines []string) int {
	basins := []int{}

	sort.Ints(basins)

	multiple := 1

	startingIdx := int(math.Max(0, float64(len(basins)-3)))

	for _, basin := range basins[startingIdx:] {
		multiple *= basin
	}
	return multiple
}

func main() {
	testCase := utils.ReadLines("./2021/day9/test_case.txt")
	input := utils.ReadLines("./2021/day9/input.txt")

	// PART 1
	fmt.Printf("PART 1 EXAMPLE: %v\n", part1(testCase))
	fmt.Printf("PART 1 RESULT: %v\n", part1(input))

	// PART 2
	fmt.Printf("PART 2 EXAMPLE: %v\n", part2(testCase))
	// fmt.Printf("PART 2 RESULT: %v\n", part2(input))
}
