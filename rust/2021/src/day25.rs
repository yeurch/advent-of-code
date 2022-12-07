pub fn part1(input: String) {
    let mut grid: Vec<Vec<char>> = input.lines().map(|l| l.chars().collect()).collect();

    let mut num_moved = usize::MAX;
    let mut result = 0;
    while num_moved > 0 {
        num_moved = do_move(&mut grid, '>', 1, 0);
        num_moved += do_move(&mut grid, 'v', 0, 1);
        result += 1;
    }

    print_grid(&grid);
    println!("Part 1: {}", result);
}

fn print_grid(grid: &Vec<Vec<char>>) {
    for line in grid {
        for c in line {
            print!("{}", c);
        }
        println!("");
    }
    println!("");
}

fn do_move(grid: &mut Vec<Vec<char>>, char_to_move: char, dx: usize, dy: usize) -> usize {
    let mut to_move = vec![];
    let y_size = grid.len();
    let x_size = grid[0].len();
    
    for (y, line) in grid.iter().enumerate() {
        for (x, c) in line.iter().enumerate() {
            if *c == char_to_move && grid[(y+dy)%y_size][(x+dx)%x_size] == '.' {
                to_move.push((x,y));
            }
        } 
    }

    let result = to_move.len();

    for (x,y) in to_move {
        grid[y][x] = '.';
        grid[(y+dy)%y_size][(x+dx)%x_size] = char_to_move;
    }
    result
}