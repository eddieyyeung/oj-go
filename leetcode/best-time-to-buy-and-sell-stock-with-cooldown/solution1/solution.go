package solution

import "math"

// dfs(i,0) 表示第i天未持有不冷冻
// dfs(i,1) 表示第i天未持有且冷冻
// dfs(i,2) 表示第i天持有

// dfs(i,0) = max(dfs(i-1,0), dfs(i-1,1))
// dfs(i,1) = max(dfs(i-1,2)+prices[i])
// dfs(i,2) = max(dfs(i-1,2), dfs(i-1,0)-prices[i])
func maxProfit(prices []int) int {
	n := len(prices)
	var dp0, dp1, dp2 int = 0, -1e9, -1e9
	for i := 0; i < n; i++ {
		new_dp0, new_dp1, new_dp2 := dp0, dp1, dp2
		dp0, dp1, dp2 = max(new_dp0, new_dp1), new_dp2+prices[i], max(new_dp2, new_dp0-prices[i])
	}
	return max(dp0, dp1)
}

func max(nums ...int) int {
	var rst int = -math.MaxInt
	for _, num := range nums {
		if num > rst {
			rst = num
		}
	}
	return rst
}
