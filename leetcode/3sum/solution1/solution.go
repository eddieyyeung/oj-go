// 15. 三数之和
// https://leetcode.cn/problems/3sum/description/
package solution

import (
	"sort"
)

// 时间复杂度 O(n^2)，其中 n 为 nums 的长度。排序 O(nlogn)。
// 外层枚举第一个数，就变成 167. 两数之和 II - 输入有序数组 了。做法是 O(n) 的双指针。
// 空间复杂度 O(1)
func threeSum(nums []int) [][]int {
	sort.Slice(nums, func(i, j int) bool {
		return nums[i] < nums[j]
	})
	n := len(nums)
	var ans [][]int
	for i := 0; i < n-2; i++ {
		// 从 i = 1 开始，跳过重复的数字
		if i > 0 && nums[i] == nums[i-1] {
			continue
		}
		// 优化 1：如果当前数和头部两个数相加比 0 大，那就跳过
		if nums[i]+nums[i+1]+nums[i+2] > 0 {
			break
		}
		// 优化 2：如果当前数和末尾两个数相加比 0 小，那就直接跳过
		if nums[i]+nums[n-2]+nums[n-1] < 0 {
			continue
		}
		l := i + 1
		r := n - 1
		for l < r {
			sum := nums[i] + nums[l] + nums[r]
			if sum == 0 {
				ans = append(ans, []int{nums[i], nums[l], nums[r]})
				for l++; l < r && nums[l] == nums[l-1]; l++ {
				}
				for r--; l < r && nums[r] == nums[r+1]; r-- {
				}
			} else if sum < 0 {
				l++
			} else {
				r--
			}
		}
	}
	return ans
}
