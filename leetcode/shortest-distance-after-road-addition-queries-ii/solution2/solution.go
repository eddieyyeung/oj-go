package solution

// https://leetcode.cn/problems/shortest-distance-after-road-addition-queries-ii/description/
func shortestDistanceAfterQueries(n int, queries [][]int) []int {
	roads := make([]int, n-1)
	for i := 0; i < len(roads); i++ {
		roads[i] = i + 1
	}
	dist := n - 1
	var ans []int = make([]int, 0, len(queries))
	for _, query := range queries {
		x, y := query[0], query[1]
		if roads[x] == -1 {
			ans = append(ans, dist)
			continue
		}
		k := roads[x]
		for k < y {
			roads[x] = roads[k]
			roads[k] = -1
			dist--
			k = roads[x]
		}
		ans = append(ans, dist)
	}
	return ans
}
