// 18. 四数之和
// https://leetcode.cn/problems/4sum/description/
package solution

import "sort"

func fourSum(nums []int, target int) [][]int {
	n := len(nums)
	sort.Ints(nums)
	var ans [][]int
	type item struct {
		a int
		b int
		c int
		d int
	}
	m := map[item]bool{}
	for i := 0; i < n-3; i++ {
		a := nums[i]
		for j := i + 1; j < n-2; j++ {
			b := nums[j]
			l, r := j+1, n-1
			for l < r {
				c, d := nums[l], nums[r]
				sum := a + b + nums[l] + nums[r]
				if sum == target {
					if _, ok := m[item{a, b, c, d}]; !ok {
						ans = append(ans, []int{a, b, c, d})
					}
					m[item{a, b, c, d}] = true
					l++
					r--
				} else if sum > target {
					r--
				} else {
					l++
				}
			}
		}
	}
	return ans
}
