require 'set'

test_case = File.open('./test_case.txt').read.split(/\n{2}/)
input = File.open('./input.txt').read.split(/\n{2}/)

p test_case

def count_any_yes_answers(groups)
  count = 0
  groups.each do |group|
    lookup = Set.new
    people = group.split(/\n/)
    people.each { |person| person.split('').each { |char| lookup.add(char) } }
    count += lookup.size
  end
  count
end

puts "Test Case: #{count_any_yes_answers(test_case)}"
puts "Input: #{count_any_yes_answers(input)}"

def count_all_yes_answers(groups)
  count = 0
  groups.each do |group|
    lookup = {}
    people = group.split(/\n/)
    people.each do |person|
      person.split('').each do |char|
        lookup[char] = 0 if !lookup.has_key?(char)
        lookup[char] += 1
      end
    end
    sub_total = 0
    lookup.values.each { |val| sub_total += 1 if val == people.size }
    count += sub_total
  end
  count
end

puts "Test Case: #{count_all_yes_answers(test_case)}"
puts "Input: #{count_all_yes_answers(input)}"
