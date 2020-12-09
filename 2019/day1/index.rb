require('../../rb_utils/load_file')

test_case = load_file_lines('./test_case.txt').map(&:to_i)
input = load_file_lines('./input.txt').map(&:to_i)

def get_required_fuel(amount)
  [(amount / 3) - 2, 0].max
end

def run_part_1(lines)
  lines.map { |line| get_required_fuel(line) }
end

puts "Test Case: #{run_part_1(test_case)}"
puts "Input: #{run_part_1(input).sum}"

def get_required_fuel_recursive(amount)
  return 0 if amount == 0
  required_amount = get_required_fuel(amount)
  required_amount + get_required_fuel_recursive(required_amount)
end

def run_part_2(lines)
  lines.map { |line| get_required_fuel_recursive(line) }
end

puts "Test Case Part 2: #{run_part_2(test_case)}"
puts "Input Part 2: #{run_part_2(input).sum}"
