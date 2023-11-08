// 198. 打家劫舍
// https://leetcode.cn/problems/house-robber/description/
package house_robber

// dfs(i) = max(dfs(i-1), dfs(i-2) + nums[i])
func rob(nums []int) int {
	dp0, dp1 := 0, nums[0]
	for i := 1; i < len(nums); i++ {
		dp0, dp1 = dp1, max(dp0+nums[i], dp1)
	}
	return dp1
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
