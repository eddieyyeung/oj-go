package solution2

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
	dp := make([]int, n+1)
	dp[0] = 1
	for i := 1; i <= n; i++ {
		ch := pressedKeys[i-1]
		count := counts[ch]
		for j := 1; j <= count; j++ {
			if i-j < 0 {
				break
			}
			dp[i] += dp[i-j] % MOD
			if i-j-1 >= 0 && pressedKeys[i-j-1] != ch {
				break
			}
		}
	}
	return dp[n] % MOD
}
