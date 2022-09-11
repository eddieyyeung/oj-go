package solution

// problem: https://leetcode.cn/problems/count-unique-characters-of-all-substrings-of-a-given-string/

func uniqueLetterString(s string) (ans int) {
	m0 := [26]int{}
	m1 := [26]int{}
	for i := 0; i < 26; i++ {
		m0[i] = -1
		m1[i] = -1
	}
	for i, c := range s {
		idx := c - 'A'
		if m1[idx] == -1 {
			m1[idx] = i
			continue
		}
		ans += (m1[idx] - m0[idx]) * (i - m1[idx])
		m0[idx] = m1[idx]
		m1[idx] = i
	}
	for idx, _ := range m1 {
		if m1[idx] != -1 {
			ans += (m1[idx] - m0[idx]) * (len(s) - m1[idx])
		}
	}
	return ans
}
