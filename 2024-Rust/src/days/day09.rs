use std::{collections::HashSet, fs::read_to_string};

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
    let free_space = digits[0] as usize;
    let p1 = p1(free_space, mem.clone());
    let p2 = p2(free_space, mem);
    (p1.to_string(), p2.to_string())
}

fn find_space(mem: &Vec<Option<usize>>, size: usize) -> Option<usize> {
    let mut c = 0;
    for (i, d) in mem.iter().enumerate() {
        if d.is_none() {
            c += 1;
            if c >= size {
                return Some(i - c + 1);
            }
        } else {
            c = 0;
        }
    }
    None
}

fn p1(mut free_space: usize, mut mem: Vec<Option<usize>>) -> usize {
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
    mem.iter()
        .map(|x| x.unwrap_or(0))
        .enumerate()
        .fold(0, |old, (i, x)| old + i * x)
}

fn p2(mut free_space: usize, mut mem: Vec<Option<usize>>) -> usize {
    let mut ignore = HashSet::new();
    while free_space < mem.len() {
        while mem.last().unwrap().is_none() {
            mem.pop();
        }

        let mut len = 0;
        let mut last = None;
        for d in mem.iter().rev() {
            if d.is_some() && !ignore.contains(&d.unwrap()) {
                if last.is_none() || last.is_some() && d.unwrap() == last.unwrap() {
                    last = *d;
                    len += 1;
                } else if last.is_some() && d.unwrap() != last.unwrap() {
                    break;
                }
            }
        }
        if last.is_none() {
            break;
        }
        ignore.insert(last.unwrap());
        let space = find_space(&mem, len);
        if space.is_some() {
            let old_pos = mem.iter().position(|x| *x == last).unwrap();
            if space.unwrap() < old_pos {
                for i in 0..len {
                    mem[old_pos + i] = None;
                }
                for i in 0..len {
                    mem[space.unwrap() + i] = last;
                }
            }
        }

        while mem[free_space].is_some() {
            free_space += 1;
            if free_space >= mem.len() {
                break;
            }
        }
    }
    mem.iter()
        .map(|x| x.unwrap_or(0))
        .enumerate()
        .fold(0, |old, (i, x)| old + i * x)
}
