package main

import (
	"fmt"

	"github.com/jsonunger/advent-of-code/utils"
)

type Grid struct {
	maxY   int
	maxX   int
	points map[int]map[int]Point
}

func (g *Grid) isInBounds(x, y int) bool {
	return x >= 0 && x < g.maxX && y >= 0 && y < g.maxY
}

type Coordinate struct {
	y int
	x int
}

type Point struct {
	Coordinate
	Char rune
}

func (p *Point) SurroundingPointCoordinates() []Coordinate {
	return []Coordinate{
		{p.y - 1, p.x - 1}, // top-left
		{p.y - 1, p.x},     // top
		{p.y - 1, p.x + 1}, // top-right
		{p.y, p.x + 1},     // right
		{p.y + 1, p.x + 1}, // bottom-right
		{p.y + 1, p.x},     // bottom
		{p.y + 1, p.x - 1}, // bottom-left
		{p.y, p.x - 1},     // left
	}
}

func buildGrid(lines []string) *Grid {
	grid := &Grid{
		maxY:   len(lines),
		maxX:   len(lines[0]),
		points: map[int]map[int]Point{},
	}
	for y, line := range lines {
		for x, char := range line {
			coord := Coordinate{y, x}
			if _, found := grid.points[coord.y]; !found {
				grid.points[coord.y] = map[int]Point{}
			}
			grid.points[coord.y][coord.x] = Point{coord, char}
		}
	}
	return grid
}

func (g *Grid) IsPointAccessible(point Point) bool {
	if point.Char != '@' {
		return false
	}
	accessiblePoints := 0
	coordinates := point.SurroundingPointCoordinates()
	for _, coordinate := range coordinates {
		if !g.isInBounds(coordinate.x, coordinate.y) {
			continue
		}
		surroundingPoint := g.points[coordinate.y][coordinate.x]
		if surroundingPoint.Char == '@' {
			accessiblePoints++
		}
	}
	return accessiblePoints < 4
}

func part1(lines []string) interface{} {
	grid := buildGrid(lines)

	accessiblePoints := 0
	for _, row := range grid.points {
		for _, point := range row {
			if grid.IsPointAccessible(point) {
				accessiblePoints++
			}
		}
	}
	return accessiblePoints
}

func iterateGrid(grid *Grid) (int, []Coordinate) {
	accessiblePoints := 0
	coordinates := []Coordinate{}
	for _, row := range grid.points {
		for _, point := range row {
			if grid.IsPointAccessible(point) {
				accessiblePoints++
				coordinates = append(coordinates, point.Coordinate)
			}
		}
	}
	return accessiblePoints, coordinates
}

func part2(lines []string) interface{} {
	grid := buildGrid(lines)
	totalAccessiblePoints := 0
	for {
		accessiblePoints, coordinates := iterateGrid(grid)
		if accessiblePoints == 0 {
			break
		}
		totalAccessiblePoints += accessiblePoints
		for _, coordinate := range coordinates {
			point := grid.points[coordinate.y][coordinate.x]
			point.Char = 'x'
			grid.points[coordinate.y][coordinate.x] = point
		}
	}
	return totalAccessiblePoints
}

func main() {
	testCase := utils.ReadLines("./2025/day4/test_case.txt")
	input := utils.ReadLines("./2025/day4/input.txt")

	// PART 1
	fmt.Printf("PART 1 EXAMPLE: %v\n", part1(testCase))
	fmt.Printf("PART 1 RESULT: %v\n", part1(input))

	// // PART 2
	fmt.Printf("PART 2 EXAMPLE: %v\n", part2(testCase))
	fmt.Printf("PART 2 RESULT: %v\n", part2(input))
}
