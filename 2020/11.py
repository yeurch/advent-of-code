from collections import *
from functools import lru_cache, reduce
import heapq
import itertools
import math
import random
import sys
import re
import time

def calc_seats(items):
    rows = len(items)
    cols = len(items[0])
    seats = []
    for ny,y in enumerate(items):
        for nx,x in enumerate(y):
            neighbours = []
            if x != '.':
                up = ny > 0
                down = ny < rows - 1
                left = nx > 0
                right = nx < cols - 1
                if up:
                    if left: neighbours.append(cols * (ny - 1) + nx - 1)
                    neighbours.append(cols * (ny - 1) + nx)
                    if right: neighbours.append(cols * (ny - 1) + nx + 1)
                if left: neighbours.append(cols * ny + nx - 1)
                if right: neighbours.append(cols * ny + nx + 1)
                if down:
                    if left: neighbours.append(cols * (ny + 1) + nx - 1)
                    neighbours.append(cols * (ny + 1) + nx)
                    if right: neighbours.append(cols * (ny + 1) + nx + 1)
            seats.append([x, neighbours])
    
    # Filter out neighbors who are floor space ('.')
    for i in seats:
        i[1] = set(filter(lambda x: seats[x][0] != '.', i[1]))

    return seats

def do_generation(seats, min_threshold):
    occupied_neighbours = []
    for s in seats:
        occupied_neighbours.append(len([n for n in s[1] if seats[n][0] == '#']))
    for i,s in enumerate(seats):
        if s[0] == 'L' and occupied_neighbours[i] == 0:
            s[0] = '#'
        elif s[0] == '#' and occupied_neighbours[i] >= 4:
            s[0] = 'L'

def main():
    results = []

    with open('11-input.txt', 'r') as f:
        items = [i.strip() for i in f.read().splitlines()]

    seats = calc_seats(items)

    # Part 1
    state_str = ''.join([s[0] for s in seats])
    while True:
        old_state_str = state_str
        do_generation(seats, 4)
        state_str = ''.join([s[0] for s in seats])
        if state_str == old_state_str:
            break
    results.append(len([1 for s in seats if s[0] == '#']))

    for i,s in enumerate(results):
        print(f'{i+1}: {s}')

if __name__ == '__main__':
    start_time = time.time()
    main()
    print("--- Executed in %s seconds ---" % (time.time() - start_time))
