// 347. 前 K 个高频元素
// https://leetcode.cn/problems/top-k-frequent-elements/description
package solution

import "container/heap"

type MinHeap []int

// Len implements heap.Interface.
func (h MinHeap) Len() int {
	return len(h)
}

// Less implements heap.Interface.
func (h MinHeap) Less(i int, j int) bool {
	return mp[h[i]] > mp[h[j]]
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
var mp map[int]int

func topKFrequent(nums []int, k int) []int {
	mp = make(map[int]int)
	for _, num := range nums {
		mp[num]++
	}
	var h MinHeap
	for k := range mp {
		heap.Push(&h, k)
	}
	var ans []int
	for i := 0; i < k; i++ {
		num := heap.Pop(&h).(int)
		ans = append(ans, num)
	}
	return ans
}
