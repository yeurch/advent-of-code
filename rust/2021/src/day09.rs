use std::collections::HashSet;
use std::collections::HashMap;
use itertools::Itertools; // Trait needed for sorted()

pub fn part1(input: String) {
    const RADIX: u32 = 10;
    let grid: Vec<Vec<u8>> = input.lines().map(|x| x.chars().map(|c| c.to_digit(RADIX).unwrap() as u8).collect()).collect();

    let width = grid[0].len();
    let height = grid.len();

    let minima = get_minima(&grid, width, height);

    let result: u32 = minima.values().map(|n| n.clone() as u32).sum::<u32>() + minima.len() as u32;
    println!("Part 1: {}", result);
}

pub fn part2(input: String) {
    const RADIX: u32 = 10;
    let grid: Vec<Vec<u8>> = input.lines().map(|x| x.chars().map(|c| c.to_digit(RADIX).unwrap() as u8).collect()).collect();

    let width = grid[0].len();
    let height = grid.len();

    let minima = get_minima(&grid, width, height);

    let mut basin_sizes = vec![];
    for p in minima.keys() {
        basin_sizes.push(get_basin_size(&grid, width, height, p.clone(), &mut HashSet::new()));
    }

    let top_3_basins: Vec<u32> = basin_sizes.into_iter().sorted().rev().take(3).collect();
    let result = top_3_basins[0] * top_3_basins[1] * top_3_basins[2];
    println!("Part 2: {}", result);
}

#[derive(PartialEq, Eq, Hash, Clone, Copy)]
struct Point {
    x: usize,
    y: usize
}

fn get_minima(grid: &Vec<Vec<u8>>, width: usize, height: usize) -> HashMap<Point, u8> {
    let mut result = HashMap::new();
    for x in 0..width {
        for y in 0..height {
            let neighbors = get_neighbors(width, height, Point{x,y});
            if !neighbors.into_iter().any(|n| grid[n.y][n.x] <= grid[y][x]) {
                result.insert(Point{x,y}, grid[y][x]);
            }
        }
    }
    result
}

fn get_basin_size(grid: &Vec<Vec<u8>>, width: usize, height: usize, point: Point, already_considered: &mut HashSet<Point>) -> u32 {
    let mut result = 1;
    let neighbors = get_neighbors(width, height, point);
    already_considered.insert(point);
    for n in neighbors {
        let n_height = grid[n.y][n.x];
        if n_height < 9 && n_height >= grid[point.y][point.x] && !already_considered.contains(&n) {
            result += get_basin_size(grid, width, height, n, already_considered);
        }
    }
    result
}

fn get_neighbors(width: usize, height: usize, p: Point) -> Vec<Point> {
    let mut neighbors = vec![];
    let x = p.x;
    let y = p.y;
    if x > 0 { neighbors.push(Point{x:x-1,y}); }
    if x < width - 1 { neighbors.push(Point{x:x+1,y}); }
    if y > 0 { neighbors.push(Point{x,y:y-1}); }
    if y < height - 1 { neighbors.push(Point{x,y:y+1}); }
    neighbors
}