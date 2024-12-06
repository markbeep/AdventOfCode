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
    fn inc(self, pos: (usize, usize)) -> (usize, usize) {
        match self {
            Dir::Up => (pos.0.wrapping_sub(1), pos.1),
            Dir::Right => (pos.0, pos.1 + 1),
            Dir::Down => (pos.0 + 1, pos.1),
            Dir::Left => (pos.0, pos.1.wrapping_sub(1)),
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
    let mut grid: Vec<Vec<char>> = f
        .trim()
        .split("\n")
        .map(|x| x.chars().collect::<Vec<char>>())
        .collect();

    let pos = f.find('^').unwrap();
    let pos = (pos / (grid[0].len() + 1), pos % (grid[0].len() + 1));
    let mut visited = HashSet::from([pos]);
    let mut empty = HashSet::new();
    let p1 = iterate_until_done(&grid, pos, &mut visited, true).unwrap();
    let mut p2 = 0;
    for (i, j) in visited.iter() {
        grid[*i][*j] = '#';
        p2 += iterate_until_done(&grid, pos, &mut empty, false).is_none() as usize;
        grid[*i][*j] = '.';
    }
    (p1.to_string(), p2.to_string())
}

/// Iterates until outside or a loop is found
fn iterate_until_done(
    grid: &Vec<Vec<char>>,
    start_pos: (usize, usize),
    visited: &mut HashSet<(usize, usize)>,
    add: bool,
) -> Option<usize> {
    let mut pos = start_pos;
    let mut dir = Dir::Up;
    let mut turns = HashSet::from([(pos, dir)]);
    let grid_w = grid[0].len();
    loop {
        for _ in 0..4 {
            let new_pos = dir.inc(pos);
            if !(new_pos.0 < grid.len() && new_pos.1 < grid_w) {
                return Some(visited.len());
            }
            match grid[new_pos.0][new_pos.1] {
                '#' => dir = dir.next(),
                _ => {
                    pos = new_pos;
                    if add {
                        visited.insert(pos);
                    }
                    if turns.contains(&(pos, dir)) {
                        return None;
                    }
                    turns.insert((pos, dir));
                    break;
                }
            }
        }
    }
}
