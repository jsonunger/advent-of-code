#!/bin/bash
# Build directory for running ruby code

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
    echo "$day_dir already exists"
    exit 1
fi

echo "Creating $day_dir directory..."

mkdir -p $day_dir

echo "Adding base files"

cd $day_dir

touch test_case.txt
touch input.txt

cat >> index.rb <<EOL
test_case = File.open('./test_case.txt').readlines
# input = File.open('./input.txt').readlines

def run_part_1(lines)
    lines
end

puts "Test Case: #{run_part_1(test_case)}"
# puts "Input: #{run_part_2(input)}"

def run_part_2(lines)
    lines
end

# puts "Test Case Part 2: #{run_part_2(test_case)}"
# puts "Input Part 2: #{run_part_2(input)}
EOL
