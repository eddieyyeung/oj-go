// 215. 数组中的第K个最大元素
// https://leetcode.cn/problems/kth-largest-element-in-an-array
package solution

import (
	"container/heap"
)

type MinHeap []int

// Len implements heap.Interface.
func (h MinHeap) Len() int {
	return len(h)
}

// Less implements heap.Interface.
func (h MinHeap) Less(i int, j int) bool {
	return h[i] < h[j]
}

// Pop implements heap.Interface.
func (h *MinHeap) Pop() any {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[0 : n-1]
	return x
}

// Push implements heap.Interface.
func (h *MinHeap) Push(x any) {
	*h = append(*h, x.(int))
}

// Swap implements heap.Interface.
func (h MinHeap) Swap(i int, j int) {
	h[i], h[j] = h[j], h[i]
}

var _ heap.Interface = &MinHeap{}

func findKthLargest(nums []int, k int) int {
	var h MinHeap = nums[:k]
	heap.Init(&h)
	for i := k; i < len(nums); i++ {
		if nums[i] > h[0] {
			heap.Pop(&h)
			heap.Push(&h, nums[i])
		}
	}
	return h[0]
}
