package main

import (
	"fmt"
	"strconv"

	"github.com/jsonunger/advent-of-code/utils"
)

func part1(lines []string) (interface{}, error) {
	frequency := 0
	for _, change := range lines {
		val, err := strconv.Atoi(change)
		utils.PanicWithMsg(err, fmt.Sprintf("converting %v", change))
		frequency += val
	}

	return frequency, nil
}

func part2(lines []string) (interface{}, error) {
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
			return frequency, nil
		}
		pastFreqs[frequency] = true
		i += 1
	}
}

func main() {
	testCaseLines := utils.ReadFile("test_case.txt")
	inputLines := utils.ReadFile("input.txt")

	// // PART 1
	// exRes1, err := part1(testCaseLines)
	// utils.PanicWithMsg(err, "part 1 example")
	// fmt.Printf("PART 1 EXAMPLE: %v\n", exRes1)

	// res1, err := part1(inputLines)
	// utils.PanicWithMsg(err, "part 1")
	// fmt.Printf("PART 1 RESULT: %v\n", res1)

	// PART 2
	exRes2, err := part2(testCaseLines)
	utils.PanicWithMsg(err, "part 2 example")
	fmt.Printf("PART 2 EXAMPLE: %v\n", exRes2)

	res2, err := part2(inputLines)
	utils.PanicWithMsg(err, "part 2")
	fmt.Printf("PART 2 RESULT: %v\n", res2)
}
