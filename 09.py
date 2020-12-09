def main():
    results = []

    with open('09-input.txt', 'r') as f:
        items = [int(i.strip()) for i in f.read().splitlines()]

    # Part 1
    preamble_size = 25
    q = items[:preamble_size]
    for i in items[preamble_size:]:
        ok = False
        for j in q:
            if i - j in q:
                ok = True
                break
        if not ok:
            results.append(i)
            break
        q = q[1:]
        q.append(i)

    # Part 2
    for i in range(len(items)):
        q = items[i:]
        sum = 0
        min = -1
        max = 0
        ok = False
        for x in q:
            sum += x
            min = x if x < min or min == -1 else min
            max = x if x > max else max
            if sum == results[0]:
                results.append(min + max)
                ok = True
                break
            elif sum > results[0]:
                break
        if ok:
            break

    for i,s in enumerate(results):
        print(f'{i+1}: {s}')

if __name__ == '__main__':
    main()
