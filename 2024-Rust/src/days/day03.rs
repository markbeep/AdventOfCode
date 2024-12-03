use regex::Regex;
use std::fs::read_to_string;

use crate::Solution;

pub fn solve() -> Solution {
    let f = read_to_string("input/03.txt").unwrap();
    let re = Regex::new(r"mul\((\d{1,3}),(\d{1,3})\)").unwrap();
    let mut p: (i64, i64) = (0, 0);
    for cap in re.captures_iter(&f) {
        let [a, b] = cap.extract().1.map(|x| x.parse::<i64>().unwrap());
        p.0 += a * b;
        let ofs = cap.get(0).unwrap().start();
        let split = f.split_at(ofs).0;
        match (split.rfind("do()"), split.rfind("don't()")) {
            (None, None) | (Some(_), None) => p.1 += a * b,
            (Some(x), Some(y)) if x > y => p.1 += a * b,
            _ => {}
        };
    }
    (p.0.to_string(), p.1.to_string())
}
