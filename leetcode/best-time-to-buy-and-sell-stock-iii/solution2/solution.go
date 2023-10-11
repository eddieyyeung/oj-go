package solution2

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func maxProfit(prices []int) int {
	dp := [4]int{-prices[0], 0, -prices[0], 0}
	for i := 1; i < len(prices); i++ {
		pricei := prices[i]
		dp[0], dp[1], dp[2], dp[3] = max(dp[0], -pricei), max(dp[1], dp[0]+pricei), max(dp[2], dp[1]-pricei), max(dp[3], dp[2]+pricei)
	}
	return dp[3]
}
