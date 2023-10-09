package solutionredblacktree

import (
	"math"

	"github.com/emirpasic/gods/trees/redblacktree"
)

// 执行用时 72 ms 击败 82.86%
// 消耗内存 6.92 MB 击败 88.57%
type MyCalendar struct {
	*redblacktree.Tree
}

func Constructor() MyCalendar {
	t := redblacktree.NewWithIntComparator()
	t.Put(math.MaxInt32, nil)
	return MyCalendar{t}
}

func (mc *MyCalendar) Book(start int, end int) bool {
	node, _ := mc.Ceiling(end)
	it := mc.IteratorAt(node)
	if !it.Prev() || it.Value().(int) <= start {
		mc.Put(start, end)
		return true
	}
	return false
}
