// 154. 寻找旋转排序数组中的最小值 II
// https://leetcode.cn/problems/find-minimum-in-rotated-sorted-array-ii/description/
package solution

// 解题思路：
// 利用跟最后一个点的对比来判断二分区间。
// 其中 to_last 表示跟最后一个点的距离内出现交错的可能的距离
// 在对比过程中，通过遍历跟最后一个点的比较来缩短 to_last
func findMin(nums []int) int {
	n := len(nums)
	left, right := 0, n-1
	last_num := nums[n-1]
	to_last := n
	for left <= right {
		mid := (left + right) / 2
		mid_num := nums[mid]
		if mid_num > last_num {
			left = mid + 1
		} else if mid_num < last_num {
			right = mid - 1
		} else {
			gap := n - 1 - mid + 1
			if gap >= to_last {
				left = mid + 1
			} else {
				for i := mid + 1; i < n-1; i++ {
					if nums[i] != last_num {
						to_last = n - 1 - i + 1
						break
					}
				}
				if gap >= to_last {
					left = mid + 1
				} else {
					right = mid - 1
				}
			}
		}
	}
	if left == n {
		return nums[left-1]
	}
	return nums[left]
}
