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
        .filter(|(sum, nums)| rec(*sum, 0, nums, 0))
        .map(|(sum, _)| sum)
        .sum();
    (p1.to_string(), "".to_string())
}

fn rec(sum: usize, cur: usize, nums: &Vec<usize>, i: usize) -> bool {
    if i == nums.len() {
        return sum == cur;
    }
    rec(sum, cur + nums[i], nums, i + 1) || rec(sum, cur * nums[i], nums, i + 1)
}
