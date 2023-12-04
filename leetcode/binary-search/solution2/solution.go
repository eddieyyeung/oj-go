// 704. 二分查找
// https://leetcode.cn/problems/binary-search/description/
package solution

import "sort"

func search(nums []int, target int) int {
	idx := sort.Search(len(nums), func(i int) bool {
		return target <= nums[i]
	})
	if idx == len(nums) || nums[idx] != target {
		return -1
	}
	return idx
}
