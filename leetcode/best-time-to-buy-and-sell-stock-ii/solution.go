// Package besttimetobuyandsellstockii ...
// https://leetcode-cn.com/problems/best-time-to-buy-and-sell-stock-ii/
package besttimetobuyandsellstockii

import "math"

func maxProfit(prices []int) int {
	if len(prices) == 0 {
		return 0
	}
	dp0 := -prices[0]
	dp1 := 0
	for i := 1; i < len(prices); i++ {
		t0 := max(dp0, dp1-prices[i])
		t1 := max(dp0+prices[i], dp1)
		dp0, dp1 = t0, t1
	}
	return dp1
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
