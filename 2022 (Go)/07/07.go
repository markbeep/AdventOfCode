package main

import (
	"fmt"
	"log"
	"os"
	"strings"
)

var hash = map[string]bool{}

type node struct {
	val int
	sub []node
}

func create(val int) node {
	return node{val: val, sub: []node{}}
}

func (n node) add(par node, val int) {
	n.sub = append(n.sub, create(val))
}

func main() {
	f, err := os.ReadFile("inp.txt")
	if err != nil {
		log.Fatal(err)
	}
	cont := strings.Split(strings.Trim(string(f), " \n"), "\n")
	c := 0

	for _, v := range cont {
		fmt.Sscanf(v, "")

	}

	_ = cont
	fmt.Println(c)
}
