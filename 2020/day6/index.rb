require 'set'

test_case = File.open('./test_case.txt').read
input = File.open('./input.txt').read

def count_any_yes_answers(data)
  groups = data.split(/\n{2}/)
  chars = groups.map { |group| Set.new(group.gsub(/\n/, '').split('')) }
  chars.reduce(0) { |sum, group_chars| sum + group_chars.size }
end

puts "Test Case: #{count_any_yes_answers(test_case)}"
puts "Input: #{count_any_yes_answers(input)}"

def count_all_yes_answers(data)
  groups = data.split(/\n{2}/)
  count = 0
  groups.each do |group|
    answers = Set.new(group.gsub(/\n/, '').split(''))
    group.split(/\n/).each do |person|
      answers = answers & Set.new(person.split(''))
    end
    count += answers.size
  end
  count
end

puts "Test Case: #{count_all_yes_answers(test_case)}"
puts "Input: #{count_all_yes_answers(input)}"
