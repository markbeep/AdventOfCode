package main

import "fmt"

func main() {
	n := create(5)
	n.add(3)
	n.add(2)
	n.add(1)
	fmt.Println(n)
}

type node[T any] struct {
	val T
	sub []*node[T]
	par *node[T]
}

func create[T any](val T) node[T] {
	return node[T]{val: val, sub: []*node[T]{}}
}

func (n *node[T]) add(val T) node[T] {
	c := create(val)
	c.par = n
	n.sub = append(n.sub, &c)
	return c
}
