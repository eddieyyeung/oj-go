// Package longestpalindromicsubstring ...
// https://leetcode-cn.com/problems/longest-palindromic-substring/
package longestpalindromicsubstring

func longestPalindrome(s string) string {
	if s == "" {
		return ""
	}
	dp := [1000][1000]int{}
	r := string(s[0])
	dpf(s, 0, len(s)-1, &dp, &r)
	return r
}

func dpf(s string, left, right int, dp *[1000][1000]int, r *string) bool {
	if right-left+1 <= 1 {
		dp[left][right] = 1
		return true
	}
	if dp[left][right] != 0 {
		return dp[left][right] == 1
	}
	if dpf(s, left+1, right-1, dp, r) && s[left] == s[right] {
		if right-left+1 > len(*r) {
			*r = s[left : right+1]
		}
		return true
	}
	dpf(s, left+1, right, dp, r)
	dpf(s, left, right-1, dp, r)
	dp[left][right] = 2
	return false
}
