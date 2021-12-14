use std::collections::HashMap;
use serde_scan::scan;

pub fn part1(input: String) {
    let polymer = input.lines().nth(0).unwrap();

    let mut mappings = HashMap::new();
    for line in input.lines().skip(2) {
        let parsed: (&str, char) = scan!("{} -> {}" <- line).unwrap();
        mappings.insert(parsed.0, parsed.1);
    }

    let mut counts = HashMap::new();
    for i in 0..polymer.len()-1 {
        *counts.entry(String::from(&polymer[i..=i+1])).or_insert(0) += 1;
    }

    counts = do_impl(counts, &mappings, 10);
    let last = polymer.chars().last().unwrap();
    println!("Part 1: {}", calc_result(last, counts));
}

pub fn part2(input: String) {
    let polymer = input.lines().nth(0).unwrap();

    let mut mappings = HashMap::new();
    for line in input.lines().skip(2) {
        let parsed: (&str, char) = scan!("{} -> {}" <- line).unwrap();
        mappings.insert(parsed.0, parsed.1);
    }

    let mut counts = HashMap::new();
    for i in 0..polymer.len()-1 {
        *counts.entry(String::from(&polymer[i..=i+1])).or_insert(0) += 1;
    }

    counts = do_impl(counts, &mappings, 40);
    let last = polymer.chars().last().unwrap();
    println!("Part 1: {}", calc_result(last, counts));
}

fn do_impl(counts: HashMap<String, u64>, mappings: &HashMap<&str, char>, num_rounds: usize) -> HashMap<String, u64> {
    let mut tmp_counts = counts;
    for _ in 0..num_rounds {
        let mut new_counts = HashMap::new();
        for (k,v) in tmp_counts.iter() {
            let inserted_char = mappings[&k as &str];
            *new_counts.entry(k.chars().nth(0).unwrap().to_string() + &inserted_char.to_string()).or_insert(0) += v;
            *new_counts.entry(inserted_char.to_string() + &k.chars().nth(1).unwrap().to_string()).or_insert(0) += v;
        }
        tmp_counts = new_counts;
    }

    tmp_counts
}

fn calc_result(last: char, counts: HashMap<String, u64>) -> u64 {
    let mut char_freq = HashMap::new();

    for (k,v) in counts {
        let c = k.chars().nth(0).unwrap();
        *char_freq.entry(c).or_insert(0) += v;
    }
    *char_freq.entry(last).or_insert(0) += 1;

    let mut max_val = 0;
    let mut min_val = u64::MAX;
    for v in char_freq.values() {
        if *v > max_val { max_val = *v; }
        if *v < min_val { min_val = *v; }
    }
    max_val - min_val
}