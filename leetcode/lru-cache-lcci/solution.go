// Package lru_cache_lcci ...
// https://leetcode-cn.com/problems/lru-cache-lcci/
package lru_cache_lcci

type Node struct {
	Key   int
	Value int
	Prev  *Node
	Next  *Node
}

type LRUCache struct {
	M    map[int]*Node
	Cap  int
	Size int
	Head *Node
	Tail *Node
}

func Constructor(capacity int) LRUCache {
	head := &Node{
		Key:   0,
		Value: 0,
		Prev:  nil,
		Next:  nil,
	}
	tail := &Node{
		Key:   0,
		Value: 0,
		Prev:  nil,
		Next:  nil,
	}
	head.Next = tail
	tail.Prev = head
	return LRUCache{
		M:    map[int]*Node{},
		Cap:  capacity,
		Size: 0,
		Head: head,
		Tail: tail,
	}
}

func (lru *LRUCache) deleteNode(node *Node) {
	node.Prev.Next = node.Next
	node.Next.Prev = node.Prev
	node.Next = nil
	node.Prev = nil
}

func (lru *LRUCache) putNodeToHead(node *Node) {
	node.Next = lru.Head.Next
	node.Prev = lru.Head
	lru.Head.Next.Prev = node
	lru.Head.Next = node
}

func (lru *LRUCache) Get(key int) int {
	node, ok := lru.M[key]
	if !ok {
		return -1
	}
	lru.deleteNode(node)
	lru.putNodeToHead(node)
	return node.Value
}

func (lru *LRUCache) Put(key int, value int) {
	if lru.Cap == 0 {
		return
	}
	node, ok := lru.M[key]
	if !ok {
		node = &Node{
			Key:   key,
			Value: value,
			Prev:  nil,
			Next:  nil,
		}
		if lru.Size >= lru.Cap {
			delete(lru.M, lru.Tail.Prev.Key)
			lru.deleteNode(lru.Tail.Prev)
		} else {
			lru.Size++
		}
		lru.M[key] = node
	} else {
		node.Value = value
		lru.deleteNode(node)
	}
	lru.putNodeToHead(node)
}

/**
 * Your LRUCache object will be instantiated and called as such:
 * obj := Constructor(capacity);
 * param_1 := obj.Get(key);
 * obj.Put(key,value);
 */
