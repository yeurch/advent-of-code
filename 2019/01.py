import time

def main():
    results = []

    with open('01-input.txt', 'r') as f:
        items = [int(i.strip()) for i in f.read().splitlines()]

    # Part 1
    results.append(sum([x // 3 - 2 for x in items]))

    # Part 2
    total_fuel = 0
    for i in items:
        fuel = i // 3 - 2
        last_fuel = fuel
        while last_fuel > 0:
            last_fuel = last_fuel // 3 - 2
            fuel += last_fuel if last_fuel > 0 else 0
        total_fuel += fuel
    results.append(total_fuel)

    for i,s in enumerate(results):
        print(f'{i+1}: {s}')

if __name__ == '__main__':
    start_time = time.time()
    main()
    print("--- Executed in {0:.3f} milliseconds ---".format((time.time() - start_time)*1000))