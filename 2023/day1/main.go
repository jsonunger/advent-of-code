package main

import (
	"fmt"
	"math"
	"strconv"
	"strings"

	"github.com/jsonunger/advent-of-code/utils"
)

func part1(lines []string) interface{} {
	sum := 0
	for _, line := range lines {
		var digits []string
		for _, char := range line {
			charString := string(char)
			if val, _ := strconv.Atoi(charString); val != 0 {
				digits = append(digits, charString)
			}
		}
		if len(digits) == 0 {
			continue
		}
		intVal, _ := strconv.Atoi(digits[0] + digits[len(digits)-1])
		sum += intVal
	}
	return sum
}

var digitWords = map[string]string{"one": "1", "two": "2", "three": "3", "four": "4", "five": "5", "six": "6", "seven": "7", "eight": "8", "nine": "9"}
var digitChars = []string{"1", "2", "3", "4", "5", "6", "7", "8", "9"}

func minAndMax(digitMap map[int]string) (int, int) {
	min, max := math.MaxInt, math.MinInt
	for idx := range digitMap {
		if idx < min {
			min = idx
		}
		if idx > max {
			max = idx
		}
	}

	return min, max
}

func part2(lines []string) interface{} {
	sum := 0
	for _, line := range lines {
		digits := map[int]string{}
		for str, val := range digitWords {
			if idx := strings.Index(line, str); idx != -1 {
				digits[idx] = val
			}
			if idx := strings.LastIndex(line, str); idx != -1 {
				digits[idx] = val
			}
		}
		for _, digit := range digitChars {
			if idx := strings.Index(line, digit); idx != -1 {
				digits[idx] = digit
			}
			if idx := strings.LastIndex(line, digit); idx != -1 {
				digits[idx] = digit
			}
		}
		min, max := minAndMax(digits)
		intVal, _ := strconv.Atoi(digits[min] + digits[max])
		sum += intVal
	}
	return sum
}

func main() {
	testCase := utils.ReadLines("./2023/day1/test_case.txt")
	input := utils.ReadLines("./2023/day1/input.txt")

	// PART 1
	fmt.Printf("PART 1 EXAMPLE: %v\n", part1(testCase))
	fmt.Printf("PART 1 RESULT: %v\n", part1(input))

	testCasePt2 := utils.ReadLines("./2023/day1/test_case_pt2.txt")

	// PART 2
	fmt.Printf("PART 2 EXAMPLE: %v\n", part2(testCasePt2))
	fmt.Printf("PART 2 RESULT: %v\n", part2(input))
}
