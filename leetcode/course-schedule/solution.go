package courseschedule

func canFinish(numCourses int, prerequisites [][]int) bool {
	outdeg := [2000]int{}
	m := [2000][]int{}

	for _, prerequisite := range prerequisites {
		x, y := prerequisite[0], prerequisite[1]
		outdeg[x]++
		m[y] = append(m[y], x)
	}

	l := []int{}
	for i := 0; i < numCourses; i++ {
		if outdeg[i] == 0 {
			l = append(l, i)
		}
	}
	count := 0
	for len(l) > 0 {
		count++
		n := l[0]
		l = l[1:]
		for _, t := range m[n] {
			outdeg[t]--
			if outdeg[t] == 0 {
				l = append(l, t)
			}
		}
	}
	return count == numCourses
}
