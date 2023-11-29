// 974. 和可被 K 整除的子数组
// https://leetcode.cn/problems/subarray-sums-divisible-by-k/description/
package solution

func subarraysDivByK(nums []int, k int) int {
	n := len(nums)
	presum := 0
	m := map[int]int{0: 1}
	var ans int
	for i := 0; i < n; i++ {
		presum += nums[i]
		k := ((presum % k) + k) % k
		ans += m[k]
		m[k]++
	}
	return ans
}
