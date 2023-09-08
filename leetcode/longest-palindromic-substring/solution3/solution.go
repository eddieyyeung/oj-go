package solution3

func longestPalindrome(s string) string {
	var traverse func(l, r int) bool
	m := s[:1]
	dp := [1000][1000]int{}
	traverse = func(l, r int) bool {
		if l >= r {
			dp[l][r] = 1
			return true
		}
		if dp[l][r] != 0 {
			return dp[l][r] == 1
		}
		if s[l] == s[r] && traverse(l+1, r-1) {
			if r-l+1 > len(m) {
				m = s[l : r+1]
			}
			dp[l][r] = 1
			return true
		}
		traverse(l+1, r)
		traverse(l, r-1)
		dp[l][r] = 2
		return false
	}
	traverse(0, len(s)-1)
	return m
}
