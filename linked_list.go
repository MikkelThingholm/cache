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

func (l *LinkedList) AddToTail(node Node) {
	l.size++
	tail := l.tail
	l.tail = &node
	if tail == nil {
		l.head = &node
	}

}

func (l *LinkedList) AddToHead(node Node) {
	l.size++
	head := l.head
	l.head = &node
	if head == nil {
		l.tail = &node
	}

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
