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
	for k, v := range nodes {
		for k2 := range children[k] {
			// nodes[k2].Sub = append(nodes[k2].Sub, nodes[k])
			nodes[k].Sub = append(nodes[k].Sub, nodes[k2])
		}
		fmt.Printf("%s: %d (%d)\n", k, len(v.Sub), v.Val.rate)
	}
	fmt.Println(rec(nodes["AA"], map[string]bool{}, 1, 0))
	fmt.Println(c)
}

func rec(n *node.Node[Duo], open map[string]bool, min, cur int) int {
	if min >= 10 {
		return cur
	}
	c++
	results := []int{}
	shouldOpen := !open[n.Val.name]
	var cpy map[string]bool
	if shouldOpen {
		cpy = util.CopyMap(open)
		cpy[n.Val.name] = true
	}
	for _, v := range n.Sub {
		if n.Val.rate > 0 && shouldOpen {
			results = append(results, rec(v, cpy, min+2, cur+(30-min)*n.Val.rate)) // open valve
		}
		results = append(results, rec(v, open, min+1, cur)) // don't open
	}
	return ints.Max(results...)
}
