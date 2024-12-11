use std::{collections::HashMap, fs::read_to_string};

use crate::Solution;

pub fn solve() -> Solution {
    let mut counts = HashMap::from_iter(
        read_to_string("input/11.txt")
            .unwrap()
            .trim()
            .split(" ")
            .map(|x| (x.parse::<u64>().unwrap(), 1u64)),
    );
    let mut p1 = 0;
    let mut p2 = 0;
    for i in 0..75 {
        let mut cpy = HashMap::new();
        for (n, v) in counts.iter() {
            if *n == 0 {
                *cpy.entry(1).or_insert(0) += *v;
            } else if n.ilog10() % 2 == 1 {
                let log = 10u64.pow(n.ilog10() / 2 + 1);
                let left = *n / log;
                let right = *n % log;
                *cpy.entry(left).or_insert(0) += *v;
                *cpy.entry(right).or_insert(0) += *v;
            } else {
                *cpy.entry(*n * 2024).or_insert(0) += *v;
            }
        }
        counts = cpy;
        if i == 24 {
            p1 = counts.values().sum();
        }
    }
    p2 = counts.values().sum();
    (p1.to_string(), p2.to_string())
}
