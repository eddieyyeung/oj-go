package solution

import "container/heap"

type Item struct {
	Val  int
	KIdx int
}

type MinHeap []Item

// Len implements heap.Interface.
func (m MinHeap) Len() int {
	return len(m)
}

// Less implements heap.Interface.
func (m MinHeap) Less(i int, j int) bool {
	return m[i].Val < m[j].Val
}

// Pop implements heap.Interface.
func (m *MinHeap) Pop() any {
	old := *m
	n := len(old)
	x := old[n-1]
	*m = old[:n-1]
	return x
}

// Push implements heap.Interface.
func (m *MinHeap) Push(x any) {
	*m = append(*m, x.(Item))
}

// Swap implements heap.Interface.
func (m MinHeap) Swap(i int, j int) {
	m[i], m[j] = m[j], m[i]
}

var _ heap.Interface = &MinHeap{}

// https://leetcode.cn/problems/smallest-range-covering-elements-from-k-lists/description/
func smallestRange(nums [][]int) []int {
	k := len(nums)
	t := make([]int, k)
	var max_num int = -1e9

	var mh MinHeap
	for i := 0; i < k; i++ {
		heap.Push(&mh, Item{
			Val:  nums[i][0],
			KIdx: i,
		})
		if nums[i][0] > max_num {
			max_num = nums[i][0]
		}
	}
	ans := []int{-1e9, 1e9}
	for {
		min_item := heap.Pop(&mh).(Item)
		if diff := max_num - min_item.Val; diff < ans[1]-ans[0] {
			ans = []int{min_item.Val, max_num}
		}
		if t[min_item.KIdx] == len(nums[min_item.KIdx])-1 {
			break
		}
		t[min_item.KIdx]++
		if num := nums[min_item.KIdx][t[min_item.KIdx]]; num > max_num {
			max_num = num
		}
		heap.Push(&mh, Item{
			Val:  nums[min_item.KIdx][t[min_item.KIdx]],
			KIdx: min_item.KIdx,
		})
	}
	return ans
}
