package solution1

func numSpecial(mat [][]int) (ans int) {
	m, n := len(mat[0]), len(mat)
	for i := 0; i < n; i++ {
		for j := 0; j < m; j++ {
			if mat[i][j] == 1 {
				isSpecial := true
				for k := 0; k < n; k++ {
					if k == i {
						continue
					}
					if mat[k][j] == 1 {
						isSpecial = false
					}
				}
				if isSpecial {
					for k := 0; k < m; k++ {
						if k == j {
							continue
						}
						if mat[i][k] == 1 {
							isSpecial = false
						}
					}
				}
				if isSpecial {
					ans++
				}
			}
		}
	}
	return
}
