// Package longest_continuous_increasing_subsequence ...
// https://leetcode-cn.com/problems/longest-continuous-increasing-subsequence/
package longest_continuous_increasing_subsequence

func findLengthOfLCIS(nums []int) int {
	if len(nums) < 1 {
		return 0
	}
	p := 0
	lcisLength := 1
	for i := 1; i < len(nums); i++ {
		if nums[i] <= nums[i-1] {
			if l := i - p; l > lcisLength {
				lcisLength = l
			}
			p = i
		}
	}
	if l := len(nums) - p; l > lcisLength {
		lcisLength = l
	}
	return lcisLength
}
