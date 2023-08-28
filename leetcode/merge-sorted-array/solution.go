package mergesortedarray

import "sort"

func merge(nums1 []int, m int, nums2 []int, n int) {
	nums1 = nums1[:m]
	nums1 = append(nums1, nums2[:n]...)
	sort.Slice(nums1, func(i, j int) bool {
		return nums1[i] < nums1[j]
	})
}
