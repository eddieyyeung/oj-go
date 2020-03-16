// Package lru_cache ...
// https://leetcode-cn.com/problems/lru-cache/
package lru_cache

import "container/list"

type LRUCache struct {
	ElementsMap map[int]*list.Element
	List        *list.List
	Capacity    int
}

type Item struct {
	Key int
	Val int
}

func Constructor(capacity int) LRUCache {
	return LRUCache{
		map[int]*list.Element{},
		list.New(),
		capacity,
	}
}

func (this *LRUCache) Get(key int) int {
	if e, ok := this.ElementsMap[key]; ok {
		v, _ := e.Value.(Item)
		this.List.Remove(e)
		this.List.PushBack(v)
		this.ElementsMap[key] = this.List.Back()
		return v.Val
	}
	return -1
}

func (this *LRUCache) Put(key int, value int) {
	if e, ok := this.ElementsMap[key]; ok {
		delete(this.ElementsMap, key)
		this.List.Remove(e)
		this.List.PushBack(Item{
			key,
			value,
		})
		this.ElementsMap[key] = this.List.Back()
		return
	}
	if this.Capacity == this.List.Len() {
		e := this.List.Front()
		v, _ := e.Value.(Item)
		item := Item{
			key,
			value,
		}
		delete(this.ElementsMap, v.Key)
		this.List.Remove(e)
		this.List.PushBack(item)
		this.ElementsMap[key] = this.List.Back()
	} else {
		item := Item{
			key,
			value,
		}
		this.List.PushBack(item)
		this.ElementsMap[key] = this.List.Back()
	}
}

/**
 * Your LRUCache object will be instantiated and called as such:
 * obj := Constructor(capacity);
 * param_1 := obj.Get(key);
 * obj.Put(key,value);
 */
