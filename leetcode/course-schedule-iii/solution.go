package coursescheduleiii

import (
	"container/heap"
	"sort"
)

// 630. 课程表 III
// 这里有 n 门不同的在线课程，按从 1 到 n 编号。给你一个数组 courses ，
// 其中 courses[i] = [durationi, lastDayi] 表示第 i 门课将会持续 上 durationi 天课，
// 并且必须在不晚于 lastDayi 的时候完成。

// 你的学期从第 1 天开始。且不能同时修读两门及两门以上的课程。

// 返回你最多可以修读的课程数目。

// 示例 1：

// 输入：courses = [[100, 200], [200, 1300], [1000, 1250], [2000, 3200]]
// 输出：3
// 解释：
// 这里一共有 4 门课程，但是你最多可以修 3 门：
// 首先，修第 1 门课，耗费 100 天，在第 100 天完成，在第 101 天开始下门课。
// 第二，修第 3 门课，耗费 1000 天，在第 1100 天完成，在第 1101 天开始下门课程。
// 第三，修第 2 门课，耗时 200 天，在第 1300 天完成。
// 第 4 门课现在不能修，因为将会在第 3300 天完成它，这已经超出了关闭日期。
// 示例 2：

// 输入：courses = [[1,2]]
// 输出：1
// 示例 3：

// 输入：courses = [[3,2],[4,3]]
// 输出：0

// 提示:

// 1 <= courses.length <= 10^4
// 1 <= durationi, lastDayi <= 10^4

type IntHeap []int

// Len implements heap.Interface.
func (h IntHeap) Len() int {
	return len(h)
}

// Less implements heap.Interface.
func (h IntHeap) Less(i int, j int) bool {
	return h[i] > h[j]
}

// Pop implements heap.Interface.
func (h *IntHeap) Pop() any {
	old := *h
	l := len(old)
	n := old[l-1]
	*h = old[:l-1]
	return n
}

// Push implements heap.Interface.
func (h *IntHeap) Push(x any) {
	*h = append(*h, x.(int))
}

// Swap implements heap.Interface.
func (h IntHeap) Swap(i int, j int) {
	h[i], h[j] = h[j], h[i]
}

var _ heap.Interface = &IntHeap{}

func scheduleCourse(courses [][]int) int {
	sort.Slice(courses, func(i, j int) bool {
		return courses[i][1] < courses[j][1]
	})

	var h IntHeap
	var total int
	for _, course := range courses {
		d, l := course[0], course[1]
		if t := total + d; t <= l {
			total = t
			heap.Push(&h, d)
		} else if h.Len() > 0 && d < h[0] {
			total = total - h[0] + d
			h[0] = d
			heap.Fix(&h, 0)
		}
	}
	return h.Len()
}
