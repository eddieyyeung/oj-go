// 131. 分割回文串
// https://leetcode.cn/problems/palindrome-partitioning/description/
package palindromepartitioning

func isPalindrome(s string, left, right int) bool {
	for left <= right {
		if s[left] != s[right] {
			return false
		}
		left++
		right--
	}
	return true
}

// 答案视角
func partition(s string) [][]string {
	n := len(s)
	var ans [][]string
	var path []string
	var dfs func(i int)
	dfs = func(i int) {
		if i == n {
			// 由于 path 长度可变，需固定答案
			ans = append(ans, append([]string{}, path...))
			return
		}
		// 题目要求不能包含重复的数字
		// 因此规定一个顺序，从当前回溯的 i 至 n 来枚举选择的数字
		for j := i; j < n; j++ {
			if isPalindrome(s, i, j) {
				path = append(path, s[i:j+1])
				dfs(j + 1)
				// 恢复现场
				path = path[:len(path)-1]
			}
		}
	}
	dfs(0)
	return ans
}
