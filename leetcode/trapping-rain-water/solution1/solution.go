// 42. 接雨水
// https://leetcode.cn/problems/trapping-rain-water/description/
package solution

// 单独看每个点所能接收的水量，为：
// water[i] = i左边高度最大值 + i右边高度最大值 - i当前高度
func trap(height []int) int {
	n := len(height)
	l_st := make([]int, n)
	r_st := make([]int, n)
	var l_max, r_max int
	for i := 0; i < n; i++ {
		l_max = max(l_max, height[i])
		l_st[i] = l_max
		r_max = max(r_max, height[n-i-1])
		r_st[n-i-1] = r_max
	}
	var ans int
	for i := 0; i < n; i++ {
		ans += min(l_st[i], r_st[i]) - height[i]
	}
	return ans
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
