// 239. 滑动窗口最大值
// https://leetcode.cn/problems/sliding-window-maximum/description/
package solution

func maxSlidingWindow(nums []int, k int) []int {
	n := len(nums)
	ans := make([]int, 0, n-k+1)
	var q []int
	for i := 0; i < n; i++ {
		// 1. 入
		for len(q) > 0 && nums[q[len(q)-1]] <= nums[i] {
			q = q[:len(q)-1]
		}
		q = append(q, i)
		// 2. 出
		if i-k == q[0] {
			q = q[1:]
		}
		// 3. 记录答案
		if i >= k-1 {
			// 由于队首到队尾单调递减，所以窗口最大值就是队首
			ans = append(ans, nums[q[0]])
		}
	}
	return ans
}
