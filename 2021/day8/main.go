package main

import (
	"fmt"
	"strings"

	"github.com/jsonunger/advent-of-code/utils"
)

func part1(lines []string) interface{} {
	count := 0
	for _, line := range lines {
		split := strings.Split(line, " | ")
		values := strings.Split(split[1], " ")
		for _, value := range values {
			switch len(value) {
			case 2, 3, 4, 7:
				count++
			}
		}
	}
	return count
}

func part2(lines []string) interface{} {
	for _, line := range lines {
		split := strings.Split(line, " | ")
		options := map[int][]string{}
		for _, pattern := range strings.Split(split[0], " ") {
			switch len(pattern) {
			case 2:
				options[1] = append(options[1], pattern)
			case 3:
				options[7] = append(options[7], pattern)
			case 4:
				options[4] = append(options[4], pattern)
			case 5:
				options[2] = append(options[2], pattern)
				options[3] = append(options[3], pattern)
				options[5] = append(options[5], pattern)
			case 6:
				options[0] = append(options[0], pattern)
				options[6] = append(options[6], pattern)
				options[9] = append(options[9], pattern)
			case 7:
				options[8] = append(options[8], pattern)
			}
		}
		fmt.Println(options)
	}
	return nil
}

func main() {
	testCase := utils.ReadLines("./2021/day8/test_case.txt")
	input := utils.ReadLines("./2021/day8/input.txt")

	// PART 1
	fmt.Printf("PART 1 EXAMPLE: %v\n", part1(testCase))
	fmt.Printf("PART 1 RESULT: %v\n", part1(input))

	// PART 2
	fmt.Printf("PART 2 EXAMPLE: %v\n", part2(testCase))
	// fmt.Printf("PART 2 RESULT: %v\n", part2(input))
}
