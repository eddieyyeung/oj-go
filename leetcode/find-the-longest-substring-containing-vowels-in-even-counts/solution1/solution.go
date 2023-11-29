// 1371. 每个元音包含偶数次的最长子字符串
// https://leetcode.cn/problems/find-the-longest-substring-containing-vowels-in-even-counts/description/
package solution

func findTheLongestSubstring(s string) int {
	var aeioumask int = 1065233
	var mask int = 0
	m := map[int]int{0: 0}
	var ans int
	for i, ch := range s {
		mask ^= aeioumask & (1 << int(ch-'a'))
		if v, ok := m[mask]; ok {
			ans = max(ans, i+1-v)
		} else {
			m[mask] = i + 1
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
