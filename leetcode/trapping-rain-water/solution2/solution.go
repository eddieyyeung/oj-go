// 42. 接雨水
// https://leetcode.cn/problems/trapping-rain-water/description/
package solution

// 从左往右遍历，并实现单调栈的同时，计算需要填的水量
func trap(height []int) int {
	n := len(height)
	var st []int
	var ans int
	for i := 0; i < n; i++ {
		for len(st) != 0 && height[st[len(st)-1]] <= height[i] {
			last := st[len(st)-1]
			st = st[:len(st)-1]
			if len(st) == 0 {
				break
			}
			pre := st[len(st)-1]
			ans += (min(height[pre], height[i]) - height[last]) * (i - pre - 1)
		}
		st = append(st, i)
	}
	return ans
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
