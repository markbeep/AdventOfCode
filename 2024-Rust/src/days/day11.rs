use std::{collections::HashMap, fs::read_to_string};

use crate::Solution;

pub fn solve() -> Solution {
    let f = read_to_string("input/11.txt").unwrap();
    let nums: Vec<u64> = f
        .trim()
        .split(" ")
        .map(|x| x.parse::<u64>().unwrap())
        .collect();

    let mut counts: HashMap<u64, u64> = HashMap::new();
    nums.iter().for_each(|x| {
        counts.insert(*x, 1);
    });
    let mut p1 = 0;
    let mut p2 = 0;
    for i in 0..75 {
        let mut cpy = HashMap::new();
        for (k, v) in counts.iter() {
            let n = *k;
            if n == 0 {
                match cpy.get_mut(&1) {
                    Some(x) => *x += *v,
                    None => drop(cpy.insert(1, *v)),
                }
            } else if n.ilog10() % 2 == 1 {
                let log = 10u64.pow(n.ilog10() / 2 + 1);
                let left = n / log;
                let right = n % log;
                match cpy.get_mut(&left) {
                    Some(x) => *x += *v,
                    None => drop(cpy.insert(left, *v)),
                }
                match cpy.get_mut(&right) {
                    Some(x) => *x += *v,
                    None => drop(cpy.insert(right, *v)),
                }
            } else {
                match cpy.get_mut(&(n * 2024)) {
                    Some(x) => *x += *v,
                    None => drop(cpy.insert(n * 2024, *v)),
                }
            }
        }
        counts = cpy;
        if i == 24 {
            p1 = counts.values().sum();
        }
        if i == 74 {
            p2 = counts.values().sum();
        }
    }

    (p1.to_string(), p2.to_string())
}
