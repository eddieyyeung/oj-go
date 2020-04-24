// Package solution2 ...
// https://leetcode-cn.com/problems/friend-circles/
package solution2

func dfs(M [][]int, visited []int, i int) {
	for j := 0; j < len(M); j++ {
		if M[i][j] == 1 && visited[j] == 0 {
			visited[j] = 1
			dfs(M, visited, j)
		}
	}
}

func findCircleNum(M [][]int) int {
	visited := make([]int, len(M))
	count := 0
	for i := 0; i < len(M); i++ {
		if visited[i] == 0 {
			visited[i] = 1
			count++
			dfs(M, visited, i)
		}
	}
	return count
}
