package zui_xiao_de_kge_shu_lcof

import (
	"errors"
	"fmt"
)

func getLeastNumbers(arr []int, k int) []int {
	mh := New(len(arr))
	for _, v := range arr {
		mh.Insert(v)
	}
	r := []int{}
	for i := 0; i < k; i++ {
		min, _ := mh.ExtractMin()
		r = append(r, min)
	}
	return r
}

type MinHeap struct {
	Heap []int
	Size int
}

func New(cap int) *MinHeap {
	return &MinHeap{
		Heap: make([]int, cap),
		Size: 0,
	}
}

func parent(i int) int {
	return (i - 1) / 2
}

func left(i int) int {
	return i*2 + 1
}

func right(i int) int {
	return i*2 + 2
}

func (mh *MinHeap) MinHeapify(i int) {
	l := left(i)
	r := right(i)
	smallest := i
	if l < mh.Size && mh.Heap[l] < mh.Heap[smallest] {
		smallest = l
	}
	if r < mh.Size && mh.Heap[r] < mh.Heap[smallest] {
		smallest = r
	}
	if i != smallest {
		mh.Heap[i], mh.Heap[smallest] = mh.Heap[smallest], mh.Heap[i]
		mh.MinHeapify(smallest)
	}
}

func (mh *MinHeap) Insert(v int) error {
	if mh.Size == len(mh.Heap) {
		return errors.New(fmt.Sprintf("cannot insert %v", v))
	}
	mh.Heap[mh.Size] = v
	i := mh.Size
	for i != 0 && mh.Heap[i] < mh.Heap[parent(i)] {
		mh.Heap[i], mh.Heap[parent(i)] = mh.Heap[parent(i)], mh.Heap[i]
		i = parent(i)
	}
	mh.Size++
	return nil
}

func (mh *MinHeap) ExtractMin() (int, error) {
	if mh.Size == 0 {
		return 0, errors.New(fmt.Sprintf("cannot ExtractMin"))
	}
	if mh.Size == 1 {
		mh.Size--
		return mh.Heap[0], nil
	}
	min := mh.Heap[0]
	mh.Heap[0] = mh.Heap[mh.Size-1]
	mh.Size--
	mh.MinHeapify(0)
	return min, nil
}

// 快排
// func getLeastNumbers(arr []int, k int) []int {
// 	sort.Slice(arr, func(i, j int) bool {
// 		return arr[i] < arr[j]
// 	})
// 	return arr[:k]
// }
