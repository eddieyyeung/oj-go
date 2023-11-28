// 862. 和至少为 K 的最短子数组
// https://leetcode.cn/problems/shortest-subarray-with-sum-at-least-k/description/
package solution

func shortestSubarray(nums []int, k int) int {
	n := len(nums)
	var ans int = 1e9
	rst_sum := make([]int, n+1)
	for i := 0; i < n; i++ {
		rst_sum[i+1] = rst_sum[i] + nums[i]
	}
	n_rst_sum := len(rst_sum)
	var q []int
	for i := 0; i < n_rst_sum; i++ {
		// 优化1：移除队尾无用数据，保持队列递增有序
		for len(q) > 0 && rst_sum[i] <= rst_sum[q[len(q)-1]] {
			q = q[:len(q)-1]
		}
		// 优化2：移除队头无用数据
		for len(q) > 0 && rst_sum[i]-rst_sum[q[0]] >= k {
			ans = min(ans, i-q[0])
			q = q[1:]
		}
		q = append(q, i)
	}
	if ans == 1e9 {
		return -1
	}
	return ans
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
