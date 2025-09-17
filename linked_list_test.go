package main

import (
	"slices"
	"testing"
)

func setupLinkedList(nodeNumber int) *LinkedList {
	if nodeNumber < 0 {
		return nil
	}
	linkedList := LinkedList{}

	for i := range nodeNumber {
		linkedList.AddToTail(Node{value: i})
	}
	return &linkedList
}

func traverseLinkedList(l LinkedList) []int {
	node := l.head
	if node == nil {
		return []int{}
	}

	value, _ := node.value.(int)
	values := []int{value}

	for node.next != nil {
		node = node.next

		value, _ = node.value.(int)
		values = append(values, value)
	}

	return values
}

func TestAddingToHeadProducesValidLinkedList(t *testing.T) {
	l := LinkedList{}

	for i := 5; i >= 0; i-- {
		l.AddToHead(i)
	}

	got := traverseLinkedList(l)
	want := []int{0, 1, 2, 3, 4, 5}

	if !slices.Equal(want, got) {
		t.Errorf("want: %v, got: %v", want, got)
	}

	if l.size != 6 {
		t.Errorf("want: %d, got: %d", 6, l.size)
	}

}
