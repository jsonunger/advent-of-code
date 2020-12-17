require('../../rb_utils/load_file')

test_case = load_file_lines('./test_case.txt')
input = load_file_lines('./input.txt')

def generate_rules(rule_lines)
  rules = {}
  rule_lines.each do |rule_line|
    rule, options = rule_line.split(': ')
    valid_numbers = []
    options.split(' or ').each do |option|
      min, max = option.split('-').map(&:to_i)
      valid_numbers += (min..max).to_a
    end
    rules[rule] = valid_numbers
  end
  rules
end

def run_part_1(lines)
  your_ticket_index = lines.index('your ticket:')
  rule_lines = lines[0...your_ticket_index - 1]
  rules = generate_rules(rule_lines)
  other_tickets =
    lines[your_ticket_index + 4..-1].map do |other_ticket_line|
      other_ticket_line.split(',').map(&:to_i)
    end
  other_tickets.flatten.select do |num|
    rules.values.all? { |valid_numbers| !valid_numbers.include?(num) }
  end.sum
end

puts "Test Case: #{run_part_1(test_case)}"
puts "Input: #{run_part_1(input)}"

test_case_part_2 = load_file_lines('./test_case_part_2.txt')

def run_part_2(lines)
  your_ticket_index = lines.index('your ticket:')
  your_ticket = lines[your_ticket_index + 1].split(',').map(&:to_i)
  rule_lines = lines[0...your_ticket_index - 1]
  rules = generate_rules(rule_lines)
  other_tickets =
    lines[your_ticket_index + 4..-1].map do |other_ticket_line|
      other_ticket_line.split(',').map(&:to_i)
    end
  valid_other_tickets =
    other_tickets.select do |other_ticket|
      other_ticket.all? do |num|
        rules.values.any? { |valid_numbers| valid_numbers.include?(num) }
      end
    end

  rules_to_indexes = {}
  rules.keys.each do |rule|
    matching_columns =
      valid_other_tickets.first.each_index.select do |index|
        valid_other_tickets.all? do |ticket|
          rules[rule].include?(ticket[index])
        end
      end
    rules_to_indexes[rule] = matching_columns
  end

  rules_to_index = {}

  until rules_to_indexes.empty?
    done, matches = rules_to_indexes.partition { |k, v| v.size == 1 }
    done.each do |rule, matching_indexes|
      rules_to_index[rule] = matching_indexes.first
      rules_to_indexes.delete(rule)
      rules_to_indexes.keys.each do |rule_to_check|
        rules_to_indexes[rule_to_check].delete(matching_indexes.first)
      end
    end
  end

  departure_indexes =
    rules_to_index.each_key.select do |key|
      key.start_with?('departure')
    end.map { |key| rules_to_index[key] }

  departure_indexes.map { |index| your_ticket[index] }.reduce(:*)
end

# puts "Test Case Part 2: #{run_part_2(test_case_part_2)}"
puts "Input Part 2: #{run_part_2(input)}"
