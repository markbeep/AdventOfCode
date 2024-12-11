use std::fs::read_to_string;

use crate::Solution;

pub fn solve() -> Solution {
    let f = read_to_string("input/11.txt").unwrap();
    let mut nums: Vec<u64> = f
        .trim()
        .split(" ")
        .map(|x| x.parse::<u64>().unwrap())
        .collect();

    for i in 0..25 {
        let mut cpy = vec![];
        for n in nums {
            if n == 0 {
                cpy.push(1);
            } else if n.ilog10() % 2 == 1 {
                let log = 10u64.pow(n.ilog10() / 2 + 1);
                let left = n / log;
                let right = n % log;
                cpy.push(left);
                cpy.push(right);
            } else {
                cpy.push(n * 2024);
            }
        }
        nums = cpy;
    }

    (nums.len().to_string(), "".to_string())
}
