// 704. 二分查找
// https://leetcode.cn/problems/binary-search/description/
package solution

func search(nums []int, target int) int {
	left, right := 0, len(nums)-1
	for left <= right {
		mid := (left + right) / 2
		if target > nums[mid] {
			left = mid + 1
		} else {
			right = mid - 1
		}
	}
	if left == len(nums) || nums[left] != target {
		return -1
	}
	return left
}
