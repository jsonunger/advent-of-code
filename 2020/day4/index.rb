test_case = [
  'ecl:gry pid:860033327 eyr:2020 hcl:#fffffd',
  'byr:1937 iyr:2017 cid:147 hgt:183cm',
  '',
  'iyr:2013 ecl:amb cid:350 eyr:2023 pid:028048884',
  'hcl:#cfa07d byr:1929',
  '',
  'hcl:#ae17e1 iyr:2013',
  'eyr:2024',
  'ecl:brn pid:760753108 byr:1931',
  'hgt:179cm',
  '',
  'hcl:#cfa07d eyr:2025 pid:166559648',
  'iyr:2011 ecl:brn hgt:59in'
]
input = File.open('./input.txt').readlines.map(&:chomp)

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

test_case_part_2 = [
  'eyr:1972 cid:100',
  'hcl:#18171d ecl:amb hgt:170 pid:186cm iyr:2018 byr:1926',
  '',
  'iyr:2019',
  'hcl:#602927 eyr:1967 hgt:170cm',
  'ecl:grn pid:012533040 byr:1946',
  '',
  'hcl:dab227 iyr:2012',
  'ecl:brn hgt:182cm pid:021572410 eyr:2020 byr:1992 cid:277',
  '',
  'pid:087499704 hgt:74in ecl:grn iyr:2012 eyr:2030 byr:1980',
  'hcl:#623a2f',
  '',
  'eyr:2029 ecl:blu cid:129 byr:1989',
  'iyr:2014 pid:896056539 hcl:#a97842 hgt:165cm',
  '',
  'hgt:59cm ecl:zzz',
  'eyr:2038 hcl:74454a iyr:2023',
  'pid:3556412378 byr:2007',
  '',
  'hcl:#888785',
  'hgt:164cm byr:2001 iyr:2015 cid:88',
  'pid:545766238 ecl:hzl',
  'eyr:2022'
]

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
