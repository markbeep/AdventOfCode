package main

import (
	"aoc/util"
	"aoc/util/node"
	"fmt"
	"strings"
)

func main() {
	f := util.ReadS("inp.txt", "\n")
	var s, e *node.Node[rune]
	ns := util.Array2[*node.Node[rune]](len(f), len(f[0]))
	for i, a := range f {
		for j, b := range a {
			ns[i][j] = node.Create(b)
			if b == 'E' {
				e = ns[i][j]
				f[i] = strings.Replace(f[i], "E", "z", 1)
			} else if b == 'S' {
				s = ns[i][j]
				f[i] = strings.Replace(f[i], "S", "a", 1)
			}
		}
	}
	for i, a := range f {
		for j, b := range a {
			if valid(f, j+1, i, b) {
				ns[i][j].Sub = append(ns[i][j].Sub, ns[i][j+1])
			}
			if valid(f, j, i+1, b) {
				ns[i][j].Sub = append(ns[i][j].Sub, ns[i+1][j])
			}
			if valid(f, j-1, i, b) {
				ns[i][j].Sub = append(ns[i][j].Sub, ns[i][j-1])
			}
			if valid(f, j, i-1, b) {
				ns[i][j].Sub = append(ns[i][j].Sub, ns[i-1][j])
			}
		}
	}
	fmt.Println(node.Bfs([]*node.Node[rune]{s}, map[*node.Node[rune]]bool{}, make(map[*node.Node[rune]]int))[e])

}

func valid(f []string, x, y int, cur rune) bool {
	return y >= 0 && y < len(f) && x >= 0 && x < len(f[y]) && int(f[y][x])-int(cur) <= 1
}
