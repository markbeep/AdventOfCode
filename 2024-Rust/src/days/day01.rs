use std::{fs::read_to_string, iter::zip};

use crate::Solution;

pub fn solve() -> Solution {
    let f = read_to_string("input/01.txt").unwrap();
    let mut l1: Vec<i64> = vec![];
    let mut l2: Vec<i64> = vec![];
    let mut count: [i64; 99999] = [0; 99999];
    for line in f.trim().split("\n") {
        let r: Vec<i64> = line
            .split("   ")
            .map(|x| x.parse::<i64>().unwrap())
            .collect();
        l1.push(r[0]);
        l2.push(r[1]);
        count[r[1] as usize] += 1;
    }
    let p2 = l1.iter().fold(0, |old, a| old + a * count[*a as usize]);
    l1.sort();
    l2.sort();
    let p1 = zip(l1, l2).fold(0, |old, (a, b)| old + (a - b).abs());
    (p1.to_string(), p2.to_string())
}
