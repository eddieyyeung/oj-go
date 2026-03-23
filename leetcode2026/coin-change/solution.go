package coin_change

import "math"

// https://leetcode.cn/problems/coin-change/description/
func coinChange(coins []int, amount int) int {
	dp := make([]int, amount+1)
	for i := range dp {
		dp[i] = -1
	}

	var dfs func(c int) int
	dfs = func(c int) int {
		if c == 0 {
			return 0
		}
		if c < 0 {
			return -1
		}
		if dp[c] != -1 {
			return dp[c]
		}
		rst := math.MaxInt
		for _, coin := range coins {
			t := dfs(c - coin)
			if t != -1 {
				rst = min(rst, t)
			}
		}
		if rst == math.MaxInt {
			dp[c] = rst
		} else {
			dp[c] = rst + 1
		}
		return dp[c]
	}
	rst := dfs(amount)
	if rst == math.MaxInt {
		return -1
	}
	return rst
}
