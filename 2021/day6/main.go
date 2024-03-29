package main

import (
	"fmt"
	"strings"

	"github.com/jsonunger/advent-of-code/utils"
)

func runner(line string, days int) interface{} {
	startingFishTimers := utils.ConvertStringsToInts(strings.Split(line, ","))

	fishes := map[int]int{0: 0, 1: 0, 2: 0, 3: 0, 4: 0, 5: 0, 6: 0, 7: 0, 8: 0}

	for _, val := range startingFishTimers {
		fishes[val]++
	}

	for i := 0; i < days; i++ {
		fishes[(i+7)%9] += fishes[i%9]
	}

	sum := 0
	for _, val := range fishes {
		sum += val
	}

	return sum
}

func main() {
	testCase := utils.ReadFile("./2021/day6/test_case.txt")
	input := utils.ReadFile("./2021/day6/input.txt")

	// PART 1
	fmt.Printf("PART 1 EXAMPLE: %v\n", runner(testCase, 18))
	fmt.Printf("PART 1 RESULT: %v\n", runner(input, 80))

	// // PART 2
	fmt.Printf("PART 2 EXAMPLE: %v\n", runner(testCase, 256))
	fmt.Printf("PART 2 RESULT: %v\n", runner(input, 256))
}
