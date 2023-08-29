package coursescheduleii

type Item struct {
	Next []int
	Deg  int
}

func findOrder(numCourses int, prerequisites [][]int) []int {
	items := make([]Item, numCourses)
	for _, prerequisite := range prerequisites {
		a := prerequisite[0]
		b := prerequisite[1]
		items[a].Deg++
		items[b].Next = append(items[b].Next, a)
	}
	var q []int
	for i, item := range items {
		if item.Deg == 0 {
			q = append(q, i)
		}
	}
	var ans []int
	visited := make([]bool, numCourses)
	for len(q) > 0 {
		i := q[0]
		q = q[1:]
		if visited[i] {
			return nil
		}
		visited[i] = true
		ans = append(ans, i)
		for _, n := range items[i].Next {
			items[n].Deg--
			if items[n].Deg == 0 {
				q = append(q, n)
			}
		}
	}
	if len(ans) != numCourses {
		return nil
	}
	return ans
}
