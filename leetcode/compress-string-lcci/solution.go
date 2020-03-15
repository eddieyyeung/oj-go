// Package compress_string_lcci ...
// https://leetcode-cn.com/problems/compress-string-lcci/
package compress_string_lcci

import (
	"strconv"
	"strings"
)

func compressString(S string) string {
	if len(S) <= 1 {
		return S
	}
	p := 0

	str := strings.Builder{}
	str.WriteString(string(S[0]))
	c := 1
	for i := 1; i < len(S); i++ {
		if S[i] == S[p] {
			c++
		} else {
			p = i
			str.WriteString(strconv.Itoa(c))
			str.WriteString(string(S[p]))
			c = 1
		}
	}
	str.WriteString(strconv.Itoa(c))
	if str.Len() >= len(S) {
		return S
	} else {
		return str.String()
	}
}
