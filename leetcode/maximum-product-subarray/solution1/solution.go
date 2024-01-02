// 152. 乘积最大子数组
// https://leetcode.cn/problems/maximum-product-subarray/description
package solution

import (
	"math"
)

func maxProduct(nums []int) int {
	n := len(nums)
	var ans int = nums[0]

	dp_max, dp_min := nums[0], nums[0]

	for i := 1; i < n; i++ {
		num := nums[i]
		dp_max, dp_min = max(dp_max*num, dp_min*num, num), min(dp_max*num, dp_min*num, num)
		ans = max(dp_max, ans)
	}
	return ans
}

func max(nums ...int) int {
	var rst int = -math.MaxInt
	for _, num := range nums {
		if num > rst {
			rst = num
		}
	}
	return rst
}

func min(nums ...int) int {
	var rst int = math.MaxInt
	for _, num := range nums {
		if num < rst {
			rst = num
		}
	}
	return rst
}
