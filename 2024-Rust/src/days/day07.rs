use std::fs::read_to_string;

use crate::Solution;

pub fn solve() -> Solution {
    let f = read_to_string("input/07.txt").unwrap();
    let rows: Vec<_> = f
        .trim()
        .split("\n")
        .map(|x| {
            let (left, right) = x.split_once(": ").unwrap();
            (
                left.parse::<usize>().unwrap(),
                right
                    .split(" ")
                    .map(|y| y.parse().unwrap())
                    .collect::<Vec<usize>>(),
            )
        })
        .collect();
    let p1: usize = rows
        .iter()
        .filter(|(sum, nums)| rec(*sum, 0, nums, 0, false))
        .map(|(sum, _)| sum)
        .sum();
    let p2: usize = rows
        .iter()
        .filter(|(sum, nums)| rec(*sum, 0, nums, 0, true))
        .map(|(sum, _)| sum)
        .sum();
    (p1.to_string(), p2.to_string())
}

fn concat(a: usize, b: usize) -> usize {
    a * (10 as usize).pow(1 + b.ilog10()) + b
}

fn rec(sum: usize, cur: usize, nums: &Vec<usize>, i: usize, c: bool) -> bool {
    if i == nums.len() {
        return sum == cur;
    }
    rec(sum, cur + nums[i], nums, i + 1, c)
        || rec(sum, cur * nums[i], nums, i + 1, c)
        || c && rec(sum, concat(cur, nums[i]), nums, i + 1, c)
}
