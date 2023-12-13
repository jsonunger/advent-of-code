package main

import (
	"fmt"
	"math"
	"regexp"
	"strconv"
	"strings"

	"github.com/jsonunger/advent-of-code/utils"
	"golang.org/x/exp/slices"
)

type Card struct {
	CardNum int
	Winners int
}

func getCard(cardStr string) Card {
	cardRegexp := regexp.MustCompile(`^Card\s+(?P<CardNum>\d+): (?P<WinningNums>[\d\s]+) \| (?P<Numbers>[\d\s]+)$`)
	matches := cardRegexp.FindStringSubmatch(cardStr)

	cardNumStr := matches[cardRegexp.SubexpIndex("CardNum")]
	cardNum, _ := strconv.Atoi(cardNumStr)
	winningNums := strings.Fields(matches[cardRegexp.SubexpIndex("WinningNums")])
	numbers := strings.Fields(matches[cardRegexp.SubexpIndex("Numbers")])

	total := 0
	for _, num := range winningNums {
		if slices.Contains(numbers, num) {
			total++
		}
	}

	return Card{cardNum, total}
}

func (c Card) getPower() float64 {
	if c.Winners == 0 {
		return 0
	}
	return math.Pow(2, float64(c.Winners-1))
}

func part1(lines []string) interface{} {
	sum := 0.0
	for _, line := range lines {
		card := getCard(line)

		sum += card.getPower()
	}
	return sum
}

func part2(lines []string) interface{} {
	sum := 0
	cards := make(map[int]int)
	for _, line := range lines {
		card := getCard(line)
		cards[card.CardNum]++
		for i := 0; i < cards[card.CardNum]; i++ {
			for j := card.CardNum + 1; j < int(math.Min(float64(card.CardNum+card.Winners+1), float64(len(lines)))); j++ {
				cards[j]++
			}
		}
	}
	for _, num := range cards {
		sum += num
	}
	return sum
}

func main() {
	testCase := utils.ReadLines("./2023/day4/test_case.txt")
	input := utils.ReadLines("./2023/day4/input.txt")

	// PART 1
	fmt.Printf("PART 1 EXAMPLE: %v\n", part1(testCase))
	fmt.Printf("PART 1 RESULT: %v\n", part1(input))

	// PART 2
	fmt.Printf("PART 2 EXAMPLE: %v\n", part2(testCase))
	fmt.Printf("PART 2 RESULT: %v\n", part2(input))
}
