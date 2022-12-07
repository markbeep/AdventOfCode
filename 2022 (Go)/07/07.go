package main

import (
	"fmt"
	"log"
	"os"
	"strings"
)

var total int = 0

func main() {
	f, err := os.ReadFile("inp.txt")
	if err != nil {
		log.Fatal(err)
	}
	cont := strings.Split(strings.Trim(string(f), " \n"), "\n")
	root := buildTree(cont)
	fmt.Println(dfs(root))
	fmt.Println(total)
}

func dfs(n *node) (int, int) {
	if !n.dir {
		if n.val <= 100_000 {
			return n.val, n.val
		}
		return 0, n.val
	}
	c := 0
	tot := 0
	for _, v := range n.sub {
		val, tot_b := dfs(v)
		tot += tot_b
		c += val
	}
	if tot <= 100_000 {
		c += tot
		total += tot
	}
	return c, tot
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
	val int
	dir bool
	sub map[string]*node
	par *node
}

func create(name string, val int, dir bool) *node {
	return &node{val: val, sub: map[string]*node{}, dir: dir}
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
