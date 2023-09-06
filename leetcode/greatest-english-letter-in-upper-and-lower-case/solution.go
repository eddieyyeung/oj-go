package greatest_english_letter_in_upper_and_lower_case

// 2309. 兼具大小写的最好英文字母

import "strings"

func greatestLetter(s string) string {
	mupper := [26]bool{}
	mlower := [26]bool{}
	greatest := ""
	for i, c := range s {
		u := c - 'A'
		l := c - 'a'
		t := strings.ToUpper(s[i : i+1])
		if u >= 0 && u < 26 {
			mupper[u] = true
			if mlower[u] && t > greatest {
				greatest = t
			}
		}
		if l >= 0 && l < 26 {
			mlower[l] = true
			if mupper[l] && t > greatest {
				greatest = t
			}
		}
	}
	return greatest
}
