package jianshengzilcof

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func cuttingRope(n int) int {
	dp := [58]int{}
	dp[0] = 1
	for i := 1; i < n; i++ {
		m := i + 1
		if i == n-1 {
			m = 1
		}
		for j := 0; j < i; j++ {
			m = max(dp[j]*(i-j), m)
		}
		dp[i] = m
	}
	return dp[n-1]
}
