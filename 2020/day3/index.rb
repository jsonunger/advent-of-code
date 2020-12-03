test_case = %w[
  ..##.......
  #...#...#..
  .#....#..#.
  ..#.#...#.#
  .#...##..#.
  ..#.##.....
  .#.#.#....#
  .#........#
  #.##...#...
  #...##....#
  .#..#...#.#
]
input = File.open('./input.txt').readlines.map(&:chomp)

def get_num_trees_hit(rows, x, y)
  curr_x = 0
  curr_y = 0
  count = 0
  while curr_y < rows.size
    row = rows[curr_y]
    count += 1 if row[curr_x % row.size] == '#'
    curr_x += x
    curr_y += y
  end
  count
end

puts 'Part 1'
puts "Test Case: #{get_num_trees_hit(test_case, 3, 1)}"
puts "Input: #{get_num_trees_hit(input, 3, 1)}"

slopes = [[1, 1], [3, 1], [5, 1], [7, 1], [1, 2]]

def get_product_of_num_trees_hit_for_different_slopes(rows, slopes)
  slopes.reduce 1 do |product, scope|
    product * get_num_trees_hit(rows, scope[0], scope[1])
  end
end

puts
puts '#######'
puts
puts 'Part 2'
puts "Test Case: #{
       get_product_of_num_trees_hit_for_different_slopes(test_case, slopes)
     }"
puts "Input: #{
       get_product_of_num_trees_hit_for_different_slopes(input, slopes)
     }"
