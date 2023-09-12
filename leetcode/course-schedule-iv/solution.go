package coursescheduleiv

// 输入：numCourses = 2, prerequisites = [[1,0]], queries = [[0,1],[1,0]]
// 输出：[false,true]
// 解释：课程 0 不是课程 1 的先修课程，但课程 1 是课程 0 的先修课程。
// 示例 2：

// 输入：numCourses = 2, prerequisites = [], queries = [[1,0],[0,1]]
// 输出：[false,false]
// 解释：没有先修课程对，所以每门课程之间是独立的。

// 输入：numCourses = 3, prerequisites = [[1,2],[1,0],[2,0]], queries = [[1,0],[1,2]]
// 输出：[true,true]

var nexts [][]int
var indegs []int
var caches [][]bool

var q []int

func checkIfPrerequisite(numCourses int, prerequisites [][]int, queries [][]int) []bool {
	nexts = make([][]int, numCourses)
	indegs = make([]int, numCourses)
	caches = make([][]bool, numCourses)
	q = []int{}

	for i := 0; i < numCourses; i++ {
		caches[i] = make([]bool, numCourses)
	}

	for _, prerequisite := range prerequisites {
		x, y := prerequisite[0], prerequisite[1]
		nexts[x] = append(nexts[x], y)
		indegs[y]++
	}
	for i, indeg := range indegs {
		if indeg == 0 {
			q = append(q, i)
		}
	}
	for len(q) > 0 {
		n := q[0]
		q = q[1:]
		for _, next := range nexts[n] {
			caches[n][next] = true
			for i := 0; i < numCourses; i++ {
				caches[i][next] = caches[i][next] || caches[i][n]
			}
			indegs[next]--
			if indegs[next] == 0 {
				q = append(q, next)
			}
		}
	}
	ans := make([]bool, 0, len(queries))
	for _, query := range queries {
		x, y := query[0], query[1]
		ans = append(ans, caches[x][y])
	}
	return ans
}
