// 300. 最长递增子序列
// https://leetcode.cn/problems/longest-increasing-subsequence/description/
package solution1

import (
	"sort"
)

// lis[i] 表示长度为 i+1 的 IS 的末尾元素最小值
// 例如:
// nums = [1,6,7,2,4,5,3]
//
//	lis = [1]
//	lis = [1,6]
//	lis = [1,6,7]
//	lis = [1,2,7]
//	lis = [1,2,4]
//	lis = [1,2,4,5]
//	lis = [1,2,3,5]
func lengthOfLIS(nums []int) int {
	n := len(nums)
	var lis []int
	for i := 0; i < n; i++ {
		idx := sort.Search(len(lis), func(k int) bool {
			return nums[i] <= lis[k]
		})
		if idx == len(lis) {
			lis = append(lis, nums[i])
		} else {
			lis[idx] = nums[i]
		}
	}
	return len(lis)
}
