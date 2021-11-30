# Advent of Code 2021

This directory hosts my solutions to the 2021 [Advent of Code](https://adventofcode.com),
using Rust.

I've never written anything beyond a Hello World in Rust before, and have spent a few hours
reading [The Rust Book](https://doc.rust-lang.org/book/) (but not practiced anything in it).
I feel that this year is going to be quite a challenge, and similar to last year, I'm unlikely
to produce any great examples of idiomatic Rust.

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
* Run `cargo run <n>` where _n_ is the daya number.
* Your code will be passed the input and timed automatically.

### Acknoweldgements

The starter template for running the solutions was provided by https://aoc-templates.util.repl.co/.

### System Requirements

I'm using the latest edition of Rust, 1.56.1, installed via `rustup` on Windows 10. You can get
Rust from [the official Rust website](https://www.rust-lang.org/learn/get-started).