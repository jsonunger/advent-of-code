require('../../rb_utils/load_file')

test_case = load_file_lines('./test_case.txt')
input = load_file_lines('./input.txt')

def run_part_1(lines)
  minutes_to_wait = lines[0].to_i
  buses = lines[1].split(',').select { |line| line != 'x' }.map(&:to_i)
  buses_and_time_to_wait =
    buses.map { |bus| [bus, bus - (minutes_to_wait % bus)] }
  result = buses_and_time_to_wait.min { |bus_a, bus_b| bus_a[1] <=> bus_b[1] }
  result[0] * result[1]
end

puts "Test Case: #{run_part_1(test_case)}"
puts "Input: #{run_part_1(input)}"

def run_part_2(lines)
  buses =
    lines[1].split(',').each_with_index.select do |bus, index|
      bus != 'x'
    end.map { |bus, index| [bus.to_i, index] }
  timestamp = 0
  step = 1
  buses.each do |bus, index|
    (timestamp..Float::INFINITY).step(step) do |n|
      if (n + index) % bus == 0
        timestamp = n.to_i
        break
      end
    end
    step *= bus
  end
  timestamp
end

puts "Test Case Part 2: #{run_part_2(test_case)}"
puts "Input Part 2: #{run_part_2(input)}"
