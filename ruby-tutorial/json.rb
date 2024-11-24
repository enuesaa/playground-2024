require 'oj'

json_data = '{"name": "Alice", "age": 30}'
parsed_data = Oj.load(json_data)

puts parsed_data["name"]  # => Alice
puts parsed_data["age"]   # => 30
