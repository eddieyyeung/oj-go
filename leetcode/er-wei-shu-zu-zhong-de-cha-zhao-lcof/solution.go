package erweishuzuzhongdechazhaolcof

// 在一个 n * m 的二维数组中，每一行都按照从左到右 非递减 的顺序排序，每一列都按照从上到下 非递减 的顺序排序。请完成一个高效的函数，输入这样的一个二维数组和一个整数，判断数组中是否含有该整数。

// 示例:

// 现有矩阵 matrix 如下：

// [
//   [1,   4,  7, 11, 15],
//   [2,   5,  8, 12, 19],
//   [3,   6,  9, 16, 22],
//   [10, 13, 14, 17, 24],
//   [18, 21, 23, 26, 30]
// ]
// 给定 target = 5，返回 true。

// 给定 target = 20，返回 false。

// 限制：

// 0 <= n <= 1000

// 0 <= m <= 1000
func findNumberIn2DArray(matrix [][]int, target int) bool {
	if len(matrix) == 0 || len(matrix[0]) == 0 {
		return false
	}
	row, column := len(matrix), len(matrix[0])

	rl := 0
	rr := row - 1
	for rl <= rr {
		rm := (rr + rl) / 2
		// fmt.Println(rl, rr, rm)
		n := matrix[rm][0]
		if n == target {
			return true
		} else if n > target {
			rr = rm - 1
		} else {
			rl = rm + 1
		}
	}
	for i := 0; i < rl; i++ {
		l, r := 0, column-1
		for l <= r {
			m := (r + l) / 2
			n := matrix[i][m]
			if n == target {
				return true
			} else if n > target {
				r = m - 1
			} else {
				l = m + 1
			}
		}
	}
	// fmt.Println("------------")
	return false
}
