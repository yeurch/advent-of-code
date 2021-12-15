use std::collections::HashSet;

pub fn part1(input: String) {
    const RADIX: u32 = 10;
    let cell_costs: Grid = input.lines().map(|x| x.chars().map(|c| c.to_digit(RADIX).unwrap() as u16).collect()).collect();

    let result = do_impl(cell_costs);
    println!("Part 1: {}", result);
}

pub fn part2(input: String) {
    const RADIX: u32 = 10;
    let base: Grid = input.lines().map(|x| x.chars().map(|c| c.to_digit(RADIX).unwrap() as u16).collect()).collect();
    let base_size = base.len();

    let mut cell_costs = vec![];
    for big_y in 0..5 {
        for y in 0..base_size {
            let mut row = vec![];
            for big_x in 0..5 {
                for x in 0..base_size {
                    let mut cell_val = get_value(&base, Point{x,y}) + big_y + big_x;
                    if cell_val > 9 {
                        cell_val = cell_val % 9;
                    }
                    row.push(cell_val);
                }
            }
            cell_costs.push(row);
        }
    }
    let result = do_impl(cell_costs);
    println!("Part 2: {}", result);
}

fn do_impl(cell_costs: Grid) -> u16 {
    let size = cell_costs.len();

    let mut grid = cell_costs.clone();
    let mut to_refresh = HashSet::new();

    let mut last = 0;
    set_value(&mut grid, Point{x:0,y:0}, 0);
    for x in 1..size {
        let p = Point{x,y:0};
        last = last + get_value(&cell_costs, p);
        set_value(&mut grid, Point{x,y:0}, last);
    }
    for x in 0..size {
        last = get_value(&grid, Point{x,y:0});
        for y in 1..size {
            let cell_value = get_value(&cell_costs, Point{x,y});
            last = last + cell_value;
            if x > 0 && last > get_value(&grid, Point{x:x-1,y}) + cell_value {
                last = get_value(&grid, Point{x:x-1,y}) + cell_value
            }
            if x > 0 {
                let cell_left = Point{x:x-1,y};
                if last + get_value(&cell_costs, cell_left) < get_value(&grid, cell_left) {
                    to_refresh.insert(cell_left);
                }
            }
            set_value(&mut grid, Point{x,y}, last);
        }
    }

    // println!("Initial refresh cells count = {}", to_refresh.len());

    while let Some(p) = to_refresh.iter().next().cloned() {
        to_refresh.remove(&p);
        refresh(size, &cell_costs, &mut grid, p, &mut to_refresh);
    }

    get_value(&grid, Point{x:size-1,y:size-1})
}

type Grid = Vec<Vec<u16>>;

#[derive(PartialEq,Eq,Clone,Copy,Hash)]
struct Point {
    x: usize,
    y: usize
}

fn refresh(size: usize, cell_costs: &Grid, grid: &mut Grid, p: Point, refresh_queue: &mut HashSet<Point>) {
    let cell_cost = get_value(cell_costs, p) as u32;
    let mut value = get_value(grid, p) as u32;

    let neighbors = get_neighbors(size, p, grid);
    let min_neighbor_val = neighbors.iter().map(|n| n.1).min().unwrap() as u32;
    if min_neighbor_val + cell_cost < value {
        value = min_neighbor_val + cell_cost;
        set_value(grid, p, value as u16);
    }
    for neighbor in neighbors {
        if neighbor.1 as u32 > get_value(cell_costs, neighbor.0) as u32 + value as u32 {
            refresh_queue.insert(neighbor.0);
        }
    }
}

fn get_neighbors(size: usize, p: Point, grid: &Grid) -> Vec<(Point,u16)> {
    let mut neighbors = vec![];
    let x = p.x;
    let y = p.y;
    if x > 0 {
        let n = Point{x:x-1,y};
        neighbors.push((n, get_value(grid, n)));
    }
    if x < size - 1 {
        let n = Point{x:x+1,y};
        neighbors.push((n, get_value(grid, n)));
    }
    if y > 0 {
        let n = Point{x,y:y-1};
        neighbors.push((n, get_value(grid, n)));
    }
    if y < size - 1 {
        let n = Point{x,y:y+1};
        neighbors.push((n, get_value(grid, n)));
    }
    neighbors
}

fn get_value(grid: &Grid, p: Point) -> u16 {
    grid[p.y][p.x]
}

fn set_value(grid: &mut Grid, p: Point, val: u16) {
    grid[p.y][p.x] = val;
}