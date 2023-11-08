// 931. 下降路径最小和
// https://leetcode.cn/problems/minimum-falling-path-sum/description/
package minimumfallingpathsum

var MAX_NUM int = 1e9

func minFallingPathSum(matrix [][]int) int {
	n := len(matrix)
	dp := make([][]int, n)
	for i := 0; i < len(dp); i++ {
		dp[i] = make([]int, n+2)
		dp[i][0], dp[i][n+1] = MAX_NUM, MAX_NUM
	}
	for i := 0; i < len(dp); i++ {
		dp[0][i+1] = matrix[0][i]
	}
	dp[0][0], dp[0][n+1] = MAX_NUM, MAX_NUM
	for i := 1; i < n; i++ {
		for j := 0; j < n; j++ {
			dp[i][j+1] = min(min(dp[i-1][j], dp[i-1][j+1]), dp[i-1][j+2]) + matrix[i][j]
		}
	}
	ans := dp[n-1][1]
	for i := 2; i < n+2; i++ {
		ans = min(ans, dp[n-1][i])
	}
	return ans
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
