use std::collections::HashMap;

pub fn part1(input: String) {
    let openers: [char; 4] = ['(', '[', '{', '<'];
    let mappings: HashMap<char, char> = [('(', ')'), ('[', ']'), ('{', '}'), ('<', '>')].iter().cloned().collect();
    let costs: HashMap<char, u32> = [(')',3), (']',57), ('}',1197), ('>',25137)].iter().cloned().collect();
    
    let mut result = 0;

    for line in input.lines() {
        let mut stack = vec![];

        for c in line.chars() {
            if openers.contains(&c) {
                stack.push(mappings[&c]);
            }
            else if let Some(stored) = stack.pop() {
                if c != stored {
                    // Corrupt
                    result += costs[&c];
                    break;
                }
            }
        }
    }
    println!("Part 1: {}", result);
}

pub fn part2(input: String) {
    let openers: [char; 4] = ['(', '[', '{', '<'];
    let mappings: HashMap<char, char> = [('(', ')'), ('[', ']'), ('{', '}'), ('<', '>')].iter().cloned().collect();
    let scores: HashMap<char, u64> = [(')',1), (']',2), ('}',3), ('>',4)].iter().cloned().collect();
    
    let mut completion_scores = vec![];

    for line in input.lines() {
        let mut stack = vec![];
        let mut is_corrupt = false;

        for c in line.chars() {
            if openers.contains(&c) {
                stack.push(mappings[&c]);
            }
            else if let Some(stored) = stack.pop() {
                if c != stored {
                    is_corrupt = true;
                    break;
                }
            }
        }
        
        if !is_corrupt {
            let mut line_score = 0;
            while let Some(stored) = stack.pop() {
                line_score = 5 * line_score + scores[&stored];
            }

            completion_scores.push(line_score);
        }
    }

    // Get the middle value from completion_scores
    completion_scores.sort();
    let mid_val = completion_scores.len() / 2;
    let result = completion_scores[mid_val];

    println!("Part 2: {}", result);
}