package main

import (
	"fmt"
	"log"
	"os"
	"sort"
	"strings"
)

var total = 0
var free_space = 0
var spaces []int

func main() {
	f, err := os.ReadFile("inp.txt")
	if err != nil {
		log.Fatal(err)
	}
	cont := strings.Split(strings.Trim(string(f), " \n"), "\n")
	root := buildTree(cont)
	tot := dfs(root)
	fmt.Println("Part 1:", total)
	free_space = tot - 40000000
	dfs(root)
	sort.Ints(spaces)
	fmt.Println("Part 2:", spaces[0])
}

func dfs(n *node) int {
	if !n.dir {
		if n.val <= 100_000 {
			return n.val
		}
		return n.val
	}
	tot := 0
	for _, v := range n.sub {
		t := dfs(v)
		tot += t
	}
	if tot <= 100_000 {
		total += tot
	}
	if free_space > 0 && tot >= free_space {
		spaces = append(spaces, tot)
	}
	return tot
}

func buildTree(cont []string) *node {
	root := create("/", 0, true)
	cur := root
parser:
	for i := 0; i < len(cont[1:]); {
		if cont[i][2:4] == "cd" {
			var name string
			fmt.Sscanf(cont[i], "$ cd %s", &name)
			if name == ".." {
				cur = cur.par
				i++
				continue parser
			}
			cur = cur.add(name, 0, true)
			i++
			continue parser
		}
		if cont[i][2:4] == "ls" {
			i++
			for i < len(cont) && cont[i][0] != '$' {
				if cont[i][0] == 'd' {
					var name string
					fmt.Sscanf(cont[i], "dir %s", &name)
					cur.add(name, 0, true)
				} else {
					var name string
					var val int
					fmt.Sscanf(cont[i], "%d %s", &val, &name)
					cur.add(name, val, false)
				}
				i++
				continue
			}
			continue parser
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

func create(name string, val int, dir bool) *node {
	return &node{name: name, val: val, sub: map[string]*node{}, dir: dir}
}

func (n *node) add(name string, val int, dir bool) *node {
	if n.sub[name] != nil {
		return n.sub[name]
	}
	c := create(name, val, dir)
	c.par = n
	n.sub[name] = c
	return c
}
