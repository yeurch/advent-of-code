import pathlib
import sys

def parse_data(puzzle_input):
    return [tuple([int(number) for number in line.split()]) for line in puzzle_input.split('\n')]

def cmp(a, b):
    return (a > b) - (a < b)

def is_safe(report):
    direction = cmp(report[1], report[0])
    if direction == 0:
        return False

    for idx in range(1, len(report)):
        delta = report[idx] - report[idx - 1]
        if cmp(delta, 0) != direction or abs(delta) < 1 or abs(delta) > 3:
            return False
    return True

def part1(data):
    safe_count = 0
    for report in data:
        safe_count += 1 if is_safe(report) else 0
    return safe_count

def part2(data):
    safe_count = 0
    for report in data:
        if is_safe(report):
            safe_count += 1
            continue
        possible_subreports = [report[:i] + report[i+1:] for i in range(len(report))]
        for possible_subreport in possible_subreports:
            if is_safe(possible_subreport):
                safe_count += 1
                break
    return safe_count

def solve(puzzle_input):
    data = parse_data(puzzle_input)
    yield part1(data)
    yield part2(data)

if __name__ == "__main__":
    for path in sys.argv[1:]:
            print(f"\n{path}:")
            solutions = solve(puzzle_input=pathlib.Path(path).read_text().rstrip())
            print("\n".join(str(solution) for solution in solutions))

