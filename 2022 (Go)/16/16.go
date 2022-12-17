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

func main() {
	f := util.ReadS("inp.txt", "\n")
	re, reRate := regexp.MustCompile(`([A-Z]){2}`), regexp.MustCompile(`(\d+)`)
	nodes := map[string]*node.Node[Duo]{}
	for _, v := range f {
		res := re.FindAllString(v, -1)
		rate := ints.SInt(reRate.FindString(v))

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
	fmt.Println("Part 1:", rec(nodes["AA"], ranges, map[string]bool{}, 0, 0))
	fmt.Println("Part 2:", rec2(nodes["AA"], nodes["AA"], ranges, map[string]bool{}, 0, 0, 0, 0))
}

func rec(n *node.Node[Duo], dist map[string]map[*node.Node[Duo]]int, vis map[string]bool, min, cur int) int {
	if min > 30 || vis[n.Val.name] {
		return cur
	}
	cpy := util.CopyMap(vis)
	cur += (30 - min) * n.Val.rate
	nw := 0
	cpy[n.Val.name] = true
	for k, v := range dist[n.Val.name] {
		if !vis[k.Val.name] {
			nw = ints.Max(nw, rec(k, dist, cpy, min+v+1, cur))
		}
	}
	return nw
}

func rec2(n1, n2 *node.Node[Duo], dist map[string]map[*node.Node[Duo]]int, vis map[string]bool, min, m1, m2, cur int) int {
	if min > 5 || vis[n1.Val.name] || vis[n2.Val.name] {
		return cur
	}
	ch1, ch2 := false, false
	if min == m1 {
		ch1 = true
		cur += (26 - min) * n1.Val.rate
		fmt.Println("kek1", cur, m1)
	}
	if min == m2 {
		ch2 = true
		cur += (26 - min) * n2.Val.rate
		fmt.Println("kek2", cur, m2)
	}
	nw := 0
	cpy := util.CopyMap(vis)
	cpy[n1.Val.name] = true
	cpy[n2.Val.name] = true
	if !ch1 && !ch2 { // stay on the same nodes, inc minute
		fmt.Printf("Change nothing: (%s, %s) [%d, %d] [%d, %d]\n", n1.Val.name, n2.Val.name, min, cur, m1, m2)
		nw = ints.Max(nw, rec2(n1, n2, dist, cpy, min+1, m1, m2, cur))
	} else if ch1 && ch2 { // change both nodes
		fmt.Printf("Change both: (%s, %s) [%d, %d]\n", n1.Val.name, n2.Val.name, min, cur)
		for k1, v1 := range dist[n1.Val.name] {
			if !vis[k1.Val.name] {
				for k2, v2 := range dist[n2.Val.name] {
					if k1 != k2 {
						nw = ints.Max(nw, rec2(k1, k2, dist, cpy, min+1, min+v1+1, min+v2+1, cur))
					}
				}
			}
		}
	} else if ch1 { // change first node
		fmt.Printf("Change ch1: (%s, %s) [%d, %d]\n", n1.Val.name, n2.Val.name, min, cur)
		for k1, v1 := range dist[n1.Val.name] {
			nw = ints.Max(nw, rec2(k1, n2, dist, cpy, min+1, min+v1+1, m2, cur))
		}
	} else if ch2 { // change second node
		fmt.Printf("Change ch2: (%s, %s) [%d, %d]\n", n1.Val.name, n2.Val.name, min, cur)
		for k2, v2 := range dist[n2.Val.name] {
			nw = ints.Max(nw, rec2(n1, k2, dist, cpy, min+1, m1, min+v2+1, cur))
		}

	}
	return nw
}

func pRanges(ranges map[string]map[*node.Node[Duo]]int) {
	for k, v := range ranges {
		fmt.Print(k, ": ")
		for k2, v2 := range v {
			fmt.Printf("%s:%d | ", k2.Val.name, v2)
		}
		fmt.Println()
	}
}
