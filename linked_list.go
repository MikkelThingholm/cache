package main

type LinkedList[K comparable, V any] struct {
	sentinel Node[K, V]
	len      int
}

type Node[K comparable, V any] struct {
	Key        K
	Value      V
	prev, next *Node[K, V]
}

func NewLinkedList[K comparable, V any]() *LinkedList[K, V] {
	l := &LinkedList[K, V]{}
	l.sentinel.next = &l.sentinel
	l.sentinel.prev = &l.sentinel
	return l
}

func (l *LinkedList[K, V]) Length() int {
	return l.len
}

func (l *LinkedList[K, V]) Head() *Node[K, V] {
	if l.len == 0 {
		return nil
	}
	return l.sentinel.next
}

func (l *LinkedList[K, V]) Tail() *Node[K, V] {
	if l.len == 0 {
		return nil
	}
	return l.sentinel.prev
}

func (l *LinkedList[K, V]) PushHead(node *Node[K, V]) {
	l.len++

	node.next = l.sentinel.next
	node.prev = &l.sentinel

	node.prev.next = node
	node.next.prev = node
}

func (l *LinkedList[K, V]) PushTail(node *Node[K, V]) {
	l.len++

	node.next = &l.sentinel
	node.prev = l.sentinel.prev

	node.next.prev = node
	node.prev.next = node
}

func (l *LinkedList[K, V]) MoveToHead(node *Node[K, V]) {
	node.prev.next = node.next
	node.next.prev = node.prev

	node.next = l.sentinel.next
	node.prev = &l.sentinel

	node.prev.next = node
	node.next.prev = node
}

func (l *LinkedList[K, V]) MoveToTail(node *Node[K, V]) {
	node.prev.next = node.next
	node.next.prev = node.prev

	node.next = &l.sentinel
	node.prev = l.sentinel.prev

	node.prev.next = node
	node.next.prev = node
}

func (l *LinkedList[K, V]) Remove(node *Node[K, V]) {
	l.len--

	node.prev.next = node.next
	node.next.prev = node.prev

	node.next = nil
	node.prev = nil

}

func (l *LinkedList[K, V]) Clear() {
	if l.Length() == 0 {
		return
	}

	for node := l.sentinel.next; node != &l.sentinel; {
		next := node.next
		node.prev = nil
		node.next = nil
		node = next
	}
	l.sentinel.next = &l.sentinel
	l.sentinel.prev = &l.sentinel
	l.len = 0
}

/*
func (l *LinkedList[K, V]) PopTail() *Node[K, V] {
	if l.len == 0 {
		return nil
	}
	l.len--

	tail := l.Tail()

	tail.prev.next = l.sentinel
	tail.next.prev = tail.prev

	return tail
}

func (l *LinkedList[K, V]) PopHead() *Node[K, V] {
	if l.len == 0 {
		return nil
	}
	l.len--

	head := l.Head()

	head.next.prev = l.sentinel
	head.prev.next = head.next

	return head
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
*/
