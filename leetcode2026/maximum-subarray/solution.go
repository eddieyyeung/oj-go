package maximum_subarray

import "slices"

func maxSubArray(nums []int) int {
	n := len(nums)
	dp := make([]int, n)
	for i := 0; i < n; i++ {
		dp[i] = nums[i]
		if i-1 >= 0 {
			dp[i] = max(dp[i], dp[i-1]+nums[i])
		}
	}
	return slices.Max(dp)
}
