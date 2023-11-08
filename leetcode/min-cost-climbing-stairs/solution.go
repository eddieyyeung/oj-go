// 746. 使用最小花费爬楼梯
// https://leetcode.cn/problems/min-cost-climbing-stairs/description/
package mincostclimbingstairs

func minCostClimbingStairs(cost []int) int {
	cost = append(cost, 0)
	dp0, dp1 := cost[0], cost[1]
	for i := 2; i < len(cost); i++ {
		dp0, dp1 = dp1, min(dp0+cost[i], dp1+cost[i])
	}
	return dp1
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
