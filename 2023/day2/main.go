package main

import (
	"fmt"
	"math"
	"regexp"
	"strconv"
	"strings"

	"github.com/jsonunger/advent-of-code/utils"
)

func part1(lines []string) interface{} {
	maxColors := map[string]int{"red": 12, "green": 13, "blue": 14}
	sum := 0

	for _, line := range lines {
		gamesRegex := regexp.MustCompile(`Game (?P<ID>\d+): (?P<Rounds>.*)`)
		matches := gamesRegex.FindStringSubmatch(line)
		roundsIndex := gamesRegex.SubexpIndex("Rounds")
		gameID := matches[gamesRegex.SubexpIndex("ID")]
		rounds := strings.Split(matches[roundsIndex], ";")
		valid := true
		for _, round := range rounds {
			cubePieces := strings.Split(strings.TrimSpace(round), ",")
			for _, cubePiece := range cubePieces {
				cubeRegex := regexp.MustCompile(`(?P<Num>\d+) (?P<Color>.*)`)
				matches := cubeRegex.FindStringSubmatch(strings.TrimSpace(cubePiece))
				num, color := matches[cubeRegex.SubexpIndex("Num")], matches[cubeRegex.SubexpIndex("Color")]
				numInt, _ := strconv.Atoi(num)
				if numInt > maxColors[color] {
					valid = false
					break
				}
			}
			if !valid {
				break
			}
		}
		if valid {
			gameIDInt, _ := strconv.Atoi(gameID)
			sum += gameIDInt
		}
	}
	return sum
}

func part2(lines []string) interface{} {

	sum := 0

	for _, line := range lines {
		gamesRegex := regexp.MustCompile(`Game (?P<ID>\d+): (?P<Rounds>.*)`)
		matches := gamesRegex.FindStringSubmatch(line)
		roundsIndex := gamesRegex.SubexpIndex("Rounds")
		rounds := strings.Split(matches[roundsIndex], ";")
		colors := map[string]int{"red": math.MinInt, "green": math.MinInt, "blue": math.MinInt}
		for _, round := range rounds {
			cubePieces := strings.Split(strings.TrimSpace(round), ",")
			for _, cubePiece := range cubePieces {
				cubeRegex := regexp.MustCompile(`(?P<Num>\d+) (?P<Color>.*)`)
				matches := cubeRegex.FindStringSubmatch(strings.TrimSpace(cubePiece))
				num, color := matches[cubeRegex.SubexpIndex("Num")], matches[cubeRegex.SubexpIndex("Color")]
				numInt, _ := strconv.Atoi(num)
				if numInt > colors[color] {
					colors[color] = numInt
				}
			}
		}
		product := 1
		for _, num := range colors {
			product *= num
		}
		sum += product
	}
	return sum
}

func main() {
	testCase := utils.ReadLines("./2023/day2/test_case.txt")
	input := utils.ReadLines("./2023/day2/input.txt")

	// PART 1
	fmt.Printf("PART 1 EXAMPLE: %v\n", part1(testCase))
	fmt.Printf("PART 1 RESULT: %v\n", part1(input))

	// PART 2
	fmt.Printf("PART 2 EXAMPLE: %v\n", part2(testCase))
	fmt.Printf("PART 2 RESULT: %v\n", part2(input))
}
