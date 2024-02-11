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
	*h = old[:n-1]
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
	var topk MinHeap
	for _, num := range nums {
		if topk.Len() == k {
			if num > topk[0] {
				topk[0] = num
				heap.Fix(&topk, 0)
			}
		} else {
			heap.Push(&topk, num)
		}
	}
	return topk[0]
}
