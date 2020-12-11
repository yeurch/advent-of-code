import time

def calc_seats(items, long_range):
    rows = len(items)
    cols = len(items[0])
    seats = []
    for ny,y in enumerate(items):
        for nx,x in enumerate(y):
            neighbours = []
            if x != '.':
                for dx in range(-1, 2):
                    for dy in range(-1, 2):
                        if dx == 0 and dy == 0:
                            continue
                        posx, posy = nx, ny
                        while True:
                            posx += dx
                            posy += dy
                            if posx < 0 or posy < 0 or posx >= cols or posy >= rows:
                                # We're out of bounds, so skip to next cardinal direction
                                break
                            if items[posy][posx] != '.':
                                # This is a seat
                                neighbours.append(cols * posy + posx)
                                break
                            if not long_range:
                                # We're only considering immediate neighbours
                                break
            seats.append([x, neighbours])
    
    return seats

def print_grid(seats, width, height):
    for y in range(height):
        print(''.join([s[0] for s in seats[y*width:(y+1)*width]]))
    print()

def do_generation(seats, min_threshold, width, height):
    occupied_neighbours = []
    for s in seats:
        occupied_neighbours.append(len([n for n in s[1] if seats[n][0] == '#']))
    for i,s in enumerate(seats):
        if s[0] == 'L' and occupied_neighbours[i] == 0:
            s[0] = '#'
        elif s[0] == '#' and occupied_neighbours[i] >= min_threshold:
            s[0] = 'L'
    # print_grid(seats, width, height)

def run_simulation(seats, min_threshold, width, height):
    state_str = ''.join([s[0] for s in seats])
    generation = 1
    while True:
        old_state_str = state_str
        do_generation(seats, min_threshold, width, height)
        state_str = ''.join([s[0] for s in seats])
        if state_str == old_state_str:
            break
        generation += 1
        if generation % 100 == 0:
            print(f'Generation {generation}')
    return len([1 for s in seats if s[0] == '#'])

def main():
    results = []

    with open('11-input.txt', 'r') as f:
        items = [i.strip() for i in f.read().splitlines()]

    # Part 1
    seats = calc_seats(items, False)
    results.append(run_simulation(seats, 4, len(items[0]), len(items)))

    # Part 2
    seats = calc_seats(items, True)
    # print_grid(seats, len(items[0]), len(items))
    results.append(run_simulation(seats, 5, len(items[0]), len(items)))

    for i,s in enumerate(results):
        print(f'{i+1}: {s}')

if __name__ == '__main__':
    start_time = time.time()
    main()
    print("--- Executed in %s seconds ---" % (time.time() - start_time))