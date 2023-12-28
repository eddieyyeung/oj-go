// 2735. 收集巧克力
// https://leetcode.cn/problems/collecting-chocolates/description/
package solution

import "slices"

func minCost(nums []int, x int) int64 {
	n := len(nums)
	s := make([]int, n)
	for i := range s {
		s[i] = i * x
	}
	for i, num := range nums {
		for j := 0; j < n; j++ {
			num = min(num, nums[(i+j)%n])
			s[j] += num
		}
	}
	return int64(slices.Min(s))
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
