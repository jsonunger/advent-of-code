package main

import (
	"fmt"
	"strings"

	"github.com/jsonunger/advent-of-code/utils"
)

type Rule struct {
	char     string
	min      int
	max      int
	password string
}

func NewRule(line string) Rule {
	splitLine := strings.Split(line, ": ")
	password := splitLine[1]
	splitLine = strings.Split(splitLine[0], " ")
	char := splitLine[1]
	minAndMax := utils.ConvertStringsToInts(strings.Split(splitLine[0], "-"))

	return Rule{
		char,
		minAndMax[0],
		minAndMax[1],
		password,
	}
}

func part1(lines []string) interface{} {
	sum := 0
	for _, line := range lines {
		rule := NewRule(line)
		charCount := strings.Count(rule.password, rule.char)
		if charCount >= rule.min && charCount <= rule.max {
			sum++
		}
	}
	return sum
}

func part2(lines []string) interface{} {
	sum := 0
	for _, line := range lines {
		rule := NewRule(line)
		minEquals := []rune(rule.password)[rule.min-1] == []rune(rule.char)[0]
		maxEquals := []rune(rule.password)[rule.max-1] == []rune(rule.char)[0]
		if minEquals != maxEquals {
			sum++
		}
	}
	return sum
}

func main() {
	testCaseLines := utils.ReadLines("./2020/day2/test_case.txt")
	inputLines := utils.ReadLines("./2020/day2/input.txt")

	// PART 1
	fmt.Printf("PART 1 EXAMPLE: %v\n", part1(testCaseLines))
	fmt.Printf("PART 1 RESULT: %v\n", part1(inputLines))

	// // PART 2
	fmt.Printf("PART 2 EXAMPLE: %v\n", part2(testCaseLines))
	fmt.Printf("PART 2 RESULT: %v\n", part2(inputLines))
}
