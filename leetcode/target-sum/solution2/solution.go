// 494. 目标和
// https://leetcode.cn/problems/target-sum/description/
package solution1

// dfs(i,c) = dfs(i-1,c)+dfs(i-1,c-nums[i])
func findTargetSumWays(nums []int, target int) int {
	n := len(nums)
	sum := 0
	for _, num := range nums {
		sum += num
	}
	p := sum + target
	if p < 0 || p%2 != 0 {
		return 0
	}
	p = p / 2

	dp := make([][]int, n+1)
	for i := 0; i < len(dp); i++ {
		dp[i] = make([]int, p+1)
	}
	dp[0][0] = 1
	for i := 0; i < n; i++ {
		for c := 0; c <= p; c++ {
			if c < nums[i] {
				dp[i+1][c] = dp[i][c]
			} else {
				dp[i+1][c] = dp[i][c] + dp[i][c-nums[i]]
			}
		}
	}
	return dp[n][p]
}
