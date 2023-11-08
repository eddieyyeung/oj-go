// 213. 打家劫舍 II
// https://leetcode.cn/problems/house-robber-ii/description/
package houserobberii

// 分类讨论，考虑是否偷 nums[0]：
// 如果偷 nums[0]，那么 nums[1] 和 nums[n−1] 不能偷，问题变成从 nums[2] 到 nums[n−2] 的非环形版本，调用 198 题的代码解决；
// 如果不偷 nums[0]，那么问题变成从 nums[1] 到 nums[n−1] 的非环形版本，同样调用 198 题的代码解决。
// 这两种方案覆盖了所有情况（毕竟 nums[0] 只有偷与不偷，没有第三种选择），所以取两种方案的最大值，即为答案。
func rob(nums []int) int {
	n := len(nums)
	return max(nums[0]+tryRob(nums, 2, n-2), tryRob(nums, 1, n-1))
}

func tryRob(nums []int, start, end int) int {
	dp0, dp1 := 0, 0
	for i := start; i <= end; i++ {
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
