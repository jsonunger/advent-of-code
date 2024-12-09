package main

import (
	"fmt"
	"slices"
	"strconv"
	"strings"

	"github.com/jsonunger/advent-of-code/utils"
)

func part1(lines []string) interface{} {
	orderedPages := map[int][]int{}

	buildOrderedPages := true
	sum := 0

	for _, line := range lines {
		if line == "" {
			buildOrderedPages = false
			continue
		}
		if buildOrderedPages {
			pages := strings.Split(line, "|")
			lhs, _ := strconv.Atoi(pages[0])
			rhs, _ := strconv.Atoi(pages[1])
			orderedPages[lhs] = append(orderedPages[lhs], rhs)
		} else {
			pages := strings.Split(line, ",")
			pageNums := []int{}
			for _, page := range pages {
				pageNum, _ := strconv.Atoi(page)
				pageNums = append(pageNums, pageNum)
			}
			correct := true
			for i, page := range pageNums {
				for _, otherPage := range pageNums[i+1:] {
					rules := orderedPages[otherPage]
					if slices.Contains(rules, page) {
						correct = false
						break
					}
				}
				if !correct {
					break
				}
			}
			if correct {
				idx := len(pageNums) / 2
				sum += pageNums[idx]
			}
		}
	}

	return sum
}

func part2(lines []string) interface{} {
	return nil
}

func main() {
	testCase := utils.ReadLines("./2024/day5/test_case.txt")
	input := utils.ReadLines("./2024/day5/input.txt")

	// PART 1
	fmt.Printf("PART 1 EXAMPLE: %v\n", part1(testCase))
	fmt.Printf("PART 1 RESULT: %v\n", part1(input))

	// // PART 2
	// fmt.Printf("PART 2 EXAMPLE: %v\n", part2(testCase))
	// fmt.Printf("PART 2 RESULT: %v\n", part2(input))
}
