// 525. 连续数组
// https://leetcode.cn/problems/contiguous-array/description/
package solution

func findMaxLength(nums []int) int {
	n := len(nums)
	presum := make([]int, n+1)
	for i := 0; i < n; i++ {
		presum[i+1] = presum[i] + (nums[i]*2 - 1)
	}
	m := map[int]int{0: 0}
	var ans int
	for i := 1; i <= n; i++ {
		if v, ok := m[presum[i]]; ok {
			ans = max(ans, i-v)
		} else {
			m[presum[i]] = i
		}
	}
	return ans
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
