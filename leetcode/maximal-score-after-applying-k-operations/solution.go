package maximalscoreafterapplyingkoperations

import (
	"container/heap"
	"sort"
)

type BigHeap struct {
	sort.IntSlice
}

func (h BigHeap) Less(i, j int) bool {
	return h.IntSlice[i] > h.IntSlice[j]
}

// Pop implements heap.Interface.
func (*BigHeap) Pop() (_ any) {
	return
}

// Push implements heap.Interface.
func (*BigHeap) Push(x any) {
}

var _ heap.Interface = &BigHeap{}

func maxKelements(nums []int, k int) int64 {
	h := BigHeap{nums}
	heap.Init(&h)
	var ans int64
	for i := 0; i < k; i++ {
		ans += int64(h.IntSlice[0])
		h.IntSlice[0] = (h.IntSlice[0] + 2) / 3
		heap.Fix(&h, 0)
	}
	return ans
}
