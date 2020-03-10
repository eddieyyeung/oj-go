// Package besttimetobuyandsellstockiv ...
// https://leetcode-cn.com/problems/best-time-to-buy-and-sell-stock-iv/
package besttimetobuyandsellstockiv

import "math"

func maxProfit(k int, prices []int) int {
	if len(prices) == 0 {
		return 0
	}
	var dp [][2]int
	var t [][2]int
	daily := false
	if k >= len(prices)/2 {
		daily = true
	} else {
		daily = false
	}
	if daily {
		dp = make([][2]int, 1)
		t = make([][2]int, 1)
		dp[0][0] = -prices[0]
	} else {
		dp = make([][2]int, k+1)
		t = make([][2]int, k+1)

		for i := 0; i < k; i++ {
			dp[i][0] = -prices[0]
		}
	}
	for i := 1; i < len(prices); i++ {
		if daily {
			t[0][0] = max(dp[0][0], dp[0][1]-prices[i])
			t[0][1] = max(dp[0][1], dp[0][0]+prices[i])
			dp[0][0], dp[0][1] = t[0][0], t[0][1]
		} else {
			for j := 0; j < k; j++ {
				t[j][0] = max(dp[j][0], dp[j+1][1]-prices[i])
				t[j][1] = max(dp[j][1], dp[j][0]+prices[i])
			}
			for j := 0; j < k; j++ {
				dp[j][0], dp[j][1] = t[j][0], t[j][1]
			}
		}
	}
	max := math.MinInt64
	if daily {
		max = dp[0][1]
	} else {
		for i := 0; i < k+1; i++ {
			if dp[i][1] > max {
				max = dp[i][1]
			}
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
