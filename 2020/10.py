import time

def combinations_from(items):
    last = items[0]
    for i in items:
        if i - last > 3:
            return 0
        last = i
        
    if len(items) == 2:
        return 1

    split = len(items) // 2
    items_except_split = items[:split] + items[split+1:]
    return combinations_from(items[:split + 1]) * combinations_from(items[split:]) + combinations_from(items_except_split)

def main():
    results = []
    with open('10-input.txt', 'r') as f:
        items = sorted([int(i.strip()) for i in f.read().splitlines()])

    # Part 1
    last = 0
    deltas = {1: 0, 2: 0, 3: 0}
    for item in items:
        deltas[item - last] += 1
        last = item
    deltas[3] += 1
    results.append(deltas[1] * deltas[3])

    # Part 2
    combinations = 0
    items.insert(0, 0)
    results.append(combinations_from(items))

    for i,s in enumerate(results):
        print(f'{i+1}: {s}')

if __name__ == '__main__':
    start_time = time.time()
    main()
    print("--- Executed in %s seconds ---" % (time.time() - start_time))