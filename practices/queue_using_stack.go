package dsalgo

import "fmt"

// Implement a queue using stack
// Solution: have input stack and output stack.
// 1. When item is added, add to input stack.
// 2. when taking an item, take an item from output stack.
// 3. When output stack is empty on taking an item, move items from input stack to output stack before taking an item (flips the order)

/*
	Preparation Code
*/
type Node struct {
	value int
	next  *Node
}

func (n *Node) GetValue() int {
	return n.value
}

func (n *Node) SetValue(value int) {
	n.value = value
}

func (n *Node) Next() *Node {
	return n.next
}

func (n *Node) SetNext(item *Node) {
	n.next = item
}

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

/*
	Main Code
*/

type QueueWithStack struct {
	inputStack  *Stack
	outputStack *Stack
	length      int
}

func (q *QueueWithStack) New() {
	q.inputStack = &Stack{}
	q.outputStack = &Stack{}
	q.length = 0
}

func (q *QueueWithStack) GetFirst() *Node {
	q.moveInputToOutputStack()
	return q.outputStack.Peek()
}

func (q *QueueWithStack) Len() int {
	return q.length
}

func (q *QueueWithStack) GetLast() *Node {
	if q.inputStack.Len() > 0 {
		return q.inputStack.top
	}

	if q.outputStack.Len() > 0 {
		return q.outputStack.bottom
	}

	return nil
}

func (q *QueueWithStack) Peek(value int) *Node {
	q.moveInputToOutputStack()
	return q.outputStack.Peek()
}

func (q *QueueWithStack) Enqueue(value int) {
	newNode := &Node{
		value: value,
	}
	q.inputStack.Push(newNode)
	q.length++
}

func (q *QueueWithStack) Dequeue() (*Node, error) {
	if q.length == 0 {
		return nil, fmt.Errorf("queue is empty")
	}
	q.moveInputToOutputStack()
	popped := q.outputStack.Pop()
	q.length--
	return popped, nil
}

func (q *QueueWithStack) moveInputToOutputStack() {
	if q.inputStack.Len() == 0 || q.outputStack.Len() > 0 {
		return
	}

	originalLength := q.inputStack.Len()
	for i := 0; i < originalLength; i++ {
		node := q.inputStack.Pop()
		q.outputStack.Push(node)
	}
}
