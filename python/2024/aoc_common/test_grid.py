import sys, os
sys.path.insert(0, os.path.join(os.path.dirname(__file__), '..'))

from aoc_common.grid import AOCGrid

data = [[1,2,3],[9,9,0]]

def test_grid_dimensions():
    g = AOCGrid(data)
    assert g.width == 3
    assert g.height == 2

def test_grid_from_chars():
    d = """
123
990
"""
    g = AOCGrid.from_chars(d)
    assert g.width == 3
    assert g.height == 2
    assert g.get_cell(1, 0) == '2'

def test_get_cell():
    g = AOCGrid(data)
    assert g.get_cell(0, 0) == 1
    assert g.get_cell(2, 1) == 0
    assert g.get_cell(9, 9) == None

def test_ray():
    g = AOCGrid(data)
    r = g.get_ray(1, 0, 1, 1, 2)
    assert r == [2, 0]

def test_omni_rays():
    g = AOCGrid(data)
    rays = g.get_omni_rays(1, 0, 2)
    assert rays == [[2,1],[2,3],[2,9],[2,9],[2,0]]
