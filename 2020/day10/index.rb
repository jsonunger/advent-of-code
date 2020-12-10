require('../../rb_utils/load_file')

test_case = load_file_lines('./test_case.txt').map(&:to_i).sort
larger_test_case = load_file_lines('./larger_test_case.txt').map(&:to_i).sort
input = load_file_lines('./input.txt').map(&:to_i).sort

def use_all_adapters(adapters)
  differences = { 1 => 0, 2 => 0, 3 => 1 }
  current_voltage = 0
  adapters.each_with_index do |adapter, index|
    differences[adapter - current_voltage] += 1
    current_voltage = adapter
  end
  differences
end

puts "Test Case: #{use_all_adapters(test_case)}"
puts "Larger Test Case: #{use_all_adapters(larger_test_case)}"
input_result = use_all_adapters(input)
puts "Input: #{input_result}, #{input_result[1] * input_result[3]}"

def count_possible_arrangements(adapters, n = adapters.last, cache = {})
  return 1 if n == 0
  return 0 if !adapters.include?(n)
  cache[n] ||=
    (1..3).sum { |diff| count_possible_arrangements(adapters, n - diff, cache) }
end

puts "Test Case Part 2: #{count_possible_arrangements(test_case)}"
puts "Larger Test Case Part 2: #{count_possible_arrangements(larger_test_case)}"
puts "Input Part 2: #{count_possible_arrangements(input)}"
