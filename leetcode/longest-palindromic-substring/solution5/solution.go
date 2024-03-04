// Package longestpalindromicsubstring ...
// https://leetcode-cn.com/problems/longest-palindromic-substring/
package solution

func longestPalindrome(s string) string {
	var ans string
	for i := 0; i < len(s); i++ {
		str := getLongestPalindrome(s, i, i)
		if len(str) > len(ans) {
			ans = str
		}
	}
	for i := 0; i < len(s)-1; i++ {
		if s[i] == s[i+1] {
			str := getLongestPalindrome(s, i, i+1)
			if len(str) > len(ans) {
				ans = str
			}
		}
	}
	return ans
}

func getLongestPalindrome(s string, l, r int) string {
	for l > 0 && r < len(s)-1 && s[l-1] == s[r+1] {
		l--
		r++
	}
	return s[l : r+1]
}
