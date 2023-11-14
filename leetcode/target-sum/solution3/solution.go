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

	dp := make([]int, p+1)
	dp[0] = 1
	for i := 0; i < n; i++ {
		// 可写成:
		// for c := p; c >= nums[i]; c-- {}
		for c := p; c >= 0; c-- {
			if c >= nums[i] {
				dp[c] = dp[c] + dp[c-nums[i]]
			}
		}
		// 如果是至少 target，则改为：
		// for c := p; c >= 0; c-- {
		// 	dp[c] = dp[c] + dp[max(c-nums[i], 0)]
		// }
	}
	return dp[p]
}
