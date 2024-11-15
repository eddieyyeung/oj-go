package countsubstringsthatsatisfykconstraintii

func countKConstraintSubstrings(s string, k int, queries [][]int) []int64 {
	n := len(s)
	presum := make([]int, n+1)
	m := [2]int{}
	j := 0
	t := make([]int, n)
	p := 0
	for i := 0; i < n; i++ {
		m[s[i]-'0']++
		for m[0] > k && m[1] > k {
			m[s[j]-'0']--
			j++
			presum[p+1] = presum[p] + i - p
			p++
		}
		t[i] = j
	}
	for p < n {
		presum[p+1] = presum[p] + n - p
		p++
	}
	ans := []int64{}
	for _, query := range queries {
		l, r := query[0], query[1]
		maxl := max(l, t[r])
		p1 := (1 + r - maxl + 1) * (r - maxl + 1) / 2
		p2 := presum[maxl] - presum[l]
		ans = append(ans, int64(p1+p2))
	}
	return ans
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
