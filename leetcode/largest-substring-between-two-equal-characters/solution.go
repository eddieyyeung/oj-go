package largest_substring_between_two_equal_characters

import "strings"

func maxLengthBetweenEqualCharacters(s string) (ans int) {
	l := len(s)
	ans = -1
	for i := 0; i < l; i++ {
		idxA := strings.Index(s, s[i:i+1])
		idxB := strings.LastIndex(s, s[i:i+1])
		if tmp := idxB - idxA - 1; tmp > ans {
			ans = tmp
		}

	}
	return
}
