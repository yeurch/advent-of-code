import time

def calc_neighbours(coordinates, state, expand=True, four_d = False):
    new_cells = dict()
    live_neighbours = 0
    w_range = range(-1, 2) if four_d else [0]
    for dx in range(-1, 2):
        for dy in range(-1, 2):
            for dz in range(-1, 2):
                for dw in w_range:
                    if dx == 0 and dy == 0 and dz == 0 and dw == 0:
                        continue # This is ourself
                    neighbour_coordinates = (coordinates[0] + dx, coordinates[1] + dy, coordinates[2] + dz, coordinates[3] + dw)
                    if neighbour_coordinates in state:
                        if state[neighbour_coordinates][0] != '.':
                            # This is a neigbouring live cube
                            live_neighbours += 1
                    elif expand and state[coordinates][0] == '#':
                        # Expand space to cover this neigbour for future generations
                        new_cells[neighbour_coordinates] = ['.', calc_neighbours(neighbour_coordinates, state, False, four_d)[0]]
    return (live_neighbours, new_cells)

def calc_all_neighbours(state, four_d):
    new = dict()
    for k in state:
        (live_neighbours, new_cells) = calc_neighbours(k, state, True, four_d)
        state[k][1] = live_neighbours
        for new_cell in new_cells:
            new[new_cell] = new_cells[new_cell]

    # Add all of the cells we're not currently tracking, but are next to an existing cell
    for i in new:
        state[i] = new[i]

def init_state(items):
    rows = len(items)
    cols = len(items[0])
    state = dict()
    for ny,y in enumerate(items):
        for nx,x in enumerate(y):
            state[(nx,ny,0,0)] = [x, 0]
    return state

def do_generation(state, four_d):
    calc_all_neighbours(state, four_d)

    # If any cells are dead, and have no live neighbours, we can remove them from the simulation for efficiency
    cells_to_remove = [k for k in state if state[k][0] == '.' and state[k][1] < 3]
    for k in cells_to_remove: del state[k]

    for k in state:
        if state[k][0] == '.' and state[k][1] == 3:
            state[k][0] = '#'
        elif state[k][0] == '#' and (state[k][1] < 2 or state[k][1] > 3):
            state[k][0] = '.'


def run_simulation(state, num_generations, four_d=False):
    for i in range(num_generations):
        do_generation(state, four_d)
        live_count = len([1 for s in state if state[s][0] == '#'])
        print(f'After generation {i+1} there are {live_count} live cells')
    return live_count

def main():
    results = []

    with open('17-input.txt', 'r') as f:
        items = [i.strip() for i in f.read().splitlines()]

    # Part 1
    print('Part 1')
    state = init_state(items)
    results.append(run_simulation(state, 6))

    # Part 2
    print('Part 2')
    state = init_state(items)
    results.append(run_simulation(state, 6, True))

    for i,s in enumerate(results):
        print(f'{i+1}: {s}')

if __name__ == '__main__':
    start_time = time.time_ns()
    main()
    print("--- Executed in {0:.3f} seconds ---".format((time.time_ns() - start_time) / (10 ** 9)))