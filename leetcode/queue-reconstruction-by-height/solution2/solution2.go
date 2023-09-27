package solution2

import "sort"

func reconstructQueue(people [][]int) [][]int {
	sort.Slice(people, func(i, j int) bool {
		a, b := people[i], people[j]
		return a[0] < b[0] || a[0] == b[0] && a[1] > b[1]
	})

	ans := make([][]int, len(people))
	for _, person := range people {
		space := person[1]
		for i := 0; i < len(ans); i++ {
			if ans[i] == nil {
				if space == 0 {
					ans[i] = person
					break
				} else {
					space--
				}
			}
		}
	}
	return ans
}
