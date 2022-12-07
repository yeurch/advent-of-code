pub fn part1(input: String) {
    let data: Vec<&str> = input.lines().collect();

    let mut snailfish = String::from(data[0]);
    for line in data.iter().skip(1) {
        snailfish = format!("[{},{}]", snailfish, line);
        //println!("After addition  : {}", snailfish);
        snailfish = reduce(snailfish);
    }

    println!("{}", calc_magnitude(&snailfish));
}

pub fn part2(input: String) {
    let data: Vec<&str> = input.lines().collect();

    let mut result = 0;

    for i in 0..data.len() {
        for j in 0..data.len() {
            if i != j {
                let mut snailfish = format!("[{},{}]", data[i], data[j]);
                snailfish = reduce(snailfish);
                let magnitude = calc_magnitude(&snailfish);
                if magnitude > result { result = magnitude; }
            }
        }
    }

    println!("{}", result);
}

fn reduce(snailfish: String) -> String {
    let mut result = snailfish;
    loop {
        let mut depth = 0;
        let mut explode_index = 0;
        for (i,c) in result.chars().enumerate() {
            if c == '[' { depth += 1; }
            if c == ']' { depth -= 1; }
            if depth == 5 {
                explode_index = i;
                break;
            }
        }
        if explode_index > 0 {
            result = explode(result, explode_index);
            continue;
        }

        let numbers = scan_numbers(&result);
        let mut did_split = false;
        for (val, start, len) in numbers {
            if val > 9 {
                result = split(result, val, start, len);
                did_split = true;
                break;
            }
        }
        if did_split { continue; }

        // If we haven't already continued, then there was nothing for us to do, so break
        break;
    }
    result
}

fn explode(snailfish: String, idx: usize) -> String {
    let mut result = String::with_capacity(snailfish.len());
    let numbers = scan_numbers(&snailfish);
    for (i,(val, start, _len)) in numbers.iter().enumerate() {
        if *start >= idx {
            if i > 0 {
                let (last_val, last_start, last_len) = numbers[i-1];
                let new_val = val + last_val;
                result.push_str(&snailfish[..last_start]);
                result.push_str(&format!("{}", new_val));
                result.push_str(&snailfish[(last_start+last_len)..*start-1]);
            }
            else {
                result.push_str(&snailfish[..*start-1]);
            }
            
            result.push_str("0");

            let (right_val, right_start, right_len) = numbers[i+1];
            if i < numbers.len() - 2 {
                let (next_val, next_start, next_len) = numbers[i+2];
                let new_val = right_val + next_val;
                result.push_str(&snailfish[(right_start+right_len+1)..next_start]);
                result.push_str(&format!("{}", new_val));
                result.push_str(&snailfish[next_start+next_len..]);
            }
            else {
                result.push_str(&snailfish[right_start+right_len+1..]);
            }

            break;
        }
    }
    //println!("After explosion : {}", result);
    result
}

fn split(snailfish: String, val: u32, start: usize, len: usize) -> String {
    let left = val / 2;
    let right = val - left;

    let mut result = String::with_capacity(snailfish.len() + 3); // "nn" -> "[n,n]"
    result.push_str(&snailfish[..start]);
    result.push_str(&format!("[{},{}]", left, right));
    result.push_str(&snailfish[start+len..]);

    //println!("After split     : {}", result);
    result
}

fn calc_magnitude(snailfish:&str) -> u64 {
    let mut result = 0;
    let numbers = scan_numbers(snailfish);
    let (val, start, len) = numbers[0];
    let mut sep_pos = 0;
    if start == 1 {
        // Left hand item is a value
        result = 3 * val as u64;
        sep_pos = start+len;
    }
    else {
        let mut depth = 1;
        for (i,c) in snailfish.chars().skip(2).enumerate() {
            if c == '[' { depth += 1; }
            if c == ']' { depth -= 1; }
            if depth == 0 {
                sep_pos = i + 3;
                result += 3 * calc_magnitude(&snailfish[1..sep_pos]);
                break;
            }
        }
    }

    for (val, start, _len) in numbers {
        if start > sep_pos {
            if start == sep_pos + 1 {
                // Right hand item is a value
                result += 2 * val as u64;
            }
            else {
                result += 2 * calc_magnitude(&snailfish[sep_pos+1..snailfish.len()-1])
            }
            break;
        }
    }
    result
}

fn scan_numbers(snailfish: &str) -> Vec<(u32, usize, usize)> { // value, start, length
    const RADIX: u32 = 10;
    let mut result = vec![];

    let mut scanning_digits = false;
    let mut value = 0;
    let mut start = 0;
    for (i,c) in snailfish.chars().enumerate() {
        if let Some(d) = c.to_digit(RADIX) {
            if !scanning_digits { start = i; }
            value *= 10;
            value += d;
            scanning_digits = true;
        }
        else {
            if scanning_digits {
                result.push((value, start, i-start));
            }
            scanning_digits = false;
            value = 0;
        }
    }
    result
}
