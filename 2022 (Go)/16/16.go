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

var c = 0

func main() {
	f := util.ReadS("inp.txt", "\n")
	re := regexp.MustCompile(`([A-Z]){2}`)
	re2 := regexp.MustCompile(`(\d+)`)
	nodes := map[string]*node.Node[Duo]{}
	children := map[string]map[string]bool{}
	dp := map[string][]int{}

	// Valve ED has flow rate=0; tunnels lead to valves PS, AW
	for _, v := range f {
		res := re.FindAllString(v, -1)
		rate := ints.SInt(re2.FindString(v))

		if nodes[res[0]] == nil {
			nodes[res[0]] = node.Create(Duo{rate: rate, name: res[0]})
		} else {
			nodes[res[0]].Val.rate = rate
		}
		children[res[0]] = map[string]bool{}
		dp[res[0]] = make([]int, 31)
		for _, k := range res[1:] {
			if nodes[k] == nil {
				nodes[k] = node.Create(Duo{name: k})
			}
			children[res[0]][k] = true
		}
	}

	for k := range nodes {
		for k2 := range children[k] {
			nodes[k].Sub = append(nodes[k].Sub, nodes[k2])
		}
	}

	ranges := map[string]map[*node.Node[Duo]]int{}
	for k, v := range nodes {
		ranges[k] = node.Bfs([]*node.Node[Duo]{v}, map[*node.Node[Duo]]bool{}, map[*node.Node[Duo]]int{})
	}
	slim := map[string]map[*node.Node[Duo]]int{}
	for k, v := range ranges {
		slim[k] = map[*node.Node[Duo]]int{}
		for k2, v2 := range v {
			if k2.Val.rate > 0 {
				slim[k][k2] = v2
			}
		}
	}

	for k, v := range slim {
		fmt.Print(k, ": ")
		for k2, v2 := range v {
			fmt.Printf("%s:%d | ", k2.Val.name, v2)
		}
		fmt.Println()
	}

	fmt.Println("Part 1:", rec(nodes["AA"], slim, map[string]bool{}, 0, 0))
	fmt.Println(c)
}

func rec(n *node.Node[Duo], dist map[string]map[*node.Node[Duo]]int, vis map[string]bool, min, cur int) int {
	if min > 30 || vis[n.Val.name] {
		return cur
	}
	cpy := util.CopyMap(vis)
	cur += (30 - min) * n.Val.rate
	c++
	nw := 0
	cpy[n.Val.name] = true
	for k, v := range dist[n.Val.name] {
		if !vis[k.Val.name] {
			nw = ints.Max(nw, rec(k, dist, cpy, min+v+1, cur))
		}
	}
	return nw
}
