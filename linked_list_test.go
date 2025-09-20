package main

import (
	"slices"
	"testing"
)

func traverseValues(l LinkedList[int, int]) []int {
	node := l.Head()
	if node == nil {
		return []int{}
	}

	result := []int{}

	for i := 0; i < l.Length(); i++ {
		result = append(result, node.Value)
		node = node.next
	}
	return result
}

/*
func backwardsTraverse(l LinkedList[int, int]) []Node[int, int] {
	node := l.Tail()
	if node == nil {
		return []Node[int, int]{}
	}

	result := []Node[int, int]{}

	for i := 0; i < l.Length(); i++ {
		result = append(result, *node)
		node = node.prev
	}
	return result
}
*/

func TestAddingToHead(t *testing.T) {
	l := NewLinkedList[int, int]()

	n0 := Node[int, int]{Value: 0}
	n1 := Node[int, int]{Value: 1}
	n2 := Node[int, int]{Value: 2}

	l.AddToHead(&n2)
	l.AddToHead(&n1)
	l.AddToHead(&n0)

	got := traverseValues(*l)
	want := []int{0, 1, 2}

	if !slices.Equal(got, want) {
		t.Errorf("got: %v, want: %v", got, want)
	}

	if l.Length() != 3 {
		t.Errorf("got: %d, want: %d", l.Length(), 3)
	}

}

func TestAddingToHeadAndTail(t *testing.T) {
	l := NewLinkedList[int, int]()

	n0 := Node[int, int]{Value: 0}
	n1 := Node[int, int]{Value: 1}
	n2 := Node[int, int]{Value: 2}

	l.AddToTail(&n0)
	l.AddToTail(&n1)
	l.AddToTail(&n2)

	got := traverseValues(*l)
	want := []int{0, 1, 2}

	if !slices.Equal(got, want) {
		t.Errorf("got: %v, want: %v", got, want)
	}

	if l.Length() != 3 {
		t.Errorf("got: %d, want: %d", l.Length(), 3)
	}

}
