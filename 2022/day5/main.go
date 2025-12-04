package main

import (
	"fmt"
	"os"
	"regexp"
	"strings"

	"github.com/jsonunger/advent-of-code/utils"
)

type Ship struct {
	Stacks      map[int]Stack
	readyToMove bool
}

func (s *Ship) MarkReady() {
	s.readyToMove = true
}

func (s *Ship) IsReady() bool {
	return s.readyToMove
}

func (s *Ship) AddCrates(line string) {
	chunks := utils.ChunkString(line, 4)
	for i, chunk := range chunks {
		chunk = strings.TrimSpace(chunk)

		crateRegexp := regexp.MustCompile(`\[([A-Z])\]`)
		matches := crateRegexp.FindStringSubmatch(chunk)
		if len(matches) == 0 {
			continue
		}

		if _, ok := s.Stacks[i+1]; !ok {
			s.Stacks[i+1] = Stack{}
		}
		stack := s.Stacks[i+1]
		stack.Unshift(matches[1])
		s.Stacks[i+1] = stack
	}
}

func (s *Ship) ExecutePart1Move(move Move) error {
	fromStack := s.Stacks[move.From]
	toStack := s.Stacks[move.To]

	for i := 0; i < move.Amount; i++ {
		crate, ok := fromStack.Pop()
		if !ok {
			return fmt.Errorf("not enought crate to move from stack %d", move.From)
		}
		toStack.Push(crate)
	}

	s.Stacks[move.From] = fromStack
	s.Stacks[move.To] = toStack

	return nil
}

func (s *Ship) ExecutePart2Move(move Move) error {
	fromStack := s.Stacks[move.From]
	toStack := s.Stacks[move.To]

	crates, ok := fromStack.PopN(move.Amount)

	if !ok {
		return fmt.Errorf("not enough crates to move from stack %d", move.From)
	}

	toStack.Push(crates...)

	s.Stacks[move.From] = fromStack
	s.Stacks[move.To] = toStack

	return nil
}

func (s *Ship) GetTopCrates() string {
	topCratesSlice := make([]string, len(s.Stacks))
	for i, stack := range s.Stacks {
		topCratesSlice[i-1] = stack[len(stack)-1]
	}
	return strings.Join(topCratesSlice, "")
}

type Stack []string

func (s *Stack) Unshift(v string) {
	*s = append([]string{v}, *s...)
}

func (s *Stack) Push(v ...string) {
	*s = append(*s, v...)
}

func (s *Stack) pop(num int) ([]string, bool) {
	if len(*s) < num {
		return []string{}, false
	}
	last := (*s)[len(*s)-num:]
	*s = (*s)[:len(*s)-num]
	return last, true
}

func (s *Stack) Pop() (string, bool) {
	crates, ok := s.pop(1)
	if !ok {
		return "", false
	}
	return crates[0], true
}

func (s *Stack) PopN(num int) ([]string, bool) {
	return s.pop(num)
}

func (s *Stack) IsEmpty() bool {
	return len(*s) == 0
}

type Move struct {
	Amount int
	From   int
	To     int
}

func parseMove(line string) Move {
	re := regexp.MustCompile(`move (\d+) from (\d+) to (\d+)`)
	matches := re.FindStringSubmatch(line)
	return Move{
		Amount: utils.ConvertStringToInt(matches[1]),
		From:   utils.ConvertStringToInt(matches[2]),
		To:     utils.ConvertStringToInt(matches[3]),
	}
}

func part1(lines []string) interface{} {
	ship := Ship{
		Stacks: map[int]Stack{},
	}

	for _, line := range lines {
		if len(line) == 0 {
			ship.MarkReady()
			continue
		}

		if !ship.IsReady() {
			ship.AddCrates(line)
			continue
		}

		move := parseMove(line)
		err := ship.ExecutePart1Move(move)
		if err != nil {
			utils.PanicWithMsg(err, "unable to execute move")
		}
	}

	return ship.GetTopCrates()
}

func part2(lines []string) interface{} {
	ship := Ship{
		Stacks: map[int]Stack{},
	}

	for _, line := range lines {
		if len(line) == 0 {
			ship.MarkReady()
			continue
		}

		if !ship.IsReady() {
			ship.AddCrates(line)
			continue
		}

		move := parseMove(line)
		err := ship.ExecutePart2Move(move)
		if err != nil {
			utils.PanicWithMsg(err, "unable to execute move")
		}
	}

	return ship.GetTopCrates()
}

func readFileWithoutTrim(fileName string) string {
	data, err := os.ReadFile(fileName)
	utils.PanicWithMsg(err, "getting working directory")
	return string(data)
}

func main() {
	testCase := strings.Split(readFileWithoutTrim("./2022/day5/test_case.txt"), "\n")
	input := strings.Split(readFileWithoutTrim("./2022/day5/input.txt"), "\n")

	// PART 1
	fmt.Printf("PART 1 EXAMPLE: %v\n", part1(testCase))
	fmt.Printf("PART 1 RESULT: %v\n", part1(input))

	// // PART 2
	fmt.Printf("PART 2 EXAMPLE: %v\n", part2(testCase))
	fmt.Printf("PART 2 RESULT: %v\n", part2(input))
}
