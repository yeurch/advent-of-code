"""Tests for AoC 9, 2024: Disk Fragmenter."""

# Standard library imports
import pathlib

# Third party imports
import aoc202409
import pytest

PUZZLE_DIR = pathlib.Path(__file__).parent


@pytest.fixture
def example1():
    puzzle_input = (PUZZLE_DIR / "example1.txt").read_text().rstrip()
    return aoc202409.parse_data(puzzle_input)


@pytest.fixture
def example2():
    puzzle_input = (PUZZLE_DIR / "example2.txt").read_text().rstrip()
    return aoc202409.parse_data(puzzle_input)


def test_parse_example1(example1):
    """Test that input is parsed properly."""
    assert True

def test_expand_blocks(example1):
    expected = [-1 if i == '.' else int(i) for i in "00...111...2...333.44.5555.6666.777.888899"]
    assert aoc202409.expand_blocks(example1) == expected

def test_part1_example1(example1):
    """Test part 1 on example input."""
    assert aoc202409.part1(example1) == 1928

def test_part2_example2(example1):
    """Test part 2 on example input."""
    assert aoc202409.part2(example1) == 2858
