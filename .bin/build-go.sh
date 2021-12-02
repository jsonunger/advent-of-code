#!/bin/bash
# Build directory for running go code

cwd=$(pwd)
year=$(date +'%Y')
day=$(date +'%-d')

while getopts d:y: flag
do
    case "${flag}" in
        y) year=${OPTARG};;
        d) day=${OPTARG};;
    esac
done

day_dir="$year/day$day"

echo -e "Checking for $day_dir..."

if [[ -d "$day_dir" ]]; then
    echo "$day_dir already exists, not creating"
else
    echo "Creating $day_dir directory..."
    mkdir -p $day_dir
fi

cd $day_dir

if [[ -f "main.go" ]]; then
    echo "main.go already exists, not creating"
    exit 0
else
    echo "Adding base files"
    touch test_case.txt
    touch input.txt
    
    cat >> main.go <<EOL
package main

import (
	"fmt"

	"github.com/jsonunger/advent-of-code/utils"
	"github.com/pkg/errors"
)

func part1(lines []string) (interface{}, error) {
	return nil, errors.New("solution not found")
}

func part2(lines []string) (interface{}, error) {
	return nil, errors.New("solution not found")
}

func main() {
	testCaseLines := utils.ReadFile("./$day_dir/test_case.txt")
	inputLines := utils.ReadFile("./$day_dir/input.txt")

	// PART 1
	exRes1, err := part1(testCaseLines)
	utils.PanicWithMsg(err, "part 1 example")
	fmt.Printf("PART 1 EXAMPLE: %v\n", exRes1)

	res1, err := part1(inputLines)
	utils.PanicWithMsg(err, "part 1")
	fmt.Printf("PART 1 RESULT: %v\n", res1)

	// // PART 2
	// exRes2, err := part2(testCaseLines)
	// utils.PanicWithMsg(err, "part 2 example")
	// fmt.Printf("PART 2 EXAMPLE: %v\n", exRes2)

	// res2, err := part2(inputLines)
	// utils.PanicWithMsg(err, "part 2")
	// fmt.Printf("PART 2 RESULT: %v\n", res2)
}
EOL
fi
