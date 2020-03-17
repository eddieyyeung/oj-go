// Package reverse_words_in_a_string ...
// https://leetcode-cn.com/problems/reverse-words-in-a-string/
package reverse_words_in_a_string

import (
	"regexp"
	"strings"
)

func reverse(s []string) []string {
	for i, j := 0, len(s)-1; i < len(s)/2; i, j = i+1, j-1 {
		s[i], s[j] = s[j], s[i]
	}
	return s
}

func reverseWords(s string) string {
	return strings.Join(reverse(regexp.MustCompile(`\s+`).Split(strings.TrimSpace(s), -1)), " ")
}
