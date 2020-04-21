// Package search_in_rotated_sorted_array ...
// https://leetcode-cn.com/problems/search-in-rotated-sorted-array/
package search_in_rotated_sorted_array

func search(nums []int, target int) int {
	left := 0
	right := len(nums) - 1
	for left <= right {
		mid := (left + right) / 2
		if nums[mid] == target {
			return mid
		}
		if nums[left] <= nums[mid] {
			if target < nums[mid] && target >= nums[left] {
				right = mid - 1
			} else {
				left = mid + 1
			}
			continue
		}
		if nums[mid] <= nums[right] {
			if target > nums[mid] && target <= nums[right] {
				left = mid + 1
			} else {
				right = mid - 1
			}
			continue
		}
		return -1
	}
	return -1
}
