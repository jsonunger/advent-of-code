package main

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/jsonunger/advent-of-code/utils"
	"github.com/zaataylor/cartesian/cartesian"
)

type Equation struct {
	testValue        int
	numbers          []int
	operationSymbols []string
}

func createEquation(equation string, operationSymbols []string) *Equation {
	sides := strings.Split(equation, ":")
	return &Equation{
		testValue:        utils.ConvertStringToInt(sides[0]),
		numbers:          utils.ConvertStringsToInts(strings.Fields(sides[1])),
		operationSymbols: operationSymbols,
	}
}

func NextIndex(ix []int, lens int) {
	for j := len(ix) - 1; j >= 0; j-- {
		ix[j]++
		if j == 0 || ix[j] < lens {
			return
		}
		ix[j] = 0
	}
}

func (e *Equation) generateOperations() [][]string {
	inputSlices := []any{}

	for i := 0; i < len(e.numbers)-1; i++ {
		inputSlices = append(inputSlices, e.operationSymbols)
	}

	v := cartesian.NewCartesianProduct(inputSlices)

	operationCombinations := [][]string{}

	for _, operations := range v.Values() {
		operationCombination := []string{}
		for _, opSymbol := range operations {
			if sym, ok := opSymbol.(string); ok {
				operationCombination = append(operationCombination, sym)
			}
		}
		operationCombinations = append(operationCombinations, operationCombination)
	}

	return operationCombinations
}

func (e *Equation) canEvaluateSuccessfully() bool {
	for _, operations := range e.generateOperations() {
		value := e.numbers[0]
		for i, num := range e.numbers[1:] {
			if operations[i] == "+" {
				value = value + num
			} else if operations[i] == "*" {
				value = value * num
			} else if operations[i] == "||" {
				lhs, rhs := strconv.Itoa(value), strconv.Itoa(num)
				strValue := lhs + rhs
				value = utils.ConvertStringToInt(strValue)
			}
		}
		if value == e.testValue {
			return true
		}
	}

	return false
}

func part1(lines []string) interface{} {
	sum := 0
	for _, equationStr := range lines {
		eq := createEquation(equationStr, []string{"+", "*"})
		if eq.canEvaluateSuccessfully() {
			sum += eq.testValue
		}
	}
	return sum
}

func part2(lines []string) interface{} {
	sum := 0
	for _, equationStr := range lines {
		eq := createEquation(equationStr, []string{"+", "*", "||"})
		if eq.canEvaluateSuccessfully() {
			sum += eq.testValue
		}
	}
	return sum
}

func main() {
	testCase := utils.ReadLines("./2024/day7/test_case.txt")
	input := utils.ReadLines("./2024/day7/input.txt")

	// PART 1
	fmt.Printf("PART 1 EXAMPLE: %v\n", part1(testCase))
	fmt.Printf("PART 1 RESULT: %v\n", part1(input))

	// PART 2
	fmt.Printf("PART 2 EXAMPLE: %v\n", part2(testCase))
	fmt.Printf("PART 2 RESULT: %v\n", part2(input))
}
