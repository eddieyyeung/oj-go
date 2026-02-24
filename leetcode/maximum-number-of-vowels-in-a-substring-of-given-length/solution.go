package maximum_number_of_vowels_in_a_substring_of_given_length

import (
	"slices"
)

func maxVowels(s string, k int) int {
	j := 0
	vowels := []byte{'a', 'e', 'i', 'o', 'u'}
	cnt := 0
	rst := 0
	for i := 0; i < len(s); i++ {
		c := s[i]
		if slices.Contains(vowels, c) {
			cnt++
		}
		for i-j+1 > k {
			if slices.Contains(vowels, s[j]) {
				cnt--
			}
			j++
		}
		if i-j+1 == k {
			rst = max(rst, cnt)
		}
	}
	return rst
}
