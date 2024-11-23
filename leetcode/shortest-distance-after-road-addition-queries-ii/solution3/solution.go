package solution

// https://leetcode.cn/problems/shortest-distance-after-road-addition-queries-ii/description/
// 并查集 Union Find
func shortestDistanceAfterQueries(n int, queries [][]int) []int {
	fa := make([]int, n-1)
	for i := range fa {
		fa[i] = i
	}
	dist := n - 1
	find := func(x int) int {
		fx := x
		for fa[fx] != fx {
			fx = fa[fx]
		}
		for fa[x] != fx {
			fa[x], x = fx, fa[x]
		}
		return fx
	}

	var ans []int = make([]int, 0, len(queries))
	for _, query := range queries {
		x, y := query[0], query[1]-1
		fy := find(y)
		for i := find(x); i < y; i = find(i + 1) {
			fa[i] = fy
			dist--
		}
		ans = append(ans, dist)
	}
	return ans
}
