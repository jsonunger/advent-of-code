require('../../rb_utils/load_file')

test_cases = [
  [0, 3, 6],
  [1, 3, 2],
  [2, 1, 3],
  [1, 2, 3],
  [2, 3, 1],
  [3, 2, 1],
  [3, 1, 2]
]
input = [7, 12, 1, 0, 16, 2]

def run(starting_numbers, number)
  numbers = {}
  last_number = starting_numbers.last
  starting_numbers.each_with_index { |num, index| numbers[num] = [index + 1] }
  (starting_numbers.size + 1..number).each do |turn|
    previous_turns = numbers.fetch(last_number)
    if previous_turns.size == 1
      zero_turns = numbers.fetch(0, [])
      zero_turns << turn
      numbers[0] = zero_turns.last(2)
      last_number = 0
    else
      diff = previous_turns[1] - previous_turns[0]
      diff_turns = numbers.fetch(diff, [])
      diff_turns << turn
      numbers[diff] = diff_turns.last(2)
      last_number = diff
    end
  end
  numbers.keys.select { |num| numbers[num].include?(number) }.first
end

test_cases.each do |test_case|
  puts "Test Case: #{test_case} => #{run(test_case, 2020)}"
end
puts "Input: #{run(input, 2020)}"

puts "Input Part 2: #{run(input, 30_000_000)}"
