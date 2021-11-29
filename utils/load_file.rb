def load_file_lines(file_name)
  File.open(file_name).readlines.map(&:chomp)
end
