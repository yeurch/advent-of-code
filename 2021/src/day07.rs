use std::collections::HashMap;

pub fn part1(input: String) {
    let values: Vec<i32> = input.split(",").map(|s| s.parse().unwrap()).collect();

    let min = values.iter().min().unwrap();
    let max = values.iter().max().unwrap();

    let mut min_cost = i32::MAX;
    for i in *min..=*max {
        let cost = get_cost(&values, i);
        if cost < min_cost {
            min_cost = cost;
        }
    }

    println!("Part 1: {}", min_cost);
}

pub fn part2(input: String) {
    let values: Vec<i32> = input.split(",").map(|s| s.parse().unwrap()).collect();

    let min = values.iter().min().unwrap();
    let max = values.iter().max().unwrap();

    let mut cache = HashMap::new();
    let mut min_cost = i32::MAX;
    for i in *min..=*max {
        let cost = get_cost2(&mut cache, &values, i);
        if cost < min_cost {
            min_cost = cost;
        }
    }

    println!("Part 2: {}", min_cost);
}

fn get_cost(values: &Vec<i32>, p: i32) -> i32 {
    let mut total_cost = 0;
    for x in values.iter() {
        total_cost += (x - p).abs();
    }
    total_cost
}

fn get_cost2(cache: &mut HashMap<i32, i32>, values: &Vec<i32>, p: i32) -> i32 {
    let mut total_cost = 0;
    for x in values.iter() {
        let dist = (x - p).abs();

        let cost;
        if cache.contains_key(&dist) {
            cost = *cache.get(&dist).unwrap();
        }
        else {
            cost = (1 ..= dist).fold(0, |a, b| a+b);
            cache.insert(dist, cost);
        }
        total_cost += cost;
    }
    total_cost
}