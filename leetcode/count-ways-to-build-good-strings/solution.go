// 2466. 统计构造好字符串的方案数
// https://leetcode.cn/problems/count-ways-to-build-good-strings/description/
package countwaystobuildgoodstrings

var mod int = 1e9 + 7

func countGoodStrings(low int, high int, zero int, one int) int {
	dp := make([]int, high+1)
	dp[0] = 1
	var ans int
	for i := 1; i <= high; i++ {
		dp1, dp2 := 0, 0
		if zidx := i - zero; zidx >= 0 {
			dp1 = dp[zidx]
		}
		if oidx := i - one; oidx >= 0 {
			dp2 = dp[oidx]
		}
		dp[i] = (dp1 + dp2) % mod
		if i >= low {
			ans = (ans + dp[i]) % mod
		}
	}
	return ans
}
