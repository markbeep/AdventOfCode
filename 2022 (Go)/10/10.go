package main

import (
	"aoc/util"
	"fmt"
	"strings"
)

func main() {
	f := util.ReadS("inp.txt", "\n")
	c := 0
	queue := make([]int, 1)
	crt := make([]string, 40)
	for i := range crt {
		crt[i] = "."
	}
	x := 1
	for i := 0; len(queue) > 0; i++ {
		po := queue[0]
		x += po
		queue = queue[1:]
		if len(f) > 0 {
			v := f[0]
			f = f[1:]
			if v == "noop" {
				queue = append(queue, 0)
			} else {
				var k int
				fmt.Sscanf(v, "addx %d", &k)
				queue = append(queue, 0, k)
			}
		}
		ind := i % 40
		crt[ind] = "."
		if ind >= x-1 && ind <= x+1 {
			crt[ind] = "#"
		}
		if (i-19)%40 == 0 {
			c += (i + 1) * x
		}
		if (i-39)%40 == 0 {
			fmt.Println(strings.Join(crt, ""))
			for i := range crt {
				crt[i] = "."
			}
		}
	}
	fmt.Println("Part 1:", c)
}
