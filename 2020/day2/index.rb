test_case = ['1-3 a: abcde', '1-3 b: cdefg', '2-9 c: ccccccccc']
input = File.open('./input.txt').readlines

REGEX = /(?<min>\d+)-(?<max>\d+) (?<letter>\w): (?<password>\w+)/

def is_valid?(input)
  matches = input.match REGEX
  return false if !matches
  num_letter = matches[:password].count(matches[:letter])
  num_letter >= matches[:min].to_i && num_letter <= matches[:max].to_i
end

def num_valid(inputs)
  inputs.select { |i| is_valid?(i) }.count
end

puts "Test Case: #{num_valid(test_case)}"
puts "Input: #{num_valid(input)}"

def is_valid_v2?(input)
  matches = input.match REGEX
  return false if !matches
  (matches[:password][matches[:min].to_i - 1] == matches[:letter]) ^
    (matches[:password][matches[:max].to_i - 1] == matches[:letter])
end

def num_valid_v2(inputs)
  inputs.select { |i| is_valid_v2?(i) }.count
end

puts "Test Case: #{num_valid_v2(test_case)}"
puts "Input: #{num_valid_v2(input)}"
