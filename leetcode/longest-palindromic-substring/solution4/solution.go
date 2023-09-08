package solution4

var m string

func expand(l, r int, s string) (int, int) {
	for ; l >= 0 && r < len(s) && s[l] == s[r]; l, r = l-1, r+1 {
	}
	return l + 1, r - 1
}

func longestPalindrome(s string) string {
	m = s[:1]
	for i := 0; i < len(s); i++ {
		l1, r1 := expand(i, i, s)
		l2, r2 := expand(i, i+1, s)
		if l := r1 - l1 + 1; l > len(m) {
			m = s[l1 : r1+1]
		}
		if l := r2 - l2 + 1; l > len(m) {
			m = s[l2 : r2+1]
		}
	}
	return m
}
