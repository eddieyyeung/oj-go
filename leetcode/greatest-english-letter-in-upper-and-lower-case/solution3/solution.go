package solution3

import (
	"unicode"
)

func greatestLetter(s string) string {
	var u, l int
	for _, c := range s {
		if unicode.IsLower(c) {
			l |= 1 << (c - 'a')
		} else {
			u |= 1 << (c - 'A')
		}
	}
	for i := 25; i >= 0; i-- {
		if l&u>>i == 1 {
			return string(rune('A' + i))
		}
	}
	return ""
}
