package minimumintervaltoincludeeachquery

import (
	"container/heap"
	"sort"
)

// 1851. 包含每个查询的最小区间

// 给你一个二维整数数组 intervals ，其中 intervals[i] = [lefti, righti] 表示第 i 个区间开始于 lefti 、结束于 righti（包含两侧取值，闭区间）。区间的 长度 定义为区间中包含的整数数目，更正式地表达是 righti - lefti + 1 。

// 再给你一个整数数组 queries 。第 j 个查询的答案是满足 lefti <= queries[j] <= righti 的 长度最小区间 i 的长度 。如果不存在这样的区间，那么答案是 -1 。

// 以数组形式返回对应查询的所有答案。

// 示例 1：

// 输入：intervals = [[1,4],[2,4],[3,6],[4,4]], queries = [2,3,4,5]
// 输出：[3,3,1,4]
// 解释：查询处理如下：
// - Query = 2 ：区间 [2,4] 是包含 2 的最小区间，答案为 4 - 2 + 1 = 3 。
// - Query = 3 ：区间 [2,4] 是包含 3 的最小区间，答案为 4 - 2 + 1 = 3 。
// - Query = 4 ：区间 [4,4] 是包含 4 的最小区间，答案为 4 - 4 + 1 = 1 。
// - Query = 5 ：区间 [3,6] 是包含 5 的最小区间，答案为 6 - 3 + 1 = 4 。
// 示例 2：

// 输入：intervals = [[2,3],[2,5],[1,8],[20,25]], queries = [2,19,5,22]
// 输出：[2,-1,4,6]
// 解释：查询处理如下：
// - Query = 2 ：区间 [2,3] 是包含 2 的最小区间，答案为 3 - 2 + 1 = 2 。
// - Query = 19：不存在包含 19 的区间，答案为 -1 。
// - Query = 5 ：区间 [2,5] 是包含 5 的最小区间，答案为 5 - 2 + 1 = 4 。
// - Query = 22：区间 [20,25] 是包含 22 的最小区间，答案为 25 - 20 + 1 = 6 。

// 提示：

// 1 <= intervals.length <= 105
// 1 <= queries.length <= 105
// intervals[i].length == 2
// 1 <= lefti <= righti <= 107
// 1 <= queries[j] <= 107

type HeapItem struct {
	Left     int
	Right    int
	Interval int
}

type Heap []HeapItem

// Len implements heap.Interface.
func (h Heap) Len() int {
	return len(h)
}

// Less implements heap.Interface.
func (h Heap) Less(i int, j int) bool {
	return h[i].Interval < h[j].Interval
}

// Pop implements heap.Interface.
func (h *Heap) Pop() any {
	old := *h
	n := old[len(old)-1]
	*h = old[:len(old)-1]
	return n
}

// Push implements heap.Interface.
func (h *Heap) Push(x any) {
	*h = append(*h, x.(HeapItem))
}

// Swap implements heap.Interface.
func (h Heap) Swap(i int, j int) {
	h[i], h[j] = h[j], h[i]
}

var _ heap.Interface = &Heap{}

func minInterval(intervals [][]int, queries []int) []int {
	qindex := make([]int, len(queries))
	for i := 0; i < len(qindex); i++ {
		qindex[i] = i
	}
	sort.Slice(intervals, func(i, j int) bool {
		return intervals[i][0] < intervals[j][0]
	})
	sort.Slice(qindex, func(i, j int) bool {
		return queries[qindex[i]] < queries[qindex[j]]
	})
	var i int
	var h Heap
	ans := make([]int, len(queries))
	for _, idx := range qindex {
		q := queries[idx]
		for ; i < len(intervals); i++ {
			l, r := intervals[i][0], intervals[i][1]
			if q >= l && q <= r {
				heap.Push(&h, HeapItem{
					Left:     l,
					Right:    r,
					Interval: r - l + 1,
				})
			} else if r < q {
				continue
			} else {
				break
			}
		}
		// fmt.Println(q, h)
		var isFound bool
		for h.Len() > 0 {
			m := h[0]
			if q >= m.Left && q <= m.Right {
				ans[idx] = m.Interval
				isFound = true
				break
			} else {
				heap.Pop(&h)
			}
		}
		if !isFound {
			ans[idx] = -1
		}
	}
	// fmt.Println("--------")
	return ans
}
