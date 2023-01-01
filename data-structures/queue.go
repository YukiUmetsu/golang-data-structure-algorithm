package dsalgo

import (
	"fmt"
)

type Queue struct {
	first  *LinkedListItem
	last   *LinkedListItem
	length int
}

func (q *Queue) Len() int { return q.length }

func (q *Queue) GetFirst() *LinkedListItem {
	return q.first
}

func (q *Queue) GetLast() *LinkedListItem {
	return q.last
}

func (q *Queue) Enqueue(value int) {
	newNode := LinkedListItem{
		value: value,
	}
	if q.length == 0 {
		q.first = &newNode
	} else {
		q.last.next = &newNode
	}
	q.last = &newNode
	q.length++
}

func (q *Queue) Dequeue() (int, error) {
	if q.length < 1 {
		return 0, fmt.Errorf("nothing to remove in the queue")
	}
	if q.length == 1 || q.first == q.last {
		q.last = nil
	}

	oldFirstValue := q.first.GetValue()
	q.first = q.first.next
	q.length--
	return oldFirstValue, nil
}

/*
	This queue stores nodes
*/
type LinkedListNode struct {
	value *Node
	next  *LinkedListNode
}

type QueueForNode struct {
	first  *LinkedListNode
	last   *LinkedListNode
	length int
}

func (li *LinkedListNode) Next() *LinkedListNode {
	return li.next
}

func (qn *QueueForNode) Len() int { return qn.length }

func (qn *QueueForNode) Enqueue(node *Node) {
	linkedListNode := &LinkedListNode{
		value: node,
	}
	if qn.length == 0 {
		qn.first = linkedListNode
	} else {
		qn.last.next = linkedListNode
	}
	qn.last = linkedListNode
	qn.length++
}

func (qn *QueueForNode) Dequeue() (*Node, error) {
	if qn.length < 1 {
		return nil, fmt.Errorf("nothing to remove in the queue")
	}
	if qn.length == 1 || qn.first == qn.last {
		qn.last = nil
	}

	oldFirstNode := qn.first.value
	qn.first = qn.first.next
	qn.length--
	return oldFirstNode, nil
}
