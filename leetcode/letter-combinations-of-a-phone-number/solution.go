// 17. 电话号码的字母组合
// https://leetcode.cn/problems/letter-combinations-of-a-phone-number/description/
package lettercombinationsofaphonenumber

func letterCombinations(digits string) []string {
	n := len(digits)
	if n == 0 {
		return nil
	}
	var ans []string
	digitMap := []string{"", "", "abc", "def", "ghi", "jkl", "mno", "pqrs", "tuv", "wxyz"}
	path := make([]byte, n)
	var dfs func(int)
	dfs = func(i int) {
		if i == n {
			ans = append(ans, string(path))
			return
		}
		for _, c := range digitMap[digits[i]-'0'] {
			path[i] = byte(c)
			dfs(i + 1)
		}
	}
	dfs(0)
	return ans
}
