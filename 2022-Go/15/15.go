package main

import (
	"aoc/util"
	"aoc/util/ints"
	"fmt"
	"runtime"
	"sync"
)

type nrange struct {
	from, to int
}
type sens struct {
	sx, sy, bx, by, dist int
}

func main() {
	f := util.ReadS("inp.txt", "\n")
	P1, P2 := 2000000, 4000000
	sensors := []sens{} // store all sensors and the beacon location
	for _, v := range f {
		var sx, sy, bx, by int
		fmt.Sscanf(v, "Sensor at x=%d, y=%d: closest beacon is at x=%d, y=%d", &sx, &sy, &bx, &by)
		sensors = append(sensors, sens{sx, sy, bx, by, ints.Abs(sx-bx) + ints.Abs(sy-by)})
	}
	c := 0
	for _, v := range single(P1, sensors) {
		c += v.to - v.from
	}
	fmt.Print("Part 1: ", c)
	NUM_ROUTS := ints.Min(runtime.GOMAXPROCS(0), runtime.NumCPU())
	SPLIT := P2 / NUM_ROUTS
	wg := new(sync.WaitGroup)
	for i := 0; i < NUM_ROUTS; i++ {
		wg.Add(1)
		go func(fr, to int, wg *sync.WaitGroup) {
			for i := fr; i < ints.Min(P2, to); i++ {
				if rng := single(i, sensors); len(rng) == 2 {
					fmt.Printf(" | Part 2: %d (%d, %d)\n", i+(ints.Min(rng[0].to, rng[1].to)+1)*4000000, i, ints.Min(rng[0].to, rng[1].to)+1)
					break
				}
			}
			wg.Done()
		}(i*SPLIT, (i+1)*SPLIT, wg)
	}
	wg.Wait()

}

func single(line int, sensors []sens) []nrange {
	xranges := []nrange{}
	for _, v := range sensors {
		if distToLine := ints.Abs(v.sy - line); distToLine <= v.dist {
			size := ints.Abs(distToLine - v.dist)
			xranges = combine(xranges, nrange{from: v.sx - size, to: v.sx + size})
		}
	}
	return xranges
}

// adds a new range and handles intersections
func combine(rngs []nrange, r nrange) []nrange {
	nw := make([]nrange, 0, 2)
	for _, v := range rngs {
		if in := intersect(v, r); len(in) == 2 { // v is disjunct
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
	if r1.to+1 < r2.from || r2.to+1 < r1.from {
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
	return []nrange{{from: ints.Min(r1.from, r2.from), to: ints.Max(r1.to, r2.to)}}
}
