use std::collections::HashSet;
use std::collections::HashMap;
use std::str::FromStr;

pub fn part1(input: String) {
    let data: Vec<Vec<&str>> = input.lines().map(|line| line.split(",").collect()).collect();
    let lines: Vec<Vec<Cmd>> = data.into_iter().map(|d| d.into_iter().map(|x| x.parse().unwrap()).collect()).collect();

    let mut grid: HashMap<Point, HashSet<u8>> = HashMap::new();

    for (line_num,line) in lines.iter().enumerate() {
        let mut pos = Point{x:0,y:0};
        for cmd in line {
            let deltas = cmd.dir.get_deltas();
            for _ in 0..cmd.distance {
                pos.translate(deltas);
                grid.entry(pos).or_insert(HashSet::new()).insert(line_num as u8);
            }
        }
    }

    let mut min_dist = u32::MAX;
    for (k,v) in grid {
        if v.len() > 1 {
            let dist = k.manhattan_dist();
            if dist < min_dist {
                min_dist = dist;
            }
        }
    }

    println!("Part 1: {}", min_dist);
}


#[derive(PartialEq,Eq,Hash,Clone,Copy)]
struct Point {
    x: i32,
    y: i32
}

impl Point {
    fn translate(&mut self, deltas: (i32, i32)) {
        self.x += deltas.0;
        self.y += deltas.1;
    }

    fn manhattan_dist(&self) -> u32 {
        (self.x.abs() + self.y.abs()) as u32
    }
}

enum Direction {
    U,
    D,
    L,
    R
}

impl Direction {
    fn get_deltas(&self) -> (i32, i32) {
        match self {
            Direction::U => (0,-1),
            Direction::D => (0,1),
            Direction::L => (-1,0),
            Direction::R => (1,0)
        }
    }
}

struct Cmd {
    dir: Direction,
    distance: u32
}

impl FromStr for Cmd {
    type Err = ();

    fn from_str(input: &str) -> Result<Cmd, Self::Err> {
        let dir = match &input[0..1] {
            "U" => Direction::U,
            "D" => Direction::D,
            "L" => Direction::L,
            "R" => Direction::R,
            _ => return Err(())
        };
        Ok(Cmd {
            dir,
            distance: input[1..].parse().unwrap()
        })
    }
}