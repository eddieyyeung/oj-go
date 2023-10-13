package juzhenzhongdelujinglcof

var v [][]int
var row, column int

func wordPuzzle(grid [][]byte, target string) bool {
	row = len(grid)
	column = len(grid[0])
	v = make([][]int, row)
	for i := 0; i < len(v); i++ {
		v[i] = make([]int, column)
	}
	for i := 0; i < row; i++ {
		for j := 0; j < column; j++ {
			if wordSearch(grid, target, i, j, 0) {
				return true
			}
		}
	}
	return false
}

func wordSearch(grid [][]byte, target string, i, j int, idx int) bool {
	if i < 0 || i >= row || j < 0 || j >= column || v[i][j] == 1 {
		return false
	}
	v[i][j] = 1
	if grid[i][j] == target[idx] {
		if idx == len(target)-1 {
			return true
		}
		if wordSearch(grid, target, i-1, j, idx+1) {
			return true
		}
		if wordSearch(grid, target, i+1, j, idx+1) {
			return true
		}
		if wordSearch(grid, target, i, j-1, idx+1) {
			return true
		}
		if wordSearch(grid, target, i, j+1, idx+1) {
			return true
		}
	}

	v[i][j] = 0
	return false
}
