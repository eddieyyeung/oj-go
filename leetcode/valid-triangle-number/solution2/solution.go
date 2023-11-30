// 611. Valid Triangle Number
// https://leetcode.cn/problems/valid-triangle-number/description/
package solution

import (
	"sort"
)

func triangleNumber(nums []int) int {
	n := len(nums)
	sort.Ints(nums)
	var ans int
	for i := 0; i < n-2; i++ {
		k := i + 2
		for j := i + 1; j < n-1; j++ {
			two_sum := nums[i] + nums[j]
			for k < n && nums[k] < two_sum {
				k++
			}
			ans += max(k-j-1, 0)
		}
	}
	return ans
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
