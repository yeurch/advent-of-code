import sys, os
sys.path.insert(0, os.path.join(os.path.dirname(__file__), '..'))

from aoc_common.math import cmp

def test_cmp_big_first():
    assert cmp(4, 6) == -1

def test_cmp_same():
    assert cmp(42, 42) == 0

def test_cmp_small_first():
    assert cmp(6, 3) == 1
