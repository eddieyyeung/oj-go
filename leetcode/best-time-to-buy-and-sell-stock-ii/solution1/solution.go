// 122. 买卖股票的最佳时机 II
// https://leetcode.cn/problems/best-time-to-buy-and-sell-stock-ii/description/
package solution

// 状态机 DP
// 定义 dfs(i,0) 表示到第 i 天结束时，未持有股票的最大利润
// 定义 dfs(i,1) 表示到第 i 天结束时，持有股票的最大利润
// =>
// dfs(i,0) = max(dfs(i-1,0), dfs(i-1,1)+prices[i])
// dfs(i-1,0): 表示在 i-1 天未持有股票，之后什么也不做
// dfs(i-1,1)+prices[i]: 表示在 i-1 天持有股票，之后卖出 prices[i]
//
// dfs(i,1) = max(dfs(i-1,1), dfs(i-1,0)-prices[i])
// dfs(i-1,1): 表示在 i-1 天持有股票，之后什么也不做
// dfs(i-1,0)-prices[i]: 表示在 i-1 天未持有股票，之后买入 prices[i]
//
// 边界条件
// dfs(-1,0) = 0
// dfs(-1,1) = -inf 不合法
func maxProfit(prices []int) int {
	n := len(prices)
	var dp0, dp1 int = 0, -1e9
	for i := 0; i < n; i++ {
		new_dp0, new_dp1 := dp0, dp1
		dp1 = max(new_dp1, new_dp0-prices[i])
		dp0 = max(new_dp0, new_dp1+prices[i])
	}
	return dp0
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
