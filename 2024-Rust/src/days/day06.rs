use std::{collections::HashSet, fs::read_to_string, usize};

use crate::Solution;

#[derive(Copy, Clone, PartialEq, Eq, Hash)]
enum Dir {
    Up,
    Right,
    Down,
    Left,
}
impl Dir {
    fn inc(self, pos: (i64, i64)) -> (i64, i64) {
        match self {
            Dir::Up => (pos.0 - 1, pos.1),
            Dir::Right => (pos.0, pos.1 + 1),
            Dir::Down => (pos.0 + 1, pos.1),
            Dir::Left => (pos.0, pos.1 - 1),
        }
    }
    fn next(self) -> Dir {
        match self {
            Dir::Up => Dir::Right,
            Dir::Right => Dir::Down,
            Dir::Down => Dir::Left,
            Dir::Left => Dir::Up,
        }
    }
}

pub fn solve() -> Solution {
    let f = read_to_string("input/06.txt").unwrap();
    let grid: Vec<Vec<char>> = f
        .trim()
        .split("\n")
        .map(|x| x.chars().collect::<Vec<char>>())
        .collect();
    let mut pos = 'brk: {
        for (i, row) in grid.iter().enumerate() {
            for (j, c) in row.iter().enumerate() {
                if *c == '^' {
                    break 'brk (i as i64, j as i64);
                }
            }
        }
        (0, 0)
    };
    let mut dir = Dir::Up;
    let mut visited = HashSet::from([pos]);
    let mut turns = HashSet::from([(pos, dir)]);
    let grid_w = grid[0].len();
    while pos.0 >= 0 && (pos.0 as usize) < grid.len() && pos.1 >= 0 && (pos.1 as usize) < grid_w {
        for _ in 0..4 {
            let new_pos = dir.inc(pos);
            if !(new_pos.0 >= 0
                && (new_pos.0 as usize) < grid.len()
                && new_pos.1 >= 0
                && (new_pos.1 as usize) < grid_w)
            {
                pos = new_pos;
                break;
            }
            if grid[new_pos.0 as usize][new_pos.1 as usize] == '.' {
                pos = new_pos;
                visited.insert(pos);
                break;
            } else {
                dir = dir.next();
            }
        }
        println!("{:?}", pos);
    }
    (visited.len().to_string(), "".to_string())
}
