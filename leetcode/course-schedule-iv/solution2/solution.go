package solution2

var v []bool
var nexts [][]int
var caches [][]bool

var l int

func dfs(x int) {
	if v[x] {
		return
	}
	v[x] = true
	for _, next := range nexts[x] {
		caches[x][next] = true
		dfs(next)
		for i := 0; i < l; i++ {
			caches[x][i] = caches[x][i] || caches[next][i]
		}
	}
}

func checkIfPrerequisite(numCourses int, prerequisites [][]int, queries [][]int) []bool {
	l = numCourses
	v = make([]bool, numCourses)
	nexts = make([][]int, numCourses)
	caches = make([][]bool, numCourses)
	for i := 0; i < numCourses; i++ {
		caches[i] = make([]bool, numCourses)
	}
	for _, prerequisite := range prerequisites {
		x, y := prerequisite[0], prerequisite[1]
		nexts[x] = append(nexts[x], y)
	}
	for i := 0; i < numCourses; i++ {
		dfs(i)
	}
	ans := make([]bool, 0, len(queries))
	for _, query := range queries {
		x, y := query[0], query[1]
		ans = append(ans, caches[x][y])
	}
	return ans
}
