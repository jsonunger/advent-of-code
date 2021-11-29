require('../../rb_utils/load_file')

test_case = load_file_lines('./test_case.txt')
input = load_file_lines('./input.txt')

ADD = '+'
MULTIPLY = '*'
OPEN_PAREN = '('
CLOSE_PAREN = ')'

def is_number?(string)
  true if Float(string)
rescue StandardError
  false
end

def evaluate(expression)
  parts = []
  result = nil
  while expression.size > 0
    e = expression.shift
    case e
    when OPEN_PAREN
      parts << evaluate(expression)
    when CLOSE_PAREN
      return result
    else
      parts << e
    end

    if parts.size == 3
      result = eval(parts.join(' '))
      parts = [result]
    end
  end
  result
end

def run_part_1(line)
  expression =
    line.gsub(' ', '').split('').map do |char|
      is_number?(char) ? char.to_i : char
    end
  evaluate(expression)
end

# test_case.each do |test_case|
#   puts "Test Case: #{test_case} = #{run_part_1(test_case)}"
# end
# puts "Input: #{input.sum { |line| run_part_1(line) }}"

def special_eval(e)
  e.gsub!(/\d+\+\d+/) { |match| eval(match) } while e.match?(/\+/)
  eval(e)
end

def complex_evaluate(expression)
  result = ''
  while expression.size > 0
    e = expression.shift
    case e
    when OPEN_PAREN
      result = result + complex_evaluate(expression)
    when CLOSE_PAREN
      return special_eval(result).to_s
    else
      result = result + e
    end
  end
  special_eval(result).to_s
end

def run_part_2(line)
  expression = line.gsub(' ', '').split('')
  complex_evaluate(expression)
end

test_case.each do |test_case|
  puts "Test Case Part 2: #{test_case} = #{run_part_2(test_case).to_i}"
end
puts "Input Part 2: #{input.sum { |line| run_part_2(line).to_i }}"
