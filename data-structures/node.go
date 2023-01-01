package dsalgo

import "fmt"

type Node struct {
	value      int
	left       *Node
	right      *Node
	bfsDataArr []int
	dfsDataArr []int
}

func (n *Node) String() string {
	return fmt.Sprintf("%d", n.value)
}

func (n *Node) SetValue(value int) {
	n.value = value
}

func (n *Node) GetValue() int {
	return n.value
}

func (n *Node) SetLeft(left *Node) {
	n.left = left
}

func (n *Node) GetLeft() *Node {
	return n.left
}

func (n *Node) SetRight(right *Node) {
	n.right = right
}

func (n *Node) GetRight() *Node {
	return n.right
}
