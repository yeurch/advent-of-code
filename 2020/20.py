from collections import *
from functools import lru_cache, reduce
import heapq
import itertools
import math
import random
import sys
import re
import time

class Tile:
    def __init__(self, id, data, image):
        self.id = id
        self.data = data
        self.image = image
        self.edges = []
        self.edges.append(data[0])
        self.edges.append(''.join([x[-1] for x in data]))
        self.edges.append(data[-1][::-1])
        self.edges.append(''.join([x[0] for x in data][::-1]))
        self.orientation = 0
        self.pos = (None, None)
        self.matches = None
    
    def calc_matches(self):
        result = [[], [], [], [], [], [], [], []]
        matches = 0
        my_edges_reverse = [x[::-1] for x in self.edges]
        for is_flip,my_edges in enumerate([self.edges, my_edges_reverse]):
            for _,(id,i) in enumerate(self.image.tiles.items()):
                if id == self.id:
                    continue
                for e in i.edges:
                    if e in my_edges:
                        j = my_edges.index(e)
                        result[is_flip*4 + j].append((id, i.edges.index(e)))
                        matches += 1
        self.matches = result
        self.num_matching_edges = len([x for x in result if len(x) > 0])
        return result

class Image:
    def __init__(self, items):
        self.tiles = dict()
        for i in range((len(items)+1) // 12):
            tile_id = int(items[i*12][5:-1])
            tile_data = items[i*12+1:i*12+11]
            tile = Tile(tile_id, tile_data, self)
            self.tiles[tile_id] = tile
    
    def get_tile(self, id):
        return self.tiles[id]

    def calc_matches(self):
        for k in self.tiles:
            self.tiles[k].calc_matches()
    
    def part1(self):
        self.calc_matches()
        result = 1
        corner_pieces = [k for k in self.tiles if self.tiles[k].num_matching_edges == 2]
        if len(corner_pieces) != 4:
            raise Exception('Difficulty detecting corners')
        for x in corner_pieces:
            result *= x
        return result

def main():
    results = []

    with open('20-input.txt', 'r') as f:
        items = [i.strip() for i in f.read().splitlines()]

    image = Image(items)

    results.append(image.part1())

    #for i in image.tiles:
    #    tile = image.tiles[i]
    #    print(f'{tile.id}: {tile.matches}')

    for i,s in enumerate(results):
        print(f'{i+1}: {s}')

if __name__ == '__main__':
    start_time = time.time_ns()
    main()
    print("--- Executed in {0:.3f} seconds ---".format((time.time_ns() - start_time) / (10 ** 9)))