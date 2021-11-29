require('../../rb_utils/load_file')

test_case = File.open('./test_case.txt').read
# input = load_file_lines('./input.txt')

def generate_rules(rules_str)
  rules_str.split(/\n/).map do |rule_line|
    rule_num, rule_def = rule_line.strip.split(': ', 2)

    {
      rule_num: rule_num.to_i,
      options:
        rule_def.split(' | ').map do |option|
          option.split(' ').map do |token|
            token.start_with?('"') ? token[1..-2] : token.to_i
          end
        end
    }
  end
end

def find_rule(rules, rule_num)
  rule = rules.find { |r| r[:rule_num] == rule_num }

  raise StandardError("No rule found for #{rule_num}") if !rule

  rule
end

NO_MATCH = 0

def test_message_against_rule(message, rules, starting_rule)
  @message = message
  @rules = rules
  def _test_at(offset, sequence, sequence_idx, loop_count)
    puts "message: #{@message}"
    if offset >= @message.length || loop_count > @message.length
      return [NO_MATCH]
    end

    token = sequence[sequence_idx]
    next_sequence_idx = sequence_idx + 1
    is_end_of_sequence = next_sequence_idx >= sequence.length

    next_offsets = []

    if token.is_a? Integer
      next_rule = find_rule(@rules, token)
      next_offsets =
        next_rule[:options].map do |next_sequence|
          _test_at(offset, next_sequence, 0, loop_count + 1)
        end.flatten
    elsif @message[offset] == token
      next_offsets = [offset + 1]
    else
      next_offsets = [NO_MATCH]
    end

    if next_sequence_idx < sequence.length
      next_offsets.select! { |consumed| consumed != NO_MATCH }.map do |off|
        _test_at(off, sequence, next_sequence_idx, 0)
      end.flatten
    end

    next_offsets.select { |consumed| consumed != NO_MATCH }
  end

  starting_rule.map { |sequence| _test_at(0, sequence, 0, 0) }.flatten
    .select { |off| off != NO_MATCH }.any? { |off| off == @message.length }
end

def run_part_1(data)
  rules_str, messages_str = data.split(/\n\n/)
  rules = generate_rules(rules_str)
  messages = messages_str.split(/\n/).map(&:strip)

  rule_0 = find_rule(rules, 0)

  messages.select do |message|
    test_message_against_rule(message, rules, rule_0)
  end.length
end

puts "Test Case: #{run_part_1(test_case)}"
# puts "Input: #{run_part_1(input)}"

def run_part_2(lines)
  lines
end

# puts "Test Case Part 2: #{run_part_2(test_case)}"
# puts "Input Part 2: #{run_part_2(input)}"
