"""Tests for AoC 1, 2024: Historian Hysteria."""

# Standard library imports
import pathlib

# Third party imports
import aoc202403
import pytest

PUZZLE_DIR = pathlib.Path(__file__).parent


@pytest.fixture
def example1():
    puzzle_input = (PUZZLE_DIR / "example1.txt").read_text().rstrip()
    return aoc202403.parse_data(puzzle_input)


@pytest.fixture
def example2():
    puzzle_input = (PUZZLE_DIR / "example2.txt").read_text().rstrip()
    return aoc202403.parse_data(puzzle_input)


def test_parse_example1(example1):
    """Test that input is parsed properly."""
    #assert example1 == [(7,6,4,2,1),(1,2,7,8,9),(9,7,6,2,1),(1,3,2,4,5),(8,6,4,4,1),(1,3,6,7,9)]


def test_part1_example1(example1):
    """Test part 1 on example input."""
    assert aoc202403.part1(example1) == 161


def test_part2_example2(example2):
    """Test part 2 on example input."""
    assert aoc202403.part2(example2) == 48
