use crate::intcode::*;

pub fn part1(input: String) {
    let program: Vec<i64> = input.split(",").map(|x| x.parse().unwrap()).collect();
    let mut cpu = IntCode::new(program);

    while *&cpu.tick() { }


}

pub fn part2(_input: String) {

}