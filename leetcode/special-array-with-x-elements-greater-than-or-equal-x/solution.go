package special_array_with_x_elements_greater_than_or_equal_x

import (
	"sort"
)

func specialArray(nums []int) int {
	sort.Slice(nums, func(i, j int) bool {
		return nums[i] > nums[j]
	})
	t := -1
	for i := 0; i < len(nums); i++ {
		if nums[i] >= i+1 {
			t = i + 1
		} else {
			if nums[i] >= t {
				t = -1
				break
			}
		}
	}
	return t
}
