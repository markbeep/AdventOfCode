use std::{cmp::Ordering, collections::HashMap, fs::read_to_string};

use crate::Solution;

pub fn solve() -> Solution {
    let f = read_to_string("input/05.txt").unwrap();
    let (rules, updates) = f.trim().split_once("\n\n").unwrap();
    let mut pre: HashMap<&str, Vec<&str>> = HashMap::new();
    for rule in rules.split("\n") {
        let (before, after) = rule.split_once("|").unwrap();
        match pre.get_mut(after) {
            None => drop(pre.insert(after, vec![before])),
            Some(x) => x.push(before),
        }
    }
    let mut p1 = 0;
    let mut p2 = 0;
    for update in updates.split("\n") {
        let update: Vec<&str> = update.split(",").collect();
        let valid = is_valid(&pre, &update);
        p1 += (valid as usize) * update[update.len() / 2].parse::<usize>().unwrap();
        if !valid {
            p2 += correct(&pre, &update);
        }
    }
    (p1.to_string(), p2.to_string())
}

fn is_valid(pre: &HashMap<&str, Vec<&str>>, update: &Vec<&str>) -> bool {
    for (i, up) in update.iter().enumerate() {
        for bef in pre.get(up).unwrap_or(&vec![]) {
            let bef_i = match update.iter().position(|x| x == bef) {
                None => continue,
                Some(x) => x,
            };
            if i < bef_i {
                return false;
            }
        }
    }
    return true;
}

fn correct(pre: &HashMap<&str, Vec<&str>>, update: &Vec<&str>) -> usize {
    let mut upd = update.clone();
    upd.sort_by(|a, b| {
        if pre.get(b).unwrap().contains(a) {
            Ordering::Less
        } else if pre.get(a).unwrap().contains(b) {
            Ordering::Greater
        } else {
            Ordering::Equal
        }
    });
    return upd[upd.len() / 2].parse::<usize>().unwrap();
}
