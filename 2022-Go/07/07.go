package main

import (
	"fmt"
	"os"
	"sort"
	"strings"
	"time"
)

var total = 0
var spaces []int

func main() {
	start := time.Now()
	f, _ := os.ReadFile("inp.txt")
	cont := strings.Split(strings.Trim(string(f), " \n"), "\n")
	root := buildTree(cont) // builds tree
	ch := make(chan int, 1)
	dfs(root, ch) // finds total root folder size
	tot := <-ch
	fmt.Println("Part 1:", total)
	sort.Ints(spaces)
	for _, v := range spaces {
		if v > tot-40000000 {
			fmt.Println("Part 2:", v)
			break
		}
	}
	fmt.Println("Time:", time.Since(start))
}

func dfs(n *node, ch chan int) {
	if !n.dir {
		ch <- n.val
		return
	}
	tot := 0
	sch := make(chan int, len(n.sub))
	for _, v := range n.sub {
		dfs(v, sch)
	}
	for range n.sub {
		tot += <-sch
	}
	if tot <= 100_000 {
		total += tot
	}
	spaces = append(spaces, tot)
	ch <- tot
}

func buildTree(cont []string) *node {
	root := &node{name: "/", dir: true, sub: map[string]*node{}}
	cur := root
	for _, v := range cont[1:] {
		var name string
		var val int
		switch v[2] {
		case 'c': // cd
			fmt.Sscanf(v, "$ cd %s", &name)
			if name == ".." {
				cur = cur.par
			} else {
				cur = cur.add(name, 0, true)
			}
		case 'l': // ls
		case 'r': // dir
			fmt.Sscanf(v, "dir %s", &name)
			cur.add(name, 0, true)
		default:
			fmt.Sscanf(v, "%d %s", &val, &name)
			cur.add(name, val, false)
		}
	}
	return root
}

type node struct {
	name string
	val  int
	dir  bool
	sub  map[string]*node
	par  *node
}

func (n *node) add(name string, val int, dir bool) *node {
	if n.sub[name] != nil {
		return n.sub[name]
	}
	c := &node{name: name, val: val, sub: map[string]*node{}, dir: dir}
	c.par = n
	n.sub[name] = c
	return c
}
