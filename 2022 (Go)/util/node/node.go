package node

type Node[T any] struct {
	val T
	par *Node[T]
	sub []*Node[T]
}

func Create[T any](val T) *Node[T] {
	return &Node[T]{val: val, sub: []*Node[T]{}}
}

func (n *Node[T]) Add(val T) *Node[T] {
	n.sub = append(n.sub, Create(val))
	c := &Node[T]{val: val, sub: []*Node[T]{}}
	c.par = n
	return c
}
