// 1542. 找出最长的超赞子字符串
// https://leetcode.cn/problems/find-longest-awesome-substring/description/
package solution

func longestAwesome(s string) int {
	var mask int = 0
	var ans int = 0
	m := map[int]int{0: 0}
	for i, ch := range s {
		mask ^= 1 << (ch - '0')
		for j := 0; j < 10; j++ {
			if v, ok := m[mask^(1<<j)]; ok {
				ans = max(ans, i+1-v)
			}
		}
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
