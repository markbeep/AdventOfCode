package main

import (
	"aoc/util"
	"aoc/util/ints"
	"aoc/util/node"
	"fmt"
	"regexp"
)

type Duo struct {
	rate int
	name string
}

type Path struct {
	vis int64
	cur int
}

var c = 0
var paths = []Path{}
var bitfield = map[string]int64{}

func main() {
	f := util.ReadS("inp.txt", "\n")
	re, reRate := regexp.MustCompile(`([A-Z]){2}`), regexp.MustCompile(`(\d+)`)
	nodes := map[string]*node.Node[Duo]{}
	for _, v := range f {
		res := re.FindAllString(v, -1)
		rate := ints.SInt(reRate.FindString(v))
		bitfield[res[0]] = 1 << len(bitfield)
		if nodes[res[0]] == nil {
			nodes[res[0]] = node.Create(Duo{rate: rate, name: res[0]})
		}
		nodes[res[0]].Val.rate = rate
		for _, k := range res[1:] {
			if nodes[k] == nil {
				nodes[k] = node.Create(Duo{name: k})
			}
			nodes[k].Sub = append(nodes[k].Sub, nodes[res[0]])
			nodes[res[0]].Sub = append(nodes[res[0]].Sub, nodes[k])
		}
	}
	// distance graph from a node to every other node
	ranges := map[string]map[*node.Node[Duo]]int{}
	for k, v := range nodes {
		ranges[k] = map[*node.Node[Duo]]int{}
		for o1, dist := range node.Bfs([]*node.Node[Duo]{v}, map[*node.Node[Duo]]bool{}, map[*node.Node[Duo]]int{}) {
			if o1.Val.rate > 0 {
				ranges[k][o1] = dist
			}
		}
	}
	part1 := rec(nodes["AA"], bitfield["AA"], 0, ranges, 0, 0)
	fmt.Println("Part 1:", part1)
	rec2(nodes["AA"], bitfield["AA"], 0, ranges, 0, 0)
	m := 0
	for _, a := range paths {
		if a.cur < part1/2 {
			continue
		}
		for _, b := range paths {
			if a.vis&b.vis == bitfield["AA"] {
				m = ints.Max(m, a.cur+b.cur)
			}
		}
	}
	fmt.Println("Part 2:", m)
}

func rec(n *node.Node[Duo], nInt, vis int64, dist map[string]map[*node.Node[Duo]]int, min, cur int) int {
	if min > 30 || nInt&vis > 0 {
		return cur
	}
	c++
	cur += (30 - min) * n.Val.rate
	nw := 0
	for k, v := range dist[n.Val.name] {
		nw = ints.Max(nw, rec(k, bitfield[k.Val.name], vis|nInt, dist, min+v+1, cur))
	}
	return nw
}

func rec2(n *node.Node[Duo], nInt, vis int64, dist map[string]map[*node.Node[Duo]]int, min, cur int) int {
	if min > 26 || nInt&vis > 0 {
		return cur
	}
	cur += (26 - min) * n.Val.rate
	paths = append(paths, Path{vis: vis | nInt, cur: cur})
	nw := 0
	for k, v := range dist[n.Val.name] {
		nw = ints.Max(nw, rec2(k, bitfield[k.Val.name], vis|nInt, dist, min+v+1, cur))
	}
	return nw
}
