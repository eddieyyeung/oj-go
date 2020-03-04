package longestcommonsubsequence

// LCS ...
func LCS(s1, s2 string) string {
	return l(s1, s2, len(s1)-1, len(s2)-1)
}

func l(s1 string, s2 string, m int, n int) string {
	if m < 0 || n < 0 {
		return ""
	}

	if s1[m] == s2[n] {
		return l(s1, s2, m-1, n-1) + string(s1[m])
	}
	l1 := l(s1, s2, m-1, n)
	l2 := l(s1, s2, m, n-1)
	if len(l1) > len(l2) {
		return l1
	}
	return l2
}
