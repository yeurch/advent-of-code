pub fn part1(input: String) {
    const RADIX: u32 = 10;
    let mut grid: Vec<Vec<u8>> = input.lines().map(|x| x.chars().map(|c| c.to_digit(RADIX).unwrap() as u8).collect()).collect();

    let mut result = 0;
    for _ in 0..100 {
        result += do_step(&mut grid);
    }

    println!("Part 1: {}", result);
}

pub fn part2(input: String) {
    const RADIX: u32 = 10;
    let mut grid: Vec<Vec<u8>> = input.lines().map(|x| x.chars().map(|c| c.to_digit(RADIX).unwrap() as u8).collect()).collect();

    let mut result = 1;
    while do_step(&mut grid) < 100 {
        result += 1;
    }

    println!("Part 2: {}", result);
}

struct Point {
    x: usize,
    y: usize
}

fn do_step(grid: &mut Vec<Vec<u8>>) -> u16 {
    let mut num_flashes = 0;

    for x in 0..10 {
        for y in 0..10 {
            grid[y][x] += 1;
            if grid[y][x] == 10 {
                num_flashes += flash(grid, Point{x,y});
            }
        }
    }

    for row in grid {
        for col in row {
            if *col > 9 { *col = 0; }
        }
    }

    num_flashes
}

fn flash(grid: &mut Vec<Vec<u8>>, p: Point) -> u16 {
    let neighbors = get_neighbors(p);
    let mut num_flashes = 1;

    for neighbor in neighbors {
        grid[neighbor.y][neighbor.x] += 1;
        if grid[neighbor.y][neighbor.x] == 10 {
            num_flashes += flash(grid, neighbor);
        }
    }
    num_flashes
}

fn get_neighbors(p: Point) -> Vec<Point> {
    let mut neighbors = vec![];
    let x = p.x;
    let y = p.y;
    if x > 0 {
        neighbors.push(Point{x:x-1,y});
        if y > 0 { neighbors.push(Point{x:x-1,y:y-1}); }
        if y < 9 { neighbors.push(Point{x:x-1,y:y+1}); }
        }
    if x < 9 {
        neighbors.push(Point{x:x+1,y});
        if y > 0 { neighbors.push(Point{x:x+1,y:y-1}); }
        if y < 9 { neighbors.push(Point{x:x+1,y:y+1}); }
        }
    if y > 0 { neighbors.push(Point{x,y:y-1}); }
    if y < 9 { neighbors.push(Point{x,y:y+1}); }
    neighbors
}