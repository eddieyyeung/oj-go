package queensthatcanattacktheking

func queensAttacktheKing(queens [][]int, king []int) [][]int {
	x, y := king[0], king[1]
	q := [8][8]int{}
	for _, queen := range queens {
		q[queen[0]][queen[1]] = 1
	}
	var ans [][]int

	findQueen := func(i, j int) bool {
		if q[i][j] == 1 {
			ans = append(ans, []int{i, j})
			return true
		}
		return false
	}
	var i, j int
	i, j = x-1, y
	for i >= 0 {
		if findQueen(i, j) {
			break
		}
		i--
	}
	i, j = x-1, y+1
	for i >= 0 && j < 8 {
		if findQueen(i, j) {
			break
		}
		i--
		j++
	}
	i, j = x, y+1
	for j < 8 {
		if findQueen(i, j) {
			break
		}
		j++
	}
	i, j = x+1, y+1
	for i < 8 && j < 8 {
		if findQueen(i, j) {
			break
		}
		i++
		j++
	}
	i, j = x+1, y
	for i < 8 {
		if findQueen(i, j) {
			break
		}
		i++
	}
	i, j = x+1, y-1
	for i < 8 && j >= 0 {
		if findQueen(i, j) {
			break
		}
		i++
		j--
	}
	i, j = x, y-1
	for j >= 0 {
		if findQueen(i, j) {
			break
		}
		j--
	}
	i, j = x-1, y-1
	for i >= 0 && j >= 0 {
		if findQueen(i, j) {
			break
		}
		i--
		j--
	}
	return ans
}
