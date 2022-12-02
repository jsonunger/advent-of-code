package main

import (
	"fmt"
	"sort"
	"strings"

	"github.com/jsonunger/advent-of-code/utils"
)

var closingOpposites = map[string]string{
	")": "(",
	"]": "[",
	"}": "{",
	">": "<",
}

var openingOpposites = map[string]string{
	"(": ")",
	"[": "]",
	"{": "}",
	"<": ">",
}

func part1(lines []string) interface{} {
	scores := map[string]int{
		")": 3,
		"]": 57,
		"}": 1197,
		">": 25137,
	}
	sum := 0

	for _, line := range lines {
		chars := []string{string(line[0])}
		for _, char := range line {
			score, ok := scores[string(char)]
			matches := chars[len(chars)-1] == closingOpposites[string(char)]
			if ok && !matches {
				sum += score
				break
			} else if ok && matches {
				chars = chars[:len(chars)-1]
			} else {
				chars = append(chars, string(char))
			}
		}
	}
	return sum
}

func part2(lines []string) interface{} {
	scoreMap := map[string]int{
		")": 1,
		"]": 2,
		"}": 3,
		">": 4,
	}
	scores := []int{}
	for _, line := range lines {
		chars := []string{string(line[0])}
		incomplete := true
		score := 0
		for _, char := range line[1:] {
			ok := strings.Contains(")]}>", string(char))
			matches := len(chars) >= 1 && chars[len(chars)-1] == closingOpposites[string(char)]
			if ok && !matches {
				incomplete = false
				break
			} else if ok && matches {
				chars = chars[:len(chars)-1]
			} else {
				chars = append(chars, string(char))
			}
		}
		if incomplete {
			for i := len(chars) - 1; i >= 0; i-- {
				char := chars[i]
				closer := openingOpposites[char]
				score *= 5
				score += scoreMap[closer]
			}
			scores = append(scores, score)
		}
	}

	sort.Ints(scores)

	return scores[len(scores)/2]
}

func main() {
	testCase := utils.ReadLines("./2021/day10/test_case.txt")
	input := utils.ReadLines("./2021/day10/input.txt")

	// PART 1
	fmt.Printf("PART 1 EXAMPLE: %v\n", part1(testCase))
	fmt.Printf("PART 1 RESULT: %v\n", part1(input))

	// PART 2
	fmt.Printf("PART 2 EXAMPLE: %v\n", part2(testCase))
	fmt.Printf("PART 2 RESULT: %v\n", part2(input))
}
