import gleeunit
import gleeunit/should
import one

pub fn main() {
  gleeunit.main()
}

pub fn d1_1_test() {
  one.one() |> should.equal(53_921)
}

pub fn d1_2_test() {
  one.two() |> should.equal(54_676)
}
