use std::collections::HashSet;

pub fn part1(input: String) {
    println!("Part 1: {}", do_impl(input, 2));
}

pub fn part2(input: String) {
    println!("Part 2: {}", do_impl(input, 50));
}

fn do_impl(input: String, n: usize) -> usize {
    let algo = build_enhancement_algorithm(input.lines().nth(0).unwrap());

    let mut grid = Grid::new();
    for (y,line) in input.lines().skip(2).enumerate() {
        for (x,c) in line.chars().enumerate() {
            if c == '#' {
                grid.add_point(Point{x: x as i64, y: y as i64});
            }
        }
    }

    for _ in 0..n {
        grid = grid.enhance(&algo);
    }
    grid.num_lit_points()
}

#[derive(PartialEq,Eq,Hash,Clone,Copy)]
struct Point {
    x: i64,
    y: i64
}

struct Grid {
    lit_points: HashSet<Point>,
    outside_val: u64,
    min_bound: Point,
    max_bound: Point
}

impl Grid {
    fn new() -> Self {
        Grid {
            lit_points: HashSet::new(),
            outside_val: 0,
            min_bound: Point{x:i64::MAX,y:i64::MAX},
            max_bound: Point{x:i64::MIN,y:i64::MIN},
        }
    }

    fn add_point(&mut self, p: Point) {
        self.lit_points.insert(p);
        if p.x > self.max_bound.x { self.max_bound = Point{x:p.x,y:self.max_bound.y}; }
        if p.x < self.min_bound.x { self.min_bound = Point{x:p.x,y:self.min_bound.y}; }
        if p.y > self.max_bound.y { self.max_bound = Point{x:self.max_bound.x,y:p.y}; }
        if p.y < self.min_bound.y { self.min_bound = Point{x:self.min_bound.x,y:p.y}; }
    }

    fn num_lit_points(&self) -> usize {
        self.lit_points.len()
    }

    fn get(&self, p: Point) -> u64 {
        if p.x < self.min_bound.x || p.x > self.max_bound.x || p.y < self.min_bound.y || p.y > self.max_bound.y {
            return self.outside_val;
        }
        if self.lit_points.contains(&p) {1} else {0}
    }

    fn enhance(self, algo: &HashSet<u64>) -> Grid {
        let mut result = Grid::new();

        let min_x = self.min_bound.x - 1;
        let min_y = self.min_bound.y - 1;
        let max_x = self.max_bound.x + 1;
        let max_y = self.max_bound.y + 1;

        for y in min_y..=max_y {
            for x in min_x..=max_x {
                let p = Point{x,y};
                let mut algo_index = 0;
                for n in get_neighbors(p) {
                    algo_index = algo_index << 1;
                    algo_index += self.get(n);
                }
                if algo.contains(&algo_index) {
                    result.add_point(p);
                }
            }
        }

        result.min_bound = Point{x:min_x, y:min_y};
        result.max_bound = Point{x:max_x, y:max_x};

        if algo.contains(&0) {
            result.outside_val = 1 - self.outside_val;
        }
        else {
            result.outside_val = self.outside_val;
        }
        result
    }
}

fn get_neighbors(p: Point) -> Vec<Point> {
    let mut neighbors = vec![];
    for y in p.y-1 ..= p.y+1 {
        for x in p.x-1 ..= p.x+1 {
            neighbors.push(Point{x,y});
        }
    }
    neighbors
}

fn build_enhancement_algorithm(data: &str) -> HashSet<u64> {
    let mut result = HashSet::new();

    for (i,c) in data.chars().enumerate() {
        if c == '#' { result.insert(i as u64); }
    }

    result
}