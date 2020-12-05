def binToDec(b):
    result, i = 0, 0
    while(b != 0): 
        digit = b % 10
        result = result + digit * pow(2, i) 
        b = b//10
        i += 1
    return result

def main():
    results = []

    with open('05-input.txt', 'r') as f:
        items = [i.strip() for i in f.read().splitlines()]

    seat_ids = [binToDec(int(i.replace('F','0').replace('B','1').replace('L','0').replace('R','1'))) for i in items]

    # Part 1
    max_seat = max(seat_ids)
    results.append(max_seat)

    # Part 2
    my_seat = [s for s in range(min(seat_ids) + 1, max(seat_ids)) if s not in seat_ids]
    results.append(my_seat[0])

    for i,s in enumerate(results):
        print(f'{i}: {s}')

if __name__ == '__main__':
    main()