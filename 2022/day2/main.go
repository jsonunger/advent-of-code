package main

import (
	"fmt"
	"strings"

	"github.com/jsonunger/advent-of-code/utils"
)

var part1Outcomes = map[string]int{
	"A X": 3 + 1,
	"A Y": 6 + 2,
	"A Z": 0 + 3,
	"B X": 0 + 1,
	"B Y": 3 + 2,
	"B Z": 6 + 3,
	"C X": 6 + 1,
	"C Y": 0 + 2,
	"C Z": 3 + 3,
}

func part1(lines []string) interface{} {
	total := 0
	for _, roundStr := range lines {
		total += part1Outcomes[roundStr]
	}
	return total
}

type OpponentMove string

const (
	Rock     OpponentMove = "A"
	Paper                 = "B"
	Scissors              = "C"
)

func (om *OpponentMove) getValue(outcome Outcome) int {
	if *om == Rock {
		if outcome == Win {
			return 2
		}
		if outcome == Lose {
			return 3
		}
		return 1
	}
	if *om == Paper {
		if outcome == Win {
			return 3
		}
		if outcome == Lose {
			return 1
		}
		return 2
	}
	if outcome == Win {
		return 1
	}
	if outcome == Lose {
		return 2
	}
	return 3
}

type Outcome string

const (
	Win  Outcome = "Z"
	Lose         = "X"
	Draw         = "Y"
)

func (o *Outcome) getValue() int {
	switch *o {
	case Win:
		return 6
	case Lose:
		return 0
	case Draw:
		return 3
	}
	return 0
}

type Round struct {
	theirMove OpponentMove
	outcome   Outcome
}

func (r *Round) getValue() int {
	return r.outcome.getValue() + r.theirMove.getValue(r.outcome)
}

func part2(lines []string) interface{} {
	total := 0
	for _, roundStr := range lines {
		vals := strings.Split(roundStr, " ")
		round := Round{
			theirMove: OpponentMove(vals[0]),
			outcome:   Outcome(vals[1]),
		}
		total += round.getValue()
	}
	return total
}

func main() {
	testCase := utils.ReadLines("./2022/day2/test_case.txt")
	input := utils.ReadLines("./2022/day2/input.txt")

	// PART 1
	fmt.Printf("PART 1 EXAMPLE: %v\n", part1(testCase))
	fmt.Printf("PART 1 RESULT: %v\n", part1(input))

	// PART 2
	fmt.Printf("PART 2 EXAMPLE: %v\n", part2(testCase))
	fmt.Printf("PART 2 RESULT: %v\n", part2(input))
}
