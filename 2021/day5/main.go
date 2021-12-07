package main

import (
	"fmt"
	"math"
	"strings"

	"github.com/jsonunger/advent-of-code/utils"
)

func maxInt(gridVal int, vals ...int) int {
	for _, val := range vals {
		gridVal = int(math.Max(float64(gridVal), float64(val)))
	}

	return gridVal
}

type VentGrid []Vent

type Vent []int

func (v Vent) IsStraight() bool {
	return v[0] == v[2] || v[1] == v[3]
}

func (v Vent) Points() [][2]int {
	points := [][2]int{}

	x, y := v[0], v[1]

	for {
		point := [2]int{x, y}
		points = append(points, point)

		if x == v[2] && y == v[3] {
			break
		}

		if v[0] > v[2] {
			x--
		} else if v[0] < v[2] {
			x++
		}

		if v[1] > v[3] {
			y--
		} else if v[1] < v[3] {
			y++
		}
	}

	return points
}

type VentMap [][]int

func (m VentMap) Sum() int {
	sum := 0
	for _, row := range m {
		for _, point := range row {
			if point >= 2 {
				sum++
			}
		}
	}

	return sum
}

func buildVentGridAndMap(lines []string) (VentGrid, VentMap) {
	ventGrid := VentGrid{}
	x, y := 0, 0
	for _, line := range lines {
		var vent Vent = utils.ConvertStringsToInts(strings.FieldsFunc(line, func(r rune) bool {
			return strings.ContainsRune(",-> ", r)
		}))

		x = maxInt(x, vent[0], vent[2])
		y = maxInt(y, vent[1], vent[3])
		ventGrid = append(ventGrid, vent)
	}

	var ventMap VentMap
	for i := 0; i <= y; i++ {
		row := []int{}
		for j := 0; j <= x; j++ {
			row = append(row, 0)
		}
		ventMap = append(ventMap, row)
	}

	return ventGrid, ventMap
}

func part1(lines []string) interface{} {
	ventGrid, ventMap := buildVentGridAndMap(lines)

	for _, vent := range ventGrid {
		if !vent.IsStraight() {
			continue
		}
		for _, point := range vent.Points() {
			ventMap[point[1]][point[0]]++
		}
	}

	return ventMap.Sum()
}

func part2(lines []string) interface{} {
	ventGrid, ventMap := buildVentGridAndMap(lines)

	for _, vent := range ventGrid {
		for _, point := range vent.Points() {
			ventMap[point[1]][point[0]]++
		}
	}

	return ventMap.Sum()
}

func main() {
	testCaseLines := utils.ReadFile("./2021/day5/test_case.txt")
	inputLines := utils.ReadFile("./2021/day5/input.txt")

	// PART 1
	fmt.Printf("PART 1 EXAMPLE: %v\n", part1(testCaseLines))
	fmt.Printf("PART 1 RESULT: %v\n", part1(inputLines))

	// // PART 2
	fmt.Printf("PART 2 EXAMPLE: %v\n", part2(testCaseLines))
	fmt.Printf("PART 2 RESULT: %v\n", part2(inputLines))
}
