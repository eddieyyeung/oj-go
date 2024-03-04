// 918. 环形子数组的最大和
// https://leetcode.cn/problems/maximum-sum-circular-subarray/description/
package solution

func maxSubarraySumCircular(nums []int) int {
	n := len(nums)
	presums := make([]int, n*2+1)
	for i := 0; i < 2*n; i++ {
		presums[i+1] = presums[i] + nums[i%n]
	}

	var ans int = -1e9
	var q []int = []int{0}
	for i := 1; i < len(presums); i++ {
		if len(q) > 0 && i-q[0] > n {
			q = q[1:]
		}
		ans = max(ans, presums[i]-presums[q[0]])
		for len(q) > 0 && presums[i] < presums[q[len(q)-1]] {
			q = q[:len(q)-1]
		}

		q = append(q, i)

	}
	return ans
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
