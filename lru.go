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

	if node, ok := lru.get(key); ok {
		lru.update(node, value)
		return
	}

	lru.add(key, value)
}

func (lru *LRU[K, V]) Get(key K) (Node[K, V], error) {
	node, ok := lru.get(key)
	if !ok {
		return Node[K, V]{}, fmt.Errorf("key not found")
	}
	return *node, nil
}

func (lru *LRU[K, V]) add(key K, value V) bool {

	if _, ok := lru.hashMap[key]; ok {
		return false
	}

	node := &Node[K, V]{
		Key:   key,
		Value: value,
	}

	if lru.evictionList.Length() == lru.cap {
		lru.remove(lru.evictionList.Tail())
	}

	lru.hashMap[key] = node
	lru.evictionList.PushHead(node)
	return true
}

func (lru *LRU[K, V]) remove(node *Node[K, V]) {
	lru.evictionList.Remove(node)
	delete(lru.hashMap, node.Key)
}

func (lru *LRU[K, V]) get(key K) (*Node[K, V], bool) {
	node, ok := lru.hashMap[key]
	if !ok {
		return nil, false
	}
	lru.evictionList.MoveToHead(node)

	return node, true
}

func (lru *LRU[K, V]) update(node *Node[K, V], newValue V) {
	node.Value = newValue
	lru.evictionList.MoveToHead(node)
}
