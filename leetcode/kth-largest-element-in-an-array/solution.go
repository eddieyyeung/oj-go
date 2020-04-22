// Package kth_largest_element_in_an_array ...
// https://leetcode-cn.com/problems/kth-largest-element-in-an-array/
package kth_largest_element_in_an_array

import "errors"

type minHeap struct {
	Arr      []int
	Size     int
	Capacity int
}

func New(cap int) *minHeap {
	return &minHeap{
		Arr:      make([]int, cap),
		Size:     0,
		Capacity: cap,
	}
}

func (h *minHeap) getParent(i int) int {
	return (i - 1) / 2
}

func (h *minHeap) getLeft(i int) int {
	return i*2 + 1
}

func (h *minHeap) getRight(i int) int {
	return i*2 + 2
}

func (h *minHeap) insert(n int) error {
	if h.Size == h.Capacity {
		return errors.New("minHeap is full")
	}
	h.Arr[h.Size] = n
	h.Size++
	i := h.Size - 1
	for i != 0 {
		parent := h.getParent(i)
		if h.Arr[parent] > h.Arr[i] {
			h.Arr[parent], h.Arr[i] = h.Arr[i], h.Arr[parent]
		} else {
			break
		}
		i = parent
	}
	return nil
}

func (h *minHeap) removeTop() int {
	top := h.Arr[0]
	h.Arr[0] = h.Arr[h.Size-1]
	h.Size--
	for i := 0; i < h.Size; {
		maxIdx := i
		if left := h.getLeft(i); left < h.Size && h.Arr[left] < h.Arr[maxIdx] {
			maxIdx = left
		}
		if right := h.getRight(i); right < h.Size && h.Arr[right] < h.Arr[maxIdx] {
			maxIdx = right
		}
		h.Arr[i], h.Arr[maxIdx] = h.Arr[maxIdx], h.Arr[i]
		if maxIdx == i {
			break
		}
		i = maxIdx
	}
	return top
}

func (h *minHeap) getTop() int {
	return h.Arr[0]
}

func findKthLargest(nums []int, k int) int {
	h := New(k)
	for i := 0; i < k; i++ {
		h.insert(nums[i])
	}
	for i := k; i < len(nums); i++ {
		min := h.getTop()
		if min > nums[i] {
			continue
		}
		h.removeTop()
		h.insert(nums[i])
	}
	return h.getTop()
}
