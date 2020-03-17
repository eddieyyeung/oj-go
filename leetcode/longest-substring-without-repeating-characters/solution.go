// Package longest_substring_without_repeating_characters ...
// https://leetcode-cn.com/problems/longest-substring-without-repeating-characters/
package longest_substring_without_repeating_characters

func lengthOfLongestSubstring(s string) int {
	hash := map[rune]int{}
	longest := 0
	start := 0
	for i, char := range s {
		if v, ok := hash[char]; ok && v >= start {
			if l := i - start; l > longest {
				longest = l
			}
			start = v + 1
			hash[char] = i
		} else {
			hash[char] = i
		}
	}
	if l := len(s) - start; l > longest {
		return l
	}
	return longest
}
