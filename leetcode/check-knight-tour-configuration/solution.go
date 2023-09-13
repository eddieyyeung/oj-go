package checkknighttourconfiguration

var nextps = [][]int{
	{-1, -2},
	{-2, -1},
	{-2, 1},
	{-1, 2},
	{1, 2},
	{2, 1},
	{2, -1},
	{1, -2},
}

func checkValidGrid(grid [][]int) bool {
	l := len(grid)
	// for _, g := range grid {
	// 	for _, n := range g {
	// 		fmt.Printf("%.2d ", n)
	// 	}
	// 	fmt.Println()
	// }
	if grid[0][0] != 0 {
		return false
	}
	total := l * l

	i := 0
	isFind := func(x, y int, next int) bool {
		if x < 0 || x >= l || y < 0 || y >= l {
			return false
		}
		if grid[x][y] == next {
			return true
		}
		return false
	}
	x, y := 0, 0
	// for i := 0; i < l; i++ {
	// 	for j := 0; j < l; j++ {
	// 		if grid[i][j] == 0 {
	// 			x, y = i, j
	// 		}
	// 	}
	// }
	for i < total-1 {
		next := i + 1
		for _, nextp := range nextps {
			nx, ny := x+nextp[0], y+nextp[1]
			if isFind(nx, ny, next) {
				x, y = nx, ny
				i++
				break
			}
		}
		if i != next {
			return false
		}
	}
	return true
}
