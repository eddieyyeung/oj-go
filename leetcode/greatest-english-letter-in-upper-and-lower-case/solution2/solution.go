package solution2

func greatestLetter(s string) string {
	m := make(map[rune]bool)
	for _, c := range s {
		m[c] = true
	}
	for u := 'Z'; u >= 'A'; u-- {
		l := rune(u - 'A' + 'a')
		if m[u] && m[l] {
			return string(u)
		}
	}
	return ""
}
