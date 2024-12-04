use std::fs::read_to_string;

use crate::Solution;

pub fn solve() -> Solution {
    let f = read_to_string("input/02.txt").unwrap();
    let reports = f.trim().split("\n").map(|x| {
        x.split(" ")
            .map(|n| n.parse::<i64>().unwrap())
            .collect::<Vec<i64>>()
    });
    let mut p1 = 0;
    let mut p2 = 0;
    for report in reports {
        if is_safe(report.to_owned()) {
            p1 += 1;
            p2 += 1;
            continue;
        }
        for i in 0..report.len() {
            let mut rep = report.clone();
            rep.remove(i);
            if is_safe(rep) {
                p2 += 1;
                break;
            }
        }
    }
    (p1.to_string(), p2.to_string())
}

fn is_safe(x: Vec<i64>) -> bool {
    x.is_sorted_by(|a, b| (-3..=-1).contains(&(b - a)))
        || x.is_sorted_by(|a, b| (1..=3).contains(&(b - a)))
}
