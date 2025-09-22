package main

import (
	"testing"
)

func TestEviction(t *testing.T) {
	lru, err := NewLRU[int, int](3)
	if err != nil {
		t.Fatalf("error while initializing LRU: %s", err)
	}
	lru.Insert(1, 1)
	lru.Insert(2, 2)
	lru.Insert(3, 3)
	lru.Insert(4, 4)

	_, err = lru.Get(1)
	if err == nil {
		t.Errorf("got nil, want error")
	}
	if got, want := lru.Count(), 3; got != want {
		t.Errorf("count: got %d, want %d", got, want)
	}

	node, err := lru.Get(2)
	if err != nil {
		t.Errorf("got %v, want nil", err)
	}
	if want := 2; node.Key != want {
		t.Errorf("node.Key: got %d, want %d", node.Key, want)
	}

	lru.Insert(5, 5)
	_, err = lru.Get(3)
	if err == nil {
		t.Errorf("got nil, want error")
	}
	if got, want := lru.Count(), 3; got != want {
		t.Errorf("count: got %d, want %d", got, want)
	}

}
