package house_robber

func rob(nums []int) int {
	n := len(nums)
	dp := make([]int, n)
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
		dp[i] = max(dfs(i-2)+nums[i], dfs(i-1))
		return dp[i]
	}
	ans := dfs(n - 1)
	return ans
}
