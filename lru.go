package main

import (
	"fmt"
)

type LRU[K comparable, V any] struct {
	hashMap      map[K]*Node[K, V]
	evictionList *LinkedList[K, V]
	cap          int
}

func NewLRU[K comparable, V any](capacity int) (*LRU[K, V], error) {
	if capacity <= 0 {
		return nil, fmt.Errorf("maxEntries have to be greater than 0")
	}

	list := NewLinkedList[K, V]()
	lru := &LRU[K, V]{
		hashMap:      map[K]*Node[K, V]{},
		evictionList: list,
		cap:          capacity,
	}

	return lru, nil
}

func (lru LRU[K, V]) Count() int {
	return lru.evictionList.Length()
}

func (lru *LRU[K, V]) Insert(key K, value V) {

	if node, ok := lru.hashMap[key]; ok {
		node.Value = value
		lru.evictionList.MoveToHead(node)
		return
	}

	node := &Node[K, V]{
		Key:   key,
		Value: value,
	}
	lru.hashMap[key] = node
	if lru.evictionList.Length() == lru.cap {
		tail := lru.evictionList.Tail()
		lru.evictionList.Remove(tail)
		delete(lru.hashMap, tail.Key)
	}
	lru.evictionList.PushHead(node)
}

func (lru *LRU[K, V]) Get(key K) (Node[K, V], error) {
	node, ok := lru.hashMap[key]
	if !ok {
		return Node[K, V]{}, fmt.Errorf("key not found")
	}
	lru.evictionList.MoveToHead(node)
	return *node, nil
}
