require('../../rb_utils/load_file')

test_case = load_file_lines('./test_case.txt')
input = load_file_lines('./input.txt')

@instruction_to_direction = {
  'N' => :lat, 'S' => :lat, 'E' => :lng, 'W' => :lng
}

@instruction_to_positivity = { 'N' => 1, 'S' => -1, 'E' => 1, 'W' => -1 }

@cardinal_directions = %w[E S W N]

def get_next_index(index, num_moves, instruction)
  set_moves = instruction == 'R' ? num_moves : 4 - num_moves
  (index + set_moves) % 4
end

def run_part_1(lines)
  facing = 'E'
  directions = { lat: 0, lng: 0 }
  lines.each do |line|
    instruction = line[0]
    direction = @instruction_to_direction[instruction]
    value = line[1..-1].to_i
    case instruction
    when 'F'
      directions[@instruction_to_direction[facing]] +=
        value * @instruction_to_positivity[facing]
    when 'L', 'R'
      current_index = @cardinal_directions.index(facing)
      num_moves = value / 90
      facing =
        @cardinal_directions[
          get_next_index(current_index, num_moves, instruction)
        ]
    else
      directions[direction] += value * @instruction_to_positivity[instruction]
    end
  end
  directions[:lat].abs + directions[:lng].abs
end

puts "Test Case: #{run_part_1(test_case)}"
puts "Input: #{run_part_1(input)}"

@intermediate_directions = %w[NE SE SW NW]

def get_intermediate_direction(waypoint)
  "#{waypoint[:lat] >= 0 ? 'N' : 'S'}#{waypoint[:lng] >= 0 ? 'E' : 'W'}"
end

def run_part_2(lines)
  waypoint = { lat: 1, lng: 10 }
  directions = { lat: 0, lng: 0 }
  lines.each do |line|
    instruction = line[0]
    direction = @instruction_to_direction[instruction]
    value = line[1..-1].to_i
    case instruction
    when 'F'
      directions[:lat] += waypoint[:lat] * value
      directions[:lng] += waypoint[:lng] * value
    when 'L', 'R'
      current_direction = get_intermediate_direction(waypoint)
      current_index = @intermediate_directions.index(current_direction)
      num_moves = value / 90
      new_direction =
        @intermediate_directions[
          get_next_index(current_index, num_moves, instruction)
        ]
      new_lat_value =
        waypoint[num_moves % 2 == 0 ? :lat : :lng].abs *
          @instruction_to_positivity[new_direction[0]]
      new_lng_value =
        waypoint[num_moves % 2 == 0 ? :lng : :lat].abs *
          @instruction_to_positivity[new_direction[1]]
      waypoint[:lat] = new_lat_value
      waypoint[:lng] = new_lng_value
    else
      waypoint[direction] += value * @instruction_to_positivity[instruction]
    end
  end
  puts directions
  directions[:lat].abs + directions[:lng].abs
end

puts "Test Case Part 2: #{run_part_2(test_case)}"
puts "Input Part 2: #{run_part_2(input)}"
