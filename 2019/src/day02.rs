use crate::intcode::*;

pub fn part1(input: String) {
    let program: Vec<i64> = input.split(",").map(|x| x.parse().unwrap()).collect();
    let mut cpu = IntCode::new(program);

    cpu.set_mem(1, 12);
    cpu.set_mem(2, 2);

    while *&cpu.tick() { }

    println!("Part 1: {}", cpu.get_mem(0));
}

pub fn part2(input: String) {
    let program: Vec<i64> = input.split(",").map(|x| x.parse().unwrap()).collect();

    for noun in 0..100 {
        for verb in 0..100 {
            let mut cpu = IntCode::new(program.clone());

            cpu.set_mem(1, noun);
            cpu.set_mem(2, verb);
        
            while *&cpu.tick() { }
        
            if cpu.get_mem(0).clone() == 19690720 {
                let result = 100 * noun + verb;
                println!("Part 2: {}", result);
                return;
            }
        }
    }
}