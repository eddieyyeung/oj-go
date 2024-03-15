package solution

import "container/list"

type LRUCache struct {
	m        map[int]*list.Element
	l        *list.List
	capacity int
}

type Item struct {
	Key int
	Val int
}

func Constructor(capacity int) LRUCache {
	return LRUCache{
		m:        make(map[int]*list.Element),
		l:        list.New(),
		capacity: capacity,
	}
}

func (this *LRUCache) Get(key int) int {
	e, ok := this.m[key]
	if ok {
		item := e.Value.(Item)
		this.l.MoveToFront(e)
		return item.Val
	}
	return -1
}

func (this *LRUCache) Put(key int, value int) {
	e, ok := this.m[key]
	if ok {
		this.l.Remove(e)
		this.l.PushFront(Item{key, value})
		this.m[key] = this.l.Front()
		return
	}
	if this.capacity == this.l.Len() {
		e = this.l.Back()
		item := e.Value.(Item)
		delete(this.m, item.Key)
		this.l.Remove(e)
	}
	this.l.PushFront(Item{key, value})
	this.m[key] = this.l.Front()
}

/**
* Your LRUCache object will be instantiated and called as such:
* obj := Constructor(capacity);
* param_1 := obj.Get(key);
* obj.Put(key,value);
 */
