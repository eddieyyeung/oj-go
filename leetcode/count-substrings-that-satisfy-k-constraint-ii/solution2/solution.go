package countsubstringsthatsatisfykconstraintii

func countKConstraintSubstrings(s string, k int, queries [][]int) []int64 {
	n := len(s)
	presum := make([]int, n+1)
	m := [2]int{}
	j := 0
	right := make([]int, n)
	ans := []int64{}
	for i := 0; i < n; i++ {
		right[i] = n - 1
	}
	for i := 0; i < n; i++ {
		m[s[i]-'0']++
		for m[0] > k && m[1] > k {
			m[s[j]-'0']--
			right[j] = i - 1
			j++
		}
		presum[i+1] = presum[i] + m[0] + m[1]
	}

	for _, query := range queries {
		l, r := query[0], query[1]
		minr := min(r, right[l])
		p1 := (minr - l + 2) * (minr - l + 1) / 2
		p2 := presum[r+1] - presum[minr+1]
		ans = append(ans, int64(p1+p2))
	}
	return ans
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
