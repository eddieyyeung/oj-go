// 516. 最长回文子序列
// https://leetcode.cn/problems/longest-palindromic-subsequence/description/
package solution

// dfs(i,j) 表示 [i,j] 区间内，最长回文子序列的长度
// dfs(i,j) 分为以下几种情况：
// 1. 如果 s[i] == s[j]，则 dfs(i,j) = dfs(i+1,j-1)+2
// 2. 如果 s[i] != s[j]，则 dfs(i,j) = max(dfs(i+1,j), dfs(i,j-1))
func longestPalindromeSubseq(s string) int {
	n := len(s)
	dp := make([][]int, n+1)
	for i := 0; i < len(dp); i++ {
		dp[i] = make([]int, n+1)
	}
	// for i := 0; i < n; i++ {
	// 	dp[i][i+1] = 1
	// }
	for i := n - 1; i >= 0; i-- {
		dp[i][i+1] = 1
		for j := 0; j < n; j++ {
			if i < j {
				if s[i] == s[j] {
					dp[i][j+1] = dp[i+1][j] + 2
				} else {
					dp[i][j+1] = max(dp[i+1][j+1], dp[i][j])
				}
			}
		}
	}
	return dp[0][n]
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
