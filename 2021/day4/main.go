package main

import (
	"fmt"
	"math"
	"strings"

	"github.com/jsonunger/advent-of-code/utils"
)

type CellState struct {
	value  int
	called bool
}

type BingoBoard struct {
	board [5][5]*CellState
}

func (b *BingoBoard) Check(draw int) {
	for _, row := range b.board {
		for _, cell := range row {
			if cell.value == draw {
				cell.called = true
			}
		}
	}
}

var inf = int(math.Inf(1))

func (b *BingoBoard) Winner() bool {
	rowResults, colResults := [5]bool{true, true, true, true, true}, [5]bool{true, true, true, true, true}

	for i, row := range b.board {
		for j, col := range row {
			if !col.called {
				rowResults[i] = false
				colResults[j] = false
			}
		}
	}

	for _, res := range rowResults {
		if res {
			return true
		}
	}

	for _, res := range colResults {
		if res {
			return true
		}
	}

	return false
}

func (b *BingoBoard) CalculateScore(draw int) int {
	sum := 0

	for _, row := range b.board {
		for _, col := range row {
			if !col.called {
				sum += col.value
			}
		}
	}

	return sum * draw
}

func newBoard(lines []string) *BingoBoard {
	board := &BingoBoard{}

	for i, line := range lines {
		lineAsInts := utils.ConvertStringsToInts(strings.Fields(line))
		for j := 0; j < 5; j++ {
			board.board[i][j] = &CellState{lineAsInts[j], false}
		}
	}

	return board
}

func part1(lines []string) interface{} {
	draws := utils.ConvertStringsToInts(strings.Split(lines[0], ","))

	boards := []*BingoBoard{}
	for i := 2; i < len(lines); i += 6 {
		boards = append(boards, newBoard(lines[i:i+5]))
	}

	for _, draw := range draws {
		for _, board := range boards {
			board.Check(draw)
			if board.Winner() {
				return board.CalculateScore(draw)
			}
		}
	}

	return nil
}

func lastBoard(boards []*BingoBoard) bool {
	sum := 0
	for _, board := range boards {
		if board == nil {
			sum += 1
		}
	}
	return sum == len(boards)-1
}

func part2(lines []string) interface{} {
	draws := utils.ConvertStringsToInts(strings.Split(lines[0], ","))

	boards := []*BingoBoard{}
	for i := 2; i < len(lines); i += 6 {
		boards = append(boards, newBoard(lines[i:i+5]))
	}

	for _, draw := range draws {
		for i, board := range boards {
			if board == nil {
				continue
			}
			board.Check(draw)
			if board.Winner() {
				if lastBoard(boards) {
					return board.CalculateScore(draw)
				}
				fmt.Println(i, len(boards))
				boards[i] = nil
			}
		}
	}

	return nil
}

func main() {
	testCaseLines := utils.ReadFile("./2021/day4/test_case.txt")
	inputLines := utils.ReadFile("./2021/day4/input.txt")

	// PART 1
	fmt.Printf("PART 1 EXAMPLE: %v\n", part1(testCaseLines))
	fmt.Printf("PART 1 RESULT: %v\n", part1(inputLines))

	// // PART 2
	fmt.Printf("PART 2 EXAMPLE: %v\n", part2(testCaseLines))
	fmt.Printf("PART 2 RESULT: %v\n", part2(inputLines))
}
