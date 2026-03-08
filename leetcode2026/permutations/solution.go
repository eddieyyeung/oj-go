package permutations

import (
	"slices"
)

// https://leetcode.cn/problems/permutations/description/
func permute(nums []int) [][]int {
	n := len(nums)
	var ans [][]int
	path := make([]int, n)
	onPath := make([]bool, n)

	var dfs func(i int)
	dfs = func(i int) {
		if i == n {
			ans = append(ans, slices.Clone(path))
			return
		}
		for j, on := range onPath {
			if !on {
				path[i] = nums[j]
				onPath[j] = true
				dfs(i + 1)
				onPath[j] = false
			}
		}
	}
	dfs(0)
	return ans
}
