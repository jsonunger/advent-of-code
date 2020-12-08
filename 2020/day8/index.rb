require('set')
require('../../rb_utils/load_file')

test_case = load_file_lines('./test_case.txt')
input = load_file_lines('./input.txt')

def get_accumulator_pre_infinite_loop(lines)
  visited_indices = Set.new
  accumulator = 0
  i = 0
  while !visited_indices.include?(i)
    visited_indices.add(i)
    command, value = lines[i].split(' ')
    case command
    when 'acc'
      accumulator += value.to_i
      i += 1
    when 'jmp'
      i += value.to_i
    else
      i += 1
    end
  end
  accumulator += lines[i].split(' ')[1].to_i if lines[i].start_with?('acc')
  accumulator
end

puts "Test Case: #{get_accumulator_pre_infinite_loop(test_case)}"
puts "Input: #{get_accumulator_pre_infinite_loop(input)}"

def get_accumulator_and_if_completed(lines)
  visited_indices = Set.new
  accumulator = 0
  i = 0
  last_i = 0
  while !visited_indices.include?(i) && i < lines.size - 1
    visited_indices.add(i)
    last_i = i
    command, value = lines[i].split(' ')
    case command
    when 'acc'
      accumulator += value.to_i
      i += 1
    when 'jmp'
      i += value.to_i
    else
      i += 1
    end
  end

  accumulator += lines[i].split(' ')[1].to_i if lines[i].start_with?('acc')
  [accumulator, i >= lines.size - 1]
end

def check_for_broken_line(lines)
  num_iterations = 0
  lines.each_with_index do |line, index|
    num_iterations += 1
    next if line.start_with?('acc')
    new_lines =
      lines.each_with_index.map do |test_line, test_index|
        if test_index != index
          test_line
        elsif test_line.start_with?('jmp')
          test_line.sub('jmp', 'nop')
        else
          test_line.sub('nop', 'jmp')
        end
      end
    accumulator, finished = get_accumulator_and_if_completed(new_lines)
    return accumulator if finished
  end
end

puts "Test Case Part 2: #{check_for_broken_line(test_case)}"
puts "Input Part 2: #{check_for_broken_line(input)}"
