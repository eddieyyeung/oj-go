// 2246. 相邻字符不同的最长路径
// https://leetcode.cn/problems/longest-path-with-different-adjacent-characters/description/
package solution

func longestPath(parent []int, s string) int {
	n := len(parent)
	tree := make([][]int, n)
	for i := 1; i < n; i++ {
		tree[parent[i]] = append(tree[parent[i]], i)
	}
	var ans int
	var dfs func(i int) int
	dfs = func(i int) int {
		node := tree[i]
		depth := 0
		for _, child := range node {
			c_depth := dfs(child) + 1
			if s[child] != s[i] {
				ans = max(ans, depth+c_depth)
				depth = max(depth, c_depth)
			}
		}
		return depth
	}
	dfs(0)
	return ans + 1
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
