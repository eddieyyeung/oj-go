// 560. 和为 K 的子数组
// https://leetcode.cn/problems/subarray-sum-equals-k/description/
package solution

func subarraySum(nums []int, k int) int {
	n := len(nums)
	presum := 0
	m := map[int]int{0: 1}
	var ans int
	for i := 0; i < n; i++ {
		presum += nums[i]
		ans += m[presum-k]
		m[presum]++
	}
	return ans
}
