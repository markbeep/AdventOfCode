package main

import (
	"fmt"
	"log"
	"os"
	"strings"
)

func main() {
	f, err := os.ReadFile("inp.txt")
	if err != nil {
		log.Fatal(err)
	}
	cont := strings.Split(strings.Trim(string(f), " \n"), "\n")
	c := 0
	hash := map[string]bool{}

	for _, v := range cont {
		fmt.Sscanf(v, "")

	}

	_ = cont
	_ = hash
	fmt.Println(c)
}

type node[T any] struct {
	val T
	par *node[T]
	sub []*node[T]
}

func create[T any](val T) *node[T] {
	return &node[T]{val: val, sub: []*node[T]{}}
}

func (n *node[T]) add(val T) *node[T] {
	n.sub = append(n.sub, create(val))
	c := &node[T]{val: val, sub: []*node[T]{}}
	c.par = n
	return c
}
