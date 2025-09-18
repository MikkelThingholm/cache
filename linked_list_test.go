package main

import (
	"slices"
	"testing"
)

/*
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
*/

func traverseLinkedList(l LinkedList[int, int]) ([]int, []int) {
	node := l.Head()
	if node == nil {
		return []int{}, []int{}
	}

	keys := []int{}
	values := []int{}

	for i := 0; i < l.len; i++ {
		keys = append(keys, node.key)
		values = append(values, node.value)
		node = node.next
	}

	return keys, values
}

func TestAddingToHeadProducesValidLinkedList(t *testing.T) {
	l := NewLinkedList[int, int]()

	for i := 5; i >= 0; i-- {
		l.AddToHead(i, i)
	}

	got_keys, got_values := traverseLinkedList(*l)
	want := []int{0, 1, 2, 3, 4, 5}

	if !slices.Equal(want, got_keys) {
		t.Errorf("want: %v, got: %v", want, got_keys)
	}

	if !slices.Equal(want, got_values) {
		t.Errorf("want: %v, got: %v", want, got_values)
	}

	if l.len != 6 {
		t.Errorf("want: %d, got: %d", 6, l.len)
	}

}
