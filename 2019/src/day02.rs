use crate::intcode::*;

pub fn part1(input: String) {
    let program: Vec<i64> = input.split(",").map(|x| x.parse().unwrap()).collect();
    let mut cpu = IntCode::new(program);

    cpu.set_mem(1, 12);
    cpu.set_mem(2, 2);

    while *&cpu.tick() { }

    println!("Part 1: {}", cpu.get_mem(0));
}

pub fn part2(_input: String) {

}