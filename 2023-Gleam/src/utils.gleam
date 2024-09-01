import gleam/int
import gleam/io
import gleam/string
import simplifile

pub fn read_file(filename) -> String {
  case simplifile.read(filename) {
    Ok(v) -> v
    Error(err) -> {
      io.debug(err)
      panic as "Invalid filename"
    }
  }
}

pub fn prepare_file(filename) -> List(String) {
  read_file(filename)
  |> string.trim
  |> string.split("\n")
}

pub fn yolo_parse(v: String) -> Int {
  case int.parse(v) {
    Ok(x) -> x
    Error(_) -> panic as "failed to yolo parse string into int"
  }
}
