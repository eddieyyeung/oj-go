package check_if_matrix_is_x_matrix

func checkXMatrix(grid [][]int) bool {
	length := len(grid)
	for x := range grid {
		for y := range grid[x] {
			if x == y || x+y == length-1 {
				if grid[x][y] == 0 {
					return false
				}
			} else {
				if grid[x][y] != 0 {
					return false
				}
			}
		}
	}
	return true
}
