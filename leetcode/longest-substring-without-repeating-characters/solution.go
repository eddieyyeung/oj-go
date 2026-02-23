// Package longest_substring_without_repeating_characters ...
// https://leetcode-cn.com/problems/longest-substring-without-repeating-characters/
package longest_substring_without_repeating_characters

func lengthOfLongestSubstring(s string) int {
	var j = 0
	m := make(map[byte]bool)
	rst := 0
	for i := 0; i < len(s); i++ {
		c := s[i]
		for m[c] && j < i {
			m[s[j]] = false
			j++
		}
		rst = max(rst, i-j+1)
		m[c] = true
	}
	return rst
}
