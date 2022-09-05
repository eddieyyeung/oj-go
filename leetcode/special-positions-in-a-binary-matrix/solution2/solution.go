package solution2

func numSpecial(mat [][]int) (ans int) {
	m, n := len(mat[0]), len(mat)
	pos := make([]int, m)
	for i := 0; i < n; i++ {
		sum := 0
		for j := 0; j < m; j++ {
			sum += mat[i][j]
		}
		for j := 0; j < m; j++ {
			if mat[i][j] == 1 {
				pos[j] += sum + 1
			}
		}
	}
	for i := 0; i < m; i++ {
		if pos[i] == 2 {
			ans++
		}
	}
	return
}
