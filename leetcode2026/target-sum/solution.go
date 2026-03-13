package target_sum

// https://leetcode.cn/problems/target-sum/description/
func findTargetSumWays(nums []int, target int) int {
	var sum int
	for _, num := range nums {
		sum += num
	}
	if (sum+target)%2 != 0 {
		return 0
	}
	t := (sum + target) / 2
	if t < 0 {
		return 0
	}
	n := len(nums)

	dp := make([][]int, n)
	for i := 0; i < len(dp); i++ {
		dp[i] = make([]int, t+1)
		for j := 0; j < len(dp[i]); j++ {
			dp[i][j] = -1
		}
	}
	var dfs func(i int, c int) int
	dfs = func(i int, c int) int {
		if c < 0 {
			return 0
		}
		if i < 0 {
			if c == 0 {
				return 1
			}
			return 0
		}
		if dp[i][c] != -1 {
			return dp[i][c]
		}
		dp[i][c] = dfs(i-1, c) + dfs(i-1, c-nums[i])
		return dp[i][c]
	}
	ans := dfs(n-1, t)
	return ans
}
