import time

def evaluate(items, n):
    last_used = {x: i for i,x in enumerate(items)}
    next = items[-1]
    used = set(items[:-1])
    for i in range(len(items)-1, n):
        last = next
        if last in used:
            next = i - last_used[last]
        else:
            next = 0
            used.add(last)
        last_used[last] = i
    return last

def main():
    results = []

    with open('15-input.txt', 'r') as f:
        items = [int(i) for i in f.read().split(',')]

    results.append(evaluate(items[:], 2020))
    results.append(evaluate(items[:], 30000000))

    for i,s in enumerate(results):
        print(f'{i+1}: {s}')

if __name__ == '__main__':
    start_time = time.time_ns()
    main()
    print("--- Executed in {0:.3f} seconds ---".format((time.time_ns() - start_time) / (10 ** 9)))