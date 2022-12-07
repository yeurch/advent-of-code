import itertools

with open('01-input.txt', 'r') as f:
    items = [int(i) for i in f.read().splitlines()]

result = [(a, b) for [a, b] in itertools.combinations(items, 2) if a + b == 2020]
(a, b) = result[0]
print(f'{a} * {b} = {a * b}')

result = [(a, b, c) for [a, b, c] in itertools.combinations(items, 3) if a + b + c == 2020]
(a, b, c) = result[0]
print(f'{a} * {b} * {c} = {a * b * c}')
