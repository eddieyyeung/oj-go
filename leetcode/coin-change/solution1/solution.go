// 322. 零钱兑换
// https://leetcode.cn/problems/coin-change/description/
package coinchange

import "sort"

// dfs(i, c) = min(dfs(i-1,c), dfs(i, c-w[i])+1)
func coinChange(coins []int, amount int) int {
	n := len(coins)
	// 先排序，把大数放前面，减少判断次数
	sort.Slice(coins, func(i, j int) bool {
		return coins[i] > coins[j]
	})
	dp := make([]int, amount+1)
	for i := 1; i < len(dp); i++ {
		dp[i] = 1e9
	}
	for i := 0; i < n; i++ {
		for j := 0; j < len(dp); j++ {
			if j >= coins[i] {
				dp[j] = min(dp[j], dp[j-coins[i]]+1)
			}
		}
	}
	if dp[amount] < 1e9 {
		return dp[amount]
	}
	return -1
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
