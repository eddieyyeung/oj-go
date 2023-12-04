// 35. 搜索插入位置
// https://leetcode.cn/problems/search-insert-position/description/
package solution

import "sort"

func searchInsert(nums []int, target int) int {
	return sort.Search(len(nums), func(i int) bool {
		return target <= nums[i]
	})
}
