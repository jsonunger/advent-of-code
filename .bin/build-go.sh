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
)

func part1(lines []string) interface{} {
	return nil
}

func part2(lines []string) interface{} {
	return nil
}

func main() {
	testCaseLines := utils.ReadFile("./$day_dir/test_case.txt")
	inputLines := utils.ReadFile("./$day_dir/input.txt")

	// PART 1
	fmt.Printf("PART 1 EXAMPLE: %v\n", part1(testCaseLines))
	fmt.Printf("PART 1 RESULT: %v\n", part1(inputLines))

	// // PART 2
	// fmt.Printf("PART 2 EXAMPLE: %v\n", part2(testCaseLines))
	// fmt.Printf("PART 2 RESULT: %v\n", part2(inputLines))
}
EOL
fi
