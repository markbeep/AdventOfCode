package main

import (
	"aoc/util"
	"aoc/util/ints"
	"fmt"
)

type nrange struct {
	from, to int
}

func main() {
	f := util.ReadS("inp.txt", "\n")
	c := 0

	for i := 0; i < 20; i++ {
	}
	for _, v := range single(2000000, f) {
		c += v.to - v.from
	}
	fmt.Println("Part 1:", c)

}

func single(line int, f []string) []nrange {
	xranges := []nrange{}
	for _, v := range f {
		var sx, sy, bx, by int
		fmt.Sscanf(v, "Sensor at x=%d, y=%d: closest beacon is at x=%d, y=%d", &sx, &sy, &bx, &by)
		dist := ints.Abs(sx-bx) + ints.Abs(sy-by)
		distToLine := ints.Abs(sy - line)
		if distToLine > dist {
			continue // ignore sensors too far away
		}
		size := ints.Abs(distToLine - dist)
		xranges = combine(xranges, nrange{from: sx - size, to: sx + size})
	}
	return xranges
}

// adds a new range and handles intersections
func combine(rngs []nrange, r nrange) []nrange {
	nw := []nrange{}
	for _, v := range rngs {
		in := intersect(v, r)
		if len(in) == 2 { // v is disjunct
			nw = append(nw, v)
		} else {
			r = in[0]
		}
	}
	nw = append(nw, r)
	return nw
}

// returns both or an intersected range
func intersect(r1, r2 nrange) []nrange {
	// dont intersect
	if r1.to < r2.from || r2.to < r1.from {
		return []nrange{r1, r2}
	}
	// contained
	if r1.from >= r2.from && r1.to <= r2.to {
		return []nrange{r2}
	}
	if r2.from >= r1.from && r2.to <= r1.to {
		return []nrange{r1}
	}
	// intersect (take min/max of both ranges)
	mx := ints.Max(r1.to, r2.to)
	mn := ints.Min(r1.from, r2.from)
	return []nrange{{from: mn, to: mx}}
}
