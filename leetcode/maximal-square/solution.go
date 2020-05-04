// Package maximal_square ...
// https://leetcode-cn.com/problems/maximal-square/
package maximal_square

import (
	"math"
)

func maximalSquare(matrix [][]byte) int {
	if len(matrix) == 0 {
		return 0
	}
	dp := make([][]int, len(matrix))
	for i := 0; i < len(matrix); i++ {
		dp[i] = make([]int, len(matrix[i]))
	}
	max := math.MinInt64
	for i := 0; i < len(matrix[0]); i++ {
		dp[0][i] = int(matrix[0][i] - byte('0'))
		if dp[0][i] > max {
			max = dp[0][i]
		}
	}
	for i := 1; i < len(matrix); i++ {
		for j := 0; j < len(matrix[i]); j++ {
			if j == 0 {
				dp[i][j] = int(matrix[i][j] - byte('0'))
				if dp[i][j] > max {
					max = dp[i][j]
				}
				continue
			}
			min := minF(dp[i][j-1], dp[i-1][j-1], dp[i-1][j])
			if matrix[i][j] == byte('1') {
				min++
			} else {
				min = 0
			}
			dp[i][j] = min
			if min > max {
				max = min
			}
		}
	}
	return max * max
}

func minF(arr ...int) int {
	min := math.MaxInt64
	for _, v := range arr {
		if v < min {
			min = v
		}
	}
	return min
}
