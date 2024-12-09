package main

import (
	"fmt"
	"strings"

	"github.com/jsonunger/advent-of-code/utils"
)

type Site struct {
	x int
	y int
}

type Direction string

const (
	Left  Direction = "left"
	Up    Direction = "up"
	Right Direction = "right"
	Down  Direction = "down"
)

func part1(lines []string) interface{} {
	var location Site
	var direction Direction
	obstacles := map[Site]bool{}
	visitedSites := map[Site]bool{}
	maxX, maxY := 0, len(lines)
	for y, line := range lines {
		chars := strings.Split(line, "")
		if maxX == 0 {
			maxX = len(chars)
		}

		for x, char := range chars {
			if char == "#" {
				site := Site{x, y}
				obstacles[site] = true
			} else if char != "." {
				location = Site{x, y}
				visitedSites[location] = true
				switch char {
				case "^":
					direction = Up
				case "v":
					direction = Down
				case ">":
					direction = Right
				case "<":
					direction = Left
				}
			}
		}
	}

	leftTheMap := false

	for !leftTheMap {
		visitedSites[location] = true
		switch direction {
		case Left:
			nextLocation := Site{location.x - 1, location.y}
			if obstacles[nextLocation] {
				direction = Up
			} else if nextLocation.x < 0 {
				leftTheMap = true
			} else {
				visitedSites[nextLocation] = true
				location = nextLocation
			}
		case Right:
			nextLocation := Site{location.x + 1, location.y}
			if obstacles[nextLocation] {
				direction = Down
			} else if nextLocation.x >= maxX {
				leftTheMap = true
			} else {
				visitedSites[nextLocation] = true
				location = nextLocation
			}
		case Up:
			nextLocation := Site{location.x, location.y - 1}
			if obstacles[nextLocation] {
				direction = Right
			} else if nextLocation.y < 0 {
				leftTheMap = true
			} else {
				visitedSites[nextLocation] = true
				location = nextLocation
			}
		case Down:
			nextLocation := Site{location.x, location.y + 1}
			if obstacles[nextLocation] {
				direction = Left
			} else if nextLocation.y >= maxY {
				leftTheMap = true
			} else {
				visitedSites[nextLocation] = true
				location = nextLocation
			}
		}
	}

	return len(visitedSites)
}

func part2(lines []string) interface{} {
	return nil
}

func main() {
	testCase := utils.ReadLines("./2024/day6/test_case.txt")
	input := utils.ReadLines("./2024/day6/input.txt")

	// PART 1
	fmt.Printf("PART 1 EXAMPLE: %v\n", part1(testCase))
	fmt.Printf("PART 1 RESULT: %v\n", part1(input))

	// // PART 2
	// fmt.Printf("PART 2 EXAMPLE: %v\n", part2(testCase))
	// fmt.Printf("PART 2 RESULT: %v\n", part2(input))
}
