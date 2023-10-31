package solution3

func findMin(nums []int) int {
	left, right := 0, len(nums)-1
	first := nums[0]
	for left <= right {
		mid := (left + right) / 2
		if nums[mid] < first {
			right = mid - 1
		} else {
			left = mid + 1
		}
	}
	if right+1 == len(nums) {
		return nums[0]
	}
	return nums[right+1]
}
