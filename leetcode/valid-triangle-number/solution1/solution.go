// 611. Valid Triangle Number
// https://leetcode.cn/problems/valid-triangle-number/description/
package solution

import (
	"slices"
	"sort"
)

func triangleNumber(nums []int) int {
	n := len(nums)
	sort.Ints(nums)
	var ans int
	for i := 0; i < n-2; i++ {
		for j := i + 1; j < n-1; j++ {
			two_sum := nums[i] + nums[j]
			idx, _ := slices.BinarySearchFunc(nums[j+1:n], two_sum, func(i1, i2 int) int {
				return i1 - i2
			})
			ans += idx
		}
	}
	return ans
}
