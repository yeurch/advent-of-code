use std::collections::HashSet;
use std::collections::HashMap;
use itertools::Itertools; // Trait needed for sorted()

pub fn part1(input: String) {
    let values: Vec<&str> = input
        .lines()
        .map(|x| x.split(" | ").last().unwrap())
        .flat_map(|x| x.split_whitespace())
        .collect();

    let valid_lens: Vec<usize> = vec![2, 3, 4, 7];

    let valid_values = values.into_iter().filter(|x| valid_lens.contains(&x.len())).count();
    
    println!("Part 1: {}", valid_values);
}

pub fn part2(input: String) {
    let mut total = 0;

    for line in input.lines() {
        let mut s1 = line.split(" | ");
        let all_vals: Vec<String> = s1.next().unwrap().split_whitespace().map(|x| {
            x.chars().sorted().collect::<String>()
        }).collect();
        let targets: Vec<String> = s1.next().unwrap().split_whitespace().map(|x| {
            x.chars().sorted().collect::<String>()
        }).rev().collect();

        let mappings = decipher(all_vals);

        let mut multiplier = 1;
        for target in targets {
            total += mappings.get(&target).unwrap() * multiplier;
            multiplier *= 10;
        }
    }

    println!("Part 2: {}", total);    
}

fn decipher(all_vals: Vec<String>) -> HashMap<String, u32> {
    let mut mappings: HashMap<String, u32> = HashMap::new();

    // Start by figuring out the digits with unique numbers of segments
    let one = all_vals.iter().find(|x| x.len() == 2).unwrap();
    let four = all_vals.iter().find(|x| x.len() == 4).unwrap();
    let seven = all_vals.iter().find(|x| x.len() == 3).unwrap();
    let eight = all_vals.iter().find(|x| x.len() == 7).unwrap();

    let set_1 = one  .chars().collect::<HashSet<char>>();
    let set_4 = four .chars().collect::<HashSet<char>>();
    let set_8 = eight.chars().collect::<HashSet<char>>();

    let seg_f = all_vals.iter()
        .filter(|x| x.len() == 6)
        .map(|x| x.chars().collect::<HashSet<char>>())
        .filter(|x| x.intersection(&set_1).count() == 1)
        .next().unwrap()
        .intersection(&set_1).next().unwrap().clone();
    let seg_c = one.replace(seg_f, "").chars().into_iter().next().unwrap();

    let set_9 = all_vals.iter()
        .filter(|x| x.len() == 6)
        .map(|x| x.chars().collect::<HashSet<char>>())
        .filter(|x| x.intersection(&set_4).count() == 4)
        .next().unwrap();
    let nine = set_9.iter().sorted().join("");
    let seg_e = (&set_8 - &set_9).into_iter().next().unwrap();

    let six = all_vals.iter()
        .filter(|x| x.len() == 6 && x.replace(seg_c, "").len() == 6)
        .next().unwrap().clone();

    let zero = all_vals.iter()
        .find(|&x| x.len() == 6 && !x.eq(&six) && !x.eq(&nine)).unwrap().clone();

    let two = all_vals.iter()
        .filter(|x| x.len() == 5 && x.replace(seg_e, "").len() == 4)
        .next().unwrap().clone();

    let five = all_vals.iter()
        .find(|&x| x.len() == 5 && x.replace(seg_c, "").len() == 5 && !x.eq(&two)).unwrap().clone();

    let three = all_vals.iter()
        .find(|&x| x.len() == 5 && !x.eq(&two) && !x.eq(&five)).unwrap().clone();

    let six = eight.replace(seg_c, "");

    mappings.insert(zero, 0);
    mappings.insert(one.to_string(), 1);
    mappings.insert(two, 2);
    mappings.insert(three, 3);
    mappings.insert(four.to_string(), 4);
    mappings.insert(five, 5);
    mappings.insert(six, 6);
    mappings.insert(seven.to_string(), 7);
    mappings.insert(eight.to_string(), 8);
    mappings.insert(nine, 9);

    mappings
}