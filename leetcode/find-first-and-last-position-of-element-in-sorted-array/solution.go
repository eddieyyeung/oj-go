package findfirstandlastpositionofelementinsortedarray

func searchRange(nums []int, target int) []int {
	start := lower_bound(nums, target)
	if start == len(nums) || nums[start] != target {
		return []int{-1, -1}
	}
	end := lower_bound(nums, target+1) - 1
	return []int{start, end}
}

// 寻找 nums 中大于或等于 target 的最小下标 index
func lower_bound(nums []int, target int) int {
	// [left, right] 左右闭合区间
	left, right := 0, len(nums)-1
	for left <= right {
		mid := (left + right) / 2
		mnum := nums[mid]
		if target <= mnum {
			right = mid - 1
		} else {
			left = mid + 1
		}
	}
	return right + 1
}
