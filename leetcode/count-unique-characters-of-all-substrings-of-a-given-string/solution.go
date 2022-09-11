package count_unique_characters_of_all_substrings_of_a_given_string

// problem: https://leetcode.cn/problems/count-unique-characters-of-all-substrings-of-a-given-string/

func uniqueLetterString(s string) (ans int) {
	letters := [26][]int{}
	for i, c := range s {
		idx := c - 'A'
		if len(letters[idx]) == 0 {
			letters[idx] = append(letters[idx], -1)
		}
		letters[idx] = append(letters[idx], i)
	}

	for _, letter := range letters {
		letter = append(letter, len(s))
		for i := 1; i < len(letter)-1; i++ {
			ans += (letter[i] - letter[i-1]) * (letter[i+1] - letter[i])
		}
	}
	return ans
}
