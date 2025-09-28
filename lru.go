package main

import (
	"fmt"
	"sync"
	"time"
)

type LRU[K comparable, V any] struct {
	mu sync.Mutex

	hashMap      map[K]*Node[K, V]
	evictionList *LinkedList[K, V]
	cap          int

	defaultTTL      time.Duration
	cleanupInterval time.Duration

	stop chan struct{}
}

func NewLRU[K comparable, V any](capacity int) (*LRU[K, V], error) {
	if capacity <= 0 {
		return nil, fmt.Errorf("maxEntries have to be greater than 0")
	}

	list := NewLinkedList[K, V]()
	lru := &LRU[K, V]{
		hashMap:         map[K]*Node[K, V]{},
		evictionList:    list,
		cap:             capacity,
		defaultTTL:      time.Hour,
		cleanupInterval: 5 * time.Minute,
		stop:            make(chan struct{}),
	}

	if lru.cleanupInterval > 0 {
		go func() {
			ticker := time.NewTicker(lru.cleanupInterval)
			defer ticker.Stop()
			for {
				select {
				case <-ticker.C:
					lru.removeExpired()
				case <-lru.stop:
					return
				}
			}
		}()
	}

	return lru, nil
}

func (lru *LRU[K, V]) Stop() {
	lru.stop <- struct{}{}
}

func (lru *LRU[K, V]) Count() int {
	return lru.evictionList.Length()
}

func (lru *LRU[K, V]) Insert(key K, value V) {
	lru.mu.Lock()
	defer lru.mu.Unlock()

	if node, ok := lru.get(key); ok {
		lru.update(node, value)
		return
	}

	lru.add(key, value)
}

func (lru *LRU[K, V]) Get(key K) (Node[K, V], bool) {
	lru.mu.Lock()
	defer lru.mu.Unlock()

	node, ok := lru.get(key)
	if !ok {
		return Node[K, V]{}, false
	}
	return *node, true
}

func (lru *LRU[K, V]) add(key K, value V) bool {

	if _, ok := lru.hashMap[key]; ok {
		return false
	}

	node := &Node[K, V]{
		Key:       key,
		Value:     value,
		ExpiresAt: time.Now().UnixMilli() + lru.defaultTTL.Milliseconds(),
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
	node.ExpiresAt = time.Now().UnixMilli() + lru.defaultTTL.Milliseconds()
	lru.evictionList.MoveToHead(node)
}

func (lru *LRU[K, V]) removeExpired() {
	lru.mu.Lock()
	defer lru.mu.Unlock()

	expiredNodes := lru.evictionList.FindExpiresBefore(time.Now().UnixMilli())
	for i := range expiredNodes {
		lru.remove(expiredNodes[i])
	}
}
