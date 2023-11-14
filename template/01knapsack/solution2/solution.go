package knapsack

// 0-1 背包
// 问题描述：
// 有 n 个物品，第 i 个物品的体积为 w[i]，价值为 v[i]，
// 每个物品至多选一个，求体积和不超过 capacity 的最大价值和。
//
// 题解：
// dp[i][c] = max(dp[i-1][c], dp[i-1][c-w[i]]+v[i])
// => dp[i+1][c] = max(dp[i][c], dp[i][c-w[i]]+v[i]) 避免数组越界
//
// 时间复杂度：O(n*c)
// 空间复杂度：O(n*c)

// capacity: 背包容量
// w[i]: 第 i 个物品的体积
// v[i]: 第 i 个物品的价值
// return: 所选物品体积不超过 capacity 的前提下，所能得到的最大价值和
func zero_one_knapsack(capacity int, w []int, v []int) int {
	n := len(w)
	dp := make([][]int, n+1)
	for i := 0; i < len(dp); i++ {
		dp[i] = make([]int, capacity+1)
	}

	for i := 0; i < n; i++ {
		for c := 0; c <= capacity; c++ {
			if c < w[i] {
				dp[i+1][c] = dp[i][c]
			} else {
				dp[i+1][c] = max(dp[i][c], dp[i][c-w[i]]+v[i])
			}
		}
	}
	return dp[n][capacity]
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
