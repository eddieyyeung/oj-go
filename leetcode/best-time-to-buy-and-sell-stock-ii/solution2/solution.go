package solution2

func maxProfit(prices []int) int {
	var dp0, dp1 int
	dp0 = -prices[0]
	for i := 1; i < len(prices); i++ {
		price := prices[i]
		t0 := max(dp1-price, dp0)
		t1 := max(dp0+price, dp1)
		dp0, dp1 = t0, t1
	}
	return dp1
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
