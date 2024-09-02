import five.{Maps, Seed}
import four
import gleam/set
import gleeunit
import gleeunit/should
import one
import three
import two

pub fn main() {
  gleeunit.main()
}

pub fn d1_1_test() {
  one.one() |> should.equal(53_921)
}

pub fn d1_2_test() {
  one.two() |> should.equal(54_676)
}

pub fn d2_1_test() {
  two.one() |> should.equal(2449)
}

pub fn d2_2_test() {
  two.two() |> should.equal(63_981)
}

pub fn d3_1_test() {
  three.one() |> should.equal(527_369)
}

pub fn d3_2_test() {
  three.two() |> should.equal(73_074_886)
}

pub fn d4_1_test() {
  four.one() |> should.equal(21_919)
}

pub fn d4_2_test() {
  four.two() |> should.equal(9_881_048)
}

pub fn d5_1_test() {
  five.one() |> should.equal(Ok(340_994_526))
}

pub fn d5_2_step_test() {
  let maps = [Maps(3000, 1000, 2000), Maps(10_000, 11_000, 12_000)]
  five.step([], maps) |> should.equal([])

  // Disjoint left & disjoint in between & disjoint right
  five.step([Seed(0, 1000), Seed(2000, 11_000), Seed(12_000, 13_000)], maps)
  |> set.from_list
  |> should.equal(
    set.from_list([Seed(0, 1000), Seed(2000, 11_000), Seed(12_000, 13_000)]),
  )

  // Overlapping left
  five.step([Seed(0, 1500)], maps)
  |> set.from_list
  |> should.equal(set.from_list([Seed(0, 1000), Seed(3000, 3500)]))

  // Overlapping right
  five.step([Seed(1500, 3000)], maps)
  |> set.from_list
  |> should.equal(set.from_list([Seed(2000, 3000), Seed(3500, 4000)]))

  five.step([Seed(1000, 12_000)], maps)
  |> set.from_list
  |> should.equal(
    set.from_list([Seed(2000, 11_000), Seed(3000, 4000), Seed(10_000, 11_000)]),
  )
}

pub fn d5_2_test() {
  five.two() |> should.equal(Ok(52_210_644))
}
