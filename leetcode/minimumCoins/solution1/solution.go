// 2944. 购买水果需要的最少金币数
// https://leetcode.cn/problems/minimum-number-of-coins-for-fruits/description/
package solution

// dfs_buy(i), dfs_nobuy(i) 分别表示第 i 个水果购买/不购买所需要的最少金币数，其中 i ∈ [1, n]
// 即返回两种状态
// 1. 第 i 个水果进行购买
// 2. 第 i 个水果不进行购买，由从 [⌈i/2⌉, i) 个水果进行购买
// =>
// dfs_buy(i) = min(dfs_buy(i-1), dfs_nobuy(i-1)) + prices[i]
// dfs_nobuy(i) = min(dfs_buy(j)...), 其中 j ∈ [⌈i/2⌉, i)
func minimumCoins(prices []int) int {
	n := len(prices)
	dp_buy := make([]int, n+1)
	dp_nobuy := make([]int, n+1)
	for i := 1; i < len(dp_buy); i++ {
		dp_buy[i] = -1
		dp_nobuy[i] = -1
	}
	dp_buy[1], dp_nobuy[1] = prices[0], 1e9
	for i := 2; i <= n; i++ {
		pre_buy, pre_nobuy := dp_buy[i-1], dp_nobuy[i-1]
		buy, nobuy := min(pre_buy, pre_nobuy)+prices[i-1], int(1e9)
		for j := (i + 1) / 2; j < i; j++ {
			j_buy := dp_buy[j]
			nobuy = min(nobuy, j_buy)
		}
		dp_buy[i], dp_nobuy[i] = buy, nobuy
	}
	return min(dp_buy[n], dp_nobuy[n])
	// var dfs func(i int) (int, int)
	// dfs = func(i int) (int, int) {
	// 	if i <= 1 {
	// 		return prices[0], 1e9
	// 	}
	// 	if dp_buy[i] != -1 {
	// 		return dp_buy[i], dp_nobuy[i]
	// 	}
	// 	pre_buy, pre_nobuy := dfs(i - 1)
	// 	buy, nobuy := min(pre_buy, pre_nobuy)+prices[i-1], int(1e9)
	// 	for j := (i + 1) / 2; j < i; j++ {
	// 		j_buy, _ := dfs(j)
	// 		nobuy = min(nobuy, j_buy)
	// 	}
	// 	dp_buy[i], dp_nobuy[i] = buy, nobuy
	// 	return buy, nobuy
	// }
	// buy, nobuy := dfs(n)
	// return min(buy, nobuy)
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
