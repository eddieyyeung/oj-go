// 216. 组合总和 III
// https://leetcode.cn/problems/combination-sum-iii/description/
package combinationsumiii

func combinationSum3(k int, n int) [][]int {
	var ans [][]int
	var path []int
	var dfs func(i int, j int)
	dfs = func(i, j int) {
		// 剪枝：判断 (path + 剩余元素的数量) 是否小于 k 长度
		// 若是，则直接返回
		if len(path)+9-i+1 < k {
			return
		}
		// 剪枝：判断 (j + 剩余元素的总和) 是否小于 n
		// 若是，则直接返回
		if (i+9)*(9-i+1)/2+j < n {
			return
		}

		if len(path) == k || j == n {
			if len(path) == k && j == n {
				ans = append(ans, append([]int{}, path...))
			}
			return
		}
		dfs(i+1, j)

		path = append(path, i)
		dfs(i+1, j+i)
		path = path[:len(path)-1]
	}
	dfs(1, 0)
	return ans
}
