package solutiondiffarray

import "github.com/emirpasic/gods/trees/redblacktree"

type MyCalendarTwo struct {
	T *redblacktree.Tree
}

func Constructor() MyCalendarTwo {
	return MyCalendarTwo{
		T: redblacktree.NewWithIntComparator(),
	}
}

func (mc *MyCalendarTwo) add(key, value int) {
	if v, ok := mc.T.Get(key); ok {
		mc.T.Put(key, v.(int)+value)
	} else {
		mc.T.Put(key, value)
	}
}

func (mc *MyCalendarTwo) Book(start int, end int) bool {
	mc.add(start, 1)
	mc.add(end, -1)

	it := mc.T.Iterator()
	maxbook := 0

	for it.Next() {
		maxbook += it.Value().(int)
		if maxbook > 2 {
			mc.add(start, -1)
			mc.add(end, 1)
			return false
		}
	}
	return true
}
