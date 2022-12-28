package main

import (
	"aoc/util"
	"aoc/util/ints"
	"fmt"
)

type Point struct {
	x int
	y int
}

func main() {
	f := util.ReadS("inp.txt", "\n")
	two := map[Point]bool{{0, 0}: true}
	nine := map[Point]bool{{0, 0}: true}
	knots := make([]Point, 10)
	for _, v := range f {
		var dir byte
		var count int
		fmt.Sscanf(v, "%c %d", &dir, &count)
		for i := 0; i < count; i++ {
			switch dir {
			case 'R':
				knots[0].x += 1
			case 'L':
				knots[0].x -= 1
			case 'U':
				knots[0].y -= 1
			case 'D':
				knots[0].y += 1
			}
			for j := 1; j < len(knots); j++ {
				horiz := ints.Abs(knots[j-1].x - knots[j].x)
				vert := ints.Abs(knots[j-1].y - knots[j].y)
				if horiz > 1 || vert > 1 {
					knots[j].x += ints.Sgn(knots[j-1].x - knots[j].x)
					knots[j].y += ints.Sgn(knots[j-1].y - knots[j].y)
					if j == 1 {
						two[knots[j]] = true
					} else if j == len(knots)-1 {
						nine[knots[j]] = true
					}
				}
			}
		}
	}
	fmt.Println("P1:", len(two), "P2:", len(nine))
}
