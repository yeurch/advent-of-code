"""Tests for AoC 1, 2024: Historian Hysteria."""

# Standard library imports
import pathlib

# Third party imports
import aoc202407
import pytest

PUZZLE_DIR = pathlib.Path(__file__).parent


@pytest.fixture
def example1():
    puzzle_input = (PUZZLE_DIR / "example1.txt").read_text().rstrip()
    return aoc202407.parse_data(puzzle_input)


@pytest.fixture
def example2():
    puzzle_input = (PUZZLE_DIR / "example2.txt").read_text().rstrip()
    return aoc202407.parse_data(puzzle_input)


def test_parse_example1(example1):
    """Test that input is parsed properly."""
    assert example1[1] == (3267, [81, 40, 27])
    assert len(example1) == 9

def test_part1_example1(example1):
    """Test part 1 on example input."""
    assert aoc202407.part1(example1) == 3749

def test_part2_example2(example1):
    """Test part 2 on example input."""
    assert aoc202407.part2(example1) == 11387
