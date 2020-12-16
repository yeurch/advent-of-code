import re
import time

def main():
    results = []

    with open('16-input.txt', 'r') as f:
        lines = [i.strip() for i in f.read().splitlines()]

    limits_regex = re.compile(r'(.+): (\d+)\-(\d+) or (\d+)\-(\d+)')

    rules = {}
    all_valid_values = set()

    # Parse input
    i = 0
    limits_match = limits_regex.fullmatch(lines[i])
    while limits_match:
        field_name = limits_match.group(1)
        min,max = int(limits_match.group(2)), int(limits_match.group(3))
        values = set(range(min, max+1))
        min,max = int(limits_match.group(4)), int(limits_match.group(5))
        values.update(range(min, max+1))
        rules[field_name] = values
        all_valid_values.update(values)
        limits_match = limits_regex.fullmatch(lines[i])
        i += 1

    my_ticket = [int(x) for x in lines[i+1].split(',')]
    nearby_tickets = [[int(y) for y in x.split(',')] for x in lines[i+4:]]

    # Part 1
    all_invalid_values = [item for sublist in nearby_tickets for item in sublist if item not in all_valid_values]
    results.append(sum(all_invalid_values))

    # Part 2
    valid_nearby_tickets = [ticket for ticket in nearby_tickets if all(map(lambda x: x in all_valid_values, ticket))]

    num_fields = len(my_ticket)
    possible_fields = [set(rules.keys()) for _ in my_ticket]
    confirmed_fields = set()
    for t in valid_nearby_tickets:
        for i,v in enumerate(t):
            for rkey,rval in rules.items():
                if not v in rval:
                    possible_fields[i].remove(rkey)

    while any(map(lambda x: len(x) > 1, possible_fields)):
        for i,f in enumerate(possible_fields):
            if len(f) == 1:
                confirmed_fields.update(f)
        for i,f in enumerate(possible_fields):
            if len(f) > 1:
                possible_fields[i] = f.difference(confirmed_fields)

    result = 1
    for i,f in enumerate(possible_fields):
        if list(f)[0].startswith('departure '):
            result *= my_ticket[i]
    results.append(result)

    for i,s in enumerate(results):
        print(f'{i+1}: {s}')

if __name__ == '__main__':
    start_time = time.time_ns()
    main()
    print("--- Executed in {0:.3f} seconds ---".format((time.time_ns() - start_time) / (10 ** 9)))