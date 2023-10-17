package solution2

import "strings"

func shortestBeautifulSubstring(s string, k int) string {
	if strings.Count(s, "1") < k {
		return ""
	}
	cnt := 0
	l := 0
	ans := s
	for r := 0; r < len(s); r++ {
		if s[r] == '1' {
			cnt++
		}
		for cnt > k || s[l] == '0' {
			if s[l] == '1' {
				cnt--
			}
			l++
		}
		if cnt == k {
			ss := s[l : r+1]
			if len(ss) < len(ans) || (len(ss) == len(ans) && ss < ans) {
				ans = ss
			}
		}
	}
	return ans
}
