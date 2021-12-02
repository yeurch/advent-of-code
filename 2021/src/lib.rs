// Days
pub mod day01;
pub mod day02;

pub fn noop(_inp: String) {}

pub type DayFn = fn(String);

pub fn get_day(day: u32) -> (DayFn, DayFn) {
    return match day {
        1 => (day01::part1, day01::part2),
        2 => (day02::part1, day02::part2),
        _ => {
            println!("Unknown day: {}", day);
            return (noop, noop);
        }
    };
}