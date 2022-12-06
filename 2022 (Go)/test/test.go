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
	sub []node[T]
}

func create[T any](val T) node[T] {
	return node[T]{val: val, sub: []node[T]{}}
}

func (n *node[T]) add(val T) {
	n.sub = append(n.sub, create(val))
}
