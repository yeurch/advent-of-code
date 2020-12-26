import time

def main():
    results = []

    with open('23-input.txt', 'r') as f:
        input = [i.strip() for i in f.read()]
    
    items = [int(c) for c in input]
    print(items)

    max_item = max(items)
    min_item = min(items)
    for i in range(100):
        target = items[0]
        while items.index(target) < 4:
            target -= 1
            if target < min_item: target = max_item
        targetpos = items.index(target)
        items = items[4:targetpos+1] + items[1:4] + items[targetpos+1:] + items[0:1]
        print(items)

    while items[0] != 1:
        i = items.pop(0)
        items.append(i)
    results.append(''.join([str(x) for x in items[1:]]))
    #results.append(0)
    #results.append(0)

    for i,s in enumerate(results):
        print(f'{i+1}: {s}')

if __name__ == '__main__':
    start_time = time.time_ns()
    main()
    print("--- Executed in {0:.3f} seconds ---".format((time.time_ns() - start_time) / (10 ** 9)))