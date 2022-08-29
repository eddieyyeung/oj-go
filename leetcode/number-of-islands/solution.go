package number_of_islands

var (
	x int
	y int
)

func numIslands(grid [][]byte) int {
	x = len(grid)
	y = len(grid[0])
	pos := make([][]bool, x)
	for i := 0; i < x; i++ {
		pos[i] = make([]bool, y)
	}
	count := 0
	for i := 0; i < x; i++ {
		for j := 0; j < y; j++ {
			if result := findIsland(i, j, grid, pos); result {
				count++
			}
		}
	}
	return count
}

func findIsland(i int, j int, grid [][]byte, pos [][]bool) bool {
	if pos[i][j] {
		return false
	}
	pos[i][j] = true
	if grid[i][j] == '0' {
		return false
	} else if grid[i][j] == '1' {
		if i+1 < x {
			findIsland(i+1, j, grid, pos)
		}
		if j+1 < y {
			findIsland(i, j+1, grid, pos)
		}
		if i-1 >= 0 {
			findIsland(i-1, j, grid, pos)
		}
		if j-1 >= 0 {
			findIsland(i, j-1, grid, pos)
		}
	}
	return true
}
