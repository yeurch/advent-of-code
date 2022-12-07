use std::str::FromStr;

pub fn part1(input: String) {
    let commands = input.lines().map(|s| {s.parse().unwrap()} );

    let mut voyage = Voyage::new();
    for c in commands {
        voyage.do_move1(&c);
    }

    println!("Final pos: x={}, depth={}, result={}", voyage.h_distance, voyage.depth, voyage.h_distance * voyage.depth);
}

pub fn part2(input: String) {
    let commands = input.lines().map(|s| {s.parse().unwrap()} );

    let mut voyage = Voyage::new();
    for c in commands {
        voyage.do_move2(&c);
    }

    println!("Final pos: x={}, depth={}, result={}", voyage.h_distance, voyage.depth, voyage.h_distance * voyage.depth);
}

enum Command {
    Forward(u32),
    Down(u32),
    Up(u32)
}

impl FromStr for Command {
    type Err = ();
    fn from_str(input: &str) -> Result<Command, Self::Err> {
        let parts = input.split_whitespace().collect::<Vec<&str>>();
        let dist = parts[1].parse().unwrap();
        match parts[0] {
            "forward" => Ok(Command::Forward(dist)),
            "down" => Ok(Command::Down(dist)),
            "up" => Ok(Command::Up(dist)),
            _ => Err(()),
        }
    }
}

struct Voyage {
    depth: u32,
    h_distance: u32,
    aim: u32
}

impl Voyage {
    fn new() -> Voyage {
        Voyage {
            depth: 0,
            h_distance: 0,
            aim: 0
        }
    }

    fn do_move1(&mut self, cmd: &Command) {
        match cmd {
            Command::Forward(n) => self.h_distance += n,
            Command::Down(n) => self.depth += n,
            Command::Up(n) => self.depth -= n
        }
    }

    fn do_move2(&mut self, cmd: &Command) {
        match cmd {
            Command::Forward(n) => {
                self.h_distance += n;
                self.depth += n * self.aim;
            }
            Command::Down(n) => self.aim += n,
            Command::Up(n) => self.aim -= n
        }
    }
}