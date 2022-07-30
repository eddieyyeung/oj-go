package solution

func wordSubsets(words1 []string, words2 []string) []string {
	var result []string
	m2 := make(map[byte]int)
	for _, word2 := range words2 {
		m := make(map[byte]int)
		for i := range word2 {
			m[word2[i]]++
		}
		for k, v := range m {
			max := MaxInt(v, m2[k])
			m2[k] = max
		}
	}
	for _, word1 := range words1 {
		m := make(map[byte]int)
		isSubset := true
		for i := range word1 {
			l := word1[i]
			m[l]++
		}
		if isSubset {
			for k, v := range m2 {
				if m[k] < v {
					isSubset = false
					break
				}
			}
		}
		if isSubset {
			result = append(result, word1)
		}
	}
	return result
}

func MaxInt(arr ...int) int {
	if len(arr) == 1 {
		return arr[0]
	}
	max := arr[0]
	for i := 1; i < len(arr); i++ {
		if arr[i] > max {
			max = arr[i]
		}
	}
	return max
}
