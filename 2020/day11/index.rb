require('../../rb_utils/load_file')

test_case = load_file_lines('./test_case.txt').map { |line| line.split('') }
input = load_file_lines('./input.txt').map { |line| line.split('') }

def count_around_occupied(rows, current_row, row_index, seat_index)
  num = 0
  [
    [-1, -1],
    [-1, 0],
    [-1, 1],
    [0, -1],
    [0, 1],
    [1, -1],
    [1, 0],
    [1, 1]
  ].each do |tuple|
    row, col = tuple
    row_index_to_check = row_index + row
    col_index_to_check = seat_index + col
    if row_index_to_check < 0 || row_index_to_check >= rows.size ||
         col_index_to_check < 0 || col_index_to_check >= current_row.size
      next
    end
    value = rows[row_index_to_check][col_index_to_check]
    num += 1 if value == '#'
  end
  num
end

def run_row(rows, row_index, row)
  row.each_with_index.map do |spot, spot_index|
    result = count_around_occupied(rows, row, row_index, spot_index)
    if spot == 'L' && result == 0
      '#'
    elsif spot == '#' && result >= 4
      'L'
    else
      spot
    end
  end
end

def run_rows(rows)
  rows.each_with_index.map { |row, row_index| run_row(rows, row_index, row) }
end

def run_part_1(rows)
  current_rows = rows
  updated_rows = []
  while true
    updated_rows = run_rows(current_rows)
    has_changed = false
    updated_rows.each_with_index do |updated_row, index|
      current_row = current_rows[index]
      if updated_row.join('') != current_row.join('')
        has_changed = true
        break
      end
    end
    break if !has_changed
    current_rows = updated_rows
  end
  current_rows.map { |row| row.join('') }.join('').count('#')
end

# puts "Test Case: #{run_part_1(test_case)}"
# puts "Input: #{run_part_1(input)}"

def update_num(current_val, direction)
  if direction == 'up' || direction == 'left'
    current_val - 1
  elsif direction == 'down' || direction == 'right'
    current_val + 1
  else
    current_val
  end
end

def count_around_occupied_part_2(
  rows, current_row, row_index, seat_index, log = false
)
  num = 0
  [
    %w[up left],
    %w[up center],
    %w[up right],
    %w[center left],
    %w[center right],
    %w[down left],
    %w[down center],
    %w[down right]
  ].each do |tuple|
    vertical, horizontal = tuple
    row_index_to_check = update_num(row_index, vertical)
    col_index_to_check = update_num(seat_index, horizontal)
    found_occupied = false
    while true
      if log
        puts "Current: #{row_index}, #{seat_index}; Directions: #{vertical}, #{
               horizontal
             }; Checking: #{row_index_to_check}, #{col_index_to_check}"
      end
      if row_index_to_check < 0 || row_index_to_check >= rows.size ||
           col_index_to_check < 0 || col_index_to_check >= current_row.size
        break
      end
      value = rows[row_index_to_check][col_index_to_check]

      if value != '.'
        found_occupied = true if value == '#'
        break
      end
      row_index_to_check = update_num(row_index_to_check, vertical)
      col_index_to_check = update_num(col_index_to_check, horizontal)
    end
    num += 1 if found_occupied
  end
  num
end

def run_row_part_2(rows, row_index, row, log = false)
  row.each_with_index.map do |spot, spot_index|
    result = count_around_occupied_part_2(rows, row, row_index, spot_index, log)
    if spot == 'L' && result == 0
      '#'
    elsif spot == '#' && result >= 5
      'L'
    else
      spot
    end
  end
end

def run_rows_part_2(rows, log = false)
  rows.each_with_index.map do |row, row_index|
    run_row_part_2(rows, row_index, row, log)
  end
end

def run_part_2(rows)
  current_rows = rows
  updated_rows = []
  while true
    updated_rows = run_rows_part_2(current_rows)
    has_changed = false
    updated_rows.each_with_index do |updated_row, index|
      current_row = current_rows[index]
      if updated_row.join('') != current_row.join('')
        has_changed = true
        break
      end
    end
    break if !has_changed
    current_rows = updated_rows
  end
  current_rows.map { |row| row.join('') }.join('').count('#')
end

puts "Test Case Part 2: #{run_part_2(test_case)}"
puts "Input Part 2: #{run_part_2(input)}"
