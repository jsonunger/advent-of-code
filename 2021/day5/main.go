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

type VentGrid struct {
	vents []Vent
	x     int
	y     int
}

type Vent struct {
	startX int
	startY int
	endX   int
	endY   int
}

func (v Vent) IsStraight() bool {
	return v.startX == v.endX || v.startY == v.endY
}

type ranges struct {
	minX int
	maxX int
	minY int
	maxY int
}

func (v Vent) startsAndEnds() ranges {
	r := ranges{
		minX: v.startX,
		maxX: v.endX,
		minY: v.startY,
		maxY: v.endY,
	}

	if v.startX > v.endX {
		r.minX, r.maxX = v.endX, v.startX
	}
	if v.startY > v.endY {
		r.minY, r.maxY = v.endY, v.startY
	}

	return r
}

func getMinAndMax(start, end int) (int, int) {
	if start > end {
		return end, start
	}
	return start, end
}

func (v Vent) Points() [][2]int {
	points := [][2]int{}

	r := v.startsAndEnds()

	if v.startX == v.endX {
		for y := r.minY; y <= r.maxY; y++ {
			point := [2]int{v.startX, y}
			points = append(points, point)
		}
		return points
	}

	if v.startY == v.endY {
		for x := r.minX; x <= r.maxX; x++ {
			point := [2]int{x, v.startY}
			points = append(points, point)
		}
		return points
	}

	currX, currY := v.startX, v.startY
	for {
		point := [2]int{currX, currY}
		points = append(points, point)
		if v.startX > v.endX {

			if currX <= v.endX {
				break
			}
			currX -= 1
		} else {
			if currX >= v.endX {
				break
			}
			currX += 1
		}
		if v.startY > v.endY {
			if currY <= v.endY {
				break
			}
			currY -= 1
		} else {
			if currY >= v.endY {
				break
			}
			currY += 1
		}
	}

	return points
}

func buildVentGrid(lines []string) VentGrid {
	ventGrid := VentGrid{}
	for _, line := range lines {
		vents := strings.Split(line, " -> ")
		aPoints := utils.ConvertStringsToInts(strings.Split(vents[0], ","))
		bPoints := utils.ConvertStringsToInts(strings.Split(vents[1], ","))
		vent := Vent{
			aPoints[0],
			aPoints[1],
			bPoints[0],
			bPoints[1],
		}
		ventGrid.x = maxInt(ventGrid.x, vent.startX, vent.endX)
		ventGrid.y = maxInt(ventGrid.y, vent.startY, vent.endY)
		ventGrid.vents = append(ventGrid.vents, vent)
	}
	return ventGrid
}

func buildGrid(x, y int) [][]int {
	grid := [][]int{}
	for i := 0; i <= y; i++ {
		row := []int{}
		for j := 0; j <= x; j++ {
			row = append(row, 0)
		}
		grid = append(grid, row)
	}

	return grid
}

func part1(lines []string) interface{} {
	ventGrid := buildVentGrid(lines)
	grid := buildGrid(ventGrid.x, ventGrid.y)

	for _, vent := range ventGrid.vents {
		if !vent.IsStraight() {
			continue
		}
		for _, point := range vent.Points() {
			grid[point[1]][point[0]] += 1
		}
	}

	sum := 0
	for _, row := range grid {
		for _, point := range row {
			if point >= 2 {
				sum += 1
			}
		}
	}

	return sum
}

func part2(lines []string) interface{} {
	ventGrid := buildVentGrid(lines)
	grid := buildGrid(ventGrid.x, ventGrid.y)

	for _, vent := range ventGrid.vents {
		for _, point := range vent.Points() {
			grid[point[1]][point[0]] += 1
		}
	}

	sum := 0
	for _, row := range grid {
		for _, point := range row {
			if point >= 2 {
				sum += 1
			}
		}
	}

	return sum
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
