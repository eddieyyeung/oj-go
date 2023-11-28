// 1438. 绝对差不超过限制的最长连续子数组
// https://leetcode.cn/problems/longest-continuous-subarray-with-absolute-diff-less-than-or-equal-to-limit/description/
package solution

func longestSubarray(nums []int, limit int) int {
	var st_max []int
	var st_min []int
	n := len(nums)
	j := 0
	var ans int
	for i := 0; i < n; i++ {
		for len(st_max) > 0 && nums[st_max[len(st_max)-1]] <= nums[i] {
			st_max = st_max[:len(st_max)-1]
		}
		for len(st_min) > 0 && nums[st_min[len(st_min)-1]] >= nums[i] {
			st_min = st_min[:len(st_min)-1]
		}
		st_max = append(st_max, i)
		st_min = append(st_min, i)
		for nums[st_max[0]]-nums[st_min[0]] > limit {
			if j == st_max[0] {
				st_max = st_max[1:]
			}
			if j == st_min[0] {
				st_min = st_min[1:]
			}
			j++
		}
		ans = max(ans, i-j+1)
	}
	return ans
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
