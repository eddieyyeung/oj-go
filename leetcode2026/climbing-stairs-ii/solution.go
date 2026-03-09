package climbing_stairs_ii

// https://leetcode.cn/problems/climbing-stairs-ii/description/
func climbStairs(n int, costs []int) int {
	dp := make([]int, n+1)
	for i := 0; i < len(dp); i++ {
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
		dp[i] = min(dfs(i-3)+costs[i]+9, dfs(i-2)+costs[i]+4, dfs(i-1)+costs[i]+1)
		return dp[i]
	}
	return dfs(n - 1)
}
