package solution2

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func maxProfit(k int, prices []int) int {
	dp := make([]int, k*2+1)
	for i := 1; i < k*2; i++ {
		if i%2 == 1 {
			dp[i] = -prices[0]
		}
	}
	for i := 1; i < len(prices); i++ {
		price := prices[i]
		for j := 1; j < len(dp); j++ {
			if j%2 == 1 {
				dp[j] = max(dp[j], dp[j-1]-price)
			} else {
				dp[j] = max(dp[j], dp[j-1]+price)
			}
		}
	}
	return dp[len(dp)-1]
}
