package besttimetobuyandsellstockwithcooldown

// maxProfit https://leetcode-cn.com/problems/best-time-to-buy-and-sell-stock-with-cooldown/
func maxProfit(prices []int) int {
	if len(prices) == 0 {
		return 0
	}
	s0 := 0
	s1 := -prices[0]
	s2 := 0
	for i := 1; i < len(prices); i++ {
		t0 := s1 + prices[i]
		t1 := max(s1, s2-prices[i])
		t2 := max(s0, s2)
		s0, s1, s2 = t0, t1, t2
	}
	return max(s0, s2)
}

func max(nums ...int) int {
	max := nums[0]
	for i := 1; i < len(nums); i++ {
		if nums[i] > max {
			max = nums[i]
		}
	}
	return max
}
