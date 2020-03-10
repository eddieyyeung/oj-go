// Package besttimetobuyandsellstock ...
// https://leetcode-cn.com/problems/best-time-to-buy-and-sell-stock/
package besttimetobuyandsellstock

func maxProfit(prices []int) int {
	max := 0
	for i := 0; i < len(prices)-1; i++ {
		for j := i + 1; j < len(prices); j++ {
			if diff := prices[j] - prices[i]; diff > max {
				max = diff
			}
		}
	}
	return max
}
