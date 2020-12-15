require('../../rb_utils/load_file')

test_case = load_file_lines('./test_case.txt').map(&:chomp)
input = load_file_lines('./input.txt').map(&:chomp)

def run_part_1(lines)
  mask = nil
  memory = {}
  lines.each do |line|
    if line.match?(/^mask =/)
      mask = line.split(' = ')[1]
      next
    end
    match = line.match(/mem\[(?<key>\d+)\] = (?<value>\d+)/)
    key = match[:key]
    value = match[:value].to_i.to_s(2).rjust(36, '0')
    (0...36).each do |index|
      next if mask[index] == 'X'
      value[index] = mask[index]
    end
    memory[key.to_sym] = value.to_i(2)
  end
  memory.values.sum
end

puts "Test Case: #{run_part_1(test_case)}"
puts "Input: #{run_part_1(input)}"

test_case_part_2 = load_file_lines('./test_case_part_2.txt').map(&:chomp)

def run_part_2(lines)
  mask = nil
  memory = {}
  lines.each do |line|
    if line.match?(/^mask =/)
      mask = line.split(' = ')[1]
      next
    end

    match = line.match(/mem\[(?<address>\d+)\] = (?<value>\d+)/)
    address = match[:address].to_i.to_s(2).rjust(36, '0')
    value = match[:value].to_i

    address_helper = []
    address.split('').zip(mask.split('')).each do |a, m|
      m == 'X' || m == a ? address_helper << m : address_helper << '1'
    end

    replaces = %w[1 0]
    clean_address = address_helper.clone
    (0...2**address_helper.count('X')).each do |num|
      bin_num = "%0#{address_helper.count('X')}d" % num.to_s(2)

      bin_num.split('').each do |char|
        address_helper[address_helper.index('X')] = char
      end

      memory[address_helper.join('')] = value
      address_helper = clean_address.clone
    end
  end
  memory.values.sum
end

puts "Test Case Part 2: #{run_part_2(test_case_part_2)}"
puts "Input Part 2: #{run_part_2(input)}"
