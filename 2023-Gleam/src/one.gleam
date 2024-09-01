import gleam/int
import gleam/list
import gleam/result
import gleam/string
import simplifile

fn read_file(filename) -> String {
  let assert Ok(x) = simplifile.read(filename)
  x
}

fn nums(line: String) -> Int {
  let is_num = fn(x) { result.is_ok(int.parse(x)) }
  let line_list = string.split(line, "")
  let assert Ok(left) = list.find(line_list, is_num)
  let assert Ok(right) = list.reverse(line_list) |> list.find(is_num)
  result.unwrap(int.parse(left), 0) * 10 + result.unwrap(int.parse(right), 0)
}

pub fn one() {
  read_file("days/day1.txt")
  |> string.trim
  |> string.split("\n")
  |> list.map(nums)
  |> int.sum
}

fn get_value(line: String, left: Bool) -> Int {
  let funcs = case left {
    True -> #(string.starts_with, string.drop_left)
    False -> #(string.ends_with, string.drop_right)
  }
  let check = fn(l, a, b, v, cb) {
    case funcs.0(l, a) || funcs.0(l, b) {
      True -> v
      False -> cb()
    }
  }

  use <- check(line, "one", "1", 1)
  use <- check(line, "two", "2", 2)
  use <- check(line, "three", "3", 3)
  use <- check(line, "four", "4", 4)
  use <- check(line, "five", "5", 5)
  use <- check(line, "six", "6", 6)
  use <- check(line, "seven", "7", 7)
  use <- check(line, "eight", "8", 8)
  use <- check(line, "nine", "9", 9)
  get_value(funcs.1(line, 1), left)
}

fn get_numbers(line) -> Int {
  get_value(line, True) * 10 + get_value(line, False)
}

pub fn two() {
  read_file("days/day1.txt")
  |> string.trim
  |> string.split("\n")
  |> list.map(get_numbers)
  |> int.sum
}
