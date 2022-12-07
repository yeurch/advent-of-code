pub fn part1(input: String) {
    let values: Vec<u64> = input.lines().map(|s| {s.parse().unwrap()} ).collect();

    let mut prev = u64::MAX;
    let mut result = 0;

    for v in values {
        if v > prev {
            result += 1;
        }
        prev = v;
    }

    println!("Part 1 result is {}.", result);
}

pub fn part2(input: String) {
    let values: Vec<u64> = input.lines().map(|s| {s.parse().unwrap()} ).collect();

    let mut prev = u64::MAX;
    let mut result = 0;

    for start in 0 .. values.len() - 2 {
        let s: u64 = (&values[start..start+3]).iter().sum();
        if s > prev {
            result += 1;
        }
        prev = s;
    }

    println!("Part 2 result is {}.", result);
}