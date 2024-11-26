mod days;

use days::{
    day00, day01, day02, day03, day04, day05, day06, day07, day08, day09, day10, day11, day12,
    day13, day14, day15, day16, day17, day18, day19, day20, day21, day22, day23, day24,
};
use std::env;

pub type Solution = (String, String);

fn main() {
    let args: Vec<String> = env::args().collect();
    if args.len() < 2 {
        panic!("Not enough arguments: cargo run <day: 0-24>")
    }

    let day: u8 = args[1].parse().unwrap();

    let day_solve = match day {
        0 => day00::solve(),
        1 => day01::solve(),
        2 => day02::solve(),
        3 => day03::solve(),
        4 => day04::solve(),
        5 => day05::solve(),
        6 => day06::solve(),
        7 => day07::solve(),
        8 => day08::solve(),
        9 => day09::solve(),
        10 => day10::solve(),
        11 => day11::solve(),
        12 => day12::solve(),
        13 => day13::solve(),
        14 => day14::solve(),
        15 => day15::solve(),
        16 => day16::solve(),
        17 => day17::solve(),
        18 => day18::solve(),
        19 => day19::solve(),
        20 => day20::solve(),
        21 => day21::solve(),
        22 => day22::solve(),
        23 => day23::solve(),
        24 => day24::solve(),
        _ => unimplemented!(),
    };

    println!(
        "\x1b[1;4mDay: {}\n\x1b[0m\x1b[1mPart 1: \x1b[0m{}\n======\n\x1b[1mPart 2: \x1b[0m{}",
        day, day_solve.0, day_solve.1
    );
}
