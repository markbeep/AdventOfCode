use std::fs::read_to_string;

use crate::Solution;

pub fn solve() -> Solution {
    let f = read_to_string("input/09.txt").unwrap();
    let digits = f
        .trim()
        .chars()
        .map(|x| x.to_digit(10).unwrap())
        .collect::<Vec<u32>>();
    let mut mem = vec![];
    for (i, d) in digits.iter().enumerate() {
        if i % 2 == 1 {
            for _ in 0..*d {
                mem.push(None);
            }
        } else {
            let x = i / 2;
            for _ in 0..*d {
                mem.push(Some(x));
            }
        }
    }
    let mut free_space = digits[0] as usize;
    while free_space < mem.len() {
        while mem.last().unwrap().is_none() {
            mem.pop();
        }
        let last = mem.pop().unwrap();
        mem[free_space] = last;
        while mem[free_space].is_some() {
            free_space += 1;
            if free_space >= mem.len() {
                break;
            }
        }
    }
    let p1 = mem
        .iter()
        .map(|x| x.unwrap_or(0))
        .enumerate()
        .fold(0, |old, (i, x)| old + i * x);
    (p1.to_string(), "".to_string())
}
