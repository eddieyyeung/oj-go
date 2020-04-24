// Package merge_intervals ...
// https://leetcode-cn.com/problems/merge-intervals/
package merge_intervals

import (
	"math"
	"sort"
)

func merge(intervals [][]int) [][]int {
	if len(intervals) <= 1 {
		return intervals
	}
	sort.Slice(intervals, func(i, j int) bool {
		return intervals[i][0] < intervals[j][0]
	})
	var result [][]int
	result = append(result, intervals[0])
	for i := 1; i < len(intervals); i++ {
		last := result[len(result)-1]
		interval := intervals[i]
		if interval[0] >= last[0] && interval[0] <= last[1] {
			last[1] = max(last[1], interval[1])
		} else {
			result = append(result, interval)
		}
	}
	return result
}

func max(arr ...int) int {
	max := math.MinInt64
	for _, num := range arr {
		if num > max {
			max = num
		}
	}
	return max
}
