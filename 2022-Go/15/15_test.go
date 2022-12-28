package main

import (
	"aoc/util"
	"aoc/util/ints"
	"fmt"
	"testing"
)

func BenchmarkRun(b *testing.B) {
	for i := 0; i < b.N; i++ {
		main()
	}
}

func BenchmarkSingle(b *testing.B) {
	f := util.ReadS("inp.txt", "\n")
	sensors := []sens{} // store all sensors and the beacon location
	for _, v := range f {
		var sx, sy, bx, by int
		fmt.Sscanf(v, "Sensor at x=%d, y=%d: closest beacon is at x=%d, y=%d", &sx, &sy, &bx, &by)
		sensors = append(sensors, sens{sx, sy, bx, by, ints.Abs(sx-bx) + ints.Abs(sy-by)})
	}
	for i := 0; i < b.N; i++ {
		single(10, sensors)
	}
}

func BenchmarkCombine(b *testing.B) {
	for i := 0; i < b.N; i++ {
		combine([]nrange{{-123123123, 5124123}, {5124523, 55124123}}, nrange{5124023, 51124123})
	}
}

func BenchmarkIntersect(b *testing.B) {
	for i := 0; i < b.N; i++ {
		intersect(nrange{-123123123, 5124123}, nrange{5124023, 51124123})
	}
}
