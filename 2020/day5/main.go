package main

import (
	"fmt"
	"sort"
	"strconv"
	"strings"

	"github.com/jsonunger/advent-of-code/utils"
)

func generateSeatIDs(lines []string) []float64 {
	seatIDs := []float64{}
	for _, line := range lines {
		row, seat := line[:7], line[7:]

		row = strings.ReplaceAll(row, "F", "0")
		row = strings.ReplaceAll(row, "B", "1")
		rowNum, err := strconv.ParseInt(row, 2, 8)
		utils.PanicWithMsg(err, "converting binary")

		seat = strings.ReplaceAll(seat, "R", "1")
		seat = strings.ReplaceAll(seat, "L", "0")
		seatNum, err := strconv.ParseInt(seat, 2, 4)
		utils.PanicWithMsg(err, "converting binary")
		seatIDs = append(seatIDs, float64(rowNum*8+seatNum))
	}
	return seatIDs
}

func part1(lines []string) int {
	highestSeatID := 0
	for _, seatID := range generateSeatIDs(lines) {
		if seatID > float64(highestSeatID) {
			highestSeatID = int(seatID)
		}
	}
	return highestSeatID
}

func part2(lines []string) float64 {
	seatIDs := generateSeatIDs(lines)
	sort.Float64s(seatIDs)
	for i, seatID := range seatIDs {
		if seatIDs[i+1] == seatID+2 {
			return seatID + 1
		}
	}
	return 0
}

func main() {
	testCaseLines := utils.ReadFile("./2020/day5/test_case.txt")
	inputLines := utils.ReadFile("./2020/day5/input.txt")

	// PART 1
	fmt.Printf("PART 1 EXAMPLE: %v\n", part1(testCaseLines))
	fmt.Printf("PART 1 RESULT: %v\n", part1(inputLines))

	// PART 2
	fmt.Printf("PART 2 RESULT: %v\n", part2(inputLines))
}
