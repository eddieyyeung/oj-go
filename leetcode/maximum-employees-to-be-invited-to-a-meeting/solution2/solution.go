package solution2

func max(x, y int) int {
	if x > y {
		return x
	}
	return y
}

func maximumInvitations(favorite []int) int {
	l := len(favorite)
	indeg := make([]int, l)
	innums := make([][]int, l)
	visisted := make([]int, l)
	maxLine := make([]int, l)
	for i, n := range favorite {
		maxLine[i] = 1
		indeg[n]++
		innums[n] = append(innums[n], i)
	}
	var q []int
	for i, deg := range indeg {
		if deg == 0 {
			q = append(q, i)
		}
	}
	for len(q) > 0 {
		n := q[0]
		q = q[1:]
		next := favorite[n]
		indeg[next]--

		if indeg[next] == 0 {
			maxLine[next] = max(maxLine[n]+1, maxLine[next])
			q = append(q, next)
		}
	}
	var findLine = func(i int) int {
		var m int
		for _, inn := range innums[i] {
			if indeg[inn] != 0 {
				continue
			}
			if maxLine[inn] > m {
				m = maxLine[inn]
			}
		}
		return m
	}
	var ans2 int
	var ans int
	for i := 0; i < l; i++ {
		if visisted[i] != 0 || indeg[i] == 0 {
			continue
		}
		count := 1
		out := favorite[i]
		visisted[i] = 1
		for i != out {
			visisted[out] = 1
			out = favorite[out]
			count++
		}
		if count == 2 {
			ans2 += 2 + findLine(i) + findLine(favorite[i])
		} else {
			if count > ans {
				ans = count
			}
		}
	}
	if ans2 > ans {
		ans = ans2
	}
	return ans
}
