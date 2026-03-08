package combination_sum_iii

import "slices"

// https://leetcode.cn/problems/combination-sum-iii/
func combinationSum3(k int, n int) [][]int {
	path := []int{}
	ans := [][]int{}
	var dfs func(i int, r int)
	dfs = func(i int, r int) {
		if (1+i)*i/2 < r {
			return
		}
		if len(path) > k {
			return
		}
		if r == 0 && len(path) == k {
			ans = append(ans, slices.Clone(path))
			return
		}
		if i == 0 {
			return
		}
		dfs(i-1, r)
		path = append(path, i)
		dfs(i-1, r-i)
		path = path[:len(path)-1]
	}
	dfs(9, n)
	return ans
}
