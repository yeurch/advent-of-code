# Advent of Code 2024

This directory hosts my solutions to the 2024 [Advent of Code](https://adventofcode.com),
using Python.

These aren't the most elegant solutions, and I'm not trying to get on the
leaderboards (the puzzles are published each day at 5am in my timezone, when
I'm very much definitely sleeping). In fact, they're probably not very good
Python at all, as it's not a language I'd consider myself to be an expert in.

To run each day's puzzle, `cd` into the correct folder for the day and run:

```bash
python aoc2024xx.py input.txt
```
where `xx` is the current day number, from 01 to 25.

### System Requirements

Running the solutions requires Python *3.9* or later.  This is required for
(at the very least) the use of the high-precision timer `time.time_ns()`.

### Optional Packages

Each day comes with a suite of unit tests (in the file prefixed with `test_`.
These are `pytest` tests, and you can run them by installing and running`pytest`:

```bash
pip install pytest
pytest
```
