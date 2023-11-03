// 77. 组合
// https://leetcode.cn/problems/combinations/description/
package combinations

// 回溯：答案视角
func combine(n int, k int) [][]int {
	var ans [][]int
	var path []int

	var dfs func(i int)

	dfs = func(i int) {
		if n-i+1+len(path) < k {
			return
		}
		if len(path) == k {
			ans = append(ans, append([]int{}, path...))
			return
		}
		if n-i+1+len(path) < k {
			return
		}
		for j := i; j <= n; j++ {
			path = append(path, j)
			dfs(j + 1)
			path = path[:len(path)-1]
		}
	}
	dfs(1)
	return ans
}
