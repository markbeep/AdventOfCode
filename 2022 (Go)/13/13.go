package main

import (
	"aoc/util"
	"aoc/util/ints"
	"aoc/util/node"
	"fmt"
	"strings"
)

func main() {
	f := util.ReadS("inp.txt", "\n\n")
	v1, _ := eval("[[2]]")
	v2, _ := eval("[[6]]")
	var p1, c1, c2 int
	for i, v := range f {
		s := strings.Split(v, "\n")
		l, _ := eval(s[0])
		r, _ := eval(s[1])
		if comp(l, r) {
			p1 += i + 1
		}
		c1 += ints.BInt(comp(r, v1))
		c1 += ints.BInt(comp(l, v1))
		c2 += ints.BInt(comp(l, v2))
		c2 += ints.BInt(comp(r, v2))
	}
	fmt.Println("Part 1:", p1, "Part 2:", (c1+1)*(c2+2))

}

func comp(l, r *node.Node[int]) bool {
	b, _ := comp_aux(l, r)
	return b
}

func comp_aux(l, r *node.Node[int]) (bool, bool) {
	if l.Val >= 0 && r.Val >= 0 {
		return l.Val < r.Val, l.Val == r.Val
	} else if l.Val >= 0 { // wrap left again
		o := node.Create(-1)
		o.Append(l)
		return comp_aux(o, r)
	} else if r.Val >= 0 { // wrap right again
		o := node.Create(-1)
		o.Append(r)
		return comp_aux(l, o)
	}
	// both are nodes
	for i := 0; i < len(l.Sub); i++ {
		if i >= len(r.Sub) {
			return false, false // left is smaller
		}
		c, v := comp_aux(l.Sub[i], r.Sub[i])
		if !v {
			return c, false
		}
	}
	return len(l.Sub) < len(r.Sub), len(l.Sub) == len(r.Sub)
}

func eval(s string) (*node.Node[int], int) {
	n := node.Create(-1)
	for i := 1; i < len(s); i++ {
		if s[i] == '[' {
			nn, ind := eval(s[i:])
			n.Append(nn)
			i += ind
		} else if s[i] == ']' {
			return n, i
		} else if s[i] != ',' {
			// is a number
			num := ints.SInt(s[i : i+1])
			// check if its 10
			if num == 1 && i < len(s)-1 && s[i+1] == '0' {
				num = 10
				i++
			}
			n.Add(num)
		}
	}
	return n, len(s) - 1
}

func pNode(n *node.Node[int], offspace int) {
	if n.Val == -1 {
		fmt.Println(strings.Repeat(" ", offspace), "n: Sub:", len(n.Sub))
	} else {
		fmt.Println(strings.Repeat(" ", offspace), "- c: Val:", n.Val)
	}
	for _, v := range n.Sub {
		pNode(v, offspace+2)
	}
}
