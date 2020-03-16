// Package lru_cache ...
// https://leetcode-cn.com/problems/lru-cache/
package lru_cache

type node struct {
	Key  int
	Val  int
	Prev *node
	Next *node
}

type LRUCache struct {
	NodesMap map[int]*node
	Head     *node
	Tail     *node
	Len      int
	Capacity int
}

func Constructor(capacity int) LRUCache {
	head := &node{}
	tail := &node{}
	head.Next = tail
	tail.Prev = head
	return LRUCache{
		NodesMap: map[int]*node{},
		Head:     head,
		Tail:     tail,
		Capacity: capacity,
	}
}

func (this *LRUCache) add(n *node) {
	n.Next = this.Head.Next
	n.Prev = this.Head
	this.Head.Next.Prev = n
	this.Head.Next = n
	this.NodesMap[n.Key] = n
	this.Len++
}

func (this *LRUCache) remove(n *node) {
	if this.Len == 0 {
		return
	}
	prev := n.Prev
	next := n.Next
	prev.Next = next
	next.Prev = prev
	delete(this.NodesMap, n.Key)
	this.Len--
}

func (this *LRUCache) pop() *node {
	if this.Len == 0 {
		return nil
	}
	n := this.Tail.Prev
	this.remove(n)
	return n
}

func (this *LRUCache) Get(key int) int {
	if n, ok := this.NodesMap[key]; ok {
		this.remove(n)
		this.add(n)
		return n.Val
	}
	return -1
}

func (this *LRUCache) Put(key int, value int) {
	if n, ok := this.NodesMap[key]; ok {
		n.Val = value
		this.remove(n)
		this.add(n)
		return
	}
	n := &node{
		Key: key,
		Val: value,
	}
	if this.Len == this.Capacity {
		this.pop()
		this.add(n)
	} else {
		this.add(n)
	}
}

// type LRUCache struct {
// 	ElementsMap map[int]*list.Element
// 	List        *list.List
// 	Capacity    int
// }
//
// type Item struct {
// 	Key int
// 	Val int
// }
//
// func Constructor(capacity int) LRUCache {
// 	return LRUCache{
// 		map[int]*list.Element{},
// 		list.New(),
// 		capacity,
// 	}
// }
//
// func (this *LRUCache) Get(key int) int {
// 	if e, ok := this.ElementsMap[key]; ok {
// 		v, _ := e.Value.(Item)
// 		this.List.Remove(e)
// 		this.List.PushBack(v)
// 		this.ElementsMap[key] = this.List.Back()
// 		return v.Val
// 	}
// 	return -1
// }
//
// func (this *LRUCache) Put(key int, value int) {
// 	if e, ok := this.ElementsMap[key]; ok {
// 		delete(this.ElementsMap, key)
// 		this.List.Remove(e)
// 		this.List.PushBack(Item{
// 			key,
// 			value,
// 		})
// 		this.ElementsMap[key] = this.List.Back()
// 		return
// 	}
// 	if this.Capacity == this.List.Len() {
// 		e := this.List.Front()
// 		v, _ := e.Value.(Item)
// 		item := Item{
// 			key,
// 			value,
// 		}
// 		delete(this.ElementsMap, v.Key)
// 		this.List.Remove(e)
// 		this.List.PushBack(item)
// 		this.ElementsMap[key] = this.List.Back()
// 	} else {
// 		item := Item{
// 			key,
// 			value,
// 		}
// 		this.List.PushBack(item)
// 		this.ElementsMap[key] = this.List.Back()
// 	}
// }

/**
 * Your LRUCache object will be instantiated and called as such:
 * obj := Constructor(capacity);
 * param_1 := obj.Get(key);
 * obj.Put(key,value);
 */
