package findminimuminrotatedsortedarray

// 二分法，比较第一个元素
func findMin(nums []int) int {
	left, right := 0, len(nums)-1
	target := nums[0]
	for left <= right {
		mid := (left + right) / 2
		if nums[mid] < target {
			right = mid - 1
		} else {
			left = mid + 1
		}
	}
	if left == len(nums) {
		return target
	}
	return nums[left]
}

// dfs 做法，不优雅
func findMin_2(nums []int) int {
	var dfs func(nums []int, left int, right int) int
	dfs = func(nums []int, left int, right int) int {
		if left > right {
			if left < len(nums) {
				return left
			} else {
				return right
			}
		}
		mid := (right + left) / 2
		mnum := nums[mid]
		lnum := nums[len(nums)-1]
		rnum := nums[0]
		if mid-1 >= 0 {
			lnum = nums[mid-1]
		}
		if mid+1 <= len(nums)-1 {
			rnum = nums[mid+1]
		}
		if lnum > mnum && mnum < rnum {
			return mid
		}
		lidx := dfs(nums, mid+1, right)
		ridx := dfs(nums, left, mid-1)
		if nums[lidx] < nums[ridx] {
			return lidx
		} else {
			return ridx
		}
	}

	return nums[dfs(nums, 0, len(nums)-1)]
}
