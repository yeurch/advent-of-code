pub fn part1(input: String) {
    let module_weights: Vec<u32> = input.lines().map(|x| x.parse().unwrap()).collect();

    let fuel_weight: u32 = module_weights.into_iter().map(|x| x / 3 - 2).sum();
    println!("Part 1: {}", fuel_weight);
}

pub fn part2(input: String) {
    let module_weights: Vec<u32> = input.lines().map(|x| x.parse().unwrap()).collect();

    let fuel_weight: u32 = module_weights.into_iter().map(|x| fuel_weight2(x)).sum();
    println!("Part 2: {}", fuel_weight);
}

fn fuel_weight2(module_weight: u32) -> u32 {
    let mut result = 0_u32;
    let mut w = module_weight as i32;
    loop {
        let fuel = w / 3 - 2;
        if fuel <= 0 {
            break;
        }
        result += fuel as u32;
        w = fuel;
    }
    result
}