use std::collections::HashSet;
use std::str::FromStr;
use serde_scan::scan;

pub fn part1(input: String) {
    let points = do_impl(input, true);
    println!("Part 1: {}", points.len());
}

pub fn part2(input: String) {
    let points = do_impl(input, false);

    let max_x = points.iter().map(|p| p.x).max().unwrap();
    let max_y = points.iter().map(|p| p.y).max().unwrap();

    for y in 0..=max_y {
        for x in 0..=max_x {
            if points.contains(&Point{x,y}) {
                print!("#");
            }
            else {
                print!(".");
            }
        }
        println!("");
    }
}

fn do_impl(input: String, is_part_one: bool) -> HashSet<Point> {
    let mut points: HashSet<Point> = HashSet::new();
    let mut commands: Vec<(FoldDirection, u32)> = vec![];

    for line in input.lines() {
        if let Ok(point) = line.parse() {
            points.insert(point);
        }
        else if line.len() > 0 {
            let parsed_line: (String, u32) = scan!("fold along {}={}" <- line).unwrap();
            commands.push((parsed_line.0.parse().unwrap(), parsed_line.1));
        }
    }

    for command in commands {
        apply_command(command.0, command.1, &mut points);

        if is_part_one {
            break;
        }
    }
    points
}

fn apply_command(dir: FoldDirection, axis_val: u32, points: &mut HashSet<Point>) {
    let mut to_add = HashSet::new();
    let mut to_drain = HashSet::new();
    for p in points.iter() {
        if dir == FoldDirection::Left && p.x > axis_val {
            let new_x = 2 * axis_val - p.x;
            to_add.insert(Point {x: new_x, y: p.y});
            to_drain.insert(Point {x: p.x, y: p.y});
        }
        else if dir == FoldDirection::Up && p.y > axis_val {
            let new_y = 2 * axis_val - p.y;
            to_add.insert(Point {x: p.x, y: new_y});
            to_drain.insert(Point {x: p.x, y: p.y});
        }
    }

    for p_add in to_add {
        let _ = &points.insert(p_add);
    }
    for p_remove in to_drain {
        let _ = &points.remove(&p_remove);
    }
}


#[derive(Eq,PartialEq,Hash,Clone,Copy)]
struct Point {
    x: u32,
    y: u32
}

impl FromStr for Point {
    type Err = ();
    fn from_str(input: &str) -> Result<Point, Self::Err> {
        if !input.contains(",") {
            return Err(());
        }
        let parts = input.split(",").collect::<Vec<&str>>();
        let xp = parts[0].parse();
        let yp = parts[1].parse();
        if let Err(_) = xp {
            return Err(());
        }
        if let Err(_) = yp {
            return Err(());
        }
        Ok(Point {
            x: xp.unwrap(),
            y: yp.unwrap()
        })
    }
}

#[derive(PartialEq)]
enum FoldDirection {
    Up,
    Left
}

impl FromStr for FoldDirection {
    type Err = ();
    fn from_str(input: &str) -> Result<FoldDirection, Self::Err> {
        match input {
            "x" => Ok(FoldDirection::Left),
            "y" => Ok(FoldDirection::Up),
            _ => Err(()),
        }
    }
}