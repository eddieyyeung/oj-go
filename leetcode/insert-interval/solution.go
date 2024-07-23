package insertinterval

import (
	"sort"
)

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func insert(intervals [][]int, newInterval []int) [][]int {
	if len(intervals) == 0 {
		return [][]int{newInterval}
	}
	idx := sort.Search(len(intervals), func(i int) bool {
		if newInterval[0] == intervals[i][0] {
			return newInterval[1] < intervals[i][1]
		}
		return newInterval[0] < intervals[i][0]
	})
	var ans [][]int
	var start, end int = intervals[0][0], intervals[0][1]
	if idx == 0 {
		start, end = newInterval[0], newInterval[1]
	}
	for i := 0; i < len(intervals); i++ {
		inter := intervals[i]
		if idx == i {
			inter = newInterval
			i--
			idx = -1
		}
		if inter[0] <= end {
			end = max(inter[1], end)
		} else {
			ans = append(ans, []int{start, end})
			start, end = inter[0], inter[1]
		}
	}
	ans = append(ans, []int{start, end})
	if idx == len(intervals) {
		if newInterval[0] <= end {
			ans[len(ans)-1][1] = max(newInterval[1], end)
		} else {
			ans = append(ans, newInterval)
		}
	}
	return ans
}
