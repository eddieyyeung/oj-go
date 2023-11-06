// 51. N 皇后
// https://leetcode.cn/problems/n-queens/description/
package nqueens

import (
	"strings"
)

func solveNQueens2(n int) [][]string {
	var ans [][]string

	path := make([]int, n)
	onPath := make([]bool, n)
	// 记录右上斜线的情况，表示为 r+c
	diag1 := make([]bool, 2*n-1)
	// 记录左上斜线的情况，表示为 r-c。由于有可能出现负数，因此进行偏移为 r-c+(n-1)
	diag2 := make([]bool, 2*n-1)
	var dfs func(i int)
	dfs = func(r int) {
		if r == n {
			rows := make([]string, 0, n)
			for _, col := range path {
				rows = append(rows, strings.Repeat(".", col)+"Q"+strings.Repeat(".", n-col-1))
			}
			ans = append(ans, rows)
			return
		}
		for c, on := range onPath {
			rc := r - c + n - 1
			// 对以下情况进行判断：
			// 1. 该列不存在
			// 2. 右上斜线不存在
			// 3. 左上斜线不存在
			if !on && !diag1[r+c] && !diag2[rc] {
				diag1[r+c], diag2[rc], onPath[c] = true, true, true
				path[r] = c
				dfs(r + 1)
				diag1[r+c], diag2[rc], onPath[c] = false, false, false
			}
		}
	}
	dfs(0)
	return ans
}
