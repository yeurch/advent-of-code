use serde_scan::scan;

pub fn part1(input: String) {
    let mut player_positions: [u32; 2] = [0,0];
    for line in input.lines() {
        let (player_num, start_pos): (usize, u32) = scan!("Player {} starting position: {}" <- line).unwrap();
        player_positions[player_num-1] = start_pos - 1; // zero indexed
    }

    let mut d100 = die(100);

    let mut cur_player = 0;
    let mut scores: [u32; 2] = [0,0];
    let mut turn_counter = 0;

    loop {
        turn_counter += 1;
        let r = roll(3, &mut d100);
        let new_pos = (player_positions[cur_player] + r) % 10;
        player_positions[cur_player] = new_pos;
        scores[cur_player] += new_pos + 1;
        if scores[cur_player] >= 1000 { break; }
        cur_player = (cur_player + 1) % 2;
    }

    let losing_score = scores.into_iter().min().unwrap();
    let result = losing_score * turn_counter * 3;
    println!("Part 1: {}", result);
}

type Die = dyn std::iter::Iterator<Item = u32>;

fn die(n: u32) -> impl std::iter::Iterator<Item = u32> {
    std::iter::successors(
        Some(1),
        move |&num| {
            if num + 1 <= n {
                Some(num + 1)
            } else {
                Some(1)
            }
        },
    )
}

fn roll(n: u8, die: &mut Die) -> u32 {
    let mut result = 0_u32;
    for _ in 0..n {
        result += die.next().unwrap();
    }
    result
}