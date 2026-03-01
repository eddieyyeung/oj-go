package count_ways_to_build_good_strings

func countGoodStrings2(low int, high int, zero int, one int) int {
	var MOD int = 1e9 + 7
	n := high
	var dfs func(int) int
	nums := []int{zero, one}
	dp := make([]int, n+1)
	for i := range dp {
		dp[i] = -1
	}

	dp[0] = 1

	dfs = func(i int) int {
		if i == 0 {
			return 1
		}
		if i < 0 {
			return 0
		}
		if dp[i] != -1 {
			return dp[i]
		}
		var rst int
		for _, num := range nums {
			rst += dfs(i - num)
		}
		dp[i] = rst % MOD
		return dp[i]
	}
	dfs(high)
	var ans int
	for i := low; i <= high; i++ {
		ans += dfs(i)
	}
	return ans % MOD
}

func countGoodStrings(low int, high int, zero int, one int) int {
	var MOD int = 1e9 + 7
	dp := make([]int, high+1)
	dp[0] = 1
	for i := 1; i <= high; i++ {
		if i >= zero {
			dp[i] += dp[i-zero] % MOD
		}
		if i >= one {
			dp[i] += dp[i-one] % MOD
		}
	}
	var ans int
	for i := low; i <= high; i++ {
		ans += dp[i]
	}
	return ans % MOD
}
