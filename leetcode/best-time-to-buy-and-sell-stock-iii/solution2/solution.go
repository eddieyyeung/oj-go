package solution2

// 状态转移公式：
// buy1 = max(buy1, -prices[i])
// sell1 = max(sell1, buy1+prices[i])
// buy2 = max(buy2, sell1-prices[i])
// sell2 = max(sell2, buy2+prices[i])
//
// 初始化：
// buy1, buy2 = -prices[0], -prices[0]
func maxProfit(prices []int) int {
	dp := [4]int{-prices[0], 0, -prices[0], 0}
	for i := 1; i < len(prices); i++ {
		pricei := prices[i]
		dp[0], dp[1], dp[2], dp[3] = max(dp[0], -pricei), max(dp[1], dp[0]+pricei), max(dp[2], dp[1]-pricei), max(dp[3], dp[2]+pricei)
	}
	return dp[3]
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
