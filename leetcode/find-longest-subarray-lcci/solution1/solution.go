// 面试题 17.05. 字母与数字
// https://leetcode.cn/problems/find-longest-subarray-lcci/description/
package solution

func findLongestSubarray(array []string) []string {
	n := len(array)
	presum := make([]int, n+1)
	for i := 0; i < n; i++ {
		presum[i+1] = presum[i] + isNumber(array[i])
	}
	m := map[int]int{0: 0}
	var ans []string
	for i := 1; i <= n; i++ {
		if v, ok := m[presum[i]]; ok {
			if i-v > len(ans) {
				ans = array[v:i]
			}
		} else {
			m[presum[i]] = i
		}
	}
	return ans
}

var mask = 1023

// number: 1
// letter: -1
func isNumber(str string) int {
	rst := mask >> (str[0] - '0') & 1
	return rst*2 - 1
}
