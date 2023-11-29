// 1590. 使数组和能被 P 整除
// https://leetcode.cn/problems/make-sum-divisible-by-p/description/
package solution

func minSubarray(nums []int, p int) int {
	n := len(nums)
	presum := make([]int, n+1)
	for i := 0; i < n; i++ {
		presum[i+1] = presum[i] + nums[i]
	}
	t := presum[n] % p
	if presum[n] < p {
		return -1
	}
	if t == 0 {
		return 0
	}
	m := map[int]int{t: 0}
	var ans int = 1e9

	for i := 0; i < n; i++ {
		k := (presum[i+1]) % p
		if v, ok := m[k]; ok {
			ans = min(ans, i+1-v)
		}
		m[(presum[i+1]+t)%p] = i + 1
	}
	if ans == 1e9 || ans == n {
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
