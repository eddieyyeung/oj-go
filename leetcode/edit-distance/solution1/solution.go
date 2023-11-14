// 72. 编辑距离
// https://leetcode.cn/problems/edit-distance/description/
package solution1

// if s[i]==t[j], dfs(i,j) = dfs(i-1,j-1) 不选
// if s[i]!=t[j], dfs(i,j) = min(<替换>dfs(i-1,j-1)+1, <删除>dfs(i-1,j)+1, <插入>dfs(i,j-1)+1)
// 边界条件
// if i < 0, return j+1
// if j < 0, return i+1
func minDistance(word1 string, word2 string) int {
	n, m := len(word1), len(word2)
	dp := make([]int, m+1)
	for i := 1; i < len(dp); i++ {
		dp[i] = i
	}
	for i := 0; i < n; i++ {
		var pre int = dp[0]
		dp[0] = i + 1
		for j := 0; j < m; j++ {
			tmp := dp[j+1]
			if word1[i] == word2[j] {
				dp[j+1] = pre
			} else {
				dp[j+1] = min(min(pre+1, dp[j+1]+1), dp[j]+1)
			}
			pre = tmp
		}
	}
	return dp[m]
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
