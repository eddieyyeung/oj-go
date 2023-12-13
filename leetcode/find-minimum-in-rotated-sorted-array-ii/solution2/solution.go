// 154. 寻找旋转排序数组中的最小值 II
// https://leetcode.cn/problems/find-minimum-in-rotated-sorted-array-ii/description/
package solution

// 解题思路：利用跟区间的最右边对比来判断二分
func findMin(nums []int) int {
	n := len(nums)
	left, right := 0, n-1
	for left <= right {
		mid := (left + right) / 2
		mid_num, right_num := nums[mid], nums[right]
		if mid_num > right_num {
			left = mid + 1
		} else if mid_num < right_num {
			// NOTE: 这里不是 mid-1，因为当前的 mid 可能就是最小值，[left, mid-1] 可能会忽略了该值
			right = mid
		} else {
			right--
		}
	}
	return nums[left]
}
