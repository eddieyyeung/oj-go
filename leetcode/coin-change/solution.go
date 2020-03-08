package coinchange

import (
	"math"
)

// coinChange https://leetcode-cn.com/problems/coin-change/
func coinChange(coins []int, amount int) int {
	dp := make([]int, amount+1)
	dp[0] = 0
	for i := 1; i < len(dp); i++ {
		min := math.MaxInt64
		for _, coin := range coins {
			if pre := i - coin; pre >= 0 && dp[pre] != -1 && dp[pre] < min {
				min = dp[pre]
			}
		}
		if min == math.MaxInt64 {
			dp[i] = -1
			continue
		}
		dp[i] = min + 1
	}
	return dp[amount]
}
