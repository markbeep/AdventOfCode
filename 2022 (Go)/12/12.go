package main

import (
	"aoc/util"
	"fmt"
	"strings"
)

type Point struct {
	x, y, d int
}

func main() {
	f := util.ReadS("inp.txt", "\n")
	e := Point{}
	s := Point{}
	aq := []*Point{&s}
	for i := range f {
		for j, v := range f[i] {
			if v == 'E' {
				f[i] = strings.Replace(f[i], "E", "z", 1)
				e.y = i
				e.x = j
			} else if v == 'S' {
				f[i] = strings.Replace(f[i], "S", "a", 1)
				s.y = i
				s.x = j
			} else if v == 'a' {
				aq = append(aq, &Point{y: i, x: j})
			}
		}
	}
	p1 := make(chan int, 1)
	p2 := make(chan int, 1)
	go bfs([]*Point{&s}, f, &e, p1)
	go bfs(aq, f, &e, p2)
	fmt.Println("Part 1:", <-p1, "Part 2:", <-p2)
}

func bfs(q []*Point, f []string, e *Point, ch chan int) {
	vis := util.Array[bool](len(f), len(f[0]))
	for len(q) > 0 {
		po := q[0]
		q = q[1:]
		cur := f[po.y][po.x]
		if po.x == e.x && po.y == e.y {
			ch <- po.d
			return
		}
		q = valid(q, f, po.x+1, po.y, po.d+1, cur, vis)
		q = valid(q, f, po.x, po.y+1, po.d+1, cur, vis)
		q = valid(q, f, po.x-1, po.y, po.d+1, cur, vis)
		q = valid(q, f, po.x, po.y-1, po.d+1, cur, vis)
	}
	ch <- -1
}

func valid(q []*Point, f []string, x, y, d int, cur byte, vis [][]bool) []*Point {
	if y >= 0 && y < len(f) && x >= 0 && x < len(f[y]) && int(f[y][x])-int(cur) <= 1 && !vis[y][x] {
		vis[y][x] = true
		return append(q, &Point{x: x, y: y, d: d})
	}
	return q
}
