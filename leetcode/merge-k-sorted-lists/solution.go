// Package merge_k_sorted_lists ...
// https://leetcode-cn.com/problems/merge-k-sorted-lists/
package merge_k_sorted_lists

type ListNode struct {
	Val  int
	Next *ListNode
}

type MinHeap struct {
	Arr      []*ListNode
	Size     int
	Capacity int
}

func (mh *MinHeap) getParent(i int) int {
	return (i - 1) / 2
}

func (mh *MinHeap) getLeft(i int) int {
	return i*2 + 1
}

func (mh *MinHeap) getRight(i int) int {
	return i*2 + 2
}

func (mh *MinHeap) Heapify() {
	i := 0
	for i < mh.Size {
		min := i
		if l := mh.getLeft(i); l < mh.Size && mh.Arr[l].Val < mh.Arr[min].Val {
			min = l
		}
		if r := mh.getRight(i); r < mh.Size && mh.Arr[r].Val < mh.Arr[min].Val {
			min = r
		}
		if min != i {
			mh.Arr[min], mh.Arr[i] = mh.Arr[i], mh.Arr[min]
			i = min
		} else {
			break
		}
	}
}

func (mh *MinHeap) PopMin() *ListNode {
	min := mh.Arr[0]
	mh.Arr[0] = mh.Arr[mh.Size-1]
	mh.Size--
	mh.Heapify()
	return min
}

func (mh *MinHeap) Add(l *ListNode) {
	mh.Arr[mh.Size] = l
	i := mh.Size
	mh.Size++
	for i > 0 {
		p := mh.getParent(i)
		if mh.Arr[i].Val < mh.Arr[p].Val {
			mh.Arr[i], mh.Arr[p] = mh.Arr[p], mh.Arr[i]
			i = p
		} else {
			break
		}
	}
}

func NewMinHeap(cap int) *MinHeap {
	return &MinHeap{
		Arr:      make([]*ListNode, cap),
		Size:     0,
		Capacity: cap,
	}
}

func mergeKLists(lists []*ListNode) *ListNode {
	cap := len(lists)
	mh := NewMinHeap(cap)
	for _, l := range lists {
		if l != nil {
			mh.Add(l)
		}
	}
	r := &ListNode{
		Val:  0,
		Next: nil,
	}
	p := r
	for mh.Size != 0 {
		min := mh.PopMin()
		p.Next = min
		p = p.Next
		if min.Next != nil {
			mh.Add(min.Next)
		}
	}
	return r.Next
}
