package dsalgo

import "fmt"

type Stack struct {
	top    *Node
	bottom *Node
	length int
	body   []*Node
}

func (s *Stack) Peek() *Node {
	return s.top
}

func (s *Stack) Push(n *Node) {
	s.top = n
	if s.length < 1 {
		s.bottom = n
	}
	s.body = append(s.body, n)
	s.length++
}

func (s *Stack) Pop() *Node {
	if s.length < 1 {
		return nil
	}

	returnNode := *s.top
	if s.length == 1 {
		s.top = nil
		s.bottom = nil
		s.length = 0
		s.body = []*Node{}
		return &returnNode
	}

	s.top = s.body[len(s.body)-2]
	s.length--
	s.body = s.body[:len(s.body)-1]
	return &returnNode
}

func (s *Stack) String() string {
	return fmt.Sprintf("%v", s.body)
}

func (s *Stack) Len() int {
	return s.length
}
