// Package max_area_of_island ...
// https://leetcode-cn.com/problems/max-area-of-island/
package max_area_of_island

var direction = [][]int{
	{-1, 0},
	{1, 0},
	{0, -1},
	{0, 1},
}
var row int
var column int

func maxAreaOfIsland(grid [][]int) int {
	if len(grid) == 0 {
		return 0
	}
	row = len(grid)
	column = len(grid[0])
	maxArea := 0
	for i := 0; i < len(grid); i++ {
		for j := 0; j < len(grid[i]); j++ {
			if grid[i][j] == 1 {
				area := 0
				bt(grid, i, j, &area, &maxArea)
			}
		}
	}
	return maxArea
}

func bt(grid [][]int, i int, j int, area *int, maxArea *int) {
	grid[i][j] = 2
	*area += 1
	if *area > *maxArea {
		*maxArea = *area
	}
	for t := 0; t < len(direction); t++ {
		x := i + direction[t][0]
		y := j + direction[t][1]
		if x >= 0 && x < row && y >= 0 && y < column && grid[x][y] == 1 {
			bt(grid, x, y, area, maxArea)
		}
	}
}
