// 2454. 下一个更大元素 IV
// https://leetcode.cn/problems/next-greater-element-iv/description/
package solution

func secondGreaterElement(nums []int) []int {
	n := len(nums)
	ans := make([]int, n)
	var s, t []int
	for i := 0; i < n; i++ {
		ans[i] = -1
		for len(t) > 0 && nums[t[len(t)-1]] < nums[i] {
			ans[t[len(t)-1]] = nums[i]
			t = t[:len(t)-1]
		}
		j := len(s)
		for j > 0 && nums[s[j-1]] < nums[i] {
			j--
		}
		t = append(t, s[j:]...)
		s = append(s[:j], i)
	}
	return ans
}
