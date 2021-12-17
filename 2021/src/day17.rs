use serde_scan::scan;

pub fn part1(input: String) {
    let s: &str = &input[..];
    let (min_x,max_x,min_y,max_y): (u32,u32,i32,i32) = scan!("target area: x={}..{}, y={}..{}" <- s).expect("Could not parse input");

    // As trench is always below submarine (i.e. negative) then at some point before reaching the target, then y = 0 (parabolic arc).
    // The upper bound of y can be found by finding which values of initial vy would result in overshooting the bottom of the target in one step.

    let max_vy0 = min_y.abs() - 1;
    let min_vy0 = min_y;

    let x0_vals = possible_x0_vals(min_x, max_x);

    for vy0 in (min_vy0..=max_vy0).rev() {
        let mut hit = false;
        for vx0 in &x0_vals {
            hit = false;
            let mut px = 0;
            let mut py = 0;
            let mut vx = *vx0;
            let mut vy = vy0;
            while px <= max_x && py >= min_y {
                px = px + vx;
                py = py + vy;
                vx = if vx == 0 {0} else {vx - 1};
                vy -= 1;
                if px >= min_x && px <= max_x && py >= min_y && py <= max_y {
                    hit = true;
                    break;
                }
                if vx == 0 && py < min_y {
                    break;
                }
            }
            if hit { break; }
        }
        if hit {
            let max_height = vy0*(vy0+1)/2;
            println!("Part 1: {}", max_height);
            return;
        }
    }
    println!("Part 1: no valid solution found");
}

pub fn part2(input: String) {
    let s: &str = &input[..];
    let (min_x,max_x,min_y,max_y): (u32,u32,i32,i32) = scan!("target area: x={}..{}, y={}..{}" <- s).expect("Could not parse input");

    // As trench is always below submarine (i.e. negative) then at some point before reaching the target, then y = 0 (parabolic arc).
    // The upper bound of y can be found by finding which values of initial vy would result in overshooting the bottom of the target in one step.

    let max_vy0 = min_y.abs() - 1;
    let min_vy0 = min_y;

    let x0_vals = possible_x0_vals(min_x, max_x);

    let mut result = 0;

    for vy0 in (min_vy0..=max_vy0).rev() {
        for vx0 in &x0_vals {
            let mut px = 0;
            let mut py = 0;
            let mut vx = *vx0;
            let mut vy = vy0;
            while px <= max_x && py >= min_y {
                px = px + vx;
                py = py + vy;
                vx = if vx == 0 {0} else {vx - 1};
                vy -= 1;
                if px >= min_x && px <= max_x && py >= min_y && py <= max_y {
                    result += 1;
                    break;
                }
                if vx == 0 && py < min_y {
                    break;
                }
            }
        }
    }
    println!("Part 2: {}", result);
}

fn possible_x0_vals(min_x: u32, max_x: u32) -> Vec<u32> {
    let mut result = vec![];
    for x in 0..=max_x {
        let mut xpos = 0;
        let mut vx = x;
        for _ in 0..=x {
            xpos += vx;
            vx -=1;
            if xpos >= min_x && xpos <= max_x {
                result.push(x);
                break;
            }
        }
    }

    result
}