// 51. N 皇后
// https://leetcode.cn/problems/n-queens/description/
package nqueens

import (
	"strings"
)

func solveNQueens(n int) [][]string {
	var ans [][]string

	var path []int

	var dfs func(i int)
	dfs = func(i int) {
		if i == n {
			rows := make([]string, 0, n)
			for _, col := range path {
				rows = append(rows, strings.Repeat(".", col)+"Q"+strings.Repeat(".", n-col-1))
			}
			ans = append(ans, rows)
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
