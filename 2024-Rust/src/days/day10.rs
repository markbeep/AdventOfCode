use std::{collections::HashSet, fs::read_to_string};

use crate::Solution;

pub fn solve() -> Solution {
    let f = read_to_string("input/10.txt").unwrap();
    let grid = f
        .trim()
        .split("\n")
        .map(|x| {
            x.chars()
                .map(|c| c.to_digit(10).unwrap())
                .collect::<Vec<u32>>()
        })
        .collect::<Vec<_>>();
    let mut p1 = 0;
    for (i, row) in grid.iter().enumerate() {
        for (j, num) in row.iter().enumerate() {
            if *num == 0 {
                p1 += dfs((i, j), &grid, &mut HashSet::new());
            }
        }
    }
    (p1.to_string(), "".to_string())
}

fn dfs((i, j): (usize, usize), grid: &Vec<Vec<u32>>, vis: &mut HashSet<(usize, usize)>) -> usize {
    if vis.contains(&(i, j)) {
        return 0;
    }
    vis.insert((i, j));

    let now = grid[i][j];
    if now == 9 {
        return 1;
    }

    let mut c = 0;
    if i > 0 && grid[i - 1][j] == now + 1 {
        c += dfs((i - 1, j), grid, vis);
    }
    if i < grid.len() - 1 && grid[i + 1][j] == now + 1 {
        c += dfs((i + 1, j), grid, vis);
    }
    if j > 0 && grid[i][j - 1] == now + 1 {
        c += dfs((i, j - 1), grid, vis);
    }
    if j < grid[0].len() - 1 && grid[i][j + 1] == now + 1 {
        c += dfs((i, j + 1), grid, vis);
    }
    c
}
