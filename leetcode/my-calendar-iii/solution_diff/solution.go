package diff

import "github.com/emirpasic/gods/trees/redblacktree"

type MyCalendarThree struct {
	RBT *redblacktree.Tree
}

func Constructor() MyCalendarThree {
	return MyCalendarThree{
		RBT: redblacktree.NewWithIntComparator(),
	}
}

func (mc *MyCalendarThree) add(key int, value int) {
	if v, ok := mc.RBT.Get(key); ok {
		mc.RBT.Put(key, v.(int)+value)
	} else {
		mc.RBT.Put(key, value)
	}
}

func (mc *MyCalendarThree) Book(startTime int, endTime int) int {
	mc.add(startTime, 1)
	mc.add(endTime, -1)
	it := mc.RBT.Iterator()
	ans := 0
	count := 0
	for it.Next() {
		count += it.Value().(int)
		if count > ans {
			ans = count
		}
	}
	return ans
}
