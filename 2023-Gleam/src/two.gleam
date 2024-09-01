import gleam/bool
import gleam/int
import gleam/list
import gleam/regex
import gleam/string
import utils

pub type Colors {
  Color(red: Int, green: Int, blue: Int)
}

fn per_sect(sect) {
  let assert Ok(re) = regex.from_string("\\d+ [rgb]")
  let sp = regex.scan(re, sect)

  list.fold(sp, Color(0, 0, 0), fn(old, m) {
    let regex.Match(st, _) = m
    case string.split(st, " ") {
      [v, "r"] -> Color(..old, red: old.red + utils.yolo_parse(v))
      [v, "g"] -> Color(..old, green: old.green + utils.yolo_parse(v))
      [v, "b"] -> Color(..old, blue: old.blue + utils.yolo_parse(v))
      _ -> old
    }
  })
}

fn per_line(line, i, part1) -> Int {
  let colors =
    case string.split(line, ": ") {
      [_, x] -> x
      _ -> ""
    }
    |> string.split(";")
    |> list.map(per_sect)

  case part1 {
    // Part 1: Return the ID (index) if valid
    True -> {
      let valid =
        list.all(colors, fn(c) { c.red <= 12 && c.green <= 13 && c.blue <= 14 })
      bool.to_int(valid) * { i + 1 }
    }

    // Part 2: Get max RGB values out of each section
    False -> {
      let c =
        list.fold(colors, Color(0, 0, 0), fn(old, c) {
          Color(
            red: int.max(old.red, c.red),
            green: int.max(old.green, c.green),
            blue: int.max(old.blue, c.blue),
          )
        })
      c.red * c.green * c.blue
    }
  }
}

pub fn one() {
  utils.prepare_file("days/2.txt")
  |> list.index_map(fn(x, i) { per_line(x, i, True) })
  |> int.sum
}

pub fn two() {
  utils.prepare_file("days/2.txt")
  |> list.index_map(fn(x, i) { per_line(x, i, False) })
  |> int.sum
}
