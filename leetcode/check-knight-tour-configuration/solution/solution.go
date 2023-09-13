package solution

// 思路：
// 1. 遍历 grid，将步骤顺序存放在队列 moves 中
// 2. 遍历 moves 队列，寻找每个步骤是否满足骑士的八种行动路线
// 3. 骑士的八种行动路线可以简化为两种情况：
//    a. |x1-x2|=1, |y1-y2|=2
//    b. |x1-x2|=2, |y1-y2|=1
//    即：(x1-x2)^2 * (y1-y2)^2 = 4 或 (x1-x2)^2 + (y1-y2)^2 = 5

func checkValidGrid(grid [][]int) bool {
	if grid[0][0] != 0 {
		return false
	}
	n := len(grid)
	moves := make([][2]int, n*n)
	for i := 0; i < n; i++ {
		for j := 0; j < n; j++ {
			moves[grid[i][j]][0] = i
			moves[grid[i][j]][1] = j
		}
	}
	for i := 1; i < n*n; i++ {
		x, y := moves[i-1][0], moves[i-1][1]
		nx, ny := moves[i][0], moves[i][1]
		if (x-nx)*(x-nx)+(y-ny)*(y-ny) != 4 {
			return false
		}
	}
	return true
}
