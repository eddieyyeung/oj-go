package solution1

var mod int = 1e9 + 7

// 递推
func numRollsToTarget(n int, k int, target int) int {
	dp := make([][]int, n+1)
	for i := 0; i < n+1; i++ {
		dp[i] = make([]int, target+1)
	}
	dp[0][0] = 1
	for i := 1; i <= n; i++ {
		for j := 1; j <= target; j++ {
			// l := max(1, j-k*(i-1))
			// r := min(k, j-i+1)
			// var total int
			// for t := l; t <= r; t++ {
			// 	total = (total + dp[i-1][j-t]) % mod
			// }
			// dp[i][j] = total
			dp[i][j] = 0
			for t := 1; t <= k; t++ {
				if j-t >= 0 {
					dp[i][j] = (dp[i][j] + dp[i-1][j-t]) % mod
				}
			}
		}
	}
	return dp[n][target]
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
