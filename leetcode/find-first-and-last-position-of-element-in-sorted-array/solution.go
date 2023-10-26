package findfirstandlastpositionofelementinsortedarray

func searchRange(nums []int, target int) []int {
	start := lower_bound(nums, target)
	if start == len(nums) || nums[start] != target {
		return []int{-1, -1}
	}
	end := lower_bound(nums, target+1) - 1
	return []int{start, end}
}

func lower_bound(nums []int, target int) int {
	left, right := 0, len(nums)-1
	for left <= right {
		mid := (left + right) >> 1
		num := nums[mid]
		if target <= num {
			right = mid - 1
		} else {
			left = mid + 1
		}
	}
	return left
}
