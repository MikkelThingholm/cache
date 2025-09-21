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

func Test_PushHeadAndTail(t *testing.T) {
	l := NewLinkedList[int, int]()

	n1 := Node[int, int]{Value: 1}
	n2 := Node[int, int]{Value: 2}
	n3 := Node[int, int]{Value: 3}
	n4 := Node[int, int]{Value: 4}
	n5 := Node[int, int]{Value: 5}

	l.PushHead(&n1)
	l.PushTail(&n2)
	l.PushTail(&n3)
	l.PushHead(&n4)
	l.PushHead(&n5)

	got := traverseValues(*l)
	want := []int{5, 4, 1, 2, 3}

	if !slices.Equal(got, want) {
		t.Errorf("got: %v, want: %v", got, want)
	}

	if l.Length() != 5 {
		t.Errorf("got: %d, want: %d", l.Length(), 5)
	}

}

func Test_RemovingNodes(t *testing.T) {
	l := NewLinkedList[int, int]()

	n1 := Node[int, int]{Value: 1}
	n2 := Node[int, int]{Value: 2}
	n3 := Node[int, int]{Value: 3}
	n4 := Node[int, int]{Value: 4}

	l.PushHead(&n1)
	l.PushHead(&n2)
	l.Remove(&n1)
	l.Remove(&n2)

	if got, want := l.Length(), 0; got != want {
		t.Errorf("Length: got %v, want %v", got, want)
	}

	if n1.prev != nil || n1.next != nil {
		t.Errorf("n1.prev = %p, n1.next = %p; both should be nil", n1.prev, n1.next)
	}
	if n2.prev != nil || n2.next != nil {
		t.Errorf("n2.prev = %p, n2.next = %p; both should be nil", n2.prev, n2.next)
	}

	l.PushHead(&n3)
	l.PushHead(&n4)

	if got, want := traverseValues(*l), []int{4, 3}; !slices.Equal(got, want) {
		t.Errorf("got: %v, want: %v", got, want)
	}

	if got, want := l.Length(), 2; got != want {
		t.Errorf("got: %v, want: %v", got, want)
	}

}

func Test_MoveToHead(t *testing.T) {
	l := NewLinkedList[int, int]()

	n1 := Node[int, int]{Value: 1}
	n2 := Node[int, int]{Value: 2}
	n3 := Node[int, int]{Value: 3}
	n4 := Node[int, int]{Value: 4}

	l.PushTail(&n1)
	l.MoveToHead(&n1)

	l.PushTail(&n2)
	l.MoveToHead(&n2)

	l.PushTail(&n3)
	l.PushTail(&n4)

	l.MoveToHead(&n4)
	l.MoveToHead(&n3)

	if got, want := traverseValues(*l), []int{3, 4, 2, 1}; !slices.Equal(got, want) {
		t.Errorf("got: %v, want: %v", got, want)
	}
	if got, want := l.Length(), 4; got != want {
		t.Errorf("got: %v, want: %v", got, want)
	}
}

func Test_Clear(t *testing.T) {
	l := NewLinkedList[int, int]()

	n1 := Node[int, int]{Value: 1}
	n2 := Node[int, int]{Value: 2}

	l.PushTail(&n1)
	l.PushTail(&n2)

	l.Clear()

	if got, want := l.Length(), 0; got != want {
		t.Errorf("Length: got %v, want %v", got, want)
	}
	if l.Head() != nil {
		t.Errorf("l.Head() = %p, want nil", l.Head())
	}
	if l.Tail() != nil {
		t.Errorf("l.Tail() = %p, want nil", l.Tail())
	}

	if n1.prev != nil || n1.next != nil {
		t.Errorf("n1.prev = %p, n1.next = %p; both should be nil", n1.prev, n1.next)
	}
	if n2.prev != nil || n2.next != nil {
		t.Errorf("n2.prev = %p, n2.next = %p; both should be nil", n2.prev, n2.next)
	}

}
