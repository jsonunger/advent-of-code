require '../../rb_utils/load_file'

test_case = load_file_lines('./test_case.txt')
input = load_file_lines('./input.txt')

@necessary_fields = %i[byr iyr eyr hgt hcl ecl pid]

def count_valid_passports(rows, validator = nil)
  passport_hashes = [{}]
  passport_hashes =
    rows.reduce([{}]) do |hashes, row|
      if row == ''
        hashes << {}
        next hashes
      end
      row.split(' ').each do |key_val_str|
        key, val = key_val_str.split(':')
        hashes[-1][key.to_sym] = val
      end
      hashes
    end

  filtered =
    passport_hashes.select do |passport_hash|
      @necessary_fields.all? do |field|
        passport_hash.key?(field) &&
          (validator.nil? || validator.call(field, passport_hash[field]))
      end
    end

  filtered.size
end

puts "Test Case: #{count_valid_passports(test_case)}"
puts "Input: #{count_valid_passports(input)}"

puts
puts '#########'
puts

test_case_part_2 = load_file_lines('./test_case_part_2.txt')

def validate(field, value)
  case field
  when :byr
    value.match?(/^\d{4}$/) && value.to_i.between?(1920, 2002)
  when :iyr
    value.match?(/^\d{4}$/) && value.to_i.between?(2010, 2020)
  when :eyr
    value.match?(/^\d{4}$/) && value.to_i.between?(2020, 2030)
  when :hgt
    match = value.match(/^(?<amount>\d+)(?<type>in|cm)$/)
    if match
      match[:amount].to_i.between?(
        *(match[:type] == 'cm' ? [150, 193] : [59, 76])
      )
    else
      false
    end
  when :hcl
    value.match?(/^#[0-9a-f]{6}$/)
  when :ecl
    %w[amb blu brn gry grn hzl oth].any? { |color| color == value }
  when :pid
    value.match?(/^\d{9}$/)
  else
    false
  end
end

puts "Test Case: #{count_valid_passports(test_case_part_2, method(:validate))}"
puts "Input: #{count_valid_passports(input, method(:validate))}"
