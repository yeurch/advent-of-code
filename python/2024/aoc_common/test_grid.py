import sys, os
sys.path.insert(0, os.path.join(os.path.dirname(__file__), '..'))

import pytest
from aoc_common.grid import AOCGrid

@pytest.fixture
def grid():
    vals = [[1,2,3],[9,9,0]]
    yield AOCGrid(vals)

def test_grid_dimensions(grid):
    assert grid.width == 3
    assert grid.height == 2

def test_grid_from_chars():
    d = """
123
990
"""
    g = AOCGrid.from_chars(d)
    assert g.width == 3
    assert g.height == 2
    assert g.get_cell(1, 0) == '2'

def test_get_cell(grid):
    assert grid.get_cell(0, 0) == 1
    assert grid.get_cell(2, 1) == 0
    assert grid.get_cell(9, 9) == None

def test_set_cell(grid):
    grid.set_cell(1, 1, 'test_val')
    assert grid.get_cell(1, 1) == 'test_val'

def test_ray(grid):
    r = grid.get_ray(1, 0, 1, 1, 2)
    assert r == [2, 0]

def test_omni_rays(grid):
    rays = grid.get_omni_rays(1, 0, 2)
    assert rays == [[2,1],[2,3],[2,9],[2,9],[2,0]]

def test_find(grid):
    assert grid.find(0) == (2,1)

def test_find_all(grid):
    assert grid.find_all(9) == [(0,1),(1,1)]

def test_count(grid):
    assert grid.count(5) == 0
    assert grid.count(1) == 1
    assert grid.count(9) == 2

def test_neighbours(grid):
    assert grid.get_neighbour_locations(1,0) == [(2,0),(1,1),(0,0)]
