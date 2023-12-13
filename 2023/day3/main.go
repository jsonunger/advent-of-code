package main

import (
	"fmt"
	"image"
	"regexp"
	"strconv"

	"github.com/jsonunger/advent-of-code/utils"
)

var numericalRegex = regexp.MustCompile(`\d+`)
var symbolRegex = regexp.MustCompile(`[^0-9\.]`)

var imageRange = []image.Point{
	{-1, -1},
	{-1, 0},
	{-1, 1},
	{0, -1},
	{0, 1},
	{1, -1},
	{1, 0},
	{1, 1},
}

func part1(lines []string) interface{} {
	sum := 0

	symbolsGrid := map[image.Point]bool{}
	for y, line := range lines {
		for x, r := range line {
			if symbolRegex.MatchString(string(r)) {
				symbolsGrid[image.Point{x, y}] = true
			}
		}
	}

	for y, line := range lines {
		numberIndices := numericalRegex.FindAllStringIndex(line, -1)
		for _, indices := range numberIndices {
			bounds := map[image.Point]bool{}
			for x := indices[0]; x < indices[1]; x++ {
				for _, d := range imageRange {
					bounds[image.Point{x, y}.Add(d)] = true
				}
			}
			num, _ := strconv.Atoi(line[indices[0]:indices[1]])
			for p := range bounds {
				if symbolsGrid[p] {
					sum += num
				}
			}
		}
	}

	return sum
}

func part2(lines []string) interface{} {
	return nil
}

func main() {
	testCase := utils.ReadLines("./2023/day3/test_case.txt")
	input := utils.ReadLines("./2023/day3/input.txt")

	// PART 1
	fmt.Printf("PART 1 EXAMPLE: %v\n", part1(testCase))
	fmt.Printf("PART 1 RESULT: %v\n", part1(input))

	// // PART 2
	// fmt.Printf("PART 2 EXAMPLE: %v\n", part2(testCase))
	// fmt.Printf("PART 2 RESULT: %v\n", part2(input))
}
