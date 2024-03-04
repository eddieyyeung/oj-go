// 32. 最长有效括号
// https://leetcode.cn/problems/longest-valid-parentheses/description/
package solution

// dp[i]
// = dp[i-2]+2, if s[i] == ')' && s[i-1] == '('
// = dp[i-1]+2+dp[i-dp[i-1]-2], if s[i] == ')' && s[i-dp[i-1]-1] == '(
// = 0, if s[i] == '('
func longestValidParentheses(s string) int {
	// dp := make([]int, len(s) + 1)
	var ans int
	dp := make([]int, len(s)+1)

	for i := 0; i < len(dp); i++ {
		dp[i] = -1
	}
	var dfs func(n int) int
	dfs = func(n int) int {
		if n <= 0 {
			return 0
		}
		if s[n] == '(' {
			return 0
		}
		if dp[n] != -1 {
			return dp[n]
		}
		var rst int
		if n-1 >= 0 && s[n-1] == '(' {
			rst = dfs(n-2) + 2
		}
		t := dfs(n - 1)
		if n-t-1 >= 0 && s[n-t-1] == '(' {
			rst = max(rst, t+2+dfs(n-t-2))
		}
		ans = max(ans, rst)
		dp[n] = rst
		return rst
	}
	for i := 1; i < len(s); i++ {
		if s[i] == ')' {
			dfs(i)
		}
	}
	return ans
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
