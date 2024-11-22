package solution1

// https://leetcode.cn/problems/shortest-distance-after-road-addition-queries-i/description/
func shortestDistanceAfterQueries(n int, queries [][]int) []int {
	graph := make([][]int, n)
	for i := 0; i < n; i++ {
		graph[i] = append(graph[i], i+1)
	}
	costs := make([]int, n)
	for i := 0; i < n; i++ {
		costs[i] = i
	}
	ans := make([]int, 0, len(queries))

	var fn func(i int)

	fn = func(i int) {
		if i == n {
			return
		}
		for _, next := range graph[i] {
			if next == n {
				continue
			}
			if costs[i]+1 < costs[next] {
				costs[next] = costs[i] + 1
				fn(next)
			}
		}
	}

	for _, query := range queries {
		x, y := query[0], query[1]
		graph[x] = append(graph[x], y)
		if costs[x]+1 < costs[y] {
			costs[y] = costs[x] + 1
			fn(y)
		}
		ans = append(ans, costs[n-1])
	}
	return ans
}
