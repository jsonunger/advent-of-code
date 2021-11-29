require('../../rb_utils/load_file')

test_case = load_file_lines('./test_case.txt')
input = load_file_lines('./input.txt')

def make_grid(default, dimensions)
  return Hash.new { |h, k| h[k] = default } if dimensions == 1

  return Hash.new { |h, k| h[k] = make_grid(default, dimensions - 1) }
end

def count_in_grid(grid)
  return grid == '#' ? 1 : 0 if !grid.is_a?(Hash)

  return grid.values.map { |subgrid| count_in_grid(subgrid) }.sum
end

def all_neighbor_indices(size)
  [-1, 0, 1].repeated_permutation(size).reject { |pos| pos == [0] * size }
end

def get_neighbors(*args)
  neighbors = []
  all_neighbor_indices(args.size).each do |indices|
    neighbors << args.zip(indices).map(&:sum)
  end
  neighbors
end

def count_neighbors_for_living_cells(grid)
  counts = make_grid(0, 3)

  grid.to_a.each do |z, layer|
    layer.to_a.each do |y, row|
      row.to_a.each do |x, cell|
        if cell == '#'
          get_neighbors(z, y, x).each { |nz, ny, nx| counts[nz][ny][nx] += 1 }
        end
      end
    end
  end

  counts
end

def generate_next_grid(grid)
  next_grid = make_grid('.', 3)
  neighbor_counts = count_neighbors_for_living_cells(grid)

  neighbor_counts.to_a.each do |z, layer|
    layer.to_a.each do |y, row|
      row.to_a.each do |x, cell_neighbor_count|
        if grid[z][y][x] == '#'
          next_grid[z][y][x] = [2, 3].include?(cell_neighbor_count) ? '#' : '.'
        else
          next_grid[z][y][x] = cell_neighbor_count == 3 ? '#' : '.'
        end
      end
    end
  end

  next_grid
end

def run_part_1(lines)
  grid = make_grid('.', 3)

  lines.each_with_index do |line, y|
    line.split('').each_with_index { |char, x| grid[0][y][x] = char }
  end

  (1..6).each do
    next_grid = generate_next_grid(grid)
    grid = next_grid
  end

  count_in_grid(grid)
end

puts "Test Case: #{run_part_1(test_case)}"
puts "Input: #{run_part_1(input)}"

def count_neighbors_for_living_cells_4d(grid)
  counts = make_grid(0, 4)

  grid.to_a.each do |w, zone|
    zone.to_a.each do |z, layer|
      layer.to_a.each do |y, row|
        row.to_a.each do |x, cell|
          if cell == '#'
            get_neighbors(w, z, y, x).each do |nw, nz, ny, nx|
              counts[nw][nz][ny][nx] += 1
            end
          end
        end
      end
    end
  end

  counts
end

def generate_next_grid_4d(grid)
  next_grid = make_grid('.', 4)
  neighbor_counts = count_neighbors_for_living_cells_4d(grid)

  neighbor_counts.to_a.each do |w, zone|
    zone.to_a.each do |z, layer|
      layer.to_a.each do |y, row|
        row.to_a.each do |x, cell_neighbor_count|
          if grid[w][z][y][x] == '#'
            next_grid[w][z][y][x] =
              [2, 3].include?(cell_neighbor_count) ? '#' : '.'
          else
            next_grid[w][z][y][x] = cell_neighbor_count == 3 ? '#' : '.'
          end
        end
      end
    end
  end

  next_grid
end

def run_part_2(lines)
  grid = make_grid('.', 4)

  lines.each_with_index do |line, y|
    line.split('').each_with_index { |char, x| grid[0][0][y][x] = char }
  end

  (1..6).each do
    next_grid = generate_next_grid_4d(grid)
    grid = next_grid
  end

  count_in_grid(grid)
end

puts "Test Case Part 2: #{run_part_2(test_case)}"
puts "Input Part 2: #{run_part_2(input)}"
