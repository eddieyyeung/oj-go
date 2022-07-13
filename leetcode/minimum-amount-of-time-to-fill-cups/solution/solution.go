package solution

import "container/heap"

var _ heap.Interface = &IntHeap{}

type IntHeap []int

func (h IntHeap) Len() int {
	return len(h)
}

func (h IntHeap) Less(i, j int) bool {
	return h[i] > h[j]
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

func fillCups(amount []int) int {
	var h IntHeap
	for i := 0; i < len(amount); i++ {
		if amount[i] != 0 {
			h = append(h, amount[i])
		}
	}
	heap.Init(&h)
	count := 0
	for len(h) != 0 {
		a := heap.Pop(&h).(int)
		if len(h) == 0 {
			count += a
			return count
		} else {
			count++
			b := heap.Pop(&h).(int)
			a -= 1
			if a != 0 {
				heap.Push(&h, a)
			}
			b -= 1
			if b != 0 {
				heap.Push(&h, b)
			}
		}
	}
	return count
}
