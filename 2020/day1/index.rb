test_case = [1721, 979, 366, 299, 675, 1456]

input = File.open('./input.txt').readlines.map(&:to_i)

def sum_2020(arr)
  (0...(arr.count - 1)).each do |i|
    ((i + 1)...arr.count).each do |j|
      puts arr[i] * arr[j] if arr[i] + arr[j] == 2020
    end
  end
end

puts 'Test Case for two numbers:'
sum_2020(test_case)
puts 'Input for two numbers:'
sum_2020(input)

def sum_2020_with_3(arr)
  (0...(arr.count - 2)).each do |i|
    ((i + 1)...(arr.count - 1)).each do |j|
      ((j + 1)...(arr.count)).each do |k|
        puts arr[i] * arr[j] * arr[k] if arr[i] + arr[j] + arr[k] == 2020
      end
    end
  end
end

puts 'Test Case for three numbers:'
sum_2020_with_3(test_case)
puts 'Input for three numbers:'
sum_2020_with_3(input)
