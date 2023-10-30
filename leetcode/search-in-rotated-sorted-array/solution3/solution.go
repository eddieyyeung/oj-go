package solution3

func search(nums []int, target int) int {
	s := nums[0]
	left, right := 0, len(nums)-1
	t := target
	for left <= right {
		mid := (left + right) / 2
		m := nums[mid]
		if m == t {
			return mid
		}
		if s <= m {
			if s <= t && t < m {
				right = mid - 1
			} else {
				left = mid + 1
			}
		} else {
			if m < t && t < s {
				left = mid + 1
			} else {
				right = mid - 1
			}
		}
	}
	return -1
}
