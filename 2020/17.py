from collections import *
from functools import lru_cache, reduce
import heapq
import itertools
import math
import random
import sys
import re
import time

def calc_neighbours(x,y,z, state, expand=True):
    new_cells = dict()
    live_neighbours = 0
    for dx in range(-1, 2):
        for dy in range(-1, 2):
            for dz in range(-1, 2):
                if dx == 0 and dy == 0 and dz == 0:
                    continue # This is ourself
                posx = x + dx
                posy = y + dy
                posz = z + dz
                if (posx,posy,posz) in state:
                    if state[posx,posy,posz][0] != '.':
                        # This is a neigbouring live cube
                        live_neighbours += 1
                elif expand:
                    # Expand space to cover this neigbour for future generations
                    new_cells[posx,posy,posz] = ['.', calc_neighbours(posx,posy,posz, state, False)[0]]
    return (live_neighbours, new_cells)

def calc_all_neighbours(state):
    new = dict()
    for (x,y,z) in state:
        (live_neighbours, new_cells) = calc_neighbours(x,y,z, state)
        state[(x,y,z)][1] = live_neighbours
        for new_cell in new_cells:
            new[new_cell] = new_cells[new_cell]
    for x in new:
        state[x] = new[x]

def init_state(items):
    rows = len(items)
    cols = len(items[0])
    state = dict()
    for ny,y in enumerate(items):
        for nx,x in enumerate(y):
            state[(nx,ny,0)] = [x, 0]
    return state

def do_generation(state):
    calc_all_neighbours(state)
    for k in state:
        if state[k][0] == '.' and state[k][1] == 3:
            state[k][0] = '#'
        elif state[k][0] == '#' and (state[k][1] < 2 or state[k][1] > 3):
            state[k][0] = '.'

def run_simulation(state, num_generations):
    for i in range(num_generations):
        do_generation(state)
        live_count = len([1 for s in state if state[s][0] == '#'])
        print(f'After generation {i+1} there are {live_count} live cells')
    return live_count

def main():
    results = []

    with open('17-input.txt', 'r') as f:
        items = [i.strip() for i in f.read().splitlines()]

    # Part 1
    state = init_state(items)
    results.append(run_simulation(state, 6))

    for i,s in enumerate(results):
        print(f'{i+1}: {s}')

if __name__ == '__main__':
    start_time = time.time_ns()
    main()
    print("--- Executed in {0:.3f} seconds ---".format((time.time_ns() - start_time) / (10 ** 9)))