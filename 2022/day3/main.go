package main

import (
	"fmt"
	"strings"

	"github.com/jsonunger/advent-of-code/utils"
)

var priorities = map[string]int{
	"a": 1,
	"b": 2,
	"c": 3,
	"d": 4,
	"e": 5,
	"f": 6,
	"g": 7,
	"h": 8,
	"i": 9,
	"j": 10,
	"k": 11,
	"l": 12,
	"m": 13,
	"n": 14,
	"o": 15,
	"p": 16,
	"q": 17,
	"r": 18,
	"s": 19,
	"t": 20,
	"u": 21,
	"v": 22,
	"w": 23,
	"x": 24,
	"y": 25,
	"z": 26,
}

func intersection(slc1, slc2 []string) []string {
	shared := []string{}
	m := map[string]bool{}

	for _, s1Val := range slc1 {
		m[s1Val] = true
	}

	for _, s2Val := range slc2 {
		if _, ok := m[s2Val]; ok {
			shared = append(shared, s2Val)
		}
	}

	return shared
}

func part1(lines []string) interface{} {
	sum := 0

	for _, line := range lines {
		length := len(line)
		first, last := strings.Split(line[0:length/2], ""), strings.Split(line[length/2:], "")

		shared := intersection(first, last)[0]
		lowerShared := strings.ToLower(shared)
		sum += priorities[lowerShared]
		if shared != lowerShared {
			sum += 26
		}
	}
	return sum
}

func part2(lines []string) interface{} {
	sum := 0

	for i := 0; i < len(lines); i += 3 {
		a, b, c := strings.Split(lines[i], ""), strings.Split(lines[i+1], ""), strings.Split(lines[i+2], "")

		abShared := intersection(a, b)
		shared := intersection(abShared, c)[0]
		lowerShared := strings.ToLower(shared)
		sum += priorities[lowerShared]
		if shared != lowerShared {
			sum += 26
		}
	}
	return sum
}

func main() {
	testCase := utils.ReadLines("./2022/day3/test_case.txt")
	input := utils.ReadLines("./2022/day3/input.txt")

	// PART 1
	fmt.Printf("PART 1 EXAMPLE: %v\n", part1(testCase))
	fmt.Printf("PART 1 RESULT: %v\n", part1(input))

	// // PART 2
	fmt.Printf("PART 2 EXAMPLE: %v\n", part2(testCase))
	fmt.Printf("PART 2 RESULT: %v\n", part2(input))
}
