package solution2

func findNumberIn2DArray(matrix [][]int, target int) bool {
	if len(matrix) == 0 || len(matrix[0]) == 0 {
		return false
	}
	row, column := len(matrix), len(matrix[0])
	i, j := 0, column-1
	for i < row && j >= 0 {
		n := matrix[i][j]
		if target > n {
			i++
		} else if target < n {
			j--
		} else {
			return true
		}
	}
	return false
}
