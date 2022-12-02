package main

import (
	"fmt"
	"math"
	"strings"

	"github.com/jsonunger/advent-of-code/utils"
)

func part1(lines []string) interface{} {
	template := strings.Split(lines[0], "")
	rules := map[[2]string]string{}

	for _, rule := range lines[2:] {
		split := strings.Split(rule, " -> ")
		key := [2]string{string(split[0][0]), string(split[0][1])}
		rules[key] = split[1]
	}

	for i := 0; i < 40; i++ {
		newTemplate := []string{}
		for j := 0; j < len(template)-1; j++ {
			key := [2]string{template[j], template[j+1]}
			newTemplate = append(newTemplate, template[j], rules[key])
		}
		newTemplate = append(newTemplate, template[len(template)-1])
		template = newTemplate
	}

	counts := map[string]int{}
	for _, char := range template {
		counts[char]++
	}

	min, max := math.MaxInt, math.MinInt

	for _, value := range counts {
		if value < min {
			min = value
		}
		if value > max {
			max = value
		}
	}

	return max - min
}

func part2(lines []string) interface{} {
	return nil
}

func main() {
	testCase := utils.ReadLines("./2021/day14/test_case.txt")
	input := utils.ReadLines("./2021/day14/input.txt")

	// PART 1
	fmt.Printf("PART 1 EXAMPLE: %v\n", part1(testCase))
	fmt.Printf("PART 1 RESULT: %v\n", part1(input))

	// // PART 2
	// fmt.Printf("PART 2 EXAMPLE: %v\n", part2(testCase))
	// fmt.Printf("PART 2 RESULT: %v\n", part2(input))
}
