package main

import (
	"aoc/util"
	"fmt"
)

func main() {
	f := util.ReadS("inp.txt", "\n")
	c := 0
	hash := map[string]bool{}

	queue := make([]int, 1)
	x := 1
	i := 0
	for ; len(queue) > 0; i++ {
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
				queue = append(queue, 0)
				queue = append(queue, k)
			}
		}
		fmt.Println(queue)
		if (i-19)%40 == 0 {
			fmt.Println(i, x)
			c += (i + 1) * x
		}

	}
	_ = f
	_ = hash
	fmt.Println(x, c)
}
