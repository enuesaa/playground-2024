import pdb

def add(a, b):
    result = a + b
    pdb.set_trace()
    return result

x = 5
y = 3
result = add(x, y)
print(f"The result is: {result}")

