package merge_strings_alternately

func mergeAlternately(word1 string, word2 string) string {
	rst := make([]byte, len(word1)+len(word2))
	i := 0
	j := 0
	for i < len(word1) && i < len(word2) {
		rst[j] = word1[i]
		j++
		rst[j] = word2[i]
		j++
		i++
	}
	for i < len(word1) {
		rst[j] = word1[i]
		i++
		j++
	}
	for i < len(word2) {
		rst[j] = word2[i]
		i++
		j++
	}
	return string(rst)
}
