# Advent of Code 2019

This directory hosts my solutions to the 2019 [Advent of Code](https://adventofcode.com),
using Rust.

I'm backfilling these puzzles while working on the 2021 challenges (I started Advent of Code in 2020).

### How to

For each day:

* Create src/dayxx.rs. Pad 0 for 1-9 so that files sort properly. 
* Write a function like `pub fn part1(input: String) {}`.
* In src/lib.rs, add a line with `pub mod dayxx;`.
* In src/lib.rs, add a case to the match, for example:

```
1 => (day01::part1, day01::part2),
2 => (day02::part1, noop),
```

Use noop whenever a part is not yet implemented.

* Create inputs/dayxx.txt and add your puzzle input
* Run `cargo run <n>` where _n_ is the day number.
* Your code will be passed the input and timed automatically.

### Acknoweldgements

The starter template for running the solutions was provided by https://aoc-templates.util.repl.co/.

### System Requirements

I'm using the latest edition of Rust, 1.56.1, installed via `rustup` on Windows 10. You can get
Rust from [the official Rust website](https://www.rust-lang.org/learn/get-started).