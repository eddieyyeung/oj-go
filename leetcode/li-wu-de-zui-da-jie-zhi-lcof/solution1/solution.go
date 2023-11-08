// LCR 166. 珠宝的最高价值
// https://leetcode.cn/problems/li-wu-de-zui-da-jie-zhi-lcof/description/
package liwudezuidajiezhilcof

func jewelleryValue(frame [][]int) int {
	n, m := len(frame), len(frame[0])
	dp := make([]int, m+1)
	for i := 0; i < n; i++ {
		for j := 0; j < m; j++ {
			dp[j+1] = max(dp[j]+frame[i][j], dp[j+1]+frame[i][j])
		}
	}
	return dp[m+1]
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
