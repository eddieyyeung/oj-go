package solution_3

import (
	"errors"
	"fmt"
)

// maxheap O(nlogk)
func getLeastNumbers(arr []int, k int) []int {
	mh := New(k)
	for i := 0; i < k; i++ {
		mh.Insert(arr[i])
	}
	for i := k; i < len(arr); i++ {
		max, _ := mh.GetMax()
		if max > arr[i] {
			mh.ExtractMax()
			mh.Insert(arr[i])
		}
	}
	return mh.Heap
}

type MaxHeap struct {
	Heap []int
	Size int
}

func New(cap int) *MaxHeap {
	return &MaxHeap{
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

func (mh *MaxHeap) MinHeapify(i int) {
	l := left(i)
	r := right(i)
	smallest := i
	if l < mh.Size && mh.Heap[l] > mh.Heap[smallest] {
		smallest = l
	}
	if r < mh.Size && mh.Heap[r] > mh.Heap[smallest] {
		smallest = r
	}
	if i != smallest {
		mh.Heap[i], mh.Heap[smallest] = mh.Heap[smallest], mh.Heap[i]
		mh.MinHeapify(smallest)
	}
}

func (mh *MaxHeap) Insert(v int) error {
	if mh.Size == len(mh.Heap) {
		return errors.New(fmt.Sprintf("cannot insert %v", v))
	}
	mh.Heap[mh.Size] = v
	i := mh.Size
	for i != 0 && mh.Heap[i] > mh.Heap[parent(i)] {
		mh.Heap[i], mh.Heap[parent(i)] = mh.Heap[parent(i)], mh.Heap[i]
		i = parent(i)
	}
	mh.Size++
	return nil
}

func (mh *MaxHeap) GetMax() (int, error) {
	if mh.Size == 0 {
		return 0, errors.New(fmt.Sprintf("size is zero"))
	}
	return mh.Heap[0], nil
}

func (mh *MaxHeap) ExtractMax() (int, error) {
	if mh.Size == 0 {
		return 0, errors.New(fmt.Sprintf("cannot ExtractMax"))
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
