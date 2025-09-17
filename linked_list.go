package main

type LinkedList struct {
	head *Node
	tail *Node
	size int
}

type Node struct {
	value any
	next  *Node
	prev  *Node
}

func (l *LinkedList) AddToTail(value any) {
	node := Node{value: value}
	tail := l.tail
	l.tail = &node
	l.size++

	if l.size == 1 {
		l.head = &node
		return
	}

	tail.next = &node
	node.prev = tail
}

func (l *LinkedList) AddToHead(value any) {
	node := Node{value: value}
	head := l.head
	l.head = &node
	l.size++

	if l.size == 1 {
		l.tail = &node
		return
	}

	head.prev = &node
	node.next = head
}

func (l *LinkedList) PopTail() *any {
	if l.size == 0 {
		return nil
	}
	l.size--
	tail := l.tail

	if l.size == 0 {
		l.head = nil
		l.tail = nil
		return &tail.value
	}

	tail.prev.next = nil

	return &tail.value
}

func (l *LinkedList) AddAfter(node *Node, val any) {
	l.size++
	newNode := Node{
		value: val,
		prev:  node,
		next:  node.next,
	}
	node.next = &newNode

	if newNode.next == nil {
		l.tail = &newNode
		return
	}

	newNode.next.prev = &newNode
}

func (l *LinkedList) AddBefore(node *Node, val any) {
	l.size++
	newNode := Node{
		value: val,
		prev:  node.prev,
		next:  node,
	}
	node.prev = &newNode

	if newNode.prev == nil {
		l.head = &newNode
	}

	newNode.prev.next = &newNode
}

func (l *LinkedList) Swap(node1 *Node, node2 *Node) {
	prev1, next1 := node1.prev, node1.next
	prev2, next2 := node2.prev, node2.next

	if prev1 == nil {
		l.head = node2
	} else if prev2 == nil {
		l.head = node1
	}

	if next1 == nil {
		l.tail = node2
	} else if next2 == nil {
		l.tail = node1
	}

}
