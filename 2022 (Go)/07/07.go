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
	f, _ := os.ReadFile("inp.txt")
	cont := strings.Split(strings.Trim(string(f), " \n"), "\n")
	start := time.Now()
	root := buildTree(cont) // builds tree
	fmt.Println("Time:", time.Since(start))
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
	for i := 0; i < len(cont[1:]); i++ {
		var name string
		if cont[i][2:4] == "cd" {
			fmt.Sscanf(cont[i], "$ cd %s", &name)
			if name == ".." {
				cur = cur.par
			} else {
				cur = cur.add(name, 0, true)
			}
		} else if cont[i][2:4] == "ls" {
			for i+1 < len(cont) && cont[i+1][0] != '$' {
				i++
				if cont[i][0] == 'd' {
					fmt.Sscanf(cont[i], "dir %s", &name)
					cur.add(name, 0, true)
				} else {
					var val int
					fmt.Sscanf(cont[i], "%d %s", &val, &name)
					cur.add(name, val, false)
				}
			}
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
