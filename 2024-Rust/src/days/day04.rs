use std::fs::read_to_string;

use crate::Solution;

pub fn solve() -> Solution {
    let f = read_to_string("input/04.txt").unwrap();
    let t = f.trim().to_string();
    let arr: Vec<Vec<char>> = t
        .split("\n")
        .map(|x| x.chars().collect::<Vec<char>>())
        .collect();
    let mut combinations = t.to_owned();
    let mut diag_down = String::new();
    let mut diag_up = String::new();
    let mut down = String::new();

    for x in 0..arr[0].len() {
        for y in 0..arr.len() {
            down.push_str(&arr[y][x].to_string());
            let xr = x + y;
            let yd = arr.len() - y - 1;
            if xr < arr[0].len() && y + 1 < arr.len() {
                diag_down.push_str(&arr[y][xr].to_string());
            }
            if xr < arr[0].len() && yd > 0 {
                diag_up.push_str(&arr[yd][xr].to_string());
            }
        }
    }
    combinations.push_str(&diag_down);
    combinations.push_str(&diag_up);
    combinations.push_str(&down);
    let backwards = combinations.chars().rev().collect::<String>();
    combinations.push_str(&backwards);
    let p1 = combinations.match_indices("XMAS").count();

    let p1 = String::from("XXMAS").match_indices("XMAS").count();
    (p1.to_string(), "".to_string())
}
