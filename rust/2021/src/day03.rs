pub fn part1(input: String) {
    let samples = input.lines().collect::<Vec<&str>>();
    let num_samples = samples.len();

    let bit_width = samples[0].len();
    let mut counters = vec![0; bit_width];
    for s in &samples {
        for (i,c) in s.chars().enumerate() {
            if c == '1' {
                counters[i] += 1;
            }
        }
    }

    let mut gamma_rate = 0;
    let mut epsilon_rate = 0;
    let base: u32 = 2;
    for i in 0..bit_width {
        
        let bit_value = base.pow((bit_width-i-1) as u32);
        if counters[i] > num_samples/2 {
            gamma_rate += bit_value;
        }
        else {
            epsilon_rate += bit_value;
        }
    }

    println!("Part 1 solution: {}", gamma_rate*epsilon_rate);
}

pub fn part2(input: String) {
    let samples = input.lines().collect::<Vec<&str>>();

    let o2_rating = get_o2_or_co2(&samples, true);
    let co2_rating = get_o2_or_co2(&samples, false);

    println!("Part 2 solution: {}", o2_rating * co2_rating);
}

fn get_o2_or_co2(samples: &Vec<&str>, is_o2: bool) -> u32 {
    let bit_width = samples[0].len();

    let mut filtered_samples = (*samples).clone();
    for i in 0..bit_width {
        let num_1s = filtered_samples.iter().filter(|s| {s.as_bytes()[i] as char == '1'}).count();
        let num_0s = filtered_samples.len() - num_1s;
        let val_to_retain: char;
        if num_0s == num_1s {
            val_to_retain = if is_o2 {'1'} else {'0'};
        }
        else {
            let most_common = if num_1s > num_0s {1} else {0};
            val_to_retain = char::from_digit(if is_o2 {most_common} else { 1 - most_common}, 10).unwrap();
        }
        filtered_samples = filtered_samples.into_iter().filter(|s| {s.as_bytes()[i] as char == val_to_retain}).collect();
        if filtered_samples.len() < 2 { break; }
    }
    let matching_sample = filtered_samples[0];
    let base: u32 = 2;
    let mut result = 0;
    for i in 0..bit_width {
        if matching_sample.as_bytes()[i] as char == '1' {
            let bit_value = base.pow((bit_width-i-1) as u32);
            result += bit_value;
        }
    }

    result
}