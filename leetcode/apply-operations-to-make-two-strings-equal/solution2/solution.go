package solution2

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

	df := make([]int, lp)
	for i := 0; i < lp; i++ {
		df[i] = -1
	}
	var dfs func(i int) int

	// dfs(i)
	// 1. dfs(i)_1 = dfs(i-1) + x
	// 2. dfs(i)_2 = dfs(i-2) + (p[i] - p[i-1]) * 2
	// dfs(i) =  min(dfs(i)_1, dfs(i)_2)
	dfs = func(i int) int {
		if i == -1 {
			return 0
		}
		if df[i] != -1 {
			return df[i]
		}
		ans1 := dfs(i-1) + x
		var ans2 int = 1e9
		if i >= 1 {
			ans2 = dfs(i-2) + (p[i]-p[i-1])*2
		}
		df[i] = min(ans1, ans2)
		return df[i]
	}
	return dfs(lp-1) / 2
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
