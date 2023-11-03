// 22. 括号生成
// https://leetcode.cn/problems/generate-parentheses/description/
package generateparenthesis

import "strings"

// 时间复杂度：O(n*C(2n, n)) Catalan
func generateParenthesis(n int) []string {
	var ans []string
	var path []string
	var dfs func(i, j int)
	dfs = func(i, j int) {
		if i == j && i == n {
			ans = append(ans, strings.Join(path, ""))
			return
		}
		// 加左括号，当左括号的个数小于 n 时才能够添加
		if i < n {
			path = append(path, "(")
			dfs(i+1, j)
			path = path[:len(path)-1]
		}

		// 加右括号，当左括号的个数大于右括号的个数时才能够添加
		if i > j {
			path = append(path, ")")
			dfs(i, j+1)
			path = path[:len(path)-1]
		}
	}
	dfs(0, 0)
	return ans
}
