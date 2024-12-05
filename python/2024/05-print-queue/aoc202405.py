import sys
import pathlib

def parse_data(puzzle_input):
    lines = puzzle_input.split('\n')
    rules = []
    page_sets = []
    for line in lines:
        if '|' in line:
            rules.append(tuple([int(i) for i in line.split('|')]))
        elif len(line.strip()) > 0:
            page_sets.append([int(i) for i in line.split(',')])
    return {
            "rules": rules,
            "page_sets": page_sets
    }

def is_valid(page_set, rules):
    for rule in rules:
        try:
            idx1 = page_set.index(rule[0])
            idx2 = page_set.index(rule[1])
        except ValueError:
            continue
        if idx2 < idx1:
            return False
    return True

def part1(data):
    result = 0
    for page_set in data["page_sets"]:
        if is_valid(page_set, data["rules"]):
            result += page_set[len(page_set) // 2]
    return result

def fix_set(page_set, rules, num_iter):
    """
    We find the first value in the page_set, assuming that it is the only one that has a
    rule with it on the left and no rules with it on the right, when we only consider
    rules that have the values of page_set in it.
    We then recurse until we get to the middle value, removing the left-most value we found,
    using num_iter as our exit condition.
    """
    rules = [x for x in rules if x[0] in page_set and x[1] in page_set]
    lead_value = [x[0] for x in rules if len([rhs for rhs in rules if rhs[1] == x[0]]) == 0][0]
    if num_iter == 0:
        return lead_value
    return fix_set([i for i in page_set if i != lead_value], rules, num_iter - 1)

def part2(data):
    result = 0
    for page_set in data["page_sets"]:
        if not is_valid(page_set, data["rules"]):
            result += fix_set(page_set, data["rules"], len(page_set) // 2)
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

