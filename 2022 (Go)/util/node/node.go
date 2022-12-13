package node

type Node[T any] struct {
	Val T
	Par *Node[T]
	Sub []*Node[T]
}

func Create[T any](val T) *Node[T] {
	return &Node[T]{Val: val, Sub: []*Node[T]{}}
}

func (n *Node[T]) Add(val T) *Node[T] {
	n.Sub = append(n.Sub, Create(val))
	c := &Node[T]{Val: val, Sub: []*Node[T]{}}
	c.Par = n
	return c
}

func (n *Node[T]) Append(o *Node[T]) *Node[T] {
	n.Sub = append(n.Sub, o)
	o.Par = n
	return o
}

func Bfs[T comparable](q []*Node[T], vis map[*Node[T]]bool, dist map[*Node[T]]int) map[*Node[T]]int {
	for len(q) > 0 {
		p := q[0]
		q = q[1:]
		for _, c := range p.Sub {
			if !vis[c] {
				vis[c] = true
				dist[c] = dist[p] + 1
				q = append(q, c)
			}
		}
	}
	return dist
}
