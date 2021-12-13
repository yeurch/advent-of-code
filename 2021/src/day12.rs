use std::collections::HashSet;
use std::collections::HashMap;

pub fn part1(input: String) {
    let graph = construct_graph(&input);

    let mut visited = vec![];

    let mut result = 0;
    visit("start", &graph, &mut visited, &mut result, false);
    println!("Part 1: {}", result);
}

pub fn part2(input: String) {
    let graph = construct_graph(&input);

    let mut visited = vec![];

    let mut result = 0;
    visit("start", &graph, &mut visited, &mut result, true);
    println!("Part 2: {}", result);
}

fn construct_graph<'a>(input: &'a String) -> HashMap<String, (bool, HashSet<&'a str>)> {
    let mut graph: HashMap<String, (bool, HashSet<&str>)> = HashMap::new();
    for line in input.lines() {
        let parts: Vec<&str> = line.split("-").collect();
        for part in parts.iter() {
            if !graph.contains_key(*part) {
                let is_small = part.chars().nth(0).unwrap().is_lowercase();
                graph.insert(String::from(*part), (is_small, HashSet::new()));
            }
        }
        graph.get_mut(parts[0]).unwrap().1.insert(parts[1].clone());
        graph.get_mut(parts[1]).unwrap().1.insert(parts[0].clone());
    }

    graph
}

fn visit<'a>(node_name: &'a str, graph: &HashMap<String, (bool, HashSet<&'a str>)>, visited: &mut Vec<&'a str>, num_paths: &mut u32, part2: bool) {
    if node_name == "end" {
        *num_paths += 1;
        return;
    }

    if node_name == "start" && visited.len() > 0 {
        // We can't go back to the start. Early exit.
        return;
    }
    
    let node = &graph[node_name];
    if node.0 && visited.contains(&node_name) && (!part2 || is_small_cave_visited_twice(visited)) {
        // This is a small node and we've already visisted it the maximum number of permitted times. That's not permitted, so early exit.
        return;
    }

    visited.push(node_name);

    let neighbors = &node.1;
    for neighbor in neighbors {
        visit(neighbor, &graph, visited, num_paths, part2);
    }    

    visited.pop();
}

fn is_small_cave_visited_twice(visited: &Vec<&str>) -> bool {
    let visited_small: Vec<&&str> = visited.iter().filter(|v| v.chars().nth(0).unwrap().is_lowercase()).collect();
    let visited_small_hash: HashSet<&&&str> = HashSet::from_iter(visited_small.iter().clone());

    visited_small.len() != visited_small_hash.len()
}