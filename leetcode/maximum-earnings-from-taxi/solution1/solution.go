// 2008. 出租车的最大盈利
// https://leetcode.cn/problems/maximum-earnings-from-taxi/description/
package solution

func maxTaxiEarnings(n int, rides [][]int) int64 {
	type item struct{ start, end, tip int }
	group := make([][]item, n+1)
	for _, ride := range rides {
		start, end, tip := ride[0], ride[1], ride[2]
		group[end] = append(group[end], item{start, end, tip})
	}
	dp := make([]int, n+1)
	for i := 1; i <= n; i++ {
		rst := dp[i-1]
		for _, item := range group[i] {
			rst = max(rst, dp[item.start]+item.end-item.start+item.tip)
		}
		dp[i] = rst
	}
	return int64(dp[n])
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
