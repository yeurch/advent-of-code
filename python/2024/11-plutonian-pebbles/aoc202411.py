import sys, os
sys.path.insert(0, os.path.join(os.path.dirname(__file__), '..'))
import pathlib

def parse_data(puzzle_input):
    values = [int(i) for i in puzzle_input.split()]
    return {x: values.count(x) for x in set(values)}

def blink(stones):
    result = dict()
    for stone in stones:
        if stone == 0:
            result[1] = result.get(1, 0) + stones[stone]
        elif len(str(stone)) % 2 == 0:
            # an even number of digits
            halfway = len(str(stone)) // 2
            first_half = int(str(stone)[halfway:])
            result[first_half] = result.get(first_half, 0) + stones[stone]
            second_half = int(str(stone)[:halfway])
            result[second_half] = result.get(second_half, 0) + stones[stone]
        else:
            result[2024*stone] = result.get(2024*stone, 0) + stones[stone]
    return result

def do_puzzle(data, num):
    stones = data
    for i in range(num):
        stones = blink(stones)
    return sum(stones.values())

def part1(data):
    return do_puzzle(data, 25)

def part2(data):
    return do_puzzle(data, 75)

def solve(puzzle_input):
    data = parse_data(puzzle_input)
    yield part1(data)
    yield part2(data)

if __name__ == "__main__":
    for path in sys.argv[1:]:
            print(f"\n{path}:")
            solutions = solve(puzzle_input=pathlib.Path(path).read_text().rstrip())
            print("\n".join(str(solution) for solution in solutions))
