"""Tests for AoC 1, 2024: Historian Hysteria."""

# Standard library imports
import pathlib

# Third party imports
import aoc202405
import pytest

PUZZLE_DIR = pathlib.Path(__file__).parent


@pytest.fixture
def example1():
    puzzle_input = (PUZZLE_DIR / "example1.txt").read_text().rstrip()
    return aoc202405.parse_data(puzzle_input)


@pytest.fixture
def example2():
    puzzle_input = (PUZZLE_DIR / "example2.txt").read_text().rstrip()
    return aoc202405.parse_data(puzzle_input)


def test_parse_example1(example1):
    """Test that input is parsed properly."""
    simple_input = """
1|3
2|4
4|3

1,3,4
2,4
"""
    assert aoc202405.parse_data(simple_input) == {
        "rules": [(1,3),(2,4),(4,3)],
        "page_sets": [[1,3,4],[2,4]]
    }


def test_part1_example1(example1):
    """Test part 1 on example input."""
    assert aoc202405.part1(example1) == 143


def test_part2_example2(example1):
    """Test part 2 on example input."""
    assert aoc202405.part2(example1) == 123
