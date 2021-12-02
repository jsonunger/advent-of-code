package main

import (
	"fmt"
	"strconv"

	"github.com/jsonunger/advent-of-code/utils"
)

func part1(lines []string) (interface{}, error) {
	count := 0
	for i, line := range lines[:len(lines)-1] {
		val, err := strconv.Atoi(line)
		utils.PanicWithMsg(err, fmt.Sprintf("converting %v", line))
		nextVal, err := strconv.Atoi(lines[i+1])
		utils.PanicWithMsg(err, fmt.Sprintf("converting %v", lines[i+1]))
		if nextVal > val {
			count += 1
		}
	}
	return count, nil
}

func sumNextThree(lines []string, startingIdx int) int {
	sum := 0
	for _, line := range lines[startingIdx : startingIdx+3] {
		val, err := strconv.Atoi(line)
		utils.PanicWithMsg(err, fmt.Sprintf("converting %v", line))
		sum += val
	}
	return sum
}

func part2(lines []string) (interface{}, error) {
	count := 0
	for i := range lines[:len(lines)-3] {
		sum := sumNextThree(lines, i)
		nextSum := sumNextThree(lines, i+1)
		if nextSum > sum {
			count += 1
		}
	}
	return count, nil
}

func main() {
	testCaseLines := utils.ReadFile("./2021/day1/test_case.txt")
	inputLines := utils.ReadFile("./2021/day1/input.txt")

	// PART 1
	exRes1, err := part1(testCaseLines)
	utils.PanicWithMsg(err, "part 1 example")
	fmt.Printf("PART 1 EXAMPLE: %v\n", exRes1)

	res1, err := part1(inputLines)
	utils.PanicWithMsg(err, "part 1")
	fmt.Printf("PART 1 RESULT: %v\n", res1)

	// PART 2
	exRes2, err := part2(testCaseLines)
	utils.PanicWithMsg(err, "part 2 example")
	fmt.Printf("PART 2 EXAMPLE: %v\n", exRes2)

	res2, err := part2(inputLines)
	utils.PanicWithMsg(err, "part 2")
	fmt.Printf("PART 2 RESULT: %v\n", res2)
}
