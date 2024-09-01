import gleam/dict
import gleam/int
import gleam/list
import gleam/option.{None, Some}
import gleam/pair
import gleam/regex
import gleam/result
import gleam/string
import utils

fn matches(win, own) {
  let assert Ok(re) = regex.from_string("\\d+")
  regex.scan(re, own)
  |> list.filter(fn(v) {
    string.contains(" " <> win <> " ", " " <> v.content <> " ")
  })
}

fn get_input() {
  utils.prepare_file("days/4.txt")
  |> list.map(string.split(_, ": "))
  |> list.map(list.last)
  |> list.map(result.unwrap(_, ""))
  |> list.map(string.split_once(_, " | "))
}

pub fn one() {
  get_input()
  |> list.fold(0, fn(old, value) {
    let assert Ok(#(win, own)) = value
    let wins =
      matches(win, own)
      |> list.fold(0, fn(total, _) { int.max(total * 2, 1) })
    wins + old
  })
}

fn add_tickets(dic, current, rem, multiplier) {
  let increment = fn(x) {
    case x {
      Some(x) -> x + multiplier
      None -> 1 + multiplier
    }
  }
  case rem {
    0 -> dic
    _ ->
      add_tickets(
        dict.upsert(dic, current, increment),
        current + 1,
        rem - 1,
        multiplier,
      )
  }
}

pub fn two() {
  get_input()
  |> list.index_fold(#(0, dict.new()), fn(old, value, index) {
    let assert Ok(#(win, own)) = value
    let wins =
      matches(win, own)
      |> list.length
    let #(total, dic) = old
    let multiplier = dict.get(dic, index) |> result.unwrap(1)

    #(
      total + wins * multiplier + 1,
      add_tickets(dic, index + 1, wins, multiplier),
    )
  })
  |> pair.first
}
