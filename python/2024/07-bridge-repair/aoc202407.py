import sys
import pathlib

def parse_data(puzzle_input):
    raw_vals = [[int(x.replace(":","")) for x in line.split()] for line in puzzle_input.rstrip().split('\n')]
    return [(v[0], v[1:]) for v in raw_vals]

def is_solvable(item, use_concat):
    target, operands = item
    if len(operands) == 1:
        return operands[0] == target
    op0 = operands[0]
    op1 = operands[1]
    remainder = operands[2:]

    return is_solvable((target, [op0+op1]+remainder), use_concat) or \
        is_solvable((target, [op0*op1]+remainder), use_concat) or \
        (use_concat and is_solvable((target, [int(str(op0)+str(op1))]+remainder), use_concat))

def part1(data):
    return sum([item[0] if is_solvable(item, False) else 0 for item in data])

def part2(data):
    return sum([item[0] if is_solvable(item, True) else 0 for item in data])

def solve(puzzle_input):
    data = parse_data(puzzle_input)
    yield part1(data)
    yield part2(data)

if __name__ == "__main__":
    for path in sys.argv[1:]:
            print(f"\n{path}:")
            solutions = solve(puzzle_input=pathlib.Path(path).read_text().rstrip())
            print("\n".join(str(solution) for solution in solutions))
