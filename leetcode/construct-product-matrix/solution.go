package constructproductmatrix

var mod int = 12345

func product(a, b, m int) int {
	return ((a % m) * (b % m)) % m
}

func constructProductMatrix(grid [][]int) [][]int {
	n, m := len(grid), len(grid[0])
	ans := make([][]int, n)
	p := 1
	for i := 0; i < n; i++ {
		ans[i] = make([]int, m)
		for j := 0; j < m; j++ {
			ans[i][j] = p
			p = product(p, grid[i][j], mod)
		}
	}
	p = 1
	for i := n - 1; i >= 0; i-- {
		for j := m - 1; j >= 0; j-- {
			ans[i][j] = product(ans[i][j], p, mod)
			p = product(p, grid[i][j], mod)
		}
	}
	return ans
}
