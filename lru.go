package main

import "fmt"

type LRU[K comparable, V any] struct {
	hashMap      map[K]*Node[K, V]
	evictionList *LinkedList[K, V]
	maxEntries   int
}

func NewLRU[K comparable, V any](maxEntries int) (*LRU[K, V], error) {
	if maxEntries <= 0 {
		return nil, fmt.Errorf("maxEntries have to be greater than 0")
	}

	list := NewLinkedList[K, V]()
	lru := &LRU[K, V]{
		hashMap:      map[K]*Node[K, V]{},
		evictionList: list,
		maxEntries:   maxEntries}

	return lru, nil
}

func (lru LRU[K, V]) Count() int {
	return lru.evictionList.Length()
}

func (lru *LRU[K, V]) Insert(key K, value V) error {
	_, ok := lru.hashMap[key]
	if ok {
		return fmt.Errorf("key already exists")
	}
	node := &Node[K, V]{Key: key, Value: value}
	lru.hashMap[key] = node
	if lru.evictionList.Length() == lru.maxEntries {
		tail := lru.evictionList.Tail()
		lru.evictionList.Remove(tail)
		delete(lru.hashMap, tail.Key)
	}
	lru.evictionList.PushHead(node)
	return nil
}

func (lru *LRU[K, V]) Get(key K) (Node[K, V], error) {
	node, ok := lru.hashMap[key]
	if !ok {
		return Node[K, V]{}, fmt.Errorf("key not found")
	}
	lru.evictionList.MoveToHead(node)
	return *node, nil
}
