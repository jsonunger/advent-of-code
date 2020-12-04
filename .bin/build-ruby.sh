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
require('../../rb_utils/load_file')

test_case = load_file_lines('./test_case.txt')
# input = load_file_lines('./input.txt')

def my_method(rows)
    puts rows
end


puts "Test Case: #{my_method(test_case)}"
# puts "Input: #{my_method(input)}"
EOL
