require('../../rb_utils/load_file')

input = load_file_lines('./input.txt')

def get_seat_id(selection)
  rows = selection[0..6]
  low_row = 0
  high_row = 127
  rows.split('').each do |row|
    next_val = (high_row + low_row) / 2.0
    if row == 'F'
      high_row = next_val.floor
    else
      low_row = next_val.ceil
    end
  end
  columns = selection[7..-1]
  low_col = 0
  high_col = 7
  columns.split('').each do |col|
    next_val = (high_col + low_col) / 2.0
    if col == 'L'
      high_col = next_val.floor
    else
      low_col = next_val.ceil
    end
  end
  (low_row * 8) + low_col
end

puts "Test Case 1: #{get_seat_id('FBFBBFFRLR')}"
puts "Test Case 2: #{get_seat_id('BFFFBBFRRR')}"
puts "Test Case 3: #{get_seat_id('FFFBBBFRRR')}"
puts "Test Case 4: #{get_seat_id('BBFFBBFRLL')}"

input_seat_ids = input.map { |row| get_seat_id(row) }

puts "Max Input: #{input_seat_ids.max}"

puts
puts '#########'
puts

(input_seat_ids.min...input_seat_ids.max).each do |val|
  puts val if !input_seat_ids.include?(val)
end
