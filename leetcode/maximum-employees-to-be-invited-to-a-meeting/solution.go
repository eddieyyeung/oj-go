package maximumemployeestobeinvitedtoameeting

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func maximumInvitations(favorite []int) int {
	l := len(favorite)
	indeg := make([]int, l)
	for _, v := range favorite {
		indeg[v]++
	}
	var q []int
	for i, v := range indeg {
		if v == 0 {
			q = append(q, i)
		}
	}
	maxDepth := make([]int, l)
	for len(q) > 0 {
		v := q[0]
		q = q[1:]
		indeg[favorite[v]]--
		maxDepth[favorite[v]] = max(maxDepth[favorite[v]], maxDepth[v]+1)
		if v1 := indeg[favorite[v]]; v1 == 0 {
			q = append(q, favorite[v])
		}
	}
	maxSize := 0
	tSize := 0
	for i, v := range indeg {
		if v != 1 {
			continue
		}
		indeg[i] = -1
		size := 1
		t := i
		for indeg[favorite[t]] != -1 {
			t = favorite[t]
			indeg[t] = -1
			size++
		}
		if size == 2 {
			maxSize += maxDepth[i] + maxDepth[favorite[i]] + 2
		} else if size > tSize {
			tSize = size
		}
	}
	return max(maxSize, tSize)
}
