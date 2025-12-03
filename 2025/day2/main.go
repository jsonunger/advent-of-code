package main

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/jsonunger/advent-of-code/utils"
)

func part1(lines []string) interface{} {
	sum := 0
	idRanges := strings.Split(lines[0], ",")
	for _, idRange := range idRanges {
		rangeParts := strings.Split(idRange, "-")
		start, end := utils.ConvertStringToInt(rangeParts[0]), utils.ConvertStringToInt(rangeParts[1])
		for id := start; id <= end; id++ {
			idStr := strconv.Itoa(id)
			mid := len(idStr) / 2
			if idStr[:mid] == idStr[mid:] {
				sum += id
			}
		}
	}
	return sum
}

func part2(lines []string) interface{} {
	sum := 0
	idRanges := strings.Split(lines[0], ",")
	for _, idRange := range idRanges {
		rangeParts := strings.Split(idRange, "-")
		start, end := utils.ConvertStringToInt(rangeParts[0]), utils.ConvertStringToInt(rangeParts[1])
		for id := start; id <= end; id++ {
			idStr := strconv.Itoa(id)

			for i := 1; i <= len(idStr)/2; i++ {
				if len(idStr)%i != 0 {
					continue
				}
				found := true
				for j := i; j < len(idStr); j += i {
					if j+i > len(idStr) {
						break
					}
					if idStr[0:i] != idStr[j:j+i] {
						found = false
						break
					}
				}
				if found {
					sum += id
					break
				}
			}
		}
	}
	return sum
}

func main() {
	testCase := utils.ReadLines("./2025/day2/test_case.txt")
	input := utils.ReadLines("./2025/day2/input.txt")

	// PART 1
	fmt.Printf("PART 1 EXAMPLE: %v\n", part1(testCase))
	fmt.Printf("PART 1 RESULT: %v\n", part1(input))

	// // PART 2
	fmt.Printf("PART 2 EXAMPLE: %v\n", part2(testCase))
	fmt.Printf("PART 2 RESULT: %v\n", part2(input))
}
