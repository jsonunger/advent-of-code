package main

import (
	"fmt"
	"strconv"

	"github.com/jsonunger/advent-of-code/utils"
)

type counts struct {
	Zero int
	One  int
}

func part1(lines []string) interface{} {
	gammaStr, epsilonStr := "", ""
	for i := 0; i < len(lines[0]); i++ {
		countZeros, countOnes := 0, 0
		for _, line := range lines {
			char := string(line[i])
			if char == "0" {
				countZeros += 1
			} else {
				countOnes += 1
			}
		}

		if countOnes > countZeros {
			gammaStr += "1"
			epsilonStr += "0"
		} else {
			gammaStr += "0"
			epsilonStr += "1"
		}
	}
	gamma, err := strconv.ParseInt(gammaStr, 2, 64)
	utils.PanicWithMsg(err, "converting gamma")
	epsilon, err := strconv.ParseInt(epsilonStr, 2, 64)
	utils.PanicWithMsg(err, "converting epsilon")
	return gamma * epsilon
}

func filter(s []string, callback func(string) bool) []string {
	filtered := []string{}
	for i := range s {
		if callback(s[i]) {
			filtered = append(filtered, s[i])
		}
	}
	return filtered
}

func part2(lines []string) interface{} {
	numChars := len(lines[0])
	oxyLines := make([]string, len(lines))
	copy(oxyLines, lines)
	for i := 0; i < numChars; i++ {
		if len(oxyLines) == 1 {
			break
		}
		countZeros, countOnes := 0, 0
		for _, line := range oxyLines {
			char := string(line[i])
			if char == "0" {
				countZeros += 1
			} else {
				countOnes += 1
			}
		}
		oxyLines = filter(oxyLines, func(line string) bool {
			if countOnes == countZeros && string(line[i]) == "1" {
				return true
			}
			if countOnes > countZeros && string(line[i]) == "1" {
				return true
			}
			if countZeros > countOnes && string(line[i]) == "0" {
				return true
			}
			return false
		})

	}
	oxy, err := strconv.ParseInt(oxyLines[0], 2, 64)
	utils.PanicWithMsg(err, "converting oxy")

	co2Lines := make([]string, len(lines))
	copy(co2Lines, lines)
	for i := 0; i < numChars; i++ {
		if len(co2Lines) == 1 {
			break
		}
		countZeros, countOnes := 0, 0
		for _, line := range co2Lines {
			char := string(line[i])
			if char == "0" {
				countZeros += 1
			} else {
				countOnes += 1
			}
		}
		co2Lines = filter(co2Lines, func(line string) bool {
			if countOnes == countZeros && string(line[i]) == "0" {
				return true
			}
			if countOnes < countZeros && string(line[i]) == "1" {
				return true
			}
			if countZeros < countOnes && string(line[i]) == "0" {
				return true
			}
			return false
		})

	}

	co2, err := strconv.ParseInt(co2Lines[0], 2, 64)
	utils.PanicWithMsg(err, "converting oxy")

	return oxy * co2
}

func main() {
	testCaseLines := utils.ReadFile("./2021/day3/test_case.txt")
	inputLines := utils.ReadFile("./2021/day3/input.txt")

	// PART 1
	// fmt.Printf("PART 1 EXAMPLE: %v\n", part1(testCaseLines))
	// fmt.Printf("PART 1 RESULT: %v\n", part1(inputLines))

	// // PART 2
	fmt.Printf("PART 2 EXAMPLE: %v\n", part2(testCaseLines))
	fmt.Printf("PART 2 RESULT: %v\n", part2(inputLines))
}
