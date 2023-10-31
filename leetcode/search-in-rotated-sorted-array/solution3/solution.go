package solution3

import "fmt"

// 二分查找 红蓝染色法
func isLeft(nums []int, mid, target int) bool {
	s, m, t := nums[0], nums[mid], target
	if s <= m {
		return s <= t && t <= m
	} else {
		return t <= m || s <= t
	}
}

func search(nums []int, target int) int {
	left, right := 0, len(nums)-1
	for left <= right {
		mid := (left + right) / 2
		if isLeft(nums, mid, target) {
			right = mid - 1
		} else {
			left = mid + 1
		}
	}
	fmt.Println("right:", right)
	if right+1 == len(nums) || nums[right+1] != target {
		return -1
	}
	return right + 1
}
