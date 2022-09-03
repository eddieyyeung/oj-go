package maximum_rows_covered_by_columns

import (
	"math/bits"
)

func maximumRows(mat [][]int, cols int) (ans int) {
	n, m := len(mat), len(mat[0])
	barr := make([]int, n)
	// 转化为二进制数组
	for i := 0; i < n; i++ {
		for j := 0; j < m; j++ {
			barr[i] |= mat[i][j] << j
		}
	}

	for mask := 0; mask < 1<<m; mask++ {
		// 在 [0, 2^m) 区间内，枚举出被选中覆盖的列的所有情况
		if bits.OnesCount(uint(mask)) != cols {
			continue
		}

		zero := 0
		for j := 0; j < n; j++ {
			// 位运算，计算出当前值是否能被选中的列所覆盖
			if barr[j] == (barr[j] & mask) {
				zero++
			}
		}
		if zero > ans {
			ans = zero
		}
	}
	return ans
}
