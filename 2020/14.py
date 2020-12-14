from collections import *
from functools import lru_cache, reduce
import heapq
import itertools
import math
import random
import sys
import re
import time
from typing import Pattern, List

set_regex: Pattern = re.compile(r'mem\[(\d+)\] = (\d+)')

def binToDec(b):
    result, i = 0, 0
    while len(b) > 0: 
        digit = int(b[-1])
        b = b[:-1]
        result = result + digit * pow(2, i) 
        i += 1
    return result

def part_1(program: List[str]):
    mask = 'XXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX'
    memory = dict()
    for instruction in program:
        set_match = set_regex.match(instruction)
        if set_match:
            addr = set_match.group(1)
            value = bin(int(set_match.group(2)))[2:].zfill(36)
            memory[addr] = ''.join([value[i] if mask[i] == 'X' else mask[i] for i in range(36)])
        else:
            mask = instruction[7:]
    return sum([binToDec(v) for _,(k,v) in enumerate(memory.items())])

def main():
    results = []

    with open('14-input.txt', 'r') as f:
        items = [i.strip() for i in f.read().splitlines()]

    # Part 1
    results.append(part_1(items))

    for i,s in enumerate(results):
        print(f'{i+1}: {s}')

if __name__ == '__main__':
    start_time = time.time()
    main()
    print("--- Executed in {0:.3f} milliseconds ---".format((time.time() - start_time)*1000))