import functools

def getLoc(items, x, y):
    row = items[y]
    return row[x % len(row)]

def count_trees(items, dx, dy):
    rows = items[::dy]
    squares = [getLoc(rows, dx*i, i) for i,_ in enumerate(rows)]
    return squares.count('#')

def main():
    with open('03-input.txt', 'r') as f:
        items = f.read().splitlines()

    test_cases = [(1,1), (3,1), (5,1), (7,1), (1,2)]
    results = [count_trees(items, x, y) for x, y in test_cases]
    print(results)
    print(f'1: {results[1]}') # puzzle 1
    print(f'2: {functools.reduce(lambda x,y: x*y, results)}') # puzzle 2

if __name__ == '__main__':
    main()