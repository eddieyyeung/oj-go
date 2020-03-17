// Package find_words_that_can_be_formed_by_characters ...
// https://leetcode-cn.com/problems/find-words-that-can-be-formed-by-characters/
package find_words_that_can_be_formed_by_characters

func countCharacters(words []string, chars string) int {
	m := map[rune]int{}
	for _, v := range chars {
		m[v]++
	}
	count := 0
	for _, word := range words {
		tm := map[rune]int{}
		isTrue := true
		for _, l := range word {
			if _, ok := m[l]; ok {
				tm[l]++
				if tm[l] > m[l] {
					isTrue = false
					continue
				}
			} else {
				isTrue = false
				continue
			}
		}
		if isTrue {
			count += len(word)
		}
	}
	return count
}
