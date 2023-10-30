package solution2

func findMin(nums []int) int {
	left, right := 0, len(nums)-1
	last := nums[len(nums)-1]
	for left <= right {
		mid := (left + right) / 2
		if nums[mid] <= last {
			right = mid - 1
		} else {
			left = mid + 1
		}
	}
	return nums[right+1]
}
