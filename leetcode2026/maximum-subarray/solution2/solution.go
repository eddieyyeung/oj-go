package maximum_subarray

import "math"

func maxSubArray(nums []int) int {
	sum := 0
	min_sum := 0
	ans := math.MinInt

	for i := 0; i < len(nums); i++ {
		sum += nums[i]
		ans = max(ans, sum-min_sum)
		min_sum = min(min_sum, sum)
	}
	return ans
}
