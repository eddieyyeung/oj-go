package knapsack

// 0-1 背包
// 问题描述：
// 有 n 个物品，第 i 个物品的体积为 w[i]，价值为 v[i]，
// 每个物品至多选一个，求体积和不超过 capacity 的最大价值和。
//
// 题解：
// dfs(i, c) = max(dfs(i-1, c), dfs(i-1, c-w[i])+v[i])

// capacity: 背包容量
// w[i]: 第 i 个物品的体积
// v[i]: 第 i 个物品的价值
// return: 所选物品体积不超过 capacity 的前提下，所能得到的最大价值和
func zero_one_knapsack(capacity int, w []int, v []int) int {
	n := len(w)
	dp := make([][]int, n)
	for i := 0; i < n; i++ {
		dp[i] = make([]int, capacity+1)
		for j := 0; j < len(dp[i]); j++ {
			dp[i][j] = -1
		}
	}

	var dfs func(i, c int) int
	dfs = func(i, c int) int {
		if i < 0 {
			return 0
		}
		if dp[i][c] != -1 {
			return dp[i][c]
		}
		var rst int
		if c < w[i] {
			rst = dfs(i-1, c)
		} else {
			rst = max(dfs(i-1, c-w[i])+v[i], dfs(i-1, c))
		}
		dp[i][c] = rst
		return rst
	}
	return dfs(n-1, capacity)
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
