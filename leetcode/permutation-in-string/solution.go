// Package permutation_in_string ...
// https://leetcode-cn.com/problems/permutation-in-string/
package permutation_in_string

func checkInclusion(s1 string, s2 string) bool {
	if len(s1) > len(s2) {
		return false
	}
	m1 := map[rune]int{}
	l := len(s1)
	dup := 0
	for _, c := range s1 {
		m1[c]++
	}
	for i := 0; i < len(s1); i++ {
		if _, ok := m1[rune(s2[i])]; ok {
			l--
			m1[rune(s2[i])]--
			if m1[rune(s2[i])] < 0 {
				dup++
			}
		}
	}
	if dup == 0 && l == 0 {
		return true
	}
	for i := len(s1); i < len(s2); i++ {
		si := s2[i-len(s1)]
		ei := s2[i]
		if _, ok := m1[rune(si)]; ok {
			l++
			m1[rune(si)]++
			if m1[rune(si)] <= 0 {
				dup--
			}
		}
		if _, ok := m1[rune(ei)]; ok {
			l--
			m1[rune(ei)]--
			if m1[rune(ei)] < 0 {
				dup++
			}
		}
		if dup == 0 && l == 0 {
			return true
		}
	}

	return false
}
