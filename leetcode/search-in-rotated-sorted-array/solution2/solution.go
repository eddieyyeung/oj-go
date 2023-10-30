package solution2

// 二分查找 红蓝染色法
func isLeft(nums []int, mid, target int) bool {
	l, m, t := nums[len(nums)-1], nums[mid], target
	if l < m {
		return l < t && t <= m
	} else {
		return t <= m || l < t
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
	if nums[right+1] != target {
		return -1
	}
	return right + 1
}
