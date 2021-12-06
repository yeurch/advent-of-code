pub fn part1(input: String) {
    do_impl(input, 80);
}

pub fn part2(input: String) {
    do_impl(input, 256)
}

fn do_impl(input: String, days: u16) {
    let pop_in: Vec<usize> = input.split(",").map(|s| s.parse().unwrap()).collect();

    // Build initial array
    let mut pop = vec![0; 9];
    for i in pop_in {
        pop[i] += 1;
    }

    for _ in 0..days {
        tick(&mut pop);
    }

    println!("Part 1: {}", pop.iter().sum::<u64>());
}

fn tick(pop: &mut Vec<u64>) {
    let tmp = pop[0];
    pop[0] = pop[1];
    pop[1] = pop[2];
    pop[2] = pop[3];
    pop[3] = pop[4];
    pop[4] = pop[5];
    pop[5] = pop[6];
    pop[6] = pop[7] + tmp;
    pop[7] = pop[8];
    pop[8] = tmp;
}