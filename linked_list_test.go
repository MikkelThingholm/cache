package main

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

/*
func TestA(t *testing.T) {
	l := setupLinkedList(5)
	t.Errorf("%+v", l)
	v := traverseLinkedList(*l)
	t.Fatalf("Test: %v", v)
}
*/
