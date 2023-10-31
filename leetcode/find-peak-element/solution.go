package findpeakelement

import "math"

func search(nums []int) int {
	left, right := 0, len(nums)-1
	for left <= right {
		mid := right - (right-left)>>1
		mnum := nums[mid]
		rnum := -math.MaxInt
		if mid+1 <= len(nums)-1 {
			rnum = nums[mid+1]
		}
		if rnum >= mnum {
			left = mid + 1
		} else {
			right = mid - 1
		}
	}
	return left
}

func findPeakElement(nums []int) int {
	return search(nums)
}
