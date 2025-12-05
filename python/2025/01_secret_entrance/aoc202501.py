import pathlib
import sys

def parse_data(puzzle_input):
    numbers = [int(line[1:]) * (-1 if line[0] == 'L' else 1) for line in puzzle_input.split('\n')]
    return numbers

def part1(data):
    pos = 50
    result = 0
    for i in data:
        pos += i
        if pos % 100 == 0:
            result += 1
    return result

def part2(data):
    pos = 50
    result = 0
    for i in data:
        dir = -1 if i < 0 else 1
        for j in range(abs(i)):
            pos += dir
            if pos % 100 == 0:
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

