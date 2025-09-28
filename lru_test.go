package main

import (
	"testing"
	"testing/synctest"
	"time"
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

	if _, ok := lru.Get(1); ok {
		t.Errorf("expected key 1 to be envicted but was found")
	}
	if got, want := lru.Count(), 3; got != want {
		t.Errorf("count: got %d, want %d", got, want)
	}

	node, ok := lru.Get(2)
	if !ok {
		t.Errorf("key 2 not found")
	}
	if want := 2; node.Key != want {
		t.Errorf("node.Key: got %d, want %d", node.Key, want)
	}

	lru.Insert(5, 5)
	_, ok = lru.Get(3)
	if ok {
		t.Errorf("found key 3")
	}
	if got, want := lru.Count(), 3; got != want {
		t.Errorf("count: got %d, want %d", got, want)
	}

}

func Test_InsertingExistingKeyUpdatesValue(t *testing.T) {
	lru, err := NewLRU[int, int](3)
	if err != nil {
		t.Fatalf("error while initializing LRU: %s", err)
	}
	lru.Insert(1, 1)
	lru.Insert(1, 2)

	node, ok := lru.Get(1)
	if !ok {
		t.Errorf("key 1 not found")
	}
	if got, want := node.Value, 2; got != want {
		t.Errorf("got %d, want %d", got, want)
	}

}

func Test_ExpiredItemsAreEvicted(t *testing.T) {
	synctest.Test(t, func(t *testing.T) {
		lru, err := NewLRU[int, int](4)
		defer lru.Stop()
		if err != nil {
			t.Fatalf("error while initializing LRU: %s", err)
		}

		lru.Insert(1, 1)
		lru.Insert(2, 2)

		time.Sleep(10 * time.Minute)

		lru.Insert(3, 3)
		lru.Insert(4, 4)
		if lru.Count() != 4 {
			t.Errorf("Expected lru count to be 4, got %d", lru.Count())
		}

		time.Sleep(50 * time.Minute)
		synctest.Wait()

		if lru.Count() != 2 {
			t.Errorf("Expected lru count to be 2, got %d", lru.Count())
		}
		if _, found := lru.Get(1); found != false {
			t.Errorf("Expected key 1 to be deleted")
		}
		if _, found := lru.Get(2); found != false {
			t.Errorf("Expected key 2 to be deleted")
		}

		time.Sleep(10 * time.Minute)
		synctest.Wait()

		if lru.Count() != 0 {
			t.Errorf("Expected lru count to be 0, got %d", lru.Count())
		}
		if _, found := lru.Get(3); found != false {
			t.Errorf("Expected key 1 to be deleted")
		}
		if _, found := lru.Get(4); found != false {
			t.Errorf("Expected key 2 to be deleted")
		}
	})
}
