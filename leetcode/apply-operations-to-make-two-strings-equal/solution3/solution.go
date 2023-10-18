package solution3

func minOperations(s1 string, s2 string, x int) int {
	var p []int
	for i := 0; i < len(s1); i++ {
		if s1[i] != s2[i] {
			p = append(p, i)
		}
	}
	lp := len(p)
	if lp == 0 {
		return 0
	}
	if lp%2 != 0 {
		return -1
	}
	// 单独算出 i == 0 的情况
	dp0, dp1 := 0, x
	for i := 1; i < lp; i++ {
		ans1 := dp1 + x
		ans2 := dp0 + (p[i]-p[i-1])*2
		dp0, dp1 = dp1, min(ans1, ans2)
	}
	return dp1 / 2
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
