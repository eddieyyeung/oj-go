// 300. 最长递增子序列
// https://leetcode.cn/problems/longest-increasing-subsequence/description/
package solution1

// 枚举 nums[i] 作为 LIS 的末尾元素，dfs(i) 表示末尾元素为 nums[i] 的 LIS 长度
// dfs(i) = max(dfs(j))+1, j<i 且 nums[j]<nums[i]
func lengthOfLIS(nums []int) int {
	n := len(nums)
	dp := make([]int, n)
	for i := 0; i < n; i++ {
		for j := 0; j <= i; j++ {
			if nums[j] < nums[i] {
				dp[i] = max(dp[i], dp[j])
			}
		}
		dp[i]++
	}
	var ans int
	for i := 0; i < len(dp); i++ {
		ans = max(ans, dp[i])
	}
	return ans
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
