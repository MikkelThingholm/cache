package main

type LinkedList[K comparable, V any] struct {
	Sentinel *Node[K, V]
	len      int
}

type Node[K comparable, V any] struct {
	key   K
	value V
	next  *Node[K, V]
	prev  *Node[K, V]
}

func NewLinkedList[K comparable, V any]() *LinkedList[K, V] {
	sentinel := Node[K, V]{}
	sentinel.next = &sentinel
	sentinel.prev = &sentinel
	return &LinkedList[K, V]{
		Sentinel: &sentinel,
		len:      0,
	}
}

func (l *LinkedList[K, V]) Head() *Node[K, V] {
	if l.len == 0 {
		return nil
	}
	return l.Sentinel.next
}

func (l *LinkedList[K, V]) Tail() *Node[K, V] {
	if l.len == 0 {
		return nil
	}
	return l.Sentinel.prev
}

func (l *LinkedList[K, V]) AddToHead(key K, value V) {
	l.len++
	node := Node[K, V]{
		key:   key,
		value: value,
		next:  l.Sentinel.next,
		prev:  l.Sentinel,
	}
	node.prev.next = &node
	node.next.prev = &node
}

func (l *LinkedList[K, V]) AddToTail(key K, value V) {
	l.len++
	node := Node[K, V]{
		key:   key,
		value: value,
		next:  l.Sentinel,
		prev:  l.Sentinel.prev,
	}
	node.next.prev = &node
	node.prev.next = &node
}

func (l *LinkedList[K, V]) PopTail() *Node[K, V] {
	if l.len == 0 {
		return nil
	}
	l.len--

	tail := l.Tail()

	tail.prev.next = l.Sentinel
	tail.next.prev = tail.prev

	return tail
}

func (l *LinkedList[K, V]) PopHead() *Node[K, V] {
	if l.len == 0 {
		return nil
	}
	l.len--

	head := l.Head()

	head.next.prev = l.Sentinel
	head.prev.next = head.next

	return head
}

/*
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
