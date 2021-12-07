package main

import (
	"fmt"
	"math"
	"sort"
	"strings"

	"github.com/jsonunger/advent-of-code/utils"
)

func part1(input string) interface{} {
	positions := utils.ConvertStringsToInts(strings.Split(input, ","))
	sort.Ints(positions)
	mid := positions[len(positions)/2]

	return utils.SumInts(positions, func(pos int) int {
		return int(math.Abs(float64(pos - mid)))
	})
}

func fuelCost(pos, avg int) int {
	distance := math.Abs(float64(pos - avg))
	return int(distance*(distance+1)) / 2
}

func part2(input string) interface{} {
	positions := utils.ConvertStringsToInts(strings.Split(input, ","))

	avg := utils.SumInts(positions, nil) / len(positions)

	min := math.Inf(1)
	for _, a := range [2]int{avg, avg + 1} {
		sum := utils.SumInts(positions, func(pos int) int {
			return fuelCost(pos, a)
		})
		if float64(sum) < min {
			min = float64(sum)
		}
	}

	return min
}

func main() {
	testCase := utils.ReadFile("./2021/day7/test_case.txt")
	input := utils.ReadFile("./2021/day7/input.txt")

	// PART 1
	fmt.Printf("PART 1 EXAMPLE: %v\n", part1(testCase))
	fmt.Printf("PART 1 RESULT: %v\n", part1(input))

	// PART 2
	fmt.Printf("PART 2 EXAMPLE: %v\n", part2(testCase))
	fmt.Printf("PART 2 RESULT: %v\n", part2(input))
}
