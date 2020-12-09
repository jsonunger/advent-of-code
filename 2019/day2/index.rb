require('../../rb_utils/load_file')

test_cases =
  load_file_lines('./test_case.txt').map { |line| line.split(',').map(&:to_i) }
input = load_file_lines('./input.txt')[0].split(',').map(&:to_i)

def run_part_1(lines)
  lines_to_manipulate = lines.map(&:clone)
  i = 0
  while true
    break if lines_to_manipulate[i] == 99
    if lines_to_manipulate[i] == 1
      first, second, third = lines_to_manipulate[i + 1, 3]
      lines_to_manipulate[third] =
        lines_to_manipulate[first] + lines_to_manipulate[second]
    elsif lines_to_manipulate[i] == 2
      first, second, third = lines_to_manipulate[i + 1, 3]
      lines_to_manipulate[third] =
        lines_to_manipulate[first] * lines_to_manipulate[second]
    end
    i += 4
  end
  lines_to_manipulate
end

test_cases.each { |test_case| puts "Test Case: #{run_part_1(test_case)}" }

input[1] = 12
input[2] = 2
puts "Input: #{run_part_1(input)[0]}"

def run_part_2(lines)
  (0..lines.size - 1).each do |i|
    (0..lines.size - 1).each do |j|
      used_lines = lines.map(&:clone)
      used_lines[1] = i
      used_lines[2] = j
      output, noun, verb = run_part_1(used_lines)
      return 100 * noun + verb if output == 19_690_720
    end
  end
end

puts "Input Part 2: #{run_part_2(input)}"
