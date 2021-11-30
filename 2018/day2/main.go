package main

import (
	"fmt"
	"strings"

	"github.com/jsonunger/advent-of-code/utils"
	"github.com/pkg/errors"
)

func part1(lines []string) (interface{}, error) {
	doubles := 0
	triples := 0
	for _, line := range lines {
		double, triple := false, false
		chars := make(map[string]int)
		for _, char := range strings.Split(line, "") {
			chars[char] += 1
		}
		for _, count := range chars {
			if count == 3 && !triple {
				triple = true
				triples += 1
				continue
			}
			if count == 2 && !double {
				double = true
				doubles += 1
			}
		}
	}
	return doubles * triples, nil
}

func part2(lines []string) (interface{}, error) {
	for i, lineI := range lines {
		for j, lineJ := range lines {
			if j <= i || len(lineI) != len(lineJ) {
				continue
			}
			diffs := make([]int, 0, len(lineI))
			for idx, _ := range strings.Split(lineI, "") {
				if lineI[idx] != lineJ[idx] {
					diffs = append(diffs, idx)
				}
			}
			if len(diffs) == 1 {
				d := []rune(lineI)
				return string(append(d[0:diffs[0]], d[diffs[0]+1:]...)), nil
			}
		}
	}
	return nil, errors.New("solution not found")
}

func main() {
	// testCaseLines := utils.ReadFile("test_case.txt")
	inputLines := utils.ReadFile("input.txt")

	// // PART 1
	// exRes1, err := part1(testCaseLines)
	// utils.PanicWithMsg(err, "part 1 example")
	// fmt.Printf("PART 1 EXAMPLE: %v\n", exRes1)

	// res1, err := part1(inputLines)
	// utils.PanicWithMsg(err, "part 1")
	// fmt.Printf("PART 1 RESULT: %v\n", res1)

	testCase2Lines := utils.ReadFile("test_case2.txt")

	// PART 2
	exRes2, err := part2(testCase2Lines)
	utils.PanicWithMsg(err, "part 2 example")
	fmt.Printf("PART 2 EXAMPLE: %v\n", exRes2)

	res2, err := part2(inputLines)
	utils.PanicWithMsg(err, "part 2")
	fmt.Printf("PART 2 RESULT: %v\n", res2)
}
