// 300. 最长递增子序列
// https://leetcode.cn/problems/longest-increasing-subsequence/description/
package solution1

func lengthOfLIS(nums []int) int {
	n := len(nums)
	dp := make([]int, n+1)
	var rst int = 0
	for i := 0; i < n; i++ {
		dp[i] = 1
		for j := 0; j < i; j++ {
			if nums[i] > nums[j] {
				dp[i] = max(dp[i], dp[j]+1)
			}
		}
		rst = max(rst, dp[i])
	}

	return rst
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
