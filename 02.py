import re

with open('02-input.txt', 'r') as f:
    items = [i for i in f.read().splitlines()]

regex = re.compile(r"(\d+)-(\d+)\W+(.):\W(.*)")

valid_a = 0
valid_b = 0
for i in items:
    m = regex.match(i)
    if m == None:
        raise Exception(f'Could not match string {i}')
    min = int(m.group(1))
    max = int(m.group(2))
    c = m.group(3)
    password = m.group(4)

    if len(c) != 1:
        raise Exception(f'Discovered multi-char rule {c}')

    # Puzzle 1
    occurrences = password.count(c)
    ok = occurrences >= min and occurrences <= max
    valid_a += 1 if ok else 0

    # Puzzle 2
    match_min = 1 if len(password) >= max and password[min-1] == c else 0
    match_max = 1 if len(password) >= max and password[max-1] == c else 0
    ok = match_min + match_max == 1
    valid_b += 1 if ok else 0

print(valid_a)
print(valid_b)


