package main

import (
	"fmt"
	"strconv"

	"github.com/jsonunger/advent-of-code/utils"
)

type Location struct {
	x, y, digit int
}

func byteToDigit(b byte) int {
	digit, _ := strconv.Atoi(string(b))
	return digit
}

func runeToDigit(r rune) int {
	digit, _ := strconv.Atoi(fmt.Sprintf("%c", r))
	return digit
}

func next(lines []string, location Location) (nextLocations []Location) {
	if byteToDigit(lines[location.y+1][location.x]) == location.digit+1 {
		nextLocations = append(nextLocations, Location{location.x, location.y, location.digit + 1})
	}
	if byteToDigit(lines[location.y-1][location.x]) == location.digit+1 {
		nextLocations = append(nextLocations, Location{location.x, location.y, location.digit + 1})
	}
	if byteToDigit(lines[location.y][location.x+1]) == location.digit+1 {
		nextLocations = append(nextLocations, Location{location.x, location.y, location.digit + 1})
	}
	if byteToDigit(lines[location.y][location.x-1]) == location.digit+1 {
		nextLocations = append(nextLocations, Location{location.x, location.y, location.digit + 1})
	}

	return nextLocations
}

func part1(lines []string) interface{} {
	trailheads := map[Location]int{}
	for y, line := range lines {
		for x, char := range line {
			loc := Location{x, y, runeToDigit(char)}
			nextLocations := next(lines, loc)
			fmt.Println(loc, nextLocations)
		}
	}

	sum := 0
	for _, score := range trailheads {
		sum += score
	}
	return sum
}

func part2(lines []string) interface{} {
	return nil
}

func main() {
	testCase := utils.ReadLines("./2024/day10/test_case.txt")
	// input := utils.ReadLines("./2024/day10/input.txt")

	// PART 1
	fmt.Printf("PART 1 EXAMPLE: %v\n", part1(testCase))
	// fmt.Printf("PART 1 RESULT: %v\n", part1(input))

	// // PART 2
	// fmt.Printf("PART 2 EXAMPLE: %v\n", part2(testCase))
	// fmt.Printf("PART 2 RESULT: %v\n", part2(input))
}
