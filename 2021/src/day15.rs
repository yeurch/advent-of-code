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

    let mut last = 0;
    set_value(&mut grid, Point{x:0,y:0}, 0);
    for x in 1..size {
        let p = Point{x,y:0};
        last = last + get_value(&cell_costs, p);
        set_value(&mut grid, Point{x,y:0}, last);
    }
    for x in 0..size {
        let _ = do_column(size, &cell_costs, &mut grid, Point{x,y:1}, true, 0);
    }

    get_value(&grid, Point{x:size-1,y:size-1})
}

fn do_column(size: usize, cell_costs: &Grid, grid: &mut Grid, p: Point, always_set: bool, init_value: u16) -> Option<usize> {
    let x = p.x;
    let mut last = get_value(grid, Point{x,y:p.y-1});
    if !always_set {
        set_value(grid, Point{x,y:p.y}, init_value);
        last = init_value - get_value(cell_costs, Point{x,y:p.y});
    }
    for y in p.y..size {
        let cell_value = get_value(cell_costs, Point{x,y});
        last = last + cell_value;
        if x > 0 {
            if last > get_value(grid, Point{x:x-1,y}) + cell_value {
                last = get_value(grid, Point{x:x-1,y}) + cell_value
            }
            let cell_left = Point{x:x-1,y};
            if last + get_value(cell_costs, cell_left) < get_value(grid, cell_left) {
                let do_result = do_column(size, cell_costs, grid, cell_left, false, last + get_value(cell_costs, cell_left));
                if let Some(row) = do_result {
                    let _ = do_column(size, cell_costs, grid, Point{x,y:row}, true, 0);
                }
            }
        }

        if always_set || last < get_value(grid, Point{x,y}) {
            set_value(grid, Point{x,y}, last);
        }
        else {
            break;
        }
    }

    if !always_set {
        last = get_value(grid, Point{x,y:p.y});
        for y in (0..p.y).rev() {
            let cell_value = get_value(cell_costs, Point{x,y});
            last = last + cell_value;
            if x > 0 {
                if last > get_value(grid, Point{x:x-1,y}) + cell_value {
                    last = get_value(grid, Point{x:x-1,y}) + cell_value
                }
                let cell_left = Point{x:x-1,y};
                if last + get_value(cell_costs, cell_left) < get_value(grid, cell_left) {
                    do_column(size, cell_costs, grid, cell_left, false, last + get_value(cell_costs, cell_left));
                }
            }

            if last < get_value(grid, Point{x,y}) {
                set_value(grid, Point{x,y}, last);
            }
            else {
                return Some(y);
            }
        }
    }
    None
}

type Grid = Vec<Vec<u16>>;

#[derive(Clone,Copy)]
struct Point {
    x: usize,
    y: usize
}

fn get_value(grid: &Grid, p: Point) -> u16 {
    grid[p.y][p.x]
}

fn set_value(grid: &mut Grid, p: Point, val: u16) {
    grid[p.y][p.x] = val;
}