package solution2

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

// 参考 123. 买卖股票的最佳时机 III
// https://leetcode.cn/problems/best-time-to-buy-and-sell-stock-iii/
//
// 状态转移公式：
// buy1 = max(buy1, -prices[i])
// sell1 = max(sell1, buy1+prices[i])
// buy2 = max(buy2, sell1-prices[i])
// sell2 = max(sell2, buy2+prices[i])
// ...
//
// 初始化：
// buy1, buy2 = -prices[0], -prices[0]
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
