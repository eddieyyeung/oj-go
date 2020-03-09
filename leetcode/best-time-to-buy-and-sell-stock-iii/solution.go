package besttimetobuyandsellstockiii

import "math"

// maxProfit https://leetcode-cn.com/problems/best-time-to-buy-and-sell-stock-iii/
func maxProfit(prices []int) int {
	if len(prices) == 0 {
		return 0
	}
	dp := make([][2]int, 3)
	t := make([][2]int, 3)
	dp[1][0] = -prices[0]
	dp[0][0] = -prices[0]
	for i := 1; i < len(prices); i++ {
		for j := 0; j < 2; j++ {
			t[j][0] = max(dp[j][0], dp[j+1][1]-prices[i])
			t[j][1] = max(dp[j][1], dp[j][0]+prices[i])
		}

		for j := 0; j < 2; j++ {
			dp[j][0], dp[j][1] = t[j][0], t[j][1]
		}
	}
	max := math.MinInt64
	for i := 0; i < 3; i++ {
		if dp[i][1] > max {
			max = dp[i][1]
		}
	}
	return max
}

func max(a ...int) int {
	max := math.MinInt64
	for _, v := range a {
		if v > max {
			max = v
		}
	}
	return max
}
