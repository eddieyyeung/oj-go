package count_number_of_texts

func countTexts(pressedKeys string) int {
	var MOD int = 1e9 + 7
	n := len(pressedKeys)
	counts := map[byte]int{
		'2': 3,
		'3': 3,
		'4': 3,
		'5': 3,
		'6': 3,
		'7': 4,
		'8': 3,
		'9': 4,
	}
	dp := make([]int, n)
	for i := 0; i < len(dp); i++ {
		dp[i] = -1
	}
	dp[0] = 1
	var dfs func(i int) int
	dfs = func(i int) int {
		if i == 0 {
			return 1
		}
		if i < 0 {
			return 1
		}
		if dp[i] != -1 {
			return dp[i]
		}
		ch := pressedKeys[i]
		count := counts[ch]
		var ans int
		for j := 1; j <= count; j++ {
			ans += dfs(i - j)
			if i-j < 0 {
				break
			}
			if pressedKeys[i-j] != ch {
				break
			}
		}
		dp[i] = ans % MOD
		return dp[i]
	}

	ans := dfs(n - 1)
	return ans % MOD
}
