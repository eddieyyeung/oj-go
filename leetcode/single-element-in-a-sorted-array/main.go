// https://leetcode.cn/problems/single-element-in-a-sorted-array
package singleelementinasortedarray

func singleNonDuplicate(nums []int) int {
	var n int = len(nums)
	var l, r int = 0, n - 1
	for l <= r {
		m := (l + r) / 2

		if m+1 < n {
			if nums[m+1] == nums[m] {
				if (n-m)&1 == 1 {
					l = m + 1
					continue
				} else {
					r = m - 1
					continue
				}
			}
		}
		if m-1 >= 0 {
			if nums[m-1] == nums[m] {
				if (n-m)&1 == 1 {
					r = m - 1
					continue
				} else {
					l = m + 1
					continue
				}
			}
		}
		l = m + 1
	}
	return nums[l-1]
}
