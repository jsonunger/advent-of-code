require('../../rb_utils/load_file')

input = load_file_lines('./input.txt')

def get_value(chars, low_letter, min, max)
  options = (min..max).to_a
  chars.split('').each do |char|
    halfway = options.size / 2
    if char == low_letter
      options = options[0..halfway]
    else
      options = options[halfway..-1]
    end
  end
  options[0]
end

def get_seat_id(selection)
  row = get_value(selection[0, 7], 'F', 0, 127)
  col = get_value(selection[7, 3], 'L', 0, 7)
  (row * 8) + col
end

puts "Test Case 1: #{get_seat_id('FBFBBFFRLR')}"
puts "Test Case 2: #{get_seat_id('BFFFBBFRRR')}"
puts "Test Case 3: #{get_seat_id('FFFBBBFRRR')}"
puts "Test Case 4: #{get_seat_id('BBFFBBFRLL')}"

input_seat_ids = input.map(&method(:get_seat_id))

puts "Max Input: #{input_seat_ids.max}"

puts
puts '#########'
puts

all_seat_ids = (input_seat_ids.min...input_seat_ids.max).to_a

puts all_seat_ids - input_seat_ids
