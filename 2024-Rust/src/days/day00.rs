use std::fs::read_to_string;

use crate::Solution;

pub fn solve() -> Solution {
    // 2022 day 1 as warmup
    let f = read_to_string("input/01.txt").unwrap();
    let mut elves = f
        .trim()
        .split("\n\n")
        .map(|foods| {
            foods
                .split("\n")
                .map(|x| x.parse::<u32>().unwrap())
                .sum::<u32>()
        })
        .collect::<Vec<u32>>();
    elves.sort_by(|a, b| b.cmp(a));

    let p2 = elves[0] + elves[1] + elves[2];

    (elves[0].to_string(), p2.to_string())
}
