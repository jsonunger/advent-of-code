package main

import (
	"fmt"
	"strings"

	"github.com/jsonunger/advent-of-code/utils"
)

type Direction string

const (
	LeftToRight          Direction = "LeftToRight"
	RightToLeft          Direction = "RightToLeft"
	TopToBottom          Direction = "TopToBottom"
	BottomToTop          Direction = "BottomToTop"
	TopLeftToBottomRight Direction = "TopLeftToBottomRight"
	BottomLeftToTopRight Direction = "BottomLeftToTopRight"
	TopRightToBottomLeft Direction = "TopRightToBottomLeft"
	BottomRightToTopLeft Direction = "BottomRightToTopLeft"
)

var directions = []Direction{
	LeftToRight,
	RightToLeft,
	TopToBottom,
	BottomToTop,
	TopLeftToBottomRight,
	BottomLeftToTopRight,
	TopRightToBottomLeft,
	BottomRightToTopLeft,
}

func reverseString(s string) (result string) {
	for _, char := range strings.Split(s, "") {
		result = char + result
	}
	return result
}

func stringFromBytes(bytes ...byte) string {
	return string(bytes)
}

func generateWord(lines []string, line string, row, col int, dir Direction) string {
	switch dir {
	case LeftToRight:
		if col > len(line)-4 {
			return ""
		}
		return line[col : col+4]
	case RightToLeft:
		if col < 3 {
			return ""
		}
		return reverseString(line[col-3 : col+1])
	case TopToBottom:
		if row > len(lines)-4 {
			return ""
		}
		return stringFromBytes(line[col], lines[row+1][col], lines[row+2][col], lines[row+3][col])
	case BottomToTop:
		if row < 3 {
			return ""
		}
		return stringFromBytes(line[col], lines[row-1][col], lines[row-2][col], lines[row-3][col])
	case TopLeftToBottomRight:
		if row > len(lines)-4 || col > len(line)-4 {
			return ""
		}
		return stringFromBytes(line[col], lines[row+1][col+1], lines[row+2][col+2], lines[row+3][col+3])
	case BottomLeftToTopRight:
		if row < 3 || col > len(line)-4 {
			return ""
		}
		return stringFromBytes(line[col], lines[row-1][col+1], lines[row-2][col+2], lines[row-3][col+3])
	case TopRightToBottomLeft:
		if row > len(lines)-4 || col < 3 {
			return ""
		}
		return stringFromBytes(line[col], lines[row+1][col-1], lines[row+2][col-2], lines[row+3][col-3])
	case BottomRightToTopLeft:
		if row < 3 || col < 3 {
			return ""
		}
		return stringFromBytes(line[col], lines[row-1][col-1], lines[row-2][col-2], lines[row-3][col-3])
	}

	return ""
}

func part1(lines []string) interface{} {
	numOccurences := 0
	for i, line := range lines {
		for j, char := range strings.Split(line, "") {
			if char != "X" {
				continue
			}

			for _, direction := range directions {
				if generateWord(lines, line, i, j, direction) == "XMAS" {
					numOccurences++
				}
			}
		}
	}
	return numOccurences
}

func part2(lines []string) interface{} {
	return nil
}

func main() {
	testCase := utils.ReadLines("./2024/day4/test_case.txt")
	input := utils.ReadLines("./2024/day4/input.txt")

	// PART 1
	fmt.Printf("PART 1 EXAMPLE: %v\n", part1(testCase))
	fmt.Printf("PART 1 RESULT: %v\n", part1(input))

	// // PART 2
	// fmt.Printf("PART 2 EXAMPLE: %v\n", part2(testCase))
	// fmt.Printf("PART 2 RESULT: %v\n", part2(input))
}
