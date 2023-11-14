// 494. 目标和
// https://leetcode.cn/problems/target-sum/description/
package solution1

// dfs(i, c) = dfs(i-1, c) + dfs(i-1, c-nums[i])
// p: 添加正号的数之和
// s-p: 添加负号的数之和
// p-(s-p) = t
// => p = (s+t)//2
// 问题转换：
// 从 nums 中选择一些数字，恰好等于 p 的方案数
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
		for j := 0; j < len(dp[i]); j++ {
			dp[i][j] = -1
		}
	}
	var dfs func(i, c int) int
	dfs = func(i, c int) int {
		if i < 0 {
			if c == 0 {
				return 1
			}
			return 0
		}
		if dp[i][c] != -1 {
			return dp[i][c]
		}
		if c < nums[i] {
			dp[i][c] = dfs(i-1, c)
			return dp[i][c]
		}
		dp[i][c] = dfs(i-1, c) + dfs(i-1, c-nums[i])
		return dp[i][c]
	}
	return dfs(n-1, p)
}
