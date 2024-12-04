use std::fs::read_to_string;

use crate::Solution;

pub fn solve() -> Solution {
    let f = read_to_string("input/04.txt").unwrap();
    let t = f.trim().to_string();
    let arr: Vec<Vec<char>> = t
        .split("\n")
        .map(|x| x.chars().collect::<Vec<char>>())
        .collect();
    let to_compare = ['X', 'M', 'A', 'S'];
    let to_compare_inv = ['S', 'A', 'M', 'X'];
    let mut p1 = 0;
    for x in 0..arr[0].len() - 3 {
        for y in 0..arr.len() {
            p1 += (arr[y][x..x + 4] == to_compare || arr[y][x..x + 4] == to_compare_inv) as usize; // right
            if y + 3 < arr.len() {
                let diag_down = [
                    arr[y][x],
                    arr[y + 1][x + 1],
                    arr[y + 2][x + 2],
                    arr[y + 3][x + 3],
                ];
                p1 += (diag_down == to_compare || diag_down == to_compare_inv) as usize;
                let left_diag_down = [
                    arr[y][arr[0].len() - x - 1],
                    arr[y + 1][arr[0].len() - x - 2],
                    arr[y + 2][arr[0].len() - x - 3],
                    arr[y + 3][arr[0].len() - x - 4],
                ];
                p1 += (left_diag_down == to_compare || left_diag_down == to_compare_inv) as usize;
            }
        }
    }
    for x in 0..arr[0].len() {
        for y in 0..arr.len() - 3 {
            let down = [arr[y][x], arr[y + 1][x], arr[y + 2][x], arr[y + 3][x]];
            p1 += (down == to_compare || down == to_compare_inv) as usize; // up down
        }
    }

    (p1.to_string(), "".to_string())
}
