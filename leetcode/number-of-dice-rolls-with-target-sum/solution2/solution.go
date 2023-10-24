package solution2

var mod int = 1e9 + 7

// 记忆化搜索
func numRollsToTarget(n int, k int, target int) int {
	dp := make([][]int, n+1)
	for i := 0; i < n+1; i++ {
		dp[i] = make([]int, target+1)
		for j := 0; j < target+1; j++ {
			dp[i][j] = -1
		}
	}
	var dfs func(a, b int) int

	dfs = func(a, b int) int {
		if a == 0 {
			if b == 0 {
				return 1
			}
			return 0
		}
		if dp[a][b] != -1 {
			return dp[a][b]
		}
		// l := max(1, b-k*(a-1))
		// r := min(k, b-a+1)
		// for i := l; i <= r; i++ {
		// 	ans = (ans + dfs(a-1, b-i)) % mod
		// }
		ans := 0
		for i := 1; i <= k; i++ {
			if b-i >= 0 {
				ans = (ans + dfs(a-1, b-i)) % mod
			}
		}
		dp[a][b] = ans
		return ans
	}
	return dfs(n, target)
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
