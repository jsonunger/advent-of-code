require 'set'
require('../../rb_utils/load_file')

test_case = load_file_lines('./test_case.txt')
input = load_file_lines('./input.txt')

def part_1(rows)
  passport_hashes =
    rows.reduce([Set[]]) do |hashes, row|
      if row == ''
        hashes << Set[]
        next hashes
      end
      row.split('').each { |val| hashes[-1].add(val) }
      hashes
    end
  passport_hashes.reduce(0) { |sum, s| sum + s.size }
end

puts "Test Case: #{part_1(test_case)}"
puts "Input: #{part_1(input)}"

def part_2(rows)
  passport_hashes =
    rows.reduce([{}]) do |hashes, row|
      if row == ''
        hashes << {}
        next hashes
      end
      hashes[-1][:num] ||= 0
      hashes[-1][:num] += 1
      row.split('').each do |val|
        hashes[-1][val] ||= 0
        hashes[-1][val] += 1
      end
      hashes
    end
  passport_hashes.reduce(0) do |sum, h|
    size = h[:num]
    h.keys.reduce(sum) do |inner_sum, key|
      if key == :num || h[key] != size
        inner_sum
      else
        inner_sum += 1
      end
    end
  end
end

puts "Test Case: #{part_2(test_case)}"
puts "Input: #{part_2(input)}"
