from functools import lru_cache, reduce
import time

@lru_cache() # without caching, adds 50% to execution time
def multiply_values(values):
    return reduce((lambda x,y: x*y), values)

def main():
    results = []

    with open('13-input.txt', 'r') as f:
        lines = [i.strip() for i in f.read().splitlines()]

    earliest_departure = int(lines[0])
    routes = lines[1].split(',')

    # Part 1
    in_service = sorted([int(n) for n in routes if n != 'x'])
    chosen_departure = (-1,in_service[-1] + earliest_departure + 1)
    for route in in_service:
        if earliest_departure % route == 0:
            # Exact match, we have our solution
            chosen_departure = (route, earliest_departure)
            break
        n = earliest_departure // route
        candidate_time = (n+1) * route
        if candidate_time < chosen_departure[1]:
            chosen_departure = (route, candidate_time)
    results.append(chosen_departure[0] * (chosen_departure[1] - earliest_departure))

    # Part 2
    requirements = [(i, int(n)) for i,n in enumerate(routes) if n != 'x']
    requirements.sort(key=lambda x: x[1], reverse=True)
    i = requirements[0]
    a = 0
    success = False
    while not success:
        for x in range(1, len(requirements)):
            j = requirements[x]
            offset = i[0] - j[0]
            if (a*i[1] - offset) % j[1] != 0:
                product = tuple([z[1] for z in requirements[:x]])
                a += multiply_values(product) // i[1]
                break
            if x == len(requirements) - 1:
                success = True
    results.append(a*i[1] - i[0])

    for i,s in enumerate(results):
        print(f'{i+1}: {s}')

if __name__ == '__main__':
    start_time = time.time()
    main()
    print("--- Executed in {0:.3f} milliseconds ---".format((time.time() - start_time)*1000))
