package main

import (
	"fmt"
	"os"
	"strings"
)

type Point struct {
	x int
	y int
}

func main() {
	f, _ := os.ReadFile("inp.txt")
	cont := strings.Split(strings.Trim(string(f), " \n"), "\n")
	two := map[Point]bool{{0, 0}: true}
	nine := map[Point]bool{{0, 0}: true}
	knots := make([]Point, 10)
	for _, v := range cont {
		var dir byte
		var count int
		fmt.Sscanf(v, "%c %d", &dir, &count)
		for i := 0; i < count; i++ {
			switch dir {
			case 'R':
				knots[0] = Point{knots[0].x + 1, knots[0].y}
			case 'L':
				knots[0] = Point{knots[0].x - 1, knots[0].y}
			case 'U':
				knots[0] = Point{knots[0].x, knots[0].y - 1}
			case 'D':
				knots[0] = Point{knots[0].x, knots[0].y + 1}
			}
			heads[knots[0]] = true
			for j := 1; j < len(knots); j++ {
				horiz := abs(knots[j-1].x - knots[j].x)
				vert := abs(knots[j-1].y - knots[j].y)
				if horiz > 1 || vert > 1 {
					sgnH := sgn(knots[j-1].x - knots[j].x)
					sgnV := sgn(knots[j-1].y - knots[j].y)
					knots[j] = Point{knots[j].x + sgnH, knots[j].y + sgnV}
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

func sgn(x int) int {
	if x < 0 {
		return -1
	} else if x == 0 {
		return 0
	}
	return 1
}

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}
