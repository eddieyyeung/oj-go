// 1039. 多边形三角剖分的最低得分
// https://leetcode.cn/problems/minimum-score-triangulation-of-polygon/description/
package solution

// dp(i,j) 表示区间 [i,j] 之间，多边形三角剖分的最低得分。n 表示 len(v)
// 在 dp(i,j) 中，枚举 k ∈ (i,j)，将 [i,j] 多边形划分为三个区域：dfs(i,k)，三角形 {i, j, k}，dfs(k,j)
// 则 dp 方程式： dp(i,j) = min(dfs(i,k) + dfs(k,j) + v[i]*v[j]*v[k]), k ∈ (i,j)
// 递归边界：dp(i,j) = 0, j <= i+1
// 递归入口：dp(0,n-1)
func minScoreTriangulation(values []int) int {
	n := len(values)
	dp := make([][]int, n)
	for i := 0; i < n; i++ {
		dp[i] = make([]int, n)
	}
	for i := n - 1; i >= 0; i-- {
		for j := i + 2; j < n; j++ {
			var rst int = 1e9
			for k := i + 1; k < j; k++ {
				rst = min(rst, dp[i][k]+dp[k][j]+values[i]*values[j]*values[k])
			}
			dp[i][j] = rst
		}
	}
	return dp[0][n-1]
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
