require 'json'

json_data = '{"name": "Alice", "age": 30}'
parsed_data = JSON.parse(json_data)

puts parsed_data['name']
puts parsed_data['age']

hash_data = { name: 'Bob', age: 25 }
puts JSON.generate(hash_data)
