use std::collections::HashMap;
use serde_scan::scan;

pub fn part1(input: String) {
    let data = parse_data(input);

    println!("{:?}", data);
}

pub fn part2(_input: String) {

}

fn parse_data(input: String) -> HashMap<u8, Vec<(i32,i32,i32)>> {
    let mut data = HashMap::new();

    let mut scanner_id = 0;
    for line in input.lines() {
        if line.len() == 0 { continue; }

        let parsed: Result<u8, _> = scan!("--- scanner {} ---" <- line);
        if let Ok(pe) = parsed {
            scanner_id = pe;
            continue;
        }

        let parsed: (i32,i32,i32) = scan!("{},{},{}" <- line).expect("Could not parse co-ordinates");
        data.entry(scanner_id).or_insert(vec![]).push(parsed);
    }
    data
}