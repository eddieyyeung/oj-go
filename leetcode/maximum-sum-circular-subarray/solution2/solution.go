// 918. 环形子数组的最大和
// https://leetcode.cn/problems/maximum-sum-circular-subarray/description/
package solution

import "math"

func maxSubarraySumCircular(nums []int) int {
	var presum int = 0
	var min_presum, max_presum int
	var max_s, min_s int = -math.MaxInt, 0
	for _, num := range nums {
		presum += num
		max_s = max(max_s, presum-min_presum)
		min_s = min(min_s, presum-max_presum)

		min_presum = min(min_presum, presum)
		max_presum = max(max_presum, presum)
	}
	if presum == min_s {
		return max_s
	}
	return max(max_s, presum-min_s)
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
