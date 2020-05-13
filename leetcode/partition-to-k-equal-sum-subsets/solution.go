// Package partition_to_k_equal_sum_subsets ...
// https://leetcode-cn.com/problems/partition-to-k-equal-sum-subsets/
package partition_to_k_equal_sum_subsets

import (
	"sort"
)

func canPartitionKSubsets(nums []int, k int) bool {
	sum := 0
	for _, v := range nums {
		sum += v
	}
	if sum%k != 0 {
		return false
	}
	target := sum / k

	sort.Slice(nums, func(i, j int) bool {
		return nums[i] < nums[j]
	})
	return divideSet(make([]int, k), nums, len(nums)-1, target)
}
func divideSet(group []int, nums []int, row int, target int) bool {
	if row < 0 {
		return true
	}
	v := nums[row]
	for i := 0; i < len(group); i++ {
		if group[i]+v <= target {
			group[i] += v
			if divideSet(group, nums, row-1, target) {
				return true
			}
			group[i] -= v
		}
		if group[i] == 0 {
			break
		}
	}
	return false
}
