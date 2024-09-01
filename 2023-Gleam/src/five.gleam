import gleam/int
import gleam/list.{Continue, Stop}
import gleam/option.{None, Some}
import gleam/result
import gleam/string
import simplifile

pub type Maps {
  Maps(destination: Int, source: Int, range: Int)
}

fn read_section(lines) {
  string.split(lines, "\n")
  |> list.drop(1)
  |> list.map(fn(line) {
    case string.split(line, " ") {
      [d, s, r] -> {
        use d <- result.try(int.parse(d))
        use s <- result.try(int.parse(s))
        use r <- result.try(int.parse(r))
        Ok(Maps(destination: d, source: s, range: r))
      }
      _ -> Error(Nil)
    }
  })
  |> result.all
}

fn path(seed: Int, sections: List(List(Maps))) -> Int {
  case list.first(sections) {
    Error(_) -> seed
    Ok(next_stage) -> {
      let cut_sections = list.drop(sections, 1)
      let new_seed =
        list.fold_until(next_stage, None, fn(_, new) {
          case new.source <= seed && new.source + new.range > seed {
            True ->
              Stop(
                Some(path(new.destination + seed - new.source, cut_sections)),
              )
            False -> Continue(None)
          }
        })
      case new_seed {
        Some(x) -> x
        None -> path(seed, cut_sections)
      }
    }
  }
}

pub fn one() {
  use file <- result.try(
    simplifile.read("days/5.txt") |> result.replace_error(Nil),
  )
  let assert [seeds, _, ..rest] = file |> string.trim |> string.split("\n")
  let assert Ok(seeds) = case string.split_once(seeds, ": ") {
    Ok(#(_, seed_values)) ->
      string.split(seed_values, " ") |> list.map(int.parse) |> result.all
    _ -> panic
  }
  use sections <- result.try(
    string.join(rest, "\n")
    |> string.split("\n\n")
    |> list.map(read_section)
    |> result.all,
  )

  let min =
    list.map(seeds, path(_, sections))
    |> list.fold(99_999_999_999_999, int.min)

  Ok(min)
}

pub fn two() {
  todo
}
