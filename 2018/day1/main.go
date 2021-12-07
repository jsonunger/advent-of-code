package main

import (
	"fmt"
	"strconv"

	"github.com/jsonunger/advent-of-code/utils"
)

func part1(lines []string) int {
	frequency := 0
	for _, change := range lines {
		val, err := strconv.Atoi(change)
		utils.PanicWithMsg(err, fmt.Sprintf("converting %v", change))
		frequency += val
	}

	return frequency
}

func part2(lines []string) int {
	frequency := 0
	pastFreqs := map[int]bool{0: true}
	i := 0

	for {
		change := lines[i%len(lines)]
		val, err := strconv.Atoi(change)
		utils.PanicWithMsg(err, fmt.Sprintf("converting %v", change))
		frequency += val
		exists := pastFreqs[frequency]
		if exists {
			return frequency
		}
		pastFreqs[frequency] = true
		i++
	}
}

func main() {
	testCaseLines := utils.ReadLines("test_case.txt")
	inputLines := utils.ReadLines("input.txt")

	// PART 1
	exRes1 := part1(testCaseLines)
	fmt.Printf("PART 1 EXAMPLE: %v\n", exRes1)

	res1 := part1(inputLines)
	fmt.Printf("PART 1 RESULT: %v\n", res1)

	// PART 2
	exRes2 := part2(testCaseLines)
	fmt.Printf("PART 2 EXAMPLE: %v\n", exRes2)

	res2 := part2(inputLines)
	fmt.Printf("PART 2 RESULT: %v\n", res2)
}
