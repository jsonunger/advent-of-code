test_case =
  File.open('./test_case.txt').readlines.map { |line| line.chomp.to_i }
input = File.open('./input.txt').readlines.map { |line| line.chomp.to_i }

def check_sum(lines, sum)
  i = 0
  j = lines.size - 1
  lines.each_with_index do |line, index|
    next if line > sum
    lines[(index + 1)..-1].each do |other_line|
      return true if line + other_line == sum
    end
  end
  false
end

def run_part_1(lines, preamble_length)
  i = preamble_length
  while i < lines.length
    lines_to_check = lines[i - preamble_length, preamble_length]
    return lines[i] if !check_sum(lines_to_check, lines[i])
    i += 1
  end
end

puts "Test Case: #{run_part_1(test_case, 5)}"
puts "Input: #{run_part_1(input, 25)}"

def run_part_2(lines, preamble_length)
  sum = run_part_1(lines, preamble_length)
  matching_set = []

  lines.each_with_index do |line, index|
    next if line > sum
    inner_sum = line
    i = index + 1
    while i < lines.size
      inner_sum += lines[i]
      break if inner_sum == sum
      i += 1
    end
    if inner_sum == sum
      matching_set = lines[index..i]
      break
    end
  end
  matching_set.min + matching_set.max
end

puts "Test Case Part 2: #{run_part_2(test_case, 5)}"
puts "Input Part 2: #{run_part_2(input, 25)}"
