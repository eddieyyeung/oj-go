package solution1

import (
	"sort"
)

func threeSum(nums []int) [][]int {
	sort.Slice(nums, func(i, j int) bool {
		return nums[i] < nums[j]
	})

	var ans [][]int

	for i := 0; i < len(nums)-2; i++ {
		if nums[i] > 0 {
			break
		}
		left := i + 1
		right := len(nums) - 1
		for right > left {
			rest := nums[i] + nums[right] + nums[left]
			if rest > 0 {
				right--
			} else if rest < 0 {
				left++
			} else {
				ans = append(ans, []int{nums[i], nums[left], nums[right]})
				for left < right && nums[left] == nums[left+1] {
					left++
				}
				for left < right && nums[right] == nums[right-1] {
					right--
				}
				left++
				right--
			}
		}
		for i < len(nums)-2 && nums[i] == nums[i+1] {
			i++
		}
	}
	return ans
}
