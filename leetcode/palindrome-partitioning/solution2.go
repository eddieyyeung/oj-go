// 131. 分割回文串
// https://leetcode.cn/problems/palindrome-partitioning/description/
package palindromepartitioning

// 输入视角
func partition2(s string) [][]string {
	n := len(s)
	ans := make([][]string, 0, 1<<n)
	path := make([]string, 0, n)

	// start 表示当前这段回文子串的开始位置
	var dfs func(i, start int)
	dfs = func(i, start int) {
		if i == n {
			// 由于 path 长度可变，需固定答案
			ans = append(ans, append([]string{}, path...))
			return
		}
		// 不选 i 和 i+1 之间的逗号（i=n-1 时一定要选）
		if i < n-1 {
			dfs(i+1, start)
		}
		// 选 i 和 i+1 之间的逗号（把 s[i] 作为子串的最后一个字符）
		if isPalindrome(s, start, i) {
			path = append(path, s[start:i+1])
			// 下一个子串从 i+1 开始
			dfs(i+1, i+1)
			path = path[:len(path)-1]
		}
	}
	dfs(0, 0)
	return ans
}
