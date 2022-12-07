use std::collections::HashMap;
use serde_scan::scan;

pub fn part1(input: String) {
    let result = do_impl(input, false);
    println!("Part 1: {}", result);
}

pub fn part2(input: String) {
    let result = do_impl(input, true);
    println!("Part 2: {}", result);
}

fn do_impl(input: String, incl_diag: bool) -> u32 {
    let line_defs = parse(&input);

    let mut grid = HashMap::new();

    for line_def in line_defs {
        if line_def.0.x == line_def.1.x {
            let range = if line_def.0.y <= line_def.1.y {line_def.0.y ..= line_def.1.y} else {line_def.1.y ..= line_def.0.y};
            for y in range {
                let p = Point { x: line_def.0.x, y };
                *grid.entry(p).or_default() += 1;
            }
        }
        else if line_def.0.y == line_def.1.y {
            let range = if line_def.0.x <= line_def.1.x {line_def.0.x ..= line_def.1.x} else {line_def.1.x ..= line_def.0.x};
            for x in range {
                let p = Point { x, y: line_def.0.y };
                *grid.entry(p).or_default() += 1;
            }
        }
        else if incl_diag {
            // Diagonal line
            let dx: i32 = if line_def.0.x < line_def.1.x {1} else {-1};
            let dy: i32 = if line_def.0.y < line_def.1.y {1} else {-1};
            let mut x = line_def.0.x;
            let mut y = line_def.0.y;
            let final_p = Point { x: line_def.1.x, y: line_def.1.y };
            let mut p = Point { x, y };
            loop {
                let last = p == final_p; // entry below takes ownership of p, so we need to check here if we're the last iteration
                *grid.entry(p).or_default() += 1;
                if last {
                    break;
                }
                x = (x as i32 + dx) as u32;
                y = (y as i32 + dy) as u32;
                p = Point { x, y };
            }
        }
    }

    grid.into_values().map(|n: u32| if n > 1 {1} else {0}).sum()
}

#[derive(Eq, PartialEq, Hash)]
struct Point {
    x: u32,
    y: u32
}

fn parse(input: &str) -> Vec<(Point, Point)> {
    input
        .lines()
        .map(|line| scan!("{},{} -> {},{}" <- line).unwrap())
        .map(|t: (u32, u32, u32, u32)| (Point{x:t.0, y:t.1}, Point{x:t.2, y:t.3}))
        .collect()
}
