package main

import (
	"aoc/util"
	"aoc/util/ints"
	"fmt"
	"math"
	"strings"
)

type Point struct {
	x int
	y int
	d int
}

func main() {
	f := util.ReadS("inp.txt", "\n")
	e := Point{}
	l := math.MaxInt
outer:
	for i := range f {
		v := strings.Index(f[i], "E")
		if v >= 0 {
			f[i] = strings.Replace(f[i], "E", "z", 1)
			e.y = i
			e.x = v
			break outer
		}
	}
	for i := range f {
		for j := range f[i] {
			if f[i][j] == 'S' || f[i][j] == 'a' {
				vis := util.Array[bool](len(f), len(f[i]))
				f[i] = strings.Replace(f[i], "S", "a", 1)
				v := bfs(f, j, i, &e, vis)
				if f[i][j] == 'S' {
					fmt.Println("Part 1:", v)
				}
				l = ints.Min(v, l)
			}
		}
	}
	fmt.Println("Part 2:", l)
}

func bfs(f []string, sx, sy int, e *Point, vis [][]bool) int {
	q := []*Point{{x: sx, y: sy}}
	for len(q) > 0 {
		po := q[0]
		q = q[1:]
		cur := f[po.y][po.x]
		vis[po.y][po.x] = true
		if po.x == e.x && po.y == e.y {
			return po.d
		}
		q = valid(q, f, po.x+1, po.y, po.d+1, cur, vis)
		q = valid(q, f, po.x, po.y+1, po.d+1, cur, vis)
		q = valid(q, f, po.x-1, po.y, po.d+1, cur, vis)
		q = valid(q, f, po.x, po.y-1, po.d+1, cur, vis)
	}
	return math.MaxInt
}

func valid(q []*Point, f []string, x, y, d int, cur byte, vis [][]bool) []*Point {
	if y >= 0 && y < len(f) && x >= 0 && x < len(f[y]) && int(f[y][x])-int(cur) <= 1 && !vis[y][x] {
		vis[y][x] = true
		return append(q, &Point{x: x, y: y, d: d})
	}
	return q
}
