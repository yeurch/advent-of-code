use std::collections::HashSet;

pub fn part1(input: String) {
    let mut lines = input.lines();

    let called_nums: Vec<i32> = (&mut lines).next().unwrap().split(",").map(|s| {s.parse().unwrap()}).collect();

    let mut earliest_win = usize::MAX;
    let mut winning_score: i32 = -1;
    loop {
        let game_data: Vec<&str> = (&mut lines).skip(1).take(5).collect();
        if game_data.len() < 5 {
            break;
        }
        let bingo = Bingo::new(game_data);

        let result = bingo.play(&called_nums);
        if result.winning_move_num < earliest_win {
            earliest_win = result.winning_move_num;
            winning_score = result.score;
        }
    }

    println!("Part 1 result is {}", winning_score);

}

pub fn part2(input: String) {
    let mut lines = input.lines();

    let called_nums: Vec<i32> = (&mut lines).next().unwrap().split(",").map(|s| {s.parse().unwrap()}).collect();

    let mut latest_win: usize = 0;
    let mut winning_score: i32 = -1;
    loop {
        let game_data: Vec<&str> = (&mut lines).skip(1).take(5).collect();
        if game_data.len() < 5 {
            break;
        }
        let bingo = Bingo::new(game_data);

        let result = bingo.play(&called_nums);
        if result.winning_move_num > latest_win {
            latest_win = result.winning_move_num;
            winning_score = result.score;
        }
    }

    println!("Part 2 result is {}", winning_score);

}

struct Bingo {
    lines: Vec<HashSet<i32>>,
    size: usize
}

struct BingoResult {
    winning_move_num: usize,
    score: i32
}

impl Bingo {
    fn new(board_data: Vec<&str>) -> Bingo {
        let mut lines = vec!();
        let mut v = vec!();

        for line in board_data {
            let nums: Vec<i32> = line.split_ascii_whitespace().map(|s| {s.parse().unwrap()}).collect();
            v.push(nums.clone());
            lines.push(HashSet::from_iter(nums));
        }
        let size = lines.len();

        let mut vertical = vec!();
        for i in 0..size {
            let vertical_line: HashSet<i32> = v.iter().map(|l| { l[i] }).collect();
            vertical.push(vertical_line);
        }
        lines.append(&mut vertical);

        Bingo {
            lines,
            size
        }
    }

    fn play(&self, called_numbers: &Vec<i32>) -> BingoResult {
        for turn in 0..called_numbers.len() {
            let mut called_so_far = HashSet::new();
            for n in called_numbers.iter().take(turn+1) {
                called_so_far.insert(n.clone());
            }
            for i in self.lines.iter() {
                if called_so_far.intersection(i).count() == self.size {
                    let last_called = called_numbers[turn];
                    let mut score = 0;

                    for line in self.lines.iter().take(self.size) {
                        for val in line {
                            if !called_so_far.contains(val) {
                                score += val;
                            }
                        }
                    }

                    score *= last_called;
                    return BingoResult {
                        winning_move_num: turn,
                        score
                    };
                }
            }
        }
        BingoResult {
            winning_move_num: usize::MAX,
            score: -1
        }
    }
}
