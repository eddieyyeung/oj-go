package solution

import "container/heap"

var _ heap.Interface = &IntHeap{}

type IntHeap []int

func (h IntHeap) Len() int {
	return len(h)
}

func (h IntHeap) Less(i, j int) bool {
	return h[i] < h[j]
}

func (h IntHeap) Swap(i, j int) {
	h[i], h[j] = h[j], h[i]
}

func (h *IntHeap) Push(x interface{}) {
	*h = append(*h, x.(int))
}

func (h *IntHeap) Pop() interface{} {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[0 : n-1]
	return x
}

type SmallestInfiniteSet struct {
	m map[int]struct{}
	h IntHeap
}

func Constructor() SmallestInfiniteSet {
	h := make(IntHeap, 1000)
	m := make(map[int]struct{})
	for i := 0; i < 1000; i++ {
		h[i] = i + 1
		m[i+1] = struct{}{}
	}
	heap.Init(&h)
	return SmallestInfiniteSet{
		m: m,
		h: h,
	}
}

func (this *SmallestInfiniteSet) PopSmallest() int {
	v := heap.Pop(&this.h).(int)
	delete(this.m, v)
	return v
}

func (this *SmallestInfiniteSet) AddBack(num int) {
	if _, ok := this.m[num]; ok {
		return
	} else {
		heap.Push(&this.h, num)
		this.m[num] = struct{}{}
	}
}

/**
 * Your SmallestInfiniteSet object will be instantiated and called as such:
 * obj := Constructor();
 * param_1 := obj.PopSmallest();
 * obj.AddBack(num);
 */
