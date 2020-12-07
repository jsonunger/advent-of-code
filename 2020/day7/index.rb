require('../../rb_utils/load_file')

test_case = load_file_lines('./test_case.txt')
input = load_file_lines('./input.txt')

def build_hash(rows)
  my_hash = {}
  rows.each do |row|
    match = row.match(/(?<bag_type>\w+ \w+) bags? contain (?<contains>.*)./)
    bag_type = match[:bag_type]
    contains = match[:contains].split(',').map(&:strip)
    contains_hash = {}
    contains.each do |sub_type|
      next if sub_type == 'no other bags'
      sub_match = sub_type.match(/(?<count>\d+) (?<bag_type>\w+ \w+) bags?/)
      contains_hash[sub_match[:bag_type]] = sub_match[:count].to_i
    end
    my_hash[match[:bag_type]] = contains_hash
  end
  my_hash
end

def get_count_available(bag_hash, bag)
  return 0 if !bag_hash.has_key?('shiny gold')

  return 1 if bag_hash[bag].has_key?('shiny gold')

  bag_hash[bag].keys.each do |sub_bag|
    return 1 if get_count_available(bag_hash, sub_bag) == 1
  end

  return 0
end

def part_1(rows)
  bag_hash = build_hash(rows)

  bag_hash.keys.map { |bag| get_count_available(bag_hash, bag) }.sum
end

puts "Test Case: #{part_1(test_case)}"
puts "Input: #{part_1(input)}"

test_case_part_2 = load_file_lines('./test_case_part_2.txt')

def get_count_sub_bags(bag_hash, bag)
  bag_hash[bag].map do |sub_bag, sub_count|
    get_count_sub_bags(bag_hash, sub_bag) * sub_count
  end.sum + 1
end

def part_2(rows)
  bag_hash = build_hash(rows)

  bag_hash['shiny gold'].map do |bag, count|
    get_count_sub_bags(bag_hash, bag) * count
  end.sum
end

puts "Test Case: #{part_2(test_case_part_2)}"
puts "Input: #{part_2(input)}"
