import gleam/bool
import gleam/dict.{type Dict}
import gleam/function
import gleam/int
import gleam/list
import gleam/string
import utils

pub type Grid {
  Number(Int, x: Int, y: Int)
  Symbol(String)
}

type Pos {
  Pos(val: Int, x: Int, y: Int)
}

/// Gets the next number
fn get_num(lines: String, acc: String) -> String {
  case string.first(lines) {
    Ok(v) -> {
      case int.parse(v) {
        Ok(_) -> get_num(string.drop_left(lines, 1), acc <> v)
        Error(_) -> acc
      }
    }
    Error(_) -> acc
  }
}

fn get_pos(lines, x, y, acc) -> Dict(#(Int, Int), Grid) {
  case string.first(lines) {
    Ok("\n") -> get_pos(string.drop_left(lines, 1), 0, y + 1, acc)
    Ok(".") -> get_pos(string.drop_left(lines, 1), x + 1, y, acc)
    Ok(v) -> {
      case int.parse(v) {
        Ok(_) -> {
          let num = get_num(lines, "")
          let i_num = Number(utils.yolo_parse(num), x: x, y: y)
          // add coordinates for all nums
          let acc =
            string.split(num, "")
            |> list.index_fold(acc, fn(old, _, i) {
              old |> dict.insert(#(x + i, y), i_num)
            })
          get_pos(
            string.drop_left(lines, string.length(num)),
            x + string.length(num),
            y,
            acc,
          )
        }
        Error(_) -> {
          // symbol
          let acc = acc |> dict.insert(#(x, y), Symbol(v))
          get_pos(string.drop_left(lines, 1), x + 1, y, acc)
        }
      }
    }
    // EOF
    Error(_) -> acc
  }
}

fn should_count(acc, x, y) -> List(Pos) {
  let ex = fn(xi, yi) {
    case dict.get(acc, #(xi, yi)) {
      Ok(Number(v, orig_x, orig_y)) -> Ok(Pos(val: v, x: orig_x, y: orig_y))
      _ -> Error(Nil)
    }
  }
  let ranges = [
    ex(x - 1, y - 1),
    ex(x, y - 1),
    ex(x + 1, y - 1),
    ex(x - 1, y),
    ex(x + 1, y),
    ex(x - 1, y + 1),
    ex(x, y + 1),
    ex(x + 1, y + 1),
  ]
  list.filter_map(ranges, function.identity)
}

pub fn one() {
  let f = utils.read_file("days/3.txt")
  let pos = get_pos(f, 0, 0, dict.new())

  let #(total, _) =
    dict.fold(pos, #(0, dict.new()), fn(old, k, v) {
      let is_number = case v {
        Number(_, _, _) -> True
        Symbol(_) -> False
      }
      use <- bool.guard(is_number, old)

      let #(count, added) = old
      let matches = should_count(pos, k.0, k.1)
      // Add up the matches
      list.fold(matches, #(count, added), fn(old, v) {
        let #(count, added) = old
        let new_count = case dict.has_key(added, #(v.x, v.y)) {
          False -> v.val + count
          True -> count
        }
        #(new_count, dict.insert(added, #(v.x, v.y), True))
      })
    })
  total
}

pub fn two() {
  let f = utils.read_file("days/3.txt")
  let pos = get_pos(f, 0, 0, dict.new())

  dict.fold(pos, 0, fn(old, k, v) {
    let is_number = case v {
      Number(_, _, _) -> True
      Symbol("*") -> False
      _ -> True
    }
    use <- bool.guard(is_number, old)

    let matches =
      should_count(pos, k.0, k.1)
      |> list.unique
    use <- bool.guard(list.length(matches) != 2, old)

    old + { list.map(matches, fn(x) { x.val }) |> list.fold(1, int.multiply) }
  })
}
