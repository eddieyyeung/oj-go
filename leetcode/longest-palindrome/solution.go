package longest_palindrome

func longestPalindrome(s string) int {
	oddCount := 0
	m := map[rune]int{}
	for _, r := range s {
		m[r]++
	}
	for _, v := range m {
		if v%2 == 1 {
			oddCount++
		}
	}
	if oddCount > 0 {
		return len(s) - oddCount + 1
	}
	return len(s)
}
