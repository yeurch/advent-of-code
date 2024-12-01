import pathlib
import sys

def parse_data(puzzle_input):
    numbers = list(zip(*[[int(number) for number in line.split()] for line in puzzle_input.split('\n')]))
    return numbers

def part1(data):
    pairs = zip(sorted(data[0]),sorted(data[1]))
    return sum([abs(i[0]-i[1]) for i in pairs])

def part2(data):
    return sum([i*data[1].count(i) for i in data[0]])

def solve(puzzle_input):
    data = parse_data(puzzle_input)
    yield part1(data)
    yield part2(data)

if __name__ == "__main__":
    for path in sys.argv[1:]:
            print(f"\n{path}:")
            solutions = solve(puzzle_input=pathlib.Path(path).read_text().rstrip())
            print("\n".join(str(solution) for solution in solutions))

