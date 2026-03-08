package n_queens

// https://leetcode.cn/problems/n-queens/description/
func solveNQueens(n int) [][]string {
	path := make([]int, n)
	isValid := func(i, j int) bool {
		for k := 0; k < i; k++ {
			if path[k] == j {
				return false
			}
			if k+path[k] == i+j {
				return false
			}
			if k-path[k] == i-j {
				return false
			}
		}
		return true
	}
	genQueen := func() []string {
		var queen []string
		for _, x := range path {
			var s string
			for j := 0; j < x; j++ {
				s += "."
			}
			s += "Q"
			for j := 0; j < n-x-1; j++ {
				s += "."
			}
			queen = append(queen, s)
		}
		return queen
	}
	var ans [][]string
	onPath := make([]bool, n)
	var dfs func(i int)
	dfs = func(i int) {
		if i == n {
			ans = append(ans, genQueen())
			return
		}
		for j := 0; j < n; j++ {
			if !onPath[j] {
				if isValid(i, j) {
					onPath[j] = true
					path[i] = j
					dfs(i + 1)
					onPath[j] = false
				}
			}
		}
	}
	dfs(0)
	return ans
}
