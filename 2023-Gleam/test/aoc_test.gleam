import five
import four
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

pub fn d5_2_test() {
  five.two() |> should.equal(9_881_048)
}
