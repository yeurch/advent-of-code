import sys, os
sys.path.insert(0, os.path.join(os.path.dirname(__file__), '..'))

import pathlib
import re
from aoc_common.grid import AOCGrid

def parse_data(puzzle_input):
    return AOCGrid.from_chars(puzzle_input)

def part1(data):
    result = 0
    for x in range(data.width):
        for y in range(data.height):
            rays = data.get_omni_rays(x, y, 4)
            result += rays.count(['X','M','A','S'])
    return result

def part2(data):
    result = 0
    for x in range(1, data.width - 1):
        for y in range(1, data.height - 1):
            val = data.get_cell(x, y)
            if val != 'A':
                continue
            tl = data.get_cell(x-1, y-1)
            tr = data.get_cell(x+1, y-1)
            bl = data.get_cell(x-1, y+1)
            br = data.get_cell(x+1, y+1)
            if (tl == 'M' and br == 'S' or tl == 'S' and br == 'M') and (tr == 'M' and bl == 'S' or tr == 'S' and bl == 'M'):
                result += 1
    return result

def solve(puzzle_input):
    data = parse_data(puzzle_input)
    yield part1(data)
    yield part2(data)

if __name__ == "__main__":
    for path in sys.argv[1:]:
            print(f"\n{path}:")
            solutions = solve(puzzle_input=pathlib.Path(path).read_text().rstrip())
            print("\n".join(str(solution) for solution in solutions))

