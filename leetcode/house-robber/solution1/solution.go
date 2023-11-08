// 198. 打家劫舍
// https://leetcode.cn/problems/house-robber/description/
package house_robber

// dfs(i) = max(dfs(i-1), dfs(i-2) + nums[i])
func rob(nums []int) int {
	n := len(nums)
	dp := make([]int, n)
	for i := 0; i < n; i++ {
		dp[i] = -1
	}
	var dfs func(i int) int
	dfs = func(i int) int {
		if i < 0 {
			return 0
		}
		if dp[i] != -1 {
			return dp[i]
		}
		ans := max(dfs(i-1), dfs(i-2)+nums[i])
		dp[i] = ans
		return ans
	}
	return dfs(n - 1)
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
