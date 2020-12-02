test_case = [1721, 979, 366, 299, 675, 1456]

input = File.open('./input.txt').readlines.map(&:to_i)

def sum_2020(arr)
  num_iter = 0
  found = false
  (0...(arr.count - 1)).each do |i|
    ((i + 1)...arr.count).each do |j|
      num_iter += 1
      if arr[i] + arr[j] == 2020
        return { product: arr[i] * arr[j], num_iter: num_iter }
        puts arr[i] * arr[j]
        found = true
        break
      end
    end
    break if found
  end
  num_iter
end

def sum_2020_optimized(arr)
  sorted_arr = arr.sort
  i = 0
  j = sorted_arr.count - 1
  num_iter = 0
  while i < j
    num_iter += 1
    sum = sorted_arr[i] + sorted_arr[j]
    case
    when sum == 2020
      return { product: sorted_arr[i] * sorted_arr[j], num_iter: num_iter }
    when sum < 2020
      i += 1
    else
      j -= 1
    end
  end
  num_iter
end

puts "Test Case for two numbers: #{sum_2020(test_case)}"
puts "Test Case for two numbers optimized: #{sum_2020_optimized(test_case)}"
puts "Input for two numbers: #{sum_2020(input)}"
puts "Input for two numbers optimized: #{sum_2020_optimized(input)}"

def sum_2020_with_3(arr)
  num_iter = 0
  (0...(arr.count - 2)).each do |i|
    ((i + 1)...(arr.count - 1)).each do |j|
      ((j + 1)...(arr.count)).each do |k|
        num_iter += 1
        if arr[i] + arr[j] + arr[k] == 2020
          return { product: arr[i] * arr[j] * arr[k], num_iter: num_iter }
        end
      end
    end
  end
end

def sum_2020_with_3_optimized(arr)
  sorted_arr = arr.sort
  num_iter = 0
  (0..sorted_arr.count).each do |i|
    val = 2020 - sorted_arr[i]
    j = i + 1
    k = sorted_arr.count - 1
    while j < k
      num_iter += 1
      sum = sorted_arr[j] + sorted_arr[k]
      case
      when sum == val
        return(
          {
            product: sorted_arr[i] * sorted_arr[j] + sorted_arr[k],
            num_iter: num_iter
          }
        )
      when sum < val
        j += 1
      else
        k -= 1
      end
    end
  end
end

puts "Test Case for three numbers: #{sum_2020_with_3(test_case)}"
puts "Test Case for three numbers optimized: #{
       sum_2020_with_3_optimized(test_case)
     }"
puts "Input for three numbers: #{sum_2020_with_3(input)}"
puts "Input for three numbers optimized: #{sum_2020_with_3_optimized(input)}"
