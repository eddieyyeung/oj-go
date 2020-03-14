// Package _sum ...
// https://leetcode-cn.com/problems/3sum/
package _sum

import (
	"sort"
)

func threeSum(nums []int) [][]int {
	var result [][]int
	sort.Slice(nums, func(i, j int) bool {
		return nums[i] < nums[j]
	})
	for i := 0; i < len(nums)-2; i++ {
		if nums[i] > 0 {
			break
		}
		if i > 0 && nums[i-1] == nums[i] {
			continue
		}
		l := i + 1
		r := len(nums) - 1
		remain := 0 - nums[i]
		for l < r {
			if nums[l]+nums[r] > remain {
				r--
			} else if nums[l]+nums[r] < remain {
				l++
			} else {
				result = append(result, []int{nums[i], nums[l], nums[r]})
				for l < r && nums[l+1] == nums[l] {
					l++
				}
				for l < r && nums[r-1] == nums[r] {
					r--
				}
				l++
				r--
			}
		}
	}
	return result
}
