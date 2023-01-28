package smallest_string_with_a_given_numeric_value

func getSmallestString(n int, k int) string {
	s := make([]byte, n)
	for i := 0; i < n; i++ {
		u := n - i - 1
		can := k - u
		if can >= 26 {
			can = 26
			s[u] = 'z'
		} else {
			s[u] = 'a' + byte(can) - 1
		}
		k -= can
	}
	return string(s)
}
