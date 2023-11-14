// 1143. 最长公共子序列
// https://leetcode.cn/problems/longest-common-subsequence/description/
package solution

// if s[i]==t[j], dfs(i,j) = dfs(i-1,j-1)+1
// if s[i]!=t[j], dfs(i,j) = max(dfs(i-1,j), dfs(i,j-1)
// =>
// if s[i]==t[j], dfs(i+1,j+1) = dfs(i,j)+1
// if s[i]!=t[j], dfs(i+1,j+1) = max(dfs(i,j+1), dfs(i+1,j)
//
// 边界条件：
// if i < 0 || j < 0, return 0
//
// 时间复杂度：O(n*m)
// 空间复杂度：O(m)
func longestCommonSubsequence(text1 string, text2 string) int {
	n, m := len(text1), len(text2)
	dp := make([]int, m+1)
	for i := 0; i < n; i++ {
		var pre int
		for j := 0; j < m; j++ {
			tmp := dp[j+1]
			if text1[i] == text2[j] {
				dp[j+1] = pre + 1
			} else {
				dp[j+1] = max(dp[j+1], dp[j])
			}
			pre = tmp
		}
	}
	return dp[m]
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
