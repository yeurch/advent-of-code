import pathlib
import sys
import re

def parse_data(puzzle_input):
    return puzzle_input

def part1(data):
    pattern = r"mul\((\d{1,3}),(\d{1,3})\)"
    matches = re.findall(pattern, data)
    return sum([int(x[0])*int(x[1]) for x in matches])

def part2(data):
    result = 0
    parts = data.split('do()')
    for part in parts:
        sub_part = part.split("don't()")[0]
        result += part1(sub_part)
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

