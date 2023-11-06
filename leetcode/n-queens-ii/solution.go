// 52. N 皇后 II
// https://leetcode.cn/problems/n-queens-ii/description/
package nqueensii

func totalNQueens(n int) int {
	var ans int
	var path []int

	var dfs func(i int)
	dfs = func(i int) {
		if i == n {
			ans++
			return
		}

		for cIdx := 0; cIdx < n; cIdx++ {
			isFound := false
			for r, c := range path {
				if i == r ||
					cIdx == c ||
					c-cIdx == r-i ||
					c-cIdx == i-r {
					isFound = true
					break
				}
			}
			if !isFound {
				path = append(path, cIdx)
				dfs(i + 1)
				path = path[:len(path)-1]
			}
		}
	}
	dfs(0)
	return ans
}
