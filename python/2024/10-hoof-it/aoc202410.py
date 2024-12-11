import sys, os
sys.path.insert(0, os.path.join(os.path.dirname(__file__), '..'))
import pathlib
import copy
from aoc_common.grid import AOCGrid

def parse_data(puzzle_input):
    return AOCGrid.from_chars(puzzle_input)

def trailhead_score(grid, trailhead):
    trail_extent = {trailhead}
    for i in range(1, 10):
        new_trail_extent = set()
        for position in trail_extent:
            neighbours = grid.get_neighbour_locations(position[0],position[1])
            uphill = [n for n in neighbours if grid.get_cell(n[0],n[1]) == str(i)]
            new_trail_extent.update(uphill)
        trail_extent = new_trail_extent
    return len(trail_extent)

def trailhead_rating(grid, trailhead):
    routes = find_next(grid, [[trailhead]])
    return len(routes)

def find_next(grid, routes):
    if len(routes) == 0:
        return None
    target_height = int(grid.get_cell(routes[0][-1][0], routes[0][-1][1]))+1
    if target_height == 10:
        return routes
    new_routes = []
    for route in routes:
        last_pos = route[-1]
        neighbours = grid.get_neighbour_locations(last_pos[0], last_pos[1])
        uphill = [n for n in neighbours if int(grid.get_cell(n[0],n[1])) == target_height]
        for new_location in uphill:
            new_routes.append(copy.deepcopy(route) + [new_location])
    return find_next(grid, new_routes)

def part1(data):
    trailheads = data.find_all('0')
    result = 0
    for trailhead in trailheads:
        result += trailhead_score(data, trailhead)
    return result

def part2(data):
    trailheads = data.find_all('0')
    result = 0
    for trailhead in trailheads:
        result += trailhead_rating(data, trailhead)
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
