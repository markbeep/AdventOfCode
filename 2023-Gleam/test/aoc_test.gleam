import gleeunit
import gleeunit/should
import one
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
