import gleam/int
import gleam/list.{Continue, Stop}
import gleam/option.{None, Some}
import gleam/result
import gleam/string
import simplifile
import utils

pub type Maps {
  Maps(destination: Int, source: Int, end: Int)
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
        Ok(Maps(destination: d, source: s, end: s + r))
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
          case new.source <= seed && new.end > seed {
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

  list.map(seeds, path(_, sections))
  |> list.reduce(int.min)
}

/// Non inclusive end
pub type Seed {
  Seed(start: Int, end: Int)
  Mapped(start: Int, end: Int)
}

pub fn step(original_seeds: List(Seed), maps: List(Maps)) -> List(Seed) {
  let new_seeds =
    list.fold(maps, original_seeds, fn(seeds_to_check, map) {
      // Takes some seeds and checks if they should be turned into more mapped seeds
      list.flat_map(seeds_to_check, fn(seed) {
        case seed {
          Mapped(_, _) as x -> [x]
          _ -> {
            // Check if in range
            case map {
              // [ () ] completely included
              Maps(d, s, e) if s <= seed.start && e >= seed.end -> {
                let offset_start = seed.start - s
                let offset_end = seed.end - s
                [Mapped(d + offset_start, d + offset_end)]
              }
              // [ ( ] ) right overlapping
              Maps(d, s, e)
                if s <= seed.start && e < seed.end && seed.start < e
              -> {
                let offset_start = seed.start - s
                let offset_end = e - s
                [Mapped(d + offset_start, d + offset_end), Seed(e, seed.end)]
              }
              // ( [ ) ] left overlapping
              Maps(d, s, e) if s > seed.start && e >= seed.end && seed.end > s -> {
                let offset_end = seed.end - s
                [Seed(seed.start, s), Mapped(d, d + offset_end)]
              }
              // ( [] ) // overlapping all
              Maps(d, s, e) if s > seed.start && e < seed.end -> {
                let offset_end = e - s
                [
                  Seed(seed.start, s),
                  Mapped(d, d + offset_end),
                  Seed(e, seed.end),
                ]
              }
              // disjoint
              _ -> [seed]
            }
          }
        }
      })
    })

  list.map(new_seeds, fn(seed) {
    case seed {
      Seed(_, _) as x -> x
      Mapped(s, e) -> Seed(s, e)
    }
  })
  |> list.unique
}

pub fn two() {
  let assert [seeds, _, ..rest] = utils.prepare_file("days/5.txt")
  let assert Ok(seeds) = case string.split_once(seeds, ": ") {
    Ok(#(_, seed_values)) ->
      string.split(seed_values, " ") |> list.map(int.parse) |> result.all
    _ -> panic
  }
  let starting_seeds =
    list.sized_chunk(seeds, 2)
    |> list.map(fn(x) {
      let assert [start, len] = x
      Seed(start: start, end: start + len)
    })
  use sections <- result.try(
    string.join(rest, "\n")
    |> string.split("\n\n")
    |> list.map(read_section)
    |> result.all,
  )

  list.fold(sections, starting_seeds, step)
  |> list.flat_map(fn(x) { [x.start, x.end] })
  |> list.reduce(int.min)
}
