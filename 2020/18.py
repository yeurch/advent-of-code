import time

def tokenize(expr):
    tokens = []
    d = ''
    reading_digit = False
    for i in expr.replace(' ',''):
        if i.isdigit():
            d += i
            reading_digit = True
        else:
            if reading_digit:
                tokens.append(int(d))
                d = ''
                reading_digit = False
            tokens.append(i)
    if reading_digit:
        tokens.append(int(d))
    return tokens

def evaluate(tokens, part):
    i = 0
    result = 0
    operator = '+'
    expanded = []
    while i < len(tokens):
        t = tokens[i]
        if t == '(':
            depth = 1
            t0 = i + 1
            while depth > 0:
                i += 1
                t = tokens[i]
                if t == '(':
                    depth += 1
                elif t == ')':
                    depth -= 1
            expanded.append(evaluate(tokens[t0:i], part))
        else:
            expanded.append(t)
        i += 1

    if part == 2:
        while '+' in expanded:
            i = expanded.index('+')
            expanded = expanded[:i-1] + [expanded[i-1] + expanded[i+1]] + expanded[i+2:]

    result = 0
    operator = '+'
    for t in expanded:
        if isinstance(t, int):
            if operator == '+':
                result += t
            elif operator == '*':
                result *= t
        else:
            operator = t
    return result

def main():
    results = []

    with open('18-input.txt', 'r') as f:
        items = [i.strip() for i in f.read().splitlines()]

    p1, p2 = 0, 0
    for i in items:
        t = tokenize(i)
        p1 += evaluate(t, 1)
        p2 += evaluate(t, 2)
    results.append(p1)
    results.append(p2)

    for i,s in enumerate(results):
        print(f'{i+1}: {s}')

if __name__ == '__main__':
    start_time = time.time_ns()
    main()
    print("--- Executed in {0:.3f} seconds ---".format((time.time_ns() - start_time) / (10 ** 9)))