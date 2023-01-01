package dsalgo

type LinkedListItem struct {
	value int
	next  *LinkedListItem
}

func (li *LinkedListItem) GetValue() int {
	return li.value
}

func (li *LinkedListItem) SetValue(value int) {
	li.value = value
}

func (li *LinkedListItem) Next() *LinkedListItem {
	return li.next
}

func (li *LinkedListItem) SetNext(item *LinkedListItem) {
	li.next = item
}

type LinkedList struct {
	head   *LinkedListItem
	tail   *LinkedListItem
	length int
}

func (l *LinkedList) Append(value int) {
	newItem := LinkedListItem{value: value}
	if l.length < 1 {
		l.head = &newItem
		l.tail = &newItem
		l.length++
		return
	}

	nextItem := l.tail
	for nextItem.Next() != nil {
		nextItem = nextItem.Next()
	}
	nextItem.SetNext(&newItem)
	l.tail = &newItem
	l.length++
}

func (l *LinkedList) Prepend(value int) {
	newItem := &LinkedListItem{value: value, next: l.head}
	if l.length < 1 {
		l.tail = newItem
	}
	l.head = newItem
	l.length++
}

func (l *LinkedList) GetLastValue() int {
	if l.length == 0 {
		return 0
	}
	if l.length == 1 {
		return l.head.GetValue()
	}

	nextIterateItem := l.tail
	for nextIterateItem.Next() != nil {
		nextIterateItem = nextIterateItem.Next()
	}
	return nextIterateItem.GetValue()
}

func (l *LinkedList) GetFirstValue() int {
	if l.head == nil {
		return 0
	}

	return l.head.GetValue()
}

func (l *LinkedList) GetValueByIndex(index int) int {
	if index == 0 {
		return l.GetFirstValue()
	}

	count := 0
	currentNode := l.head
	for currentNode.Next() != nil && count < index {
		currentNode = currentNode.Next()
		count++
	}
	return currentNode.GetValue()
}
