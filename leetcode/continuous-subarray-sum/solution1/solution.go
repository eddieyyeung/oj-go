// 523. 连续的子数组和
// https://leetcode.cn/problems/continuous-subarray-sum/description/
package solution

func checkSubarraySum(nums []int, k int) bool {
	n := len(nums)
	m := map[int]int{0: 0}
	var presum int = 0
	for i := 1; i <= n; i++ {
		presum += nums[i-1]
		if ii, ok := m[presum%k]; ok {
			if i-ii >= 2 {
				return true
			}
		} else {
			m[presum%k] = i
		}
	}
	return false
}
