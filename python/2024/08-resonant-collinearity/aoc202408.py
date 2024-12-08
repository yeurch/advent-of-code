import sys, os
sys.path.insert(0, os.path.join(os.path.dirname(__file__), '..'))
import pathlib

from aoc_common.grid import AOCGrid

def parse_data(puzzle_input):
    return AOCGrid.from_chars(puzzle_input)

def part1(data):
    frequencies = "ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz0123456789"
    antinodes = set()
    for freq in frequencies:
        antennae = data.find_all(freq)
        for pos in antennae:
            # We're iterating over all antennae of the given frequency
            for other in antennae:
                if other == pos:
                    continue
                # and iterating over all other antennae of the same frequency
                vec = (other[0]-pos[0], other[1]-pos[1])
                antinode = (pos[0]-vec[0], pos[1]-vec[1])
                antinodes.add(antinode)
    return len([n for n in antinodes if n[0] >= 0 and n[1] >= 0 and n[0] < data.width and n[1] < data.height])

def part2(data):
    frequencies = "ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz0123456789"
    antinodes = set()
    for freq in frequencies:
        antennae = data.find_all(freq)
        for pos in antennae:
            # We're iterating over all antennae of the given frequency
            for other in antennae:
                if other == pos:
                    continue
                # and iterating over all other antennae of the same frequency
                vec = (other[0]-pos[0], other[1]-pos[1])
                antinode = pos
                while True:
                    antinodes.add(antinode)
                    antinode = (antinode[0] - vec[0], antinode[1] - vec[1])
                    if antinode[0] < 0 or antinode[1] < 0 or antinode[0] >= data.width or antinode[1] >= data.height:
                        break
                antinode = other
                while True:
                    antinodes.add(antinode)
                    antinode = (antinode[0] + vec[0], antinode[1] + vec[1])
                    if antinode[0] < 0 or antinode[1] < 0 or antinode[0] >= data.width or antinode[1] >= data.height:
                        break
    return len(antinodes)

    return len([n for n in antinodes if n[0] >= 0 and n[1] >= 0 and n[0] < data.width and n[1] < data.height])


def solve(puzzle_input):
    data = parse_data(puzzle_input)
    yield part1(data)
    yield part2(data)

if __name__ == "__main__":
    for path in sys.argv[1:]:
            print(f"\n{path}:")
            solutions = solve(puzzle_input=pathlib.Path(path).read_text().rstrip())
            print("\n".join(str(solution) for solution in solutions))