import sys, os
sys.path.insert(0, os.path.join(os.path.dirname(__file__), '..'))

import copy
import pathlib
from aoc_common.grid import AOCGrid

def parse_data(puzzle_input):
    return AOCGrid.from_chars(puzzle_input)

class Guard():
    def __init__(self, grid):
        self.__grid = grid
        self.pos = grid.find('^')
        if self.pos is None:
            for y in range(self.__grid.height):
                for x in range(self.__grid.width):
                    print(self.__grid.get_cell(x, y), end="")
                print()
        self.__grid.set_cell(self.pos[0], self.pos[1], '*') # visit the first cell
        self.barriers = grid.find_all('#')
        self.facing = (0, -1)

    def move(self):
        while True:
            next_pos = (self.pos[0] + self.facing[0], self.pos[1] + self.facing[1])
            next_tile = self.__grid.get_cell(next_pos[0], next_pos[1])
            if next_tile is None:
                return False
            elif next_tile != '#':
                break
            # We hit an obstacle, turn to the right and try again
            # The following transformation is a right turn
            self.facing = (0 - self.facing[1], self.facing[0])
        self.pos = next_pos
        self.__grid.set_cell(self.pos[0], self.pos[1], '*')
        return True

    def num_squares_visited(self):
        return self.__grid.count('*')

    def get_state(self):
        return f"p=({self.pos[0]},{self.pos[1]});f=({self.facing[0]},{self.facing[1]})"

def part1(data):
    guard = Guard(copy.deepcopy(data))
    while guard.move():
        pass
    return guard.num_squares_visited()

def part2(data):
    # There are width * height * 4 (facings) possible states for the guard.
    # If any repeat, it's a loop.
    result = 0
    for y in range(data.height):
        for x in range(data.width):
            if data.get_cell(x, y) != '.':
                continue # we can only add an obstruction in an empty cell
            visited_states = set()
            grid = copy.deepcopy(data)
            grid.set_cell(x, y, '#')
            guard = Guard(grid)
            should_abort = False
            while True:
                current_state = guard.get_state()
                if current_state in visited_states:
                    result += 1
                    should_abort = True
                    break
                visited_states.add(current_state)
                if not guard.move():
                    break
            if should_abort:
                continue
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

