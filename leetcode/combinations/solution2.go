// 77. 组合
// https://leetcode.cn/problems/combinations/description/
package combinations

// 回溯：输入视角
func combine2(n int, k int) [][]int {
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

		dfs(i + 1)
		path = append(path, i)
		dfs(i + 1)
		path = path[:len(path)-1]
	}
	dfs(1)
	return ans
}
