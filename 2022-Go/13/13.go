package main

import (
	"aoc/util"
	"aoc/util/ints"
	"aoc/util/node"
	"fmt"
	"runtime"
	"strings"
)

func main() {
	f := util.ReadS("inp.txt", "\n\n")
	v1, _ := eval("[[2]]")
	v2, _ := eval("[[6]]")
	NUM_ROUTS := ints.Min(runtime.GOMAXPROCS(0), runtime.NumCPU())
	split := len(f) / NUM_ROUTS
	ch1 := make(chan int, NUM_ROUTS)
	ch2 := make(chan int, NUM_ROUTS)
	ch3 := make(chan int, NUM_ROUTS)
	for r := 0; r < NUM_ROUTS; r++ {
		if r < NUM_ROUTS-1 {
			go calc(f[r*split:(r+1)*split], ch1, ch2, ch3, r*split, v1, v2)
		} else {
			go calc(f[r*split:], ch1, ch2, ch3, r*split, v1, v2)

		}
	}
	var p1, c1, c2 int
	for i := 0; i < NUM_ROUTS; i++ {
		p1 += <-ch1
		c1 += <-ch2
		c2 += <-ch3
	}
	fmt.Println("Part 1:", p1, "Part 2:", (c1+1)*(c2+2))

}

func calc(f []string, ch1, ch2, ch3 chan<- int, si int, v1, v2 *node.Node[int]) {
	var p1, c1, c2 int
	for i, v := range f {
		s := strings.Split(v, "\n")
		l, _ := eval(s[0])
		r, _ := eval(s[1])
		p1 += ints.BInt(comp(l, r)) * (i + 1 + si)
		c1 += ints.BInt(comp(r, v1)) + ints.BInt(comp(l, v1))
		c2 += ints.BInt(comp(l, v2)) + ints.BInt(comp(r, v2))
	}
	ch1 <- p1
	ch2 <- c1
	ch3 <- c2
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
		switch s[i] {
		case '[':
			nn, ind := eval(s[i:])
			n.Append(nn)
			i += ind
		case ']':
			return n, i
		case ',':
		default:
			// is a number
			num := ints.CInt(s[i])
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
